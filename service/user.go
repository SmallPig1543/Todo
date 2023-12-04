package service

import (
	"Todo/pkg/ctl"
	"Todo/pkg/e"
	"Todo/pkg/util"
	"Todo/repository/db/dao"
	"Todo/repository/db/model"
	"Todo/types"
	"context"
	"errors"
	"gorm.io/gorm"
	"sync"
	"time"
)

var UserServiceOnce sync.Once
var UserServiceIns *UserService

type UserService struct {
}

func GetUserService() *UserService {
	UserServiceOnce.Do(func() {
		UserServiceIns = &UserService{}
	})
	return UserServiceIns
}
func (s *UserService) Register(ctx context.Context, req *types.UserServiceReq) (interface{}, error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)
	//没找到就创建新用户
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &model.User{
			UserName: req.UserName,
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		//密码加密
		if err = user.SetPassword(req.Password); err != nil {
			util.LogrusObj.Info(err)
			code := e.SetPasswordFail
			return ctl.RespError(err, code), err
		}

		//存入数据库
		if err = userDao.CreateUser(user); err != nil {
			util.LogrusObj.Info(err)
			code := e.ErrorDatabase
			return ctl.RespError(err, code), err
		}
		return ctl.RespSuccess(), nil
	}
	//找到报错返回
	code := e.ErrorUserExist
	err = errors.New("user exists")
	return ctl.RespError(err, code), nil
}

func (s *UserService) Login(ctx context.Context, req *types.UserServiceReq) (interface{}, error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)
	//没找到
	if err != nil {
		code := e.ErrorUserNotExist
		return ctl.RespError(err, code), err
	}
	//校验密码
	if !user.CheckPassword(req.Password) {
		err = errors.New("密码错误")
		util.LogrusObj.Info(err)
		code := e.ErrorPassword
		return ctl.RespError(err, code), err
	}
	//生成token
	token, err := util.GenerateToken(user.ID, user.UserName)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.TokenGeneratedFail
		return ctl.RespError(err, code), err
	}
	//返回数据
	u := &types.UserResp{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreateAt,
	}
	data := &types.TokenData{
		User:  u,
		Token: token,
	}
	return ctl.RespSuccessWithData(data), nil
}
