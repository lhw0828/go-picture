package dao

import (
	"database/sql"
	"picture/common/errorx"
	"picture/rpc/space-rpc/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SpaceDao struct {
	conn sqlx.SqlConn
}

func NewSpaceDao(conn sqlx.SqlConn) *SpaceDao {
	return &SpaceDao{conn: conn}
}

func (d *SpaceDao) Insert(space *model.Space) (sql.Result, error) {
	query := `insert into space(spaceName, spaceLevel, spaceType, maxSize, maxCount, totalSize, totalCount,
		userId, createTime, editTime, updateTime) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	return d.conn.Exec(query,
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

func (d *SpaceDao) UpdateUsage(spaceId int64, size int64, operation string) error {
	var query string
	if operation == "add" {
		query = `update space set totalSize = totalSize + ?, totalCount = totalCount + 1 
                 where id = ? and maxSize >= totalSize + ?`
	} else {
		query = `update space set totalSize = totalSize - ?, totalCount = totalCount - 1 
                 where id = ? and totalSize >= ?`
	}

	// 打印 SQL 语句和参数，方便调试
	logx.Infof("SQL: %s, params: [size=%d, spaceId=%d]", query, size, spaceId)

	result, err := d.conn.Exec(query, size, spaceId, size)
	if err != nil {
		logx.Errorf("执行更新失败: %v", err)
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		logx.Errorf("获取影响行数失败: %v", err)
		return err
	}

	if affected == 0 {
		return errorx.NewCodeError(errorx.SpaceNotEnough, "空间容量不足或操作无效")
	}

	return nil
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
