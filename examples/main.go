//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/symbol
//

package main

import (
	"fmt"
	"strconv"

	"github.com/kshard/atom"
)

const n = 1000000

func main() {
	var err error

	atoms := atom.New(atom.NewEphemeralMap())
	ids := make([]atom.Atom, n)

	for i := 0; i < n; i++ {
		v := "https://pkg.go.dev/hash/fnv@go1.20." + strconv.Itoa(i)
		ids[i], err = atoms.Atom(v)
		if err != nil {
			panic(err)
		}
	}

	for _, x := range ids {
		fmt.Printf("==> %08x %s\n", uint32(x), atoms.String(x))
	}
}
