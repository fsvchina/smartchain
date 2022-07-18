package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


const (
	Minter  = "minter"
	Burner  = "burner"
	Staking = "staking"
)


type PermissionsForAddress struct {
	permissions []string
	address     sdk.AccAddress
}


func NewPermissionsForAddress(name string, permissions []string) PermissionsForAddress {
	return PermissionsForAddress{
		permissions: permissions,
		address:     NewModuleAddress(name),
	}
}


func (pa PermissionsForAddress) HasPermission(permission string) bool {
	for _, perm := range pa.permissions {
		if perm == permission {
			return true
		}
	}
	return false
}


func (pa PermissionsForAddress) GetAddress() sdk.AccAddress {
	return pa.address
}


func (pa PermissionsForAddress) GetPermissions() []string {
	return pa.permissions
}


func validatePermissions(permissions ...string) error {
	for _, perm := range permissions {
		if strings.TrimSpace(perm) == "" {
			return fmt.Errorf("module permission is empty")
		}
	}
	return nil
}
