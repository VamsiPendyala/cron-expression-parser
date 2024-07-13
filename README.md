### `README.md` ###

# Cron Parser #

This is a simple Go command-line application that parses a cron string and expands each field to show the times at which it will run. It only considers the standard cron format with five time fields (minute, hour, day of month, month, and day of week) plus a command. It does not handle special time strings such as "@yearly".

## Project Structure ##

```
cron_parser/
├── cmd/
│   └── cron_parser/
│       └── main.go
├── pkg/
│   └── parser/
│       ├── parser.go
│       └── parser_test.go
├── go.mod
└── go.sum
```

## Installation

1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd cron_parser
    ```

2. Initialize the Go module:
    ```sh
    go mod init cron_parser
    ```

3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

To run the program, use the following command:

```sh
go run ./cmd/cron_parser -cron="*/15 0 1,15 * 1-5 /usr/bin/find"
```

### Example

Input:
```sh
*/15 0 1,15 * 1-5 /usr/bin/find
```

Output:
```
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```

## Testing

To run the tests, use the following command:

```sh
go test ./pkg/parser
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
```

This `README.md` provides an overview of the project, installation instructions, usage examples, and testing instructions. Replace `<repository-url>` with the actual URL of your repository.