# go-api-template
The main target for this project is to provide a go lang starting template with the most used features that we will need to have for a project development.

- [go-template](#go-template)
  - [Dependencies](#dependencies)
    - [oapi-codegen](#oapi-codegen)
  - [Makefile](#makefile)
  - [Git hooks](#git-hooks)


## Dependencies

- **golang**: the project uses [go1.19](https://go.dev/dl/) so that is the bare minimum that should be installed
- **oapi-codegen**: to generate the code for the api from openapi spec and for using the gin framework (if not installed the makefile will intall it)

### oapi-codegen

To install the dependency just run the following code:

```
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4
```

To generate the code we run the script at the top of the file `internal/api/server.go` and the output will be placed at `internal/api/openapi` or use the makefile.

For more info go to oapi-codegen [repo](https://github.com/deepmap/oapi-codegen).

## Makefile

The project comes with a makefile as an helper for some actions.

To install dependencies:

```shell
make install-dependencies
```

To run the project:

```shell
make run
```

To run the project tests:

```shell
make tests
```

To run the linter:

```shell
make lint
```

To run go generate to generate code:

```shell
make generate
```

To do a local build:

```shell
make build
```

## Git Hooks

In order to validate the commit message and run the go lint when we are working locally, we need something that performs a verification before the commit is done.

For that, you must do the following steps:

1. Install [golangci-lint](https://golangci-lint.run/usage/install/)

```
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
```

2. Make the hooks executable

```
chmod +x '.git-hooks/commit-msg'
```

```
chmod +x '.git-hooks/pre-commit'
```

3. Define git hook directory

```
git config core.hooksPath '.git-hooks'
```

After that, the hook will be executed before the commit is done.