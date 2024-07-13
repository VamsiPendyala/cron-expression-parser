# Cron Parser CLI

The Cron Parser CLI is a command-line application that parses a cron string and expands each field to show the times at which it will run. The application supports the standard cron format with five time fields (minute, hour, day of month, month, and day of week) plus a command.

## Features

- Parses standard cron format strings
- Expands each field into a list of times
- Displays the results in a formatted table

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/cron_parser.git
    cd cron_parser
    ```

2. Build the CLI application:

    ```sh
    go build -o cron_parser cmd/cron_parser/main.go
    ```

## Usage

To use the Cron Parser CLI, run the application with the `--cron` flag followed by the cron string you want to parse. 

Example:

```sh
./cron_parser --cron "*/15 0 1,15 * 1-5 /usr/bin/find"
```

The output will be formatted as a table showing the expanded fields:

```sh
minute         0 15 30 45
hour           0
day of month   1 15
month          1 2 3 4 5 6 7 8 9 10 11 12
day of week    1 2 3 4 5
command        /usr/bin/find
```

## Testing

The project includes unit tests for the parser. To run the tests, use the following command:

```sh
go test ./pkg/parser
```

## Project Structure

```
cron_parser/
├── cmd/
│   └── cron_parser/
│       └── main.go
├── pkg/
│   └── parser/
│       ├── parser.go
│       └── parser_test.go
└── go.mod
```

- `cmd/cron_parser/main.go`: The main entry point for the CLI application.
- `pkg/parser/parser.go`: The parser logic for parsing and expanding cron strings.
- `pkg/parser/parser_test.go`: Unit tests for the parser.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## Contact

If you have any questions, feel free to reach out at vamsi.pendyala8@example.com.

