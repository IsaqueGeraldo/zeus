# Zeus - Environment Variable Manager ⚡

Zeus is an environment variable manager that uses a database to store them persistently. With Zeus, you can define, retrieve, remove, list, and export environment variables efficiently, making it an essential tool for managing your app's environment configuration.

## Features 🛠️

- **Set**: Define an environment variable.
- **Get**: Retrieve the value of an environment variable.
- **Unset**: Remove an environment variable.
- **Clear**: Remove all environment variables.
- **List**: List all environment variables.
- **Export**: Export all environment variables to a file.

## Installation 🔧

To install Zeus, you can use the following methods:

### Option 1: Using `go install` (recommended)

```bash
go install github.com/IsaqueGeraldo/zeus@latest
```

This will download, compile, and install the `zeus` binary directly to your `$GOPATH/bin` directory.

### Option 2: Manual Installation

If you prefer to clone and compile the project manually, follow these steps:

1. Clone the repository to your machine:

   ```bash
   git clone https://github.com/IsaqueGeraldo/zeus.git
   cd zeus
   ```

2. Install the Go dependencies:

   ```bash
   go mod tidy
   ```

3. Compile the project:

   ```bash
   go build -o zeus
   ```

Alternatively, you can directly run the following to download and compile:

```bash
go install
```

## Usage 🚀

Once you've installed `zeus`, you can use it via the terminal with the `zeus` command.

### Example Commands:

#### 1. Set an environment variable:

```bash
zeus set MY_VARIABLE value
```

#### 2. Get the value of an environment variable:

```bash
zeus get MY_VARIABLE
```

#### 3. Remove an environment variable:

```bash
zeus unset MY_VARIABLE
```

#### 4. Remove all environment variables:

```bash
zeus clear
```

#### 5. List all environment variables:

```bash
zeus list
```

#### 6. Export all environment variables to a file:

```bash
zeus export -o .env
```

## Flags 📌

- `-s`, `--source`: Specifies the path to the environment file (if not provided, it defaults to the `ZEUS_SOURCE_DIR` environment variable).
- `-o`, `--output`: Specifies the output file for the export command (default is `.env`).

### Example usage with the `source` and `output` flags:

```bash
zeus -s /path/to/env/file export -o /path/to/output/file.env
```

## How It Works 🧠

Zeus uses the [odin](https://github.com/IsaqueGeraldo/odin) package to manage environment variables. The environment file can be passed via the `--source` flag, which specifies the path to the file. If the flag is not provided, the program will use the default environment file defined by the `ZEUS_SOURCE_DIR` environment variable.

## License 📄

This package is open-source and available under the MIT License.
