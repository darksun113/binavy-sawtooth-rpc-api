package main

import (
	"github.com/binavy/sawtooth-rpc-go/boot"
	"github.com/binavy/sawtooth-rpc-go/rpc/consumer/evaluator"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger
var c * BinaSawCommonAPIConsumer

type BinaSawCommonAPIConsumer struct {
	evaluator.Evaluator
}

func (b BinaSawCommonAPIConsumer) Name() string {
	return "Common-API-Consumer"
}

func (b BinaSawCommonAPIConsumer) Version() string {
	return "v1.0"
}

func configureAPI(api *iris.APIContainer) {
	c := Controller{}
	api.Post("common/getblockbybatchid", c.GetBlockByBatchID)
}


func main()  {
	c = &BinaSawCommonAPIConsumer{}

	boot.RegisterConsumer(c)
	boot.StartConsumeWithConfig("config/config.yaml")

	zapLogger = initLogger("./api.log", "debug", true)

	app := iris.New()
	app.Logger().SetLevel("debug")

	app.ConfigureContainer(configureAPI)
	app.Listen(":16899")
}