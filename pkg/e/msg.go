package e

var MsgFlags = map[int]string{
	SUCCESS: "操作成功",
	ERROR:   "操作失败",

	InvalidParams: "请求参数有误",

	ErrorUserExist:    "用户已存在",
	ErrorUserNotExist: "用户不存在",

	//token
	ErrorTokenFail:     "Token有误",
	ErrorTokenTimeout:  "Token超时",
	TokenGeneratedFail: "token生成失败",

	ErrorDatabase: "数据库操作有误",
	ErrorRedis:    "redis操作有误",

	SetPasswordFail: "密码生成失败",

	//user
	ErrorUserCreate:  "用户数据库存储失败",
	ErrorPassword:    "密码错误",
	ErrorGetUserInfo: "获取用户信息失败",

	//task
	ErrorTaskCreate:    "任务数据库存储失败",
	ErrorTaskNotExists: "任务不存在",
	ErrorTaskUpdate:    "任务更新错误",
}

// GetMsg 获取错误码对应的信息
func GetMsg(code int) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[ERROR]
}
