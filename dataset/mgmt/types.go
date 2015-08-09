// This file was generated by nomgen.
// To regenerate, run `go generate` in this package.

package mgmt

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// Dataset

type Dataset struct {
	m types.Map
}

func NewDataset() Dataset {
	return Dataset{
		types.NewMap(types.NewString("$name"), types.NewString("Dataset")),
	}
}

func DatasetFromVal(v types.Value) Dataset {
	return Dataset{v.(types.Map)}
}

// TODO: This was going to be called Value() but it collides with root.value. We need some other place to put the built-in fields like Value() and Equals().
func (s Dataset) NomsValue() types.Map {
	return s.m
}

func (s Dataset) Equals(p Dataset) bool {
	return s.m.Equals(p.m)
}

func (s Dataset) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Dataset) Id() types.String {
	return types.StringFromVal(s.m.Get(types.NewString("id")))
}

func (s Dataset) SetId(p types.String) Dataset {
	return DatasetFromVal(s.m.Set(types.NewString("id"), p))
}

func (s Dataset) Heads() types.Value {
	return (s.m.Get(types.NewString("heads")))
}

func (s Dataset) SetHeads(p types.Value) Dataset {
	return DatasetFromVal(s.m.Set(types.NewString("heads"), p))
}

// SetOfDataset

type SetOfDataset struct {
	s types.Set
}

type SetOfDatasetIterCallback (func(p Dataset) (stop bool))

func NewSetOfDataset() SetOfDataset {
	return SetOfDataset{types.NewSet()}
}

func SetOfDatasetFromVal(p types.Value) SetOfDataset {
	return SetOfDataset{p.(types.Set)}
}

func (s SetOfDataset) NomsValue() types.Set {
	return s.s
}

func (s SetOfDataset) Equals(p SetOfDataset) bool {
	return s.s.Equals(p.s)
}

func (s SetOfDataset) Ref() ref.Ref {
	return s.s.Ref()
}

func (s SetOfDataset) Empty() bool {
	return s.s.Empty()
}

func (s SetOfDataset) Len() uint64 {
	return s.s.Len()
}

func (s SetOfDataset) Has(p Dataset) bool {
	return s.s.Has(p.NomsValue())
}

func (s SetOfDataset) Iter(cb SetOfDatasetIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(DatasetFromVal(v))
	})
}

func (s SetOfDataset) Insert(p ...Dataset) SetOfDataset {
	return SetOfDataset{s.s.Insert(s.fromElemSlice(p)...)}
}

func (s SetOfDataset) Remove(p ...Dataset) SetOfDataset {
	return SetOfDataset{s.s.Remove(s.fromElemSlice(p)...)}
}

func (s SetOfDataset) Union(others ...SetOfDataset) SetOfDataset {
	return SetOfDataset{s.s.Union(s.fromStructSlice(others)...)}
}

func (s SetOfDataset) Subtract(others ...SetOfDataset) SetOfDataset {
	return SetOfDataset{s.s.Subtract(s.fromStructSlice(others)...)}
}

func (s SetOfDataset) Any() Dataset {
	return DatasetFromVal(s.s.Any())
}

func (s SetOfDataset) fromStructSlice(p []SetOfDataset) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfDataset) fromElemSlice(p []Dataset) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}
