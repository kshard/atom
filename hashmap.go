//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/symbol
//

package atom

import "sync"

type hashmap struct {
	sync.RWMutex
	kv map[Atom]string
}

func NewMemMap() HashMap {
	return &hashmap{
		kv: make(map[Atom]string),
	}
}

func (m *hashmap) Get(key Atom) (string, error) {
	m.RLock()
	val, has := m.kv[key]
	m.RUnlock()

	if !has {
		return "", nil
	}
	return val, nil
}

func (m *hashmap) Put(key Atom, val string) error {
	m.Lock()
	m.kv[key] = val
	m.Unlock()

	return nil
}
