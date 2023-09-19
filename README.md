# Zeus

Zeus is a Go package that provides database connectivity and functions for managing environment settings. It uses the Gorm library for database operations and is designed to simplify working with environment configurations.

## Installation

To use the Zeus package in your Go project, you can import it as follows:

```go
import (
    "github.com/IsaqueGeraldo/zeus"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)
```

You should also have the Gorm library and a SQLite database driver installed in your project.

## Usage

### Initializing the Database Connection

The `Bootstrap` function is used to initialize the database connection. It takes the path to the SQLite database file as an argument.

```go
db, err := zeus.Bootstrap("my_database.db")
if err != nil {
    // Handle the error
}
```

### Managing Environment Settings

Zeus provides functions to get, set, and list environment settings.

#### Get an Environment Setting

To retrieve an environment setting by its key, use the `Getenv` function:

```go
value, err := zeus.Getenv("my_key")
if err != nil {
    // Handle the error
}
```

#### Set an Environment Setting

To set an environment setting, use the `Setenv` function:

```go
err := zeus.Setenv("my_key", "my_value")
if err != nil {
    // Handle the error
}
```

#### List All Environment Settings

To retrieve a list of all environment settings, use the `Environ` function:

```go
envSettings, err := zeus.Environ()
if err != nil {
    // Handle the error
}
```

## Error Handling

Zeus functions return errors to handle various scenarios, such as database connection issues and record not found errors. Ensure that you handle errors appropriately in your code.

## Contributing

Contributions to the Zeus package are welcome! If you have ideas for improvements or new features, please open an issue or submit a pull request on the [GitHub repository](https://github.com/IsaqueGeraldo/zeus).

## License

Zeus is released under the MIT License. See the [LICENSE](LICENSE) file for details.

```
Please make sure to replace any placeholder URLs or paths with the actual URLs and paths relevant to your project, such as the link to your GitHub repository and the license details.
```
