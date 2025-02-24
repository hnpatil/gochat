# gochat

GoChat is a lightweight yet powerful chat server built with Go and the Gofr framework, designed for high scalability. It facilitates efficient message exchange over TCP, supporting multiple concurrent clients with minimal overhead.

## Features

- **Send Messages to Multiple Recipients**: Send messages to one or more recipients efficiently.
- **Spaces for Grouped Messages**: Organize related messages within designated spaces for better conversation management.
- **High Write Throughput**: Optimized for handling a large volume of messages efficiently.

## Prerequisites

- **Go**: Ensure that Go is installed on your system. You can download it from the [official Go website](https://golang.org/dl/).

## Getting Started

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/hnpatil/gochat.git
   cd gochat
2. **Install Dependencies**:
    ```bash
   go mod tidy
3. **Set Configuration**
   <br>Before running the application, set the required environment variables:
    ```bash
   export CASS_HOST=<Cassandra Host>
   export CASS_KEYSPACE=<Cassandra Keyspace>
   export CASS_PORT=<Cassandra Port>
   export API_KEYS=<Server API Key>
