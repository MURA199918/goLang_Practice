package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "currency/protos"
)


type Currency struct {
	log hclog.Logger
}

func (c*Currency) GetRate(ctx context.Context, *protos.RateRequest) (*protos.RateResponse, error) {

}