# Tonic Web Framework

Tonic is a minimalist Web framework implemented (at least planned) in Go. The framework is currently under construction and aims to provide a lightweight structure for handling HTTP requests and responses (to start with).

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)
- [TODO](#todo)

## Features

- Minimalistic HTTP framework.
- Handle HTTP GET and POST requests.
- Define routes and their handlers seamlessly.
- Simple request and response parsing.
- Basic Content-Type and status code handling.

## Installation

To install and run Tonic, you will need to have Go installed on your machine. Follow the steps below to clone the repository, build, and run the server:

```sh
git clone https://github.com/yourusername/tonic.git
cd tonic
go build ./cmd/main.go
./main
```

## Usage

Once the server is running, you can make a GET request to the `/phones` endpoint:

```sh
curl http://localhost:8080/phones
```

This will return a simple HTML response.

## Project Structure

```
tonic/
├── cmd/
│   └── main.go         # Entry point for the application
├── internal/
│   ├── handlers/
│   │   ├── get.go      # Handler for GET requests
│   │   ├── post.go     # Handler for POST requests
│   │   └── request.go  # Defines Request and Response structs and parsing logic
│   └── server/
│       └── server.go   # Server logic to handle incoming requests
│       └── response.go # Logic to create and send responses
├── README.md           # Project Documentation
├── go.mod              # Go module file
├── go.sum              # Go module dependencies
└── flattened.go        # A flattened version of the project (incomplete)
```

## Contributing

We welcome contributions to enhance the functionality and capabilities of Tonic. Feel free to fork the repository, create a new branch, and submit a pull request.

1. Fork the repo
2. Create a new branch: `git checkout -b feature-branch`
3. Make your changes and commit: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin feature-branch`
5. Open a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## TODO

- [ ] Implement POST request handling.
- [ ] Improve error handling and logging mechanism.
- [ ] Add support for more HTTP methods (PUT, DELETE, etc.).
- [ ] Enhance routing capabilities with parameters and middleware support.
- [ ] Add unit tests and integration tests for core functionalities.
- [ ] Improve configuration options for the server.
- [ ] Enhance documentation with more examples and detailed explanations.
- [ ] Implement support for HTTPS.
- [ ] Add more sample handler functions and routes for demonstration.
- [ ] Create benchmarks and performance tests.

---

Thank you for checking out Tonic! If you have any questions or suggestions, please feel free to open an issue or reach out.
