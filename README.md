# zeus

The `zeus` package is a lightweight Go library for managing environment variables stored in a SQLite database. It provides functions to bootstrap a database connection, retrieve, set, clear, and list environment variables.

## Installation

To use the `zeus` package, make sure you have Go installed on your machine. You can import the package into your Go project using the following command:

```bash
go get github.com/IsaqueGeraldo/zeus
```

## Usage

### Bootstrap

Before using the functions to manage environment variables, you need to establish a connection to the SQLite database:

```go
package main

import (
	"github.com/IsaqueGeraldo/zeus"
)

func main() {
	zeus.Bootstrap() // Initializes the database connection
}
```

### Functions

#### `Getenv(key string) (Environment, error)`

Retrieves an environment variable by its key.

**Parameters:**

- `key`: The key of the environment variable to retrieve.

**Returns:**

- `Environment`: The environment variable if found.
- `error`: An error if the key is not found or any other error occurs.

**Example:**

```go
env, err := zeus.Getenv("MY_KEY")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println("Value:", env.Value)
}
```

#### `Setenv(key, value string) error`

Sets or updates an environment variable in the database.

**Parameters:**

- `key`: The key of the environment variable.
- `value`: The value of the environment variable.

**Returns:**

- `error`: An error if saving the variable fails.

**Example:**

```go
err := zeus.Setenv("MY_KEY", "my_value")
if err != nil {
    fmt.Println(err)
}
```

#### `Clearenv(key string) error`

Removes an environment variable by its key.

**Parameters:**

- `key`: The key of the environment variable to remove.

**Returns:**

- `error`: An error if the key is not found or removal fails.

**Example:**

```go
err := zeus.Clearenv("MY_KEY")
if err != nil {
    fmt.Println(err)
}
```

#### `Environ() ([]Environment, error)`

Lists all environment variables stored in the database.

**Returns:**

- `[]Environment`: A slice of all environment variables.
- `error`: An error if the listing fails.

**Example:**

```go
envs, err := zeus.Environ()
if err != nil {
    fmt.Println(err)
} else {
    for _, env := range envs {
        fmt.Printf("Key: %s, Value: %s\n", env.Key, env.Value)
    }
}
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
