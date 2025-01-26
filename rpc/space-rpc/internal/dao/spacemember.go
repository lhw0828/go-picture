package dao

import (
	"database/sql"
	"picture/rpc/space-rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SpaceMemberDao struct {
	conn sqlx.SqlConn
}

func NewSpaceMemberDao(conn sqlx.SqlConn) *SpaceMemberDao {
	return &SpaceMemberDao{conn: conn}
}

func (d *SpaceMemberDao) Insert(member *model.SpaceMember) (sql.Result, error) {
	query := `insert into space_user(spaceId, userId, spaceRole, createTime, updateTime) 
		values (?, ?, ?, ?, ?)`
	return d.conn.Exec(query, member.SpaceId, member.UserId, member.SpaceRole,
		member.CreateTime, member.UpdateTime)
}

func (d *SpaceMemberDao) FindBySpaceIdAndUserId(spaceId, userId int64) (*model.SpaceMember, error) {
	var member model.SpaceMember
	query := `select * from space_user where spaceId = ? and userId = ? limit 1`
	err := d.conn.QueryRow(&member, query, spaceId, userId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (d *SpaceMemberDao) FindBySpaceId(spaceId int64) ([]*model.SpaceMember, error) {
	var members []*model.SpaceMember
	query := `select * from space_user where spaceId = ?`
	err := d.conn.QueryRows(&members, query, spaceId)
	return members, err
}

func (d *SpaceMemberDao) CheckMember(spaceId, userId int64) (*model.SpaceMember, error) {
	var member model.SpaceMember
	query := `select * from space_member where spaceId = ? and userId = ? and isDelete = 0 limit 1`
	err := d.conn.QueryRow(&member, query, spaceId, userId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &member, nil
}
