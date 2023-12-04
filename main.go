package main

import (
	"Todo/pkg/util"
	"Todo/repository/cache"
	"Todo/repository/db/dao"
	"Todo/route"
)

func main() {
	util.InitLog()
	dao.InitMySQL()
	cache.LinkRedis()

	r := route.NewRouter()
	_ = r.Run(":9090")
}
