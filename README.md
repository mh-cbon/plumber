# emd

[![travis Status](https://travis-ci.org//mh-cbon/emd.svg?branch=master)](https://travis-ci.org//mh-cbon/emd) [![Appveyor Status](https://ci.appveyor.com/api/projects/status//github/mh-cbon/emd?branch=master&svg=true)](https://ci.appveyor.com/projects//mh-cbon/emd) [![Go Report Card](https://goreportcard.com/badge/github.com/mh-cbon/emd)](https://goreportcard.com/report/github.com/mh-cbon/emd) [![GoDoc](https://godoc.org/github.com/mh-cbon/emd?status.svg)](http://godoc.org/github.com/mh-cbon/emd) [![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Package plumber builds pipes to transform a data stream.


# TOC
- [Install](#install)
- [Generator](#generator)
  - [Usage](#usage)
    - [$ plumber -help](#-plumber--help)
  - [Cli examples](#cli-examples)
- [API example](#api-example)
  - [> main.go](#-maingo)
- [Recipes](#recipes)
  - [Release the project](#release-the-project)
- [History](#history)

# Install
```sh
go get github.com/mh-cbon/emd
```

# Generator

To help you to deal with the step of interface implementation `plumber`
comes with a command line program to generate your own typed pipes.

## Usage

#### $ plumber -help
```sh
plumber 0.0.0

Usage

	plumber [out] [pkg] [...types]

	out: 	Output destination of the results, use '-' for stdout.
	pkg: 	The package name of the generated code.
	types:	A list of fully qualified types such as
	     	'[]byte', 'semver.Version' or '*my.PointerType'.
```

## Cli examples

```sh
# Create a pipe of *tomate.SuperStruct in the package mysuperpkg
plumber - mysuperpkg *tomate.SuperStruct
```
# API example

Demonstrates how you can take advantage of this API to stream process the data

#### > main.go

```go
//+build ignore

//Package cmd implement a cli tool to manipulate Versions.
package main

import (
	"os"

	"github.com/mh-cbon/semver/cmd/stream"
)

func main() {

	src := os.Stdin

	pipeSrc := stream.NewByteReader(src)
	pipe := pipeSrc.
		Pipe(stream.NewBytesSplitter(' ', '\n')).
		Pipe(&stream.BytesTrimer{}).
		Pipe(&stream.VersionFromByte{SkipInvalid: true}).
		Pipe(&stream.VersionSorter{Asc: true}).
		Pipe(&stream.LastVersionOnly{}).
		Pipe(&stream.VersionToByte{}).
		Pipe(stream.NewBytesPrefixer("- ", "\n"))

	pipe.Sink(stream.NewByteSink(dest))

	if err := pipeSrc.Consume(); err != nil {
		panic(err)
	}
	os.Exit(0)
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
