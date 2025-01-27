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

func (d *UserDao) FindById(id int64) (*model.User, error) {
	var user model.User
	query := `select * from user where id = ? and isDelete = 0 limit 1`
	err := d.conn.QueryRow(&user, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserDao(conn sqlx.SqlConn) *UserDao {
	return &UserDao{
		conn: conn,
	}
}

func (d *UserDao) FindByUserAccount(userAccount string) (*model.User, error) {
	var user model.User
	query := `select * from user where userAccount = ? and isDelete = 0 limit 1`
	err := d.conn.QueryRow(&user, query, userAccount)
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

// 查询总数
func (d *UserDao) Count(ctx context.Context) (int64, error) {
    var count int64
    query := `select count(*) from user where isDelete = 0`
    err := d.conn.QueryRowCtx(ctx, &count, query)
    if err != nil {
        return 0, err
    }
    return count, nil
}

// 分页查询
func (d *UserDao) FindByPage(ctx context.Context, offset, pageSize int64) ([]*model.User, error) {
    var users []*model.User
    query := `select * from user where isDelete = 0 order by createTime desc limit ?, ?`
    err := d.conn.QueryRowsCtx(ctx, &users, query, offset, pageSize)
    if err != nil {
        return nil, err
    }
    return users, nil
}
