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

//go:generate plumber semver_gen.go "github.com/mh-cbon/semver/*semver.Version"

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
