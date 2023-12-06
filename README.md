# Domino REST API Golang SDK

Domino REST Go SDK is a package build designed to assist developers in integrating Domino with their applications.

&copy; 2023 HCL America Inc. Apache-2.0 license [https://www.apache.org/licenses/LICENSE-2.0](https://www.apache.org/licenses/LICENSE-2.0)

## 📔 Documentation

- [Domino REST API documentation](https://opensource.hcltechsw.com/Domino-rest-api/index.html)
- [Using Domino REST API Go SDK examples](/examples/)

## ⬇️ Installation

First step is to install and configure Golang. Below is the link to follow the installation and configuration process.
```
https://go.dev/doc/install
```

## ⭐ Highlights

- Supports both JavaScript and TypeScript.
- Has built-in methods for the following Domino REST API calls:
  - **Basis**:
    - `/document`
    - `/document/{unid}`
    - `/bulk/create`
    - `/bulk/{unid}`
    - `/bulk/update`
    - `/bulk/delete`
    - `/query`
    - `/lists`
    - `/lists/{name}`
    - `/listspivot/{name}`
  - **Setup**:
    - `/design/{designType}/{designName}`
    - `/admin/scope`
    - `/admin/scopes`

## 📦 Importing

You can import the whole SDK by:
```
// Using golang get command
go get github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go
```

You can also import within your Golang file:
```
import "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
```

## 🔬 Overview

![Domino REST API Go SDK Model](/docs/sdk-model.png)

Domino REST API Go SDK has four moving parts:

- `Access`
- `Server`
- `Connector`
- `Session`

### ℹ️ Access

`Access` is a function that facilitates your access to the Domino REST API server. It takes in a `baseUrl`, which is your Idp provider, as well as your credentials, such as your `username`, `password`, `scope` and `type` (the authentication type: `basic` or `oauth`).

### ℹ️ Server

`Server` is a function that gets information on what APIs are available on your current server. It takes in a url to your Domino REST API server as a parameter. This class produces a `ConnectorMethods` interface base on your chosen API.

### ℹ️ Connector

`Connector` is the function that does the actual communication between the Domino REST API Go SDK and your Domino REST API server.

### ℹ️ Session

`UserSession` is a class that contains all the operation you can perform on your Domino REST API server. It includes built-in methods, and a generic request method if you want to execute an operation on your own.

### 🎮 Running a Domino REST API operation using the SDK

Here is an example of how to use the four moving parts mentioned above in order to execute one Domino REST API Node SDK.

```Go
import (
	"fmt"
	gosdk "https://github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

config := new(gosdk.Config)
config.BaseUrl = "http://localhost:8880"
config.Credentials.Scope = "$DATA"
config.Credentials.Type = "BASIC"
config.Credentials.UserName = "username"
config.Credentials.Password = "password"

access, err := config.DominoAccess()
if err != nil {
    fmt.Println(err.Error())
}

server, err := gosdk.Server(access.GetBaseUrl())
if err != nil {
    fmt.Println(err.Error())
}

connector, err := server.GetConnector("basis")
if err != nil {
    fmt.Println(err.Error())
}

session := new(gosdk.SessionConfig)
session.AccessMethods = access
session.ConnectorMethods = connector

session := session.UserSession()

// Get a domino document example
result, err := session.GetDocument("parameter1", "parameter2", ...)
if err != nil {
    fmt.Println(err)
}
fmt.Println(result)
```

For other examples, please go to our [examples](/examples/operations/).