# Redis-Go

**Redis-Go** is a Redis clone written in Go. It is compliant with RESP (Redis Serialization Protocol) version 2.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Make sure you have Go installed on your machine. You can download and install Go from [here](https://golang.org/dl/).

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/redis-go.git
    ```

2. Navigate into the project directory:

    ```sh
    cd redis-go
    ```

3. Build the project:

    ```sh
    cd src
    go build -o redis-go
    ```

### Running Redis-Go

To run the Redis-Go server:

```sh
./redis-go
```

### Testing Redis-Go

To test if the server is running correctly, open a new terminal window and use the `redis-cli`:

```sh
redis-cli ping
```
You should expect a PONG response if everything is working correctly.

