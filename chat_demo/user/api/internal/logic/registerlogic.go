package logic

import (
	"chat_demo/user/model"
	"context"

	"chat_demo/user/api/internal/svc"
	"chat_demo/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) error {
	_, err := l.svcCtx.UserModel.FindOneByName(req.Username)
	if err == nil {
		return errorDuplicateUsername
	}

	_, err = l.svcCtx.UserModel.FindOneByMobile(req.Mobile)
	if err == nil {
		return errorDuplicateMobile
	}

	_, err = l.svcCtx.UserModel.Insert(model.User{
		Name:     req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		Gender:   req.Gender,
		Nickname: req.Nickname,
	})

	return err
}
