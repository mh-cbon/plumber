---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
Name: emd
---
# {{.Name}}

{{template "badge/travis" .}} {{template "badge/appveyor" .}} {{template "badge/goreport" .}} {{template "badge/godoc" .}} {{template "license/shields" .}}

{{pkgdoc "plumber.go"}}

# {{toc 5}}

# Install
{{template "go/install" .}}

# Generator

To help you to deal with the step of interface implementation `plumber`
comes with a command line program to generate your own typed pipes.

## Usage

#### $ {{exec "plumber" "-help" | color "sh"}}

## Cli examples

```sh
# Create a pipe of *tomate.SuperStruct in the package mysuperpkg
plumber - mysuperpkg *tomate.SuperStruct
```
# API example

Demonstrates how you can take advantage of this API to stream process the data

#### > main.go

{{read "demo.go" | color "go"}}

# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)
