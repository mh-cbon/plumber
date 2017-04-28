# plumber

[![travis Status](https://travis-ci.org//mh-cbon/plumber.svg?branch=master)](https://travis-ci.org//mh-cbon/plumber) [![Appveyor Status](https://ci.appveyor.com/api/projects/status//github/mh-cbon/plumber?branch=master&svg=true)](https://ci.appveyor.com/projects//mh-cbon/plumber) [![Go Report Card](https://goreportcard.com/badge/github.com/mh-cbon/plumber)](https://goreportcard.com/report/github.com/mh-cbon/plumber) [![GoDoc](https://godoc.org/github.com/mh-cbon/plumber?status.svg)](http://godoc.org/github.com/mh-cbon/plumber) [![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Package plumber builds pipes to transform a data stream.


# TOC
- [Install](#install)
- [Generator](#generator)
  - [Usage](#usage)
    - [$ plumber -help](#-plumber--help)
  - [Cli examples](#cli-examples)
- [API example](#api-example)
  - [> demo/main.go](#-demomaingo)
  - [> demo/version.go](#-demoversiongo)
  - [> demo/semver_gen.go](#-demosemver_gengo)
- [Recipes](#recipes)
  - [Release the project](#release-the-project)
- [History](#history)

# Install
```sh
go get github.com/mh-cbon/plumber
```

# Generator

To help you to deal with the step of interface implementation,
 `plumber` comes with a command line program to generate your own typed pipes.

## Usage

#### $ plumber -help
```sh
plumber 0.0.0

Usage

	plumber [out] [pkg] [...types]

	out: 	Output destination of the results, use '-' for stdout.
	pkg: 	The package name of the generated code.
	types:	A list of fully qualified types such as
	     	'[]byte', 'semver.Version', '*my.PointerType'
	     	or 'github.com/mh-cbon/semver/*my.PointerType'.
```

## Cli examples

```sh
# Create a pipe of *tomate.SuperStruct in the package mysuperpkg
plumber - mysuperpkg *tomate.SuperStruct
```
# API example

Following example reads a `source` of `[]byte`, os.Stdin,
as a list of versions, one per line,
manipulates and transforms the chunks
until the data is written on the `sink`, os.Stdout.

#### > demo/main.go
```go
//Package cmd implement a cli tool to manipulate Versions.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/mh-cbon/plumber"
)

//go:generate plumber semver_gen.go main "github.com/mh-cbon/semver/*semver.Version"

func main() {

	flag.Parse()

	src := plumber.NewByteReader(nargsOrStdin())
	dest := plumber.NewByteSink(os.Stdout)

	src.
		Pipe(plumber.NewBytesSplitter(' ', '\n')).
		Pipe(&plumber.BytesTrimer{}).
		Pipe(&VersionFromByte{SkipInvalid: true}).
		Pipe(&VersionSorter{Asc: !true}).
		Pipe(&VersionToByte{}).
		Pipe(plumber.NewBytesPrefixer("- ", "\n")).
		Pipe(&plumber.LastChunkOnly{}).
		Sink(dest)

	if err := src.Consume(); err != nil {
		panic(err)
	}
	os.Exit(0)
}

// arguments are provided via stdin or os.Args ?
// in both case, return an io.Reader complying with the pipe.
func nargsOrStdin() io.Reader {
	if flag.NArg() > 0 {
		var b bytes.Buffer
		ret := &b
		for _, arg := range flag.Args() {
			fmt.Fprintf(&b, "%v\n", arg)
		}
		return ret
	}
	return os.Stdin
}
```

Following code is the implementation of various
pipe `transformer` that works with `*semver.Version` type.

#### > demo/version.go
```go
package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/mh-cbon/plumber"
	"github.com/mh-cbon/semver"
)

// VersionFromByte receives bytes encoded *Version, pushes *Version
type VersionFromByte struct {
	VersionStream
	SkipInvalid bool
}

// Write receive a chunk of []byte, writes a *Version on the connected Pipes.
func (p *VersionFromByte) Write(d []byte) error {
	s, err := semver.NewVersion(string(d))
	if err != nil {
		err := fmt.Errorf("Invalid version %q", string(d))
		if p.SkipInvalid {
			err = nil
		}
		return err
	}
	return p.VersionStream.Write(s)
}

// VersionSorter receives *Version, buffer them until flush, order all *Versions, writes all *Version to the connected Pipes.
type VersionSorter struct {
	VersionStream
	all []*semver.Version
	Asc bool
}

// Write *Version to the buffer.
func (p *VersionSorter) Write(v *semver.Version) error {
	p.all = append(p.all, v)
	return nil
}

// Flush sorts all buffered *Version, writes all *Version to the connected Pipes.
func (p *VersionSorter) Flush() error {
	if p.Asc {
		sort.Sort(semver.Collection(p.all))
	} else {
		sort.Sort(sort.Reverse(semver.Collection(p.all)))
	}
	for _, v := range p.all {
		p.VersionStream.Write(v)
	}
	p.all = p.all[:0]
	return p.VersionStream.Flush()
}

// VersionJsoner receives *Version, buffer them until flush, json encode *Versions, writes bytes to the connected Pipes.
type VersionJsoner struct {
	plumber.ByteStream
	all []*semver.Version
}

// Write *Version to the buffer.
func (p *VersionJsoner) Write(v *semver.Version) error {
	p.all = append(p.all, v)
	return nil
}

// Flush sorts all buffered *Version, writes all *Version to the connected Pipes.
func (p *VersionJsoner) Flush() error {
	blob, err := json.Marshal(p.all)
	if err != nil {
		return err
	}
	err = p.ByteStream.Write(blob)
	if err != nil {
		return err
	}
	return p.ByteStream.Flush()
}

// InvalidVersionFromByte receives bytes chunks of *Version, when it fails to decode it as a *Version, writes the chunk on the connected Pipes.
type InvalidVersionFromByte struct {
	plumber.ByteStream
}

// Write a chunk of bytes, when it is not a valid *Version, writes the chunk on the connected Pipes.
func (p *InvalidVersionFromByte) Write(d []byte) error {
	_, err := semver.NewVersion(string(d))
	if err == nil {
		return nil
	}
	return p.ByteStream.Write(d)
}

// VersionToByte receives *Version, writes bytes chunks to the connection Pipes.
type VersionToByte struct {
	plumber.ByteStream
}

// Write encode *Version to a byte chunk, writes the chunk to the connected Pipes.
func (p *VersionToByte) Write(d *semver.Version) error {
	return p.ByteStream.Write([]byte(d.String()))
}
```

Following is the generated code to build pipes
to work with `*semver.Version` values.

#### > demo/semver_gen.go
```go
// Package main implements pipes for a stream of *semver.Version
package main

import (
	"fmt"

	"github.com/mh-cbon/plumber"
	"github.com/mh-cbon/semver"
)

// This file was automatically generated by
// github.com/mh-cbon/plumber
// To not edit.

// VersionPipeWriter receives *semver.Version
type VersionPipeWriter interface {
	plumber.Flusher
	plumber.Sinker
	Write(*semver.Version) error
}

// VersionStream receives *semver.Version, writes it to the connected Pipes.
type VersionStream struct {
	Streams []VersionPipeWriter
}

// Pipe connects a Pipe, returns the connected Pipe left-end.
func (p *VersionStream) Pipe(s plumber.Piper) plumber.Piper {
	p.Sink(s)
	return s
}

// Sink connects an ending Piper.
func (p *VersionStream) Sink(s plumber.Flusher) {
	x, ok := s.(VersionPipeWriter)
	if !ok {
		panic(
			fmt.Errorf("Cannot Pipe a %T on %T", s, p),
		)
	}
	p.Streams = append(p.Streams, x)
}

// Unpipe disconnect a connected Pipe.
func (p *VersionStream) Unpipe(s plumber.Flusher) {
	// todo: add sync
	x, ok := s.(VersionPipeWriter)
	if !ok {
		panic(
			fmt.Errorf("Cannot Pipe a %T on %T", s, p),
		)
	}
	i := -1
	for e, pp := range p.Streams {
		if pp == x {
			i = e
			break
		}
	}
	if i > -1 {
		p.Streams = append(p.Streams[:i], p.Streams[i+1:]...)
	}
}

// Flush flushes the connected Pipes.
func (p *VersionStream) Flush() error {
	for _, pp := range p.Streams {
		if err := pp.Flush(); err != nil {
			return err
		}
	}
	return nil
}

// Write a *semver.Version on the connected Pipes.
func (p *VersionStream) Write(d *semver.Version) error {
	for _, pp := range p.Streams {
		if err := pp.Write(d); err != nil {
			return err
		}
	}
	return nil
}
```

# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)
