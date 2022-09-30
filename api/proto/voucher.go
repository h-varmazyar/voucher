package voucherApi

import (
	"database/sql/driver"
	"encoding/json"
)

func (x *VoucherType) InRange(v interface{}) bool {
	_, ok := VoucherType_value[v.(VoucherType).String()]
	return ok
}
func (x *VoucherType) Scan(value interface{}) error {
	*x = VoucherType(VoucherType_value[value.(string)])
	return nil
}
func (x *VoucherType) Value() (driver.Value, error) {
	return x.String(), nil
}
func (x *VoucherType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*x = VoucherType(VoucherType_value[str])
	return nil
}
func (x *VoucherType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
