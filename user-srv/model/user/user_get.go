package user

import (
	"github.com/eopenio/basic/db"
	proto "github.com/eopenio/template/user-srv/proto/user"
	log "github.com/micro/go-micro/v2/logger"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`

	// 获取数据库
	o := db.GetDB()

	ret = &proto.User{}

	// 查询
	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Errorf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}

func (s *service) QueryUserById(userId string) (ret *proto.User, err error) {
	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_id = ?`

	// 获取数据库
	o := db.GetDB()

	ret = &proto.User{}

	// 查询
	err = o.QueryRow(queryString, userId).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Errorf("[QueryUserById] 查询数据失败，err：%s", err)
		return
	}
	return
}
