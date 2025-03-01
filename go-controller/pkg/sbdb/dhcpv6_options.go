// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package sbdb

import "github.com/ovn-org/libovsdb/model"

type (
	DHCPv6OptionsType = string
)

var (
	DHCPv6OptionsTypeIpv6 DHCPv6OptionsType = "ipv6"
	DHCPv6OptionsTypeMAC  DHCPv6OptionsType = "mac"
	DHCPv6OptionsTypeStr  DHCPv6OptionsType = "str"
)

// DHCPv6Options defines an object in DHCPv6_Options table
type DHCPv6Options struct {
	UUID string            `ovsdb:"_uuid"`
	Code int               `ovsdb:"code"`
	Name string            `ovsdb:"name"`
	Type DHCPv6OptionsType `ovsdb:"type"`
}

func (a *DHCPv6Options) DeepCopyInto(b *DHCPv6Options) {
	*b = *a
}

func (a *DHCPv6Options) DeepCopy() *DHCPv6Options {
	b := new(DHCPv6Options)
	a.DeepCopyInto(b)
	return b
}

func (a *DHCPv6Options) CloneModelInto(b model.Model) {
	c := b.(*DHCPv6Options)
	a.DeepCopyInto(c)
}

func (a *DHCPv6Options) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *DHCPv6Options) Equals(b *DHCPv6Options) bool {
	return a.UUID == b.UUID &&
		a.Code == b.Code &&
		a.Name == b.Name &&
		a.Type == b.Type
}

func (a *DHCPv6Options) EqualsModel(b model.Model) bool {
	c := b.(*DHCPv6Options)
	return a.Equals(c)
}

var _ model.CloneableModel = &DHCPv6Options{}
var _ model.ComparableModel = &DHCPv6Options{}
