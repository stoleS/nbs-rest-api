# nbs-rest

`nbs-rest` is a Go-based REST proxy that interfaces with the National Bank of Serbia (NBS) SOAP API. This project aims to provide a more accessible and modern RESTful interface for interacting with the NBS services.

## Features

- **RESTful API**: Translates SOAP requests to RESTful endpoints.
- **Go-based**: Written in Go for performance and reliability.
- **Easy Integration**: Simplifies the process of integrating with the NBS API.

## Installation

To get started with `nbs-rest`, you need to have Go installed on your system. Then, you can clone the repository and build the project.

```sh
git clone https://github.com/stoleS/nbs-rest.git
cd nbs-rest
go build
```

## Usage

Run the `nbs-rest` server using the following command:

```sh
./nbs-rest
```

By default, the server will start on port 8000. You can customize the port by setting the `PORT` environment variable.

```sh
PORT=9090 ./nbs-rest
```

## Endpoints

`nbs-rest` exposes several RESTful endpoints corresponding to the NBS SOAP API functionalities. Below are some examples:

### Get Exchange Rates

```
GET /CoreService/GetCompanyCount
```

Fetch the number of registered companies

#### Request

```json
{
  "City": "Novi Sad"
}
```

#### Response

```json
{
  "GetCompanyCountResult": "70662"
}
```

## Configuration

Configuration options are managed through environment variables:

- `PORT`: Specifies the port on which the server runs (default is 8080).
- `NBS_API_URL`: The base URL for the NBS SOAP API.
- `LOG_LEVEL`: Sets the logging level (e.g., DEBUG, INFO, WARN, ERROR).

## Development

If you wish to contribute or modify the project, follow these steps to set up the development environment:

1. Fork the repository and clone it to your local machine.
2. Ensure you have Go installed (version 1.16 or later).
3. Install dependencies:

   ```sh
   go mod tidy
   ```

4. Run dev mode:

   ```sh
   go run cmd/api/main.go
   ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or support, please open an issue on the GitHub repository or contact the maintainer at [predragstosic29@gmail.com](mailto:predragstosic29@gmail.com).

---

Thank you for using `nbs-rest`! We hope it simplifies your integration with the National Bank of Serbia's services.
