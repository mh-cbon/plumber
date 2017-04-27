// Package plumber builds pipes to transform a data stream.
package plumber

// Flusher can Flush.
type Flusher interface {
	Flush() error
}

// Sinker can Sink to a Flusher.
type Sinker interface {
	Sink(Flusher)
}

// Piper is a Flusher+Sinker that can Pipe/Unpipe.
type Piper interface {
	Flusher
	Sinker
	Pipe(Piper) Piper
	Unpipe(Flusher)
}
