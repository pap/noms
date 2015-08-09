// This file was generated by nomgen.
// To regenerate, run `go generate` in this package.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// Photo

type Photo struct {
	m types.Map
}

func NewPhoto() Photo {
	return Photo{
		types.NewMap(types.NewString("$name"), types.NewString("Photo")),
	}
}

func PhotoFromVal(v types.Value) Photo {
	return Photo{v.(types.Map)}
}

// TODO: This was going to be called Value() but it collides with root.value. We need some other place to put the built-in fields like Value() and Equals().
func (s Photo) NomsValue() types.Map {
	return s.m
}

func (s Photo) Equals(p Photo) bool {
	return s.m.Equals(p.m)
}

func (s Photo) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Photo) Title() types.String {
	return types.StringFromVal(s.m.Get(types.NewString("title")))
}

func (s Photo) SetTitle(p types.String) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("title"), p))
}

func (s Photo) Id() types.String {
	return types.StringFromVal(s.m.Get(types.NewString("id")))
}

func (s Photo) SetId(p types.String) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("id"), p))
}

func (s Photo) Tags() SetOfString {
	return SetOfStringFromVal(s.m.Get(types.NewString("tags")))
}

func (s Photo) SetTags(p SetOfString) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("tags"), p.NomsValue()))
}

func (s Photo) Image() types.Blob {
	return types.BlobFromVal(s.m.Get(types.NewString("image")))
}

func (s Photo) SetImage(p types.Blob) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("image"), p))
}

func (s Photo) Url() types.String {
	return types.StringFromVal(s.m.Get(types.NewString("url")))
}

func (s Photo) SetUrl(p types.String) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("url"), p))
}

// SetOfString

type SetOfString struct {
	s types.Set
}

type SetOfStringIterCallback (func(p types.String) (stop bool))

func NewSetOfString() SetOfString {
	return SetOfString{types.NewSet()}
}

func SetOfStringFromVal(p types.Value) SetOfString {
	return SetOfString{p.(types.Set)}
}

func (s SetOfString) NomsValue() types.Set {
	return s.s
}

func (s SetOfString) Equals(p SetOfString) bool {
	return s.s.Equals(p.s)
}

func (s SetOfString) Ref() ref.Ref {
	return s.s.Ref()
}

func (s SetOfString) Empty() bool {
	return s.s.Empty()
}

func (s SetOfString) Len() uint64 {
	return s.s.Len()
}

func (s SetOfString) Has(p types.String) bool {
	return s.s.Has(p)
}

func (s SetOfString) Iter(cb SetOfStringIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(types.StringFromVal(v))
	})
}

func (s SetOfString) Insert(p ...types.String) SetOfString {
	return SetOfString{s.s.Insert(s.fromElemSlice(p)...)}
}

func (s SetOfString) Remove(p ...types.String) SetOfString {
	return SetOfString{s.s.Remove(s.fromElemSlice(p)...)}
}

func (s SetOfString) Union(others ...SetOfString) SetOfString {
	return SetOfString{s.s.Union(s.fromStructSlice(others)...)}
}

func (s SetOfString) Subtract(others ...SetOfString) SetOfString {
	return SetOfString{s.s.Subtract(s.fromStructSlice(others)...)}
}

func (s SetOfString) Any() types.String {
	return types.StringFromVal(s.s.Any())
}

func (s SetOfString) fromStructSlice(p []SetOfString) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfString) fromElemSlice(p []types.String) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

// MapOfStringToSetOfPhoto

type MapOfStringToSetOfPhoto struct {
	m types.Map
}

type MapOfStringToSetOfPhotoIterCallback (func(k types.String, v SetOfPhoto) (stop bool))

func NewMapOfStringToSetOfPhoto() MapOfStringToSetOfPhoto {
	return MapOfStringToSetOfPhoto{types.NewMap()}
}

func MapOfStringToSetOfPhotoFromVal(p types.Value) MapOfStringToSetOfPhoto {
	return MapOfStringToSetOfPhoto{p.(types.Map)}
}

func (m MapOfStringToSetOfPhoto) NomsValue() types.Map {
	return m.m
}

func (m MapOfStringToSetOfPhoto) Equals(p MapOfStringToSetOfPhoto) bool {
	return m.m.Equals(p.m)
}

func (m MapOfStringToSetOfPhoto) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToSetOfPhoto) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToSetOfPhoto) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToSetOfPhoto) Has(p types.String) bool {
	return m.m.Has(p)
}

func (m MapOfStringToSetOfPhoto) Get(p types.String) SetOfPhoto {
	return SetOfPhotoFromVal(m.m.Get(p))
}

func (m MapOfStringToSetOfPhoto) Set(k types.String, v SetOfPhoto) MapOfStringToSetOfPhoto {
	return MapOfStringToSetOfPhotoFromVal(m.m.Set(k, v.NomsValue()))
}

// TODO: Implement SetM?

func (m MapOfStringToSetOfPhoto) Remove(p types.String) MapOfStringToSetOfPhoto {
	return MapOfStringToSetOfPhotoFromVal(m.m.Remove(p))
}

func (m MapOfStringToSetOfPhoto) Iter(cb MapOfStringToSetOfPhotoIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(types.StringFromVal(k), SetOfPhotoFromVal(v))
	})
}

// SetOfPhoto

type SetOfPhoto struct {
	s types.Set
}

type SetOfPhotoIterCallback (func(p Photo) (stop bool))

func NewSetOfPhoto() SetOfPhoto {
	return SetOfPhoto{types.NewSet()}
}

func SetOfPhotoFromVal(p types.Value) SetOfPhoto {
	return SetOfPhoto{p.(types.Set)}
}

func (s SetOfPhoto) NomsValue() types.Set {
	return s.s
}

func (s SetOfPhoto) Equals(p SetOfPhoto) bool {
	return s.s.Equals(p.s)
}

func (s SetOfPhoto) Ref() ref.Ref {
	return s.s.Ref()
}

func (s SetOfPhoto) Empty() bool {
	return s.s.Empty()
}

func (s SetOfPhoto) Len() uint64 {
	return s.s.Len()
}

func (s SetOfPhoto) Has(p Photo) bool {
	return s.s.Has(p.NomsValue())
}

func (s SetOfPhoto) Iter(cb SetOfPhotoIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(PhotoFromVal(v))
	})
}

func (s SetOfPhoto) Insert(p ...Photo) SetOfPhoto {
	return SetOfPhoto{s.s.Insert(s.fromElemSlice(p)...)}
}

func (s SetOfPhoto) Remove(p ...Photo) SetOfPhoto {
	return SetOfPhoto{s.s.Remove(s.fromElemSlice(p)...)}
}

func (s SetOfPhoto) Union(others ...SetOfPhoto) SetOfPhoto {
	return SetOfPhoto{s.s.Union(s.fromStructSlice(others)...)}
}

func (s SetOfPhoto) Subtract(others ...SetOfPhoto) SetOfPhoto {
	return SetOfPhoto{s.s.Subtract(s.fromStructSlice(others)...)}
}

func (s SetOfPhoto) Any() Photo {
	return PhotoFromVal(s.s.Any())
}

func (s SetOfPhoto) fromStructSlice(p []SetOfPhoto) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfPhoto) fromElemSlice(p []Photo) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}
