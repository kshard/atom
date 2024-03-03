//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/symbol
//

package atom

type Atom = uint32

// Getter interface abstract hash-table required to achieve probing
type Getter interface{ Get(Atom) (string, error) }

// Setter interface abstract hash-table
type Putter interface{ Put(Atom, string) error }

// HashMap interface
type HashMap interface {
	Getter
	Putter
}

// abstraction of permanent storage
type Store interface {
	Get([]byte) ([]byte, error)
	Put([]byte, []byte) error
}
