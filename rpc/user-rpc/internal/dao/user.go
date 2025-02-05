package dao

import (
	"context"
	"database/sql"
	"picture/rpc/user-rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserDao struct {
	conn sqlx.SqlConn
}

func NewUserDao(conn sqlx.SqlConn) *UserDao {
	return &UserDao{
		conn: conn,
	}
}

// 添加分页查询方法
func (d *UserDao) FindByPage(ctx context.Context, offset, limit int64) ([]*model.User, error) {
	var users []*model.User
	query := `select * from user where isDelete = 0 order by id desc limit ?, ?`
	err := d.conn.QueryRowsCtx(ctx, &users, query, offset, limit)
	return users, err
}

// 添加统计方法
func (d *UserDao) Count(ctx context.Context) (int64, error) {
	var count int64
	query := `select count(*) from user where isDelete = 0`
	err := d.conn.QueryRowCtx(ctx, &count, query)
	return count, err
}

func (d *UserDao) FindById(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	query := `select * from user where id = ? and isDelete = 0 limit 1`
	err := d.conn.QueryRowCtx(ctx, &user, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserDao) FindOne(ctx context.Context, id int64) (*model.User, error) {
	return d.FindById(ctx, id)
}

func (d *UserDao) FindByUserAccount(ctx context.Context, userAccount string) (*model.User, error) {
	var user model.User
	query := `select * from user where userAccount = ? and isDelete = 0 limit 1`
	err := d.conn.QueryRowCtx(ctx, &user, query, userAccount)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserDao) Insert(ctx context.Context, user *model.User) (sql.Result, error) {
	query := `insert into user(userAccount, userPassword, userName, userRole, userAvatar, userProfile, createTime, updateTime) 
              values (?, ?, ?, ?, ?, ?, ?, ?)`
	return d.conn.ExecCtx(ctx, query,
		user.UserAccount,
		user.UserPassword,
		user.UserName,
		user.UserRole,
		user.UserAvatar,
		user.UserProfile,
		user.CreateTime,
		user.UpdateTime)
}

func (d *UserDao) Update(ctx context.Context, user *model.User) error {
	query := `update user set userName=?, userAvatar=?, userProfile=?, updateTime=? where id=?`
	_, err := d.conn.ExecCtx(ctx, query,
		user.UserName,
		user.UserAvatar,
		user.UserProfile,
		user.UpdateTime,
		user.Id)
	return err
}

func (d *UserDao) Delete(ctx context.Context, id int64) error {
	query := `update user set isDelete=1 where id=?`
	_, err := d.conn.ExecCtx(ctx, query, id)
	return err
}
