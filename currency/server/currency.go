package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/MURA199918/goLang_Practice/tree/master/currency/protos"
)


type Currency struct {
	log hclog.Logger
}

func (c*Currency) GetRate(ctx context.Context, *protos.RateRequest) (*protos.RateResponse, error) {

}