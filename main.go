package main

import (
	"Todo/conf"
	"Todo/pkg/util"
	"Todo/repository/cache"
	"Todo/repository/db/dao"
	"Todo/route"
)

// 接口文档：https://apifox.com/apidoc/shared-cac07b03-37fd-47b8-b0fb-0d4a8939adfc
func main() {
	conf.LoadMysql()
	conf.LoadRedis()
	util.InitLog()
	dao.InitMySQL()
	cache.LinkRedis()

	r := route.NewRouter()
	_ = r.Run(":9090")
}
