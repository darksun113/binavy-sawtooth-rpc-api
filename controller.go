package main

import "C"
import (
	"github.com/binavy/sawtooth-rpc-go/rpc/context"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
)

type Controller struct {
}

func (ctl *Controller) GetBlockByBatchID(req *GetBlockInfoRequest) (rsp * GetBlockInfoResponse, err error) {
	zapLogger.Info("CreateSale req:", zap.Any("req", req))

	rsp = new(GetBlockInfoResponse)

	ctx := context.NewConsumerContext()
	rspRaw, err := c.ClientBlockGetByBatchId(ctx, req.BatchID)
	if err != nil {
		rsp.Code = 500
		rsp.ErrorMsg = err.Error()
		zapLogger.Error("GetBlockByBatchID Error:", zap.String("err",rsp.ErrorMsg))
		return
	}

	err = proto.Unmarshal(rspRaw.Block.Header, &rsp.BlockHeader)

	if err != nil {
		rsp.Code = 500
		rsp.ErrorMsg = err.Error()
		zapLogger.Error("GetBlockByBatchID Error:", zap.String("err",rsp.ErrorMsg))
		return
	}
	rsp.Code = 0
	err = nil



	zapLogger.Info("CreateSale rsp:", zap.Any("rsp", rsp))
	return
}

