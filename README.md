# Zeus

Zeus is a lightweight command-line tool for managing environment variables. It allows you to set, get, list, remove, or clear environment variables stored in a SQLite database. This tool is designed to simplify working with environment variables in your development environment.

## Features

- Set an environment variable with a key and value.
- Get the value of an environment variable by specifying its key.
- Rename an environment variable.
- List all environment variables.
- Remove an environment variable by specifying its key.
- Clear all environment variables.

## Installation

To install Zeus, you can use the `go get` command:

```bash
go get github.com/IsaqueGeraldo/zeus
```

## Usage

Zeus provides several subcommands to interact with environment variables:

### Set an Environment Variable

```bash
zeus set [key] [value]
```

Example:

```bash
zeus set API_KEY my-secret-key
```

### Get the Value of an Environment Variable

```bash
zeus get [key]
```

Example:

```bash
zeus get API_KEY
```

### Rename an Environment Variable

```bash
zeus rename [oldkey] [newkey]
```

Example:

```bash
zeus rename OLD_KEY NEW_KEY
```

### List All Environment Variables

```bash
zeus environ
```

### Remove an Environment Variable

```bash
zeus unset [key]
```

Example:

```bash
zeus unset API_KEY
```

### Clear All Environment Variables

```bash
zeus clearenv
```

### Search for Environment Variables

```bash
zeus find [key]
```

Example:

```bash
zeus find API
```

## Database

Zeus stores environment variables in a SQLite database named `odin.db` in the current working directory. It will be created automatically if it doesn't exist.

## Note

- Environment variable keys are case-insensitive and will be converted to uppercase.
- Only alphanumeric characters are allowed in keys; any other characters will be replaced with underscores.

## License

Zeus is released under the MIT License. See [LICENSE](LICENSE) for more information.

---

Feel free to contribute to this project or report any issues on [GitHub](https://github.com/IsaqueGeraldo/zeus).
