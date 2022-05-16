# depdep

`depdep` is a package for static analysis
to find dependencies to avoid.

You can definde the rule in a yaml file like:

```yaml
blocked:
  - from: ^example\.com/foo$
    to:
      - ^example\.com/baz$
  - from: ^example\.com/foo$
    to:
      - ^example\.com/bar$
      - ^example\.com/qux.*
```

`depdep` reports the import declarations which violate the rule.

## Install

```bash
go install github.com/shota3506/depdep/cmd/depdep@latest
```

## How to use

You can use `depdep` with go vet command.

```bash
go vet -vettool=$(which depdep) -depdep.config=${PATH_TO_YAML_FILE} ./...
```
