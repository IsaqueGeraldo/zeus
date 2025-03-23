# Zeus - Environment Variable Manager ⚡

Zeus is an environment variable manager that uses a database to store them persistently. With it, you can define, retrieve, remove, and list environment variables in a simple and efficient way.

## Features 🛠️

- **Set**: Define an environment variable.
- **Get**: Retrieve the value of an environment variable.
- **Unset**: Remove an environment variable.
- **Clear**: Remove all environment variables.
- **List**: List all environment variables.

## Installation 🔧

Clone the repository to your machine:

```bash
git clone https://github.com/your-username/zeus.git
cd zeus
```

Install the Go dependencies:

```bash
go mod tidy
```

Compile the project:

```bash
go build -o zeus
```

Or, if you want to just download and compile it:

```bash
go install
```

## Usage 🚀

After compiling the program, you can use `zeus` directly from the terminal.

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

## Flags 📌

- `-s`, `--source`: Specifies the path to the environment file.

## Example usage with the `source` flag:

```bash
zeus -s /path/to/file set MY_VARIABLE value
```

## How It Works 🧠

Zeus uses the [odin](https://github.com/IsaqueGeraldo/odin) package to manage the environment variables. The environment file can be passed via the `--source` flag, which specifies the path to the file. If the flag is not provided, the program uses the default `./odin.db` directory.

## License 📄

This package is open-source and available under the MIT License.
