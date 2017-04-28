---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
---
# {{.Name}}

{{template "badge/travis" .}} {{template "badge/appveyor" .}} {{template "badge/goreport" .}} {{template "badge/godoc" .}} {{template "license/shields" .}}

{{pkgdoc "plumber.go"}}

# {{toc 5}}

# Install
{{template "go/install" .}}

# Generator

To help you to deal with the step of interface implementation,
 `plumber` comes with a command line program to generate your own typed pipes.

## Usage

#### $ {{exec "plumber" "-help" | color "sh"}}

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

#### > {{cat "demo/main.go" | color "go"}}

Following code is the implementation of various
pipe `transformer` that works with `*semver.Version` type.

#### > {{cat "demo/version.go" | color "go"}}

Following is the generated code to build pipes
to work with `*semver.Version` values.

#### > {{cat "demo/semver_gen.go" | color "go"}}

# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)
