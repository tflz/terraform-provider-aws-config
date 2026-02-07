# OpenTofu Provider: OpenTofu Landing Zone for AWS Configuration Loader

_This provider is built on the [Terraform Provider Scaffolding](https://github.com/hashicorp/terraform-provider-scaffolding-framework)_

## Overview

This OpenTofu provider enables users to load and manage HCL-defined configurations for OpenTofu Landing Zone for AWS. It provides a seamless integration between OpenTofu and the Landing Zone infrastructure-as-code patterns.

## Features

- Load HCL configuration files
- Manage OpenTofu Landing Zone for AWS resources
- Integrated with OpenTofu ecosystem

## Requirements

- [OpenTofu](https://opentofu.org/docs/intro/install) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.24

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to the provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
