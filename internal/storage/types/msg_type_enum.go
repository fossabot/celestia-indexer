// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.7
// Revision: bf63e108589bbd2327b13ec2c5da532aad234029
// Build Date: 2023-07-25T23:27:55Z
// Built By: goreleaser

package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

const (
	// MsgUnknown is a MsgType of type MsgUnknown.
	MsgUnknown MsgType = "MsgUnknown"
	// MsgWithdrawValidatorCommission is a MsgType of type MsgWithdrawValidatorCommission.
	MsgWithdrawValidatorCommission MsgType = "MsgWithdrawValidatorCommission"
	// MsgWithdrawDelegatorReward is a MsgType of type MsgWithdrawDelegatorReward.
	MsgWithdrawDelegatorReward MsgType = "MsgWithdrawDelegatorReward"
	// MsgEditValidator is a MsgType of type MsgEditValidator.
	MsgEditValidator MsgType = "MsgEditValidator"
	// MsgBeginRedelegate is a MsgType of type MsgBeginRedelegate.
	MsgBeginRedelegate MsgType = "MsgBeginRedelegate"
	// MsgCreateValidator is a MsgType of type MsgCreateValidator.
	MsgCreateValidator MsgType = "MsgCreateValidator"
	// MsgDelegate is a MsgType of type MsgDelegate.
	MsgDelegate MsgType = "MsgDelegate"
	// MsgUndelegate is a MsgType of type MsgUndelegate.
	MsgUndelegate MsgType = "MsgUndelegate"
	// MsgUnjail is a MsgType of type MsgUnjail.
	MsgUnjail MsgType = "MsgUnjail"
	// MsgSend is a MsgType of type MsgSend.
	MsgSend MsgType = "MsgSend"
	// MsgCreateVestingAccount is a MsgType of type MsgCreateVestingAccount.
	MsgCreateVestingAccount MsgType = "MsgCreateVestingAccount"
	// MsgCreatePeriodicVestingAccount is a MsgType of type MsgCreatePeriodicVestingAccount.
	MsgCreatePeriodicVestingAccount MsgType = "MsgCreatePeriodicVestingAccount"
	// MsgPayForBlobs is a MsgType of type MsgPayForBlobs.
	MsgPayForBlobs MsgType = "MsgPayForBlobs"
	// MsgGrantAllowance is a MsgType of type MsgGrantAllowance.
	MsgGrantAllowance MsgType = "MsgGrantAllowance"
	// MsgRegisterEVMAddress is a MsgType of type MsgRegisterEVMAddress.
	MsgRegisterEVMAddress MsgType = "MsgRegisterEVMAddress"
	// MsgSetWithdrawAddress is a MsgType of type MsgSetWithdrawAddress.
	MsgSetWithdrawAddress MsgType = "MsgSetWithdrawAddress"
)

var ErrInvalidMsgType = errors.New("not a valid MsgType")

// MsgTypeValues returns a list of the values for MsgType
func MsgTypeValues() []MsgType {
	return []MsgType{
		MsgUnknown,
		MsgWithdrawValidatorCommission,
		MsgWithdrawDelegatorReward,
		MsgEditValidator,
		MsgBeginRedelegate,
		MsgCreateValidator,
		MsgDelegate,
		MsgUndelegate,
		MsgUnjail,
		MsgSend,
		MsgCreateVestingAccount,
		MsgCreatePeriodicVestingAccount,
		MsgPayForBlobs,
		MsgGrantAllowance,
		MsgRegisterEVMAddress,
		MsgSetWithdrawAddress,
	}
}

// String implements the Stringer interface.
func (x MsgType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x MsgType) IsValid() bool {
	_, err := ParseMsgType(string(x))
	return err == nil
}

var _MsgTypeValue = map[string]MsgType{
	"MsgUnknown":                      MsgUnknown,
	"MsgWithdrawValidatorCommission":  MsgWithdrawValidatorCommission,
	"MsgWithdrawDelegatorReward":      MsgWithdrawDelegatorReward,
	"MsgEditValidator":                MsgEditValidator,
	"MsgBeginRedelegate":              MsgBeginRedelegate,
	"MsgCreateValidator":              MsgCreateValidator,
	"MsgDelegate":                     MsgDelegate,
	"MsgUndelegate":                   MsgUndelegate,
	"MsgUnjail":                       MsgUnjail,
	"MsgSend":                         MsgSend,
	"MsgCreateVestingAccount":         MsgCreateVestingAccount,
	"MsgCreatePeriodicVestingAccount": MsgCreatePeriodicVestingAccount,
	"MsgPayForBlobs":                  MsgPayForBlobs,
	"MsgGrantAllowance":               MsgGrantAllowance,
	"MsgRegisterEVMAddress":           MsgRegisterEVMAddress,
	"MsgSetWithdrawAddress":           MsgSetWithdrawAddress,
}

// ParseMsgType attempts to convert a string to a MsgType.
func ParseMsgType(name string) (MsgType, error) {
	if x, ok := _MsgTypeValue[name]; ok {
		return x, nil
	}
	return MsgType(""), fmt.Errorf("%s is %w", name, ErrInvalidMsgType)
}

// MarshalText implements the text marshaller method.
func (x MsgType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *MsgType) UnmarshalText(text []byte) error {
	tmp, err := ParseMsgType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errMsgTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *MsgType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = MsgType("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case string:
		*x, err = ParseMsgType(v)
	case []byte:
		*x, err = ParseMsgType(string(v))
	case MsgType:
		*x = v
	case *MsgType:
		if v == nil {
			return errMsgTypeNilPtr
		}
		*x = *v
	case *string:
		if v == nil {
			return errMsgTypeNilPtr
		}
		*x, err = ParseMsgType(*v)
	default:
		return errors.New("invalid type for MsgType")
	}

	return
}

// Value implements the driver Valuer interface.
func (x MsgType) Value() (driver.Value, error) {
	return x.String(), nil
}
