package handler

import (
	"context"

	inv "github.com/eopenio/template/inventory-srv/model/inventory"
	proto "github.com/eopenio/template/inventory-srv/proto/inventory"
	log "github.com/micro/go-micro/v2/logger"
)

var (
	invService inv.Service
)

type Service struct {
}

// Init 初始化handler
func Init() {
	invService, _ = inv.GetService()
}

// Sell 库存销存
func (e *Service) Sell(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	id, err := invService.Sell(req.BookId, req.UserId)
	if err != nil {
		log.Infof("[Sell] 销存失败，bookId：%d，userId: %d，%s", req.BookId, req.UserId, err)
		rsp.Success = false
		return
	}

	rsp.InvH = &proto.InvHistory{
		Id: id,
	}

	rsp.Success = true
	return nil
}

// Confirm 库存销存 确认
func (e *Service) Confirm(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	err = invService.Confirm(req.HistoryId, int(req.HistoryState))
	if err != nil {
		log.Infof("[Confirm] 确认销存失败，%s", err)
		rsp.Success = false
		return
	}

	rsp.Success = true
	return nil
}
