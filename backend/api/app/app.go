package app

import (
	"context"

	v1 "backend/api/app/v1"
)

type IAppV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Install(ctx context.Context, req *v1.InstallReq) (res *v1.InstallRes, err error)
	Uninstall(ctx context.Context, req *v1.UninstallReq) (res *v1.UninstallRes, err error)
	Start(ctx context.Context, req *v1.StartReq) (res *v1.StartRes, err error)
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
}
