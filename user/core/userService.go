package core

import (
	"context"
	"user/dao"
	"user/model"
	pb "user/service"
)

type UserService struct {
}

func (userService *UserService) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserDetailResponse) (err error) {
	resp.Code = 200
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)

	if err != nil {
		resp.Code = 500
		return
	}

	if !user.CheckPassword(req.Password) {
		resp.Code = 300
		return
	}
	resp.UserDetail = BuildUser(user)
	return
}

func BuildUser(item *model.User) *pb.UserModel {
	userModel := pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}
