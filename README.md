# GoRetro

GoRetro is a CLI tool inspired by Refit and Retrofit, designed to simplify the creation of Go API clients. 
It automatically generate HTTP client implementations from struct definitions to generate code that can be used to
interact with HTTP APIs.

## Features
- Zero dependencies.
- Simple and easy to use.
- Generates code in your project. No runtime reflection.
- Automatically generate HTTP clients from interface definitions.
- Supports common HTTP methods (GET, POST, etc.).
- Handles query parameters, path parameters, and request bodies.
- Supports basic, API key, and bearer token authentication.


## Installation

To install GoRetro, use `go install`:

```sh
go install github.com/AlterNayte/go-retro@latest
```

## Usage
Define Your API Interface
Create an interface that describes your API endpoints. Use struct tags to specify the HTTP method, path, and other details.

```go
package cats

import "time"

type Fact struct {
    Id        string    `json:"_id"`
    V         int       `json:"__v"`
    Text      string    `json:"text"`
    UpdatedAt time.Time `json:"updatedAt"`
    Deleted   bool      `json:"deleted"`
    Source    string    `json:"source"`
}

type CatFactsAPI struct {
    GetFacts        func() ([]Fact, error)                   `method:"GET" path:"/facts"`
    AnimalFacts     func(animal_type string) ([]Fact, error) `method:"GET" path:"/facts" query:"animal_type"`
    Fact            func(id string) (*Fact, error)           `method:"GET" path:"/facts/{id}"`
}
```

## Generate the Client
Run GoRetro to generate the client code. This can be done by running the goretro command in your terminal.

```sh
go-retro
```
By default, GoRetro will look for any api definition files throughout your project and generate the client code in the 
generated folder ('/go-retro/generated' by default). You can also specify the output directory using the -output flag.


## Options
- -dir: Specify the directory to search for API interface definitions.
- -output: Specify the output directory for the generated client code.



## Use the Generated Client
Use the generated client in your application.

```go
package main

import (
    goretro "your-app/go-retro/generated"
    "context"
    "fmt"
)

func main() {
    client := goretro.NewCatsFactsAPIClient("https://api.example.com")

    // Example of calling a GET method
    data, err := client.GetFacts(context.Background())
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Data: %+v\n", data)
}
```

## Authentication
GoRetro supports three types of authentication: None, Basic, API Key, and Bearer Token. You can specify the 
authentication type in your API interface definition using the *auth* tag.

## Supported Authentication Types
- None: No authentication required.
- Basic: Basic authentication using a username and password.
- API Key: API key authentication.
- Bearer: Bearer token authentication.


```go
// omitted code...
type CatFactsAPI struct {
	GetFacts       func() ([]Fact, error)                   `method:"GET" path:"/facts" auth:"Bearer`
	AnimalFacts func(animal_type string) ([]Fact, error) `method:"GET" path:"/facts" query:"animal_type" auth:"API`
	Fact        func(id string) (*Fact, error)           `method:"GET" path:"/facts/{id}" auth:"Basic`
}
```

## Setting Authentication
You can set the authentication credentials when creating the client.



## Example Project Structure

```
myproject/
├── api/
│   └── api.go
├── generated/
│   └── binance_client_gen.go
├── main.go
└── go.mod
```

In this example, api/api.go contains your API interface definitions, and running goretro generates the client code in generated/binance_client_gen.go.

## Common Code
GoRetro generates a base file containing common code, such as authentication types and helper functions. Ensure this file is included in your project to avoid conflicts.

```go
package goretro

type AuthType int

const (
    AuthNone AuthType = iota
    AuthBasic
    AuthAPIKey
    AuthBearer
)

// Other common code...
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request on GitHub.

# License
GoRetro is released under the MIT License. See [LICENSE](LICENSE) for details.
