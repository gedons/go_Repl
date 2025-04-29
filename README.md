# Go-REPL API

Go-REPL is a lightweight Go API for building and interacting with a Read-Eval-Print Loop (REPL) environment. This project is designed to help developers experiment with Go code dynamically.

## Features

- Execute Go code snippets in a REPL environment.
- Evaluate expressions and return results in real-time.
- Lightweight and easy to integrate into other Go projects.

## Installation

To install Go-REPL, use `go get`:

```bash
go get github.com/gedons/go_Repl
```

## Usage

Here’s a basic example of how to use Go-REPL:

```go
package main

import (
    "github.com/gedons/go_Repl"
)

func main() {
    repl := gorepl.New()
    repl.Start()
}
```

## API Endpoints

If your Go-REPL includes an HTTP API, here are some example endpoints:

- `POST /execute` - Execute a Go code snippet.
  - **Request Body**: `{ "code": "fmt.Println(\"Hello, World!\")" }`
  - **Response**: `{ "output": "Hello, World!\n" }`

- `GET /history` - Retrieve the history of executed commands.
  - **Response**: `[ "fmt.Println(\"Hello, World!\")", "2 + 2" ]`

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

Special thanks to the Go community for their support and inspiration.
