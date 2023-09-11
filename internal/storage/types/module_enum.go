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
	// ModuleNameAuth is a ModuleName of type auth.
	ModuleNameAuth ModuleName = "auth"
	// ModuleNameBlob is a ModuleName of type blob.
	ModuleNameBlob ModuleName = "blob"
	// ModuleNameCrisis is a ModuleName of type crisis.
	ModuleNameCrisis ModuleName = "crisis"
	// ModuleNameDistribution is a ModuleName of type distribution.
	ModuleNameDistribution ModuleName = "distribution"
	// ModuleNameIndexer is a ModuleName of type indexer.
	ModuleNameIndexer ModuleName = "indexer"
	// ModuleNameGov is a ModuleName of type gov.
	ModuleNameGov ModuleName = "gov"
	// ModuleNameSlashing is a ModuleName of type slashing.
	ModuleNameSlashing ModuleName = "slashing"
	// ModuleNameStaking is a ModuleName of type staking.
	ModuleNameStaking ModuleName = "staking"
)

var ErrInvalidModuleName = errors.New("not a valid ModuleName")

// ModuleNameValues returns a list of the values for ModuleName
func ModuleNameValues() []ModuleName {
	return []ModuleName{
		ModuleNameAuth,
		ModuleNameBlob,
		ModuleNameCrisis,
		ModuleNameDistribution,
		ModuleNameIndexer,
		ModuleNameGov,
		ModuleNameSlashing,
		ModuleNameStaking,
	}
}

// String implements the Stringer interface.
func (x ModuleName) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ModuleName) IsValid() bool {
	_, err := ParseModuleName(string(x))
	return err == nil
}

var _ModuleNameValue = map[string]ModuleName{
	"auth":         ModuleNameAuth,
	"blob":         ModuleNameBlob,
	"crisis":       ModuleNameCrisis,
	"distribution": ModuleNameDistribution,
	"indexer":      ModuleNameIndexer,
	"gov":          ModuleNameGov,
	"slashing":     ModuleNameSlashing,
	"staking":      ModuleNameStaking,
}

// ParseModuleName attempts to convert a string to a ModuleName.
func ParseModuleName(name string) (ModuleName, error) {
	if x, ok := _ModuleNameValue[name]; ok {
		return x, nil
	}
	return ModuleName(""), fmt.Errorf("%s is %w", name, ErrInvalidModuleName)
}

// MarshalText implements the text marshaller method.
func (x ModuleName) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ModuleName) UnmarshalText(text []byte) error {
	tmp, err := ParseModuleName(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errModuleNameNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *ModuleName) Scan(value interface{}) (err error) {
	if value == nil {
		*x = ModuleName("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case string:
		*x, err = ParseModuleName(v)
	case []byte:
		*x, err = ParseModuleName(string(v))
	case ModuleName:
		*x = v
	case *ModuleName:
		if v == nil {
			return errModuleNameNilPtr
		}
		*x = *v
	case *string:
		if v == nil {
			return errModuleNameNilPtr
		}
		*x, err = ParseModuleName(*v)
	default:
		return errors.New("invalid type for ModuleName")
	}

	return
}

// Value implements the driver Valuer interface.
func (x ModuleName) Value() (driver.Value, error) {
	return x.String(), nil
}