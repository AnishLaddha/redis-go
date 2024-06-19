# Redis-Go

**Redis-Go** is a Redis clone written in Go. It is compliant with RESP (Redis Serialization Protocol) version 2.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Make sure you have Go installed on your machine. You can download and install Go from [here](https://golang.org/dl/).

You will also need redis-cli to test/interact with the server. You can download it [here](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/)*

*_\*Note: that this will also download redis's server software._*

To use redis-go, you need to ensure nothing else is running on the standard redis port, 6379

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/redis-go.git
    ```

2. Navigate into the project directory:

    ```sh
    cd redis-go
    ```


### Running Redis-Go

To build the Redis-Go server:

```sh
make build
```

To run the Redis-Go server:

```sh
make run
```

### Testing Redis-Go

To test if the server is running correctly, open a new terminal window and use the `redis-cli`:

```sh
redis-cli ping
```
You should expect a PONG response if everything is working correctly.

## Using redis-go
You can find out more about the resp2 protocol [here](https://redis.io/docs/latest/develop/reference/protocol-spec/)

redis-go currently supports: Ping, Echo, Set (with expiry), and Get. It will support Del very shortly