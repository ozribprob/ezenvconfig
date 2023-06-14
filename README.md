# ezenvconfig

## Introduction

The `ezenvconfig` package is a simple, easy-to-use Go library that provides functionality to extract values from environment variables with additional fallbacks to predefined defaults or aliases. 

## Purpose

The primary purpose of this library is to simplify the way applications handle environment variables. This is particularly helpful in containerized applications or serverless functions where configurations are commonly injected through environment variables.

## Usage

### Installation

```shell
go get github.com/problem-company-toolkit/ezenvconfig
```

### Basic Usage

```go
package main

import (
	"fmt"
	"os"
	"github.com/problem-company-toolkit/ezenvconfig"
)

func main() {
	// Define an entry
	entry := ezenvconfig.Entry{
		Name: "YOUR_VARIABLE",
		Aliases: []string{"YOUR_VARIABLE_ALIAS1", "YOUR_VARIABLE_ALIAS2"},
		Default: "default_value",
		Optional: true,
	}

	// Extract the environment variable
	value, err := ezenvconfig.ExtractFromEnv(entry)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value: %s\n", value)
}
```

In this example, `ezenvconfig` will try to extract the value of "YOUR_VARIABLE" from the environment. If not found, it will try the aliases in order. If it still doesn't find a value, it will fall back to "default_value".

### Handling Not Found

You can define a custom function to be executed when a variable is not found and no default value is provided:

```go
entry := ezenvconfig.Entry{
	Name: "YOUR_VARIABLE",
	Aliases: []string{"YOUR_VARIABLE_ALIAS1", "YOUR_VARIABLE_ALIAS2"},
	OnNotFound: func() {
		fmt.Println("Variable not found!")
		os.Exit(1)
	},
	Optional: false,
}
```

In this case, the function prints a message and exits the program if the variable is not found.

## Testing

Tests are written using the Ginkgo BDD testing framework. Run them with `go test`, or if you have Ginkgo installed, you can use `ginkgo -r`.