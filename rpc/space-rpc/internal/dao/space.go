package dao

import (
	"context"
	"database/sql"
	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SpaceDao struct {
	conn sqlx.SqlConn
}

func NewSpaceDao(conn sqlx.SqlConn) *SpaceDao {
	return &SpaceDao{conn: conn}
}

func (d *SpaceDao) Insert(ctx context.Context, space *model.Space) (sql.Result, error) {
	query := `insert into space(spaceName, spaceLevel, spaceType, maxSize, maxCount, totalSize, totalCount,
		userId, createTime, editTime, updateTime) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	return d.conn.ExecCtx(ctx, query,
		space.SpaceName,
		space.SpaceLevel,
		space.SpaceType,
		space.MaxSize,
		space.MaxCount,
		space.TotalSize,
		space.TotalCount,
		space.UserId,
		space.CreateTime,
		space.EditTime,
		space.UpdateTime)
}

func (d *SpaceDao) FindById(id int64) (*model.Space, error) {
	var space model.Space

	// 打印 SQL 语句和参数，方便调试
	query := `SELECT * FROM space WHERE id = ? AND isDelete = 0`
	logx.Infof("SQL: %s, params: %d", query, id)

	err := d.conn.QueryRow(&space, query, id)
	if err != nil {
		if err == sqlx.ErrNotFound {
			logx.Infof("空间不存在, id: %d", id)
			return nil, nil
		}
		logx.Errorf("查询空间失败: %v", err)
		return nil, err
	}

	// 打印查询结果，方便调试
	logx.Infof("查询结果: %+v", space)

	return &space, nil
}

func (d *SpaceDao) GetUsage(spaceId int64) (int64, int64, error) {
	var space struct {
		MaxSize   int64 `db:"maxSize"`
		TotalSize int64 `db:"totalSize"`
	}

	query := `select maxSize, totalSize from space where id = ? and isDelete = 0`
	logx.Infof("SQL: %s, params: %d", query, spaceId)

	err := d.conn.QueryRow(&space, query, spaceId)
	if err != nil {
		logx.Errorf("获取空间使用量失败: %v", err)
		return 0, 0, err
	}

	logx.Infof("查询结果: maxSize=%d, totalSize=%d", space.MaxSize, space.TotalSize)
	return space.MaxSize, space.TotalSize, nil
}

// ExistsByUserIdAndType 检查用户是否已有同类型空间
func (d *SpaceDao) ExistsByUserIdAndType(ctx context.Context, userId int64, spaceType int32) (bool, error) {
	query := `select count(*) from space where userId = ? and spaceType = ? and isDelete = 0`

	var count int
	err := d.conn.QueryRowCtx(ctx, &count, query, userId, spaceType)
	if err != nil {
		logx.Errorf("查询用户空间失败: userId=%d, spaceType=%d, err=%v", userId, spaceType, err)
		return false, err
	}

	return count > 0, nil
}

// 在 SpaceDao 中添加删除方法
func (d *SpaceDao) Delete(ctx context.Context, id int64) error {
	query := `update space set isDelete = 1, updateTime = ? where id = ?`
	result, err := d.conn.ExecCtx(ctx, query, time.Now(), id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errorx.NewCodeError(errorx.DeleteSpaceFailed, "删除空间失败")
	}
	return nil
}

// Update 更新空间信息
func (d *SpaceDao) Update(ctx context.Context, space *model.Space) error {
	query := `update space set spaceName = ?, spaceLevel = ?, maxSize = ?, maxCount = ?,
		editTime = ?, updateTime = ? where id = ? and isDelete = 0`

	// 打印 SQL 语句和参数，方便调试
	logx.Infof("SQL: %s, params: %+v", query, space)

	result, err := d.conn.ExecCtx(ctx, query,
		space.SpaceName,
		space.SpaceLevel,
		space.MaxSize,
		space.MaxCount,
		time.Now(),
		time.Now(),
		space.Id)
	if err != nil {
		logx.Errorf("更新空间失败: %v", err)
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errorx.NewCodeError(errorx.UpdateSpaceFailed, "更新空间失败")
	}

	return nil
}

// Count 获取空间总数
func (d *SpaceDao) Count(ctx context.Context, spaceName string, spaceType int32) (int64, error) {
	query := "select count(*) from space where isDelete = 0"
	params := make([]interface{}, 0)

	if spaceName != "" {
		query += " and spaceName like ?"
		params = append(params, "%"+spaceName+"%")
	}
	if spaceType >= 0 {
		query += " and spaceType = ?"
		params = append(params, spaceType)
	}

	logx.Infof("SQL: %s, params: %v", query, params)

	var count int64
	err := d.conn.QueryRowCtx(ctx, &count, query, params...)
	if err != nil {
		logx.Errorf("获取空间总数失败: %v", err)
		return 0, err
	}

	return count, nil
}

// List 分页查询空间列表
func (d *SpaceDao) List(ctx context.Context, current, pageSize int64, spaceName string, spaceType int32) ([]*model.Space, error) {
	query := "select * from space where isDelete = 0"
	params := make([]interface{}, 0)

	if spaceName != "" {
		query += " and spaceName like ?"
		params = append(params, "%"+spaceName+"%")
	}
	if spaceType >= 0 {
		query += " and spaceType = ?"
		params = append(params, spaceType)
	}

	query += " order by createTime desc limit ?, ?"
	params = append(params, (current-1)*pageSize, pageSize)

	logx.Infof("SQL: %s, params: %v", query, params)

	var spaces []*model.Space
	err := d.conn.QueryRowsCtx(ctx, &spaces, query, params...)
	if err != nil {
		logx.Errorf("分页查询空间列表失败: %v", err)
		return nil, err
	}

	return spaces, nil
}

// GetUserPermissions 获取用户在空间中的权限
func (d *SpaceDao) GetUserPermissions(ctx context.Context, spaceId, userId int64) ([]string, error) {
	// 1. 获取空间信息
	space, err := d.FindById(spaceId)
	if err != nil {
		return nil, err
	}
	if space == nil {
		return nil, errorx.NewCodeError(errorx.NotFoundError, "空间不存在")
	}

	// 2. 获取用户角色
	var userRole string
	err = d.conn.QueryRowCtx(ctx, &userRole, "select userRole from user where id = ? and isDelete = 0", userId)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, errorx.NewCodeError(errorx.NotFoundError, "用户不存在")
		}
		return nil, err
	}

	// 3. 获取权限列表
	permissions := make([]string, 0)

	// 3.1 如果是管理员，拥有所有权限
	if userRole == "admin" {
		permissions = append(permissions,
			"space:view",     // 查看空间
			"space:edit",     // 编辑空间
			"space:delete",   // 删除空间
			"space:upload",   // 上传文件
			"space:download", // 下载文件
			"space:share",    // 分享文件
			"space:manage",   // 空间管理
		)
		return permissions, nil
	}

	// 3.2 如果是空间所有者，拥有除管理权限外的所有权限
	if space.UserId == userId {
		permissions = append(permissions,
			"space:view",     // 查看空间
			"space:edit",     // 编辑空间
			"space:delete",   // 删除空间
			"space:upload",   // 上传文件
			"space:download", // 下载文件
			"space:share",    // 分享文件
		)
		return permissions, nil
	}

	// 3.3 如果是团队空间，查询用户在空间中的权限
	if space.SpaceType == 1 {
		var spaceRole string
		err = d.conn.QueryRowCtx(ctx, &spaceRole,
			"select spaceRole from space_user where spaceId = ? and userId = ? and isDelete = 0",
			spaceId, userId)
		if err != nil && err != sqlx.ErrNotFound {
			return nil, err
		}
		if err == nil {
			switch spaceRole {
			case "admin":
				permissions = append(permissions,
					"space:view",
					"space:edit",
					"space:upload",
					"space:download",
					"space:share",
					"space:manage",
				)
			case "editor":
				permissions = append(permissions,
					"space:view",
					"space:edit",
					"space:upload",
					"space:download",
					"space:share",
				)
			case "viewer":
				permissions = append(permissions,
					"space:view",
					"space:download",
				)
			}
			return permissions, nil
		}
	}

	// 3.4 如果是私有空间且不是所有者，只有查看权限
	if space.SpaceType == 0 {
		permissions = append(permissions, "space:view")
	}

	return permissions, nil
}
