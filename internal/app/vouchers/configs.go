package vouchers

import "time"

type Configs struct {
	WalletServiceAddress  string        `yaml:"wallet_service_address"`
	VoucherServiceAddress string        `yaml:"voucher_service_address"`
	ApplyCreditTimeout    time.Duration `yaml:"apply_credit_timeout"`
}
