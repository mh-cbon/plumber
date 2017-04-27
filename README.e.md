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
# Init a basic emd file to get started.
plumber - mysuperpkg *tomate.SuperStruct
```
# API example

Demonstrates how you can take advantage of this API to build data transformation

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
