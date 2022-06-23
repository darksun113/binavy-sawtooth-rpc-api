package main

import (
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/block_pb2"
)

type GetBlockInfoResponse struct {
	Code     int    `json:"code"`
	ErrorMsg string `json:"errorMsg"`
	block_pb2.BlockHeader
}

