package atom_test

import (
	"bytes"
	"strconv"
	"testing"
	"time"

	"github.com/kshard/atom"
)

const (
	sa = "String interning"
	sb = "String"
	sc = "interning"
)

func TestEphemeralPut(t *testing.T) {
	s := atom.New(atom.NewEphemeralMap())

	for val, expected := range map[string]uint32{
		sa: 1247594388,
		sb: 3572195896,
		sc: 1304336027,
	} {
		sym, err := s.Atom(val)
		if err != nil {
			t.Errorf("failed to assign symbol: %s", err)
		}
		if sym != expected {
			t.Errorf("failed to assign symbol: %d, expected %d", sym, expected)
		}
		if val != s.String(sym) {
			t.Errorf("failed to lookup string")
		}
	}
}

func TestPermanentPut(t *testing.T) {
	s := atom.New(atom.NewPermanentMap(&none{}))

	for val, expected := range map[string]uint32{
		sa: 1247594388,
		sb: 3572195896,
		sc: 1304336027,
	} {
		sym, err := s.Atom(val)
		if err != nil {
			t.Errorf("failed to assign symbol: %s", err)
		}
		if sym != expected {
			t.Errorf("failed to assign symbol: %d, expected %d", sym, expected)
		}
		if val != s.String(sym) {
			t.Errorf("failed to lookup string")
		}
	}
}

// ---------------------------------------------------------------

// go test -run=^$ -bench=. -cpu=1 -benchtime=10s -cpuprofile profile.out
func BenchmarkEphemeralPut(b *testing.B) {
	s := atom.New(atom.NewEphemeralMap())

	b.ReportAllocs()
	b.ResetTimer()

	t := time.Now().Nanosecond()

	for n := 0; n < b.N; n++ {
		s.Atom("https://pkg.go.dev/hash/fnv@go1.20." + strconv.Itoa(t+n))
	}
}

func BenchmarkPermanentPut(b *testing.B) {
	s := atom.New(atom.NewPermanentMap(&none{}))

	b.ReportAllocs()
	b.ResetTimer()

	t := time.Now().Nanosecond()

	for n := 0; n < b.N; n++ {
		s.Atom("https://pkg.go.dev/hash/fnv@go1.20." + strconv.Itoa(t+n))
	}
}

// ---------------------------------------------------------------

// go test -fuzz=FuzzSymbolOf
func FuzzSymbolOf(f *testing.F) {
	s := atom.New(atom.NewEphemeralMap())

	f.Add("abc")

	f.Fuzz(func(t *testing.T, x string) {
		_, err := s.Atom(x)
		if err != nil {
			t.Errorf("failed: %s", err)
		}
	})
}

// ---------------------------------------------------------------

type none struct {
	key []byte
	val []byte
}

func (n *none) Get(key []byte) ([]byte, error) {
	if bytes.Equal(n.key, key) {
		return n.val, nil
	}

	return nil, nil
}

func (n *none) Put(key []byte, val []byte) error {
	n.key = key
	n.val = val
	return nil
}
