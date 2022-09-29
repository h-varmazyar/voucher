package main

import (
	"github.com/h-varmazyar/voucher/internal/app/vouchers"
	"github.com/h-varmazyar/voucher/internal/app/wallets"
	"github.com/h-varmazyar/voucher/pkg/netext"
)

type Configs struct {
	DSN      string      `yaml:"dsn"`
	GRPCPort netext.Port `yaml:"grpc_port"`
	HttpPort netext.Port `yaml:"http_port"`

	TransactionsConfigs *vouchers.Configs
	WalletsConfigs      *wallets.Configs
}
