//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/symbol
//

package atom

import (
	"encoding/binary"
	"sync"
	"unsafe"
)

//------------------------------------------------------------------------------

type ephemeral struct {
	sync.RWMutex
	kv map[Atom]string
}

func NewEphemeralMap() HashMap {
	return &ephemeral{
		kv: make(map[Atom]string),
	}
}

func (m *ephemeral) Get(key Atom) (string, error) {
	m.RLock()
	val, has := m.kv[key]
	m.RUnlock()

	if !has {
		return "", nil
	}
	return val, nil
}

func (m *ephemeral) Put(key Atom, val string) error {
	m.Lock()
	m.kv[key] = val
	m.Unlock()

	return nil
}

//------------------------------------------------------------------------------

type permanent struct {
	store Store
}

func NewPermanentMap(store Store) HashMap {
	return &permanent{store: store}
}

func (m *permanent) Get(key Atom) (string, error) {
	var bkey [5]byte
	bkey[0] = ':'
	binary.LittleEndian.PutUint32(bkey[1:], key)

	val, err := m.store.Get(bkey[:])
	if err != nil {
		return "", err
	}

	// This is copied from runtime. It relies on the string
	// header being a prefix of the slice header!
	str := *(*string)(unsafe.Pointer(&val))

	return str, nil
}

func (m *permanent) Put(key Atom, val string) error {
	var bkey [5]byte
	bkey[0] = ':'
	binary.LittleEndian.PutUint32(bkey[1:], key)

	// This is copied from runtime. It relies on the string
	// header being a prefix of the slice header!
	bval := *(*[]byte)(unsafe.Pointer(&val))

	err := m.store.Put(bkey[:], bval)
	if err != nil {
		return err
	}

	return nil
}
