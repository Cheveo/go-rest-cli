# go-rest-template-cli
A CLI Tool to create domain driven REST Services

REST-CLI is a command-line tool built with Golang and Cobra that simplifies the process of creating a whole Golang REST API project skeleton based on a specified domain and Go module name.

## Installation

To install `rest-cli`, you can use `go get`:

```bash
go get github.com/Cheveo/go-rest-template-cli
```

## Usage

Once installed, you can use `rest-cli` to create a new Golang REST API project skeleton in a specific directory with the following flags:

- `--domain`: Specify the domain object for your REST API.
- `--directory`: Specify the directory where you want to create the project.
- `--goModName`: Specify the Go module name for your project.

### Example

```bash
rest-cli create --domain user --directory my-api-project --goModName example.com/test
```

This command will create a new Golang REST API project skeleton in the `my-api-project` directory, with the Go module name set to `example.com/test` and the domain set to `user`.

## License

This project is licensed under the [MIT License](LICENSE).

## Contribution

Contributions are welcome! Feel free to open an issue or submit a pull request.

## Author

- Danyal iqbal
- GitHub: [github.com/Sagato](https://github.com/Sagato)
- Website: [cheveo.de](https://cheveo.de)
