# atom

> In computer science, string interning is a method of storing only one copy of each distinct string value, which must be immutable. Interning strings makes some string processing tasks more time-efficient or space-efficient at the cost of requiring more time when the string is created or interned. The distinct values are stored in a string intern pool.
- https://en.wikipedia.org/wiki/String_interning

This Golang library of defines an `atom` type that does string interning.

[![Version](https://img.shields.io/github/v/tag/kshard/atom?label=version)](https://github.com/kshard/atom/releases)
[![Documentation](https://pkg.go.dev/badge/github.com/kshard/atom)](https://pkg.go.dev/github.com/kshard/atom)
[![Build Status](https://github.com/kshard/atom/workflows/build/badge.svg)](https://github.com/kshard/atom/actions/)
[![Git Hub](https://img.shields.io/github/last-commit/kshard/atom.svg)](https://github.com/kshard/atom)
[![Coverage Status](https://coveralls.io/repos/github/kshard/atom/badge.svg?branch=main)](https://coveralls.io/github/kshard/atom?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/kshard/atom)](https://goreportcard.com/report/github.com/kshard/atom)


## Inspiration

In Standard ML, the `atom` type is used to ensure that each string literal is stored exactly once in memory, regardless of how many times it is used, which optimizes both memory usage and comparison times. This idea is brought into Golang through a custom library that mimics the behavior of atoms. The library maintains a global map that holds references to unique strings, effectively interning them.

This library intentionally chosen double hashing algorithm to calculate atoms for strings instead of maintaining sequence id as other libraries are doing. This technique requires usage of `map` to preserve association `atom ⟼ string`.

When a string is interned in Golang using this library, it checks if the string already exists in the global map. If the string is new, it's added to the map, and its reference is returned; if it already exists, the existing reference is returned. This process ensures that all interned strings are unique in memory, allowing for constant-time string comparisons. Overall, the adaptation involves careful consideration of Golang's concurrency model and memory management practices to ensure that the atom type operates efficiently and safely in a multi-threaded environment.


## Getting started

The latest version of the library is available at `main` branch of this repository. All development, including new features and bug fixes, take place on the `main` branch using forking and pull requests as described in contribution guidelines. The stable version is available via Golang modules.

```go
import "github.com/kshard/atom"

// Create new atoms table 
atoms := atom.New(atom.NewMemMap())

// Convert string to atom
code := atoms.Atom("String interning")

// Convert atom back to string
atoms.String(code)
```

## How To Contribute

The library is [MIT](LICENSE) licensed and accepts contributions via GitHub pull requests:

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

The build and testing process requires [Go](https://golang.org) version 1.13 or later.

**build** and **test** library.

```bash
git clone https://github.com/kshard/atom
cd atom
go test
```

### commit message

The commit message helps us to write a good release note, speed-up review process. The message should address two question what changed and why. The project follows the template defined by chapter [Contributing to a Project](http://git-scm.com/book/ch5-2.html) of Git book.

### bugs

If you experience any issues with the library, please let us know via [GitHub issues](https://github.com/kshard/atom/issue). We appreciate detailed and accurate reports that help us to identity and replicate the issue. 


## License

[![See LICENSE](https://img.shields.io/github/license/kshard/atom.svg?style=for-the-badge)](LICENSE)


## References

1. https://en.wikipedia.org/wiki/String_interning
2. [String Interning in JVM](https://hg.openjdk.org/jdk7/jdk7/hotspot/file/tip/src/share/vm/classfile/symbolTable.cpp)
3. Another approach on [String interning in Go](https://artem.krylysov.com/blog/2018/12/12/string-interning-in-go/) 
