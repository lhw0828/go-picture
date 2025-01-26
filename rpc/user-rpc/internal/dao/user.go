package dao

import (
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

func (d *UserDao) Insert(user *model.User) (sql.Result, error) {
	query := `insert into user(userAccount, userPassword, userName, userRole) values (?, ?, ?, ?)`
	return d.conn.Exec(query, user.UserAccount, user.UserPassword, user.UserName, user.UserRole)
}
