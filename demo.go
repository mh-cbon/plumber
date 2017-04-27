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
