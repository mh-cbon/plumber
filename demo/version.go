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
