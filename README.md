# Terraform Provider Aria (Terraform Plugin Framework)

This is the [Terraform](https://www.terraform.io) provider for VMWare's Aria Automation Platform.

The provider is [published here](https://registry.terraform.io/providers/davidfischer-ch/aria/latest).

It has been developped by the CSC Team from the IT department of the State of Geneva (Switzerland).

Please be aware that Broadcom is not responsible neither involved on this project.

_This provider is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework). See [Which SDK Should I Use?](https://developer.hashicorp.com/terraform/plugin/framework-benefits) in the Terraform documentation for additional information._

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21

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

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

Fill this in for each provider

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate ./...`.
To format the code `make fmt` or `find . -name "*.go" -exec gofmt -s -w {} \;`

You have to set the following environment variables `ARIA_HOST` and `ARIA_REFRESH_TOKEN` before running tests. For the unit tests you can set those to dummy values.

To run the full suite of Unit tests, run `go test ./...`.

For running the acceptance tests you also have to set additionnal environment variables:

* `TF_VAR_test_org_id` to the organization you are targeting
* `TF_VAR_test_project_id` to an already provisioned writable project
* `TF_VAR_test_project_ids` to already provisioned writable projects (comma separated)
* `TF_VAR_test_icon_id` to an already provisioned Icon
* `TF_VAR_test_secret_id` to an already provisioned Secret
* `TF_VAR_test_abx_action_id` to an already provisioned ABX action

Resources generated by the acceptance tests will be generated "inside" given project.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
export ARIA_HOST=https://some-aria-host.net
export ARIA_INSECURE=false
export ARIA_REFRESH_TOKEN=*****
export ARIA_ACCESS_TOKEN=***** # If you have one, not required

export TF_VAR_test_org_id=2817c6e5-7408-449f-a86d-8f511105e5ba
export TF_VAR_test_project_id=2e34b115-dd18-48b3-a6af-f794469e5e0d
export TF_VAR_test_project_ids=8f274902-94dc-40fd-98b5-f06c68ae1237,a9441e75-57c0-46fa-9262-c06a47acb1a9,2e34b115-dd18-48b3-a6af-f794469e5e0d
export TF_VAR_test_icon_id=72a9a2c7-494e-31d7-afe8-cd27479c407e
export TF_VAR_test_secret_id=a9af6450-a0c6-42cf-921e-14f7f8db50b3
export TF_VAR_test_abx_action_id=8a7480d38e535332018e857e0d4f3437

make testacc
```
