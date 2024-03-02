//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/symbol
//

package atom

import (
	"strconv"
)

type Pool struct {
	hash    Hash
	hashmap HashMap
}

// Create instance of in-memory symbols table
func New(hashmap HashMap) *Pool {
	return &Pool{
		hash:    NewHash(hashmap),
		hashmap: hashmap,
	}
}

// Cast string to Symbol
func (p *Pool) Atom(s string) (Atom, error) {
	hash, err := p.hash.String(s)
	if err != nil {
		return 0, err
	}

	if err := p.hashmap.Put(hash, s); err != nil {
		return 0, err
	}

	return hash, nil
}

// Cast Symbol to string
func (p *Pool) String(s Atom) string {
	val, err := p.hashmap.Get(s)
	if err != nil {
		return ":" + strconv.Itoa(int(s))
	}

	return val
}
