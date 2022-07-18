package types

import (
	"fmt"
	"sort"

	yaml "gopkg.in/yaml.v2"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)



func NewCapability(index uint64) *Capability {
	return &Capability{Index: index}
}




func (ck *Capability) String() string {
	return fmt.Sprintf("Capability{%p, %d}", ck, ck.Index)
}

func NewOwner(module, name string) Owner {
	return Owner{Module: module, Name: name}
}


func (o Owner) Key() string {
	return fmt.Sprintf("%s/%s", o.Module, o.Name)
}

func (o Owner) String() string {
	bz, _ := yaml.Marshal(o)
	return string(bz)
}

func NewCapabilityOwners() *CapabilityOwners {
	return &CapabilityOwners{Owners: make([]Owner, 0)}
}




func (co *CapabilityOwners) Set(owner Owner) error {
	i, ok := co.Get(owner)
	if ok {

		return sdkerrors.Wrapf(ErrOwnerClaimed, owner.String())
	}


	co.Owners = append(co.Owners, Owner{})
	copy(co.Owners[i+1:], co.Owners[i:])
	co.Owners[i] = owner

	return nil
}



func (co *CapabilityOwners) Remove(owner Owner) {
	if len(co.Owners) == 0 {
		return
	}

	i, ok := co.Get(owner)
	if ok {

		co.Owners = append(co.Owners[:i], co.Owners[i+1:]...)
	}
}




func (co *CapabilityOwners) Get(owner Owner) (int, bool) {

	i := sort.Search(len(co.Owners), func(i int) bool { return co.Owners[i].Key() >= owner.Key() })
	if i < len(co.Owners) && co.Owners[i].Key() == owner.Key() {

		return i, true
	}

	return i, false
}
