# Domino REST API Golang SDK Examples

This folder contains example usages of Domino REST API Golang SDK.

## Setup

Create a file named `env.json` inside `pkg/resources` directory. It should have the following format:

```json
{
  "baseUrl": "https://myserver.io",
  "credentials": {
    "userName": "user",
    "passWord": "pass",
    "scope": "$DATA,$SETUP",
    "type": "BASIC"
  }
}
```

## Trying it out

For `access` and `connector` examples, you can run it directly via:

```shell
go run access/access.go
```

For examples inside `operations` folder, edit `operation.go` to use your desired operation from `basis` and `setup` folders, then run it via:

```shell
go run operations/operation.go
```
