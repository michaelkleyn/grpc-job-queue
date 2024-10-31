# gRPC Job Queue Service

A simple gRPC-based Job Queue Service implemented in Go. This service allows clients to submit jobs and retrieve their statuses. It's designed as a foundational project that can be expanded into a Distributed Task Queue system.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Project Structure](#project-structure)
- [Usage](#usage)
  - [Running the Server](#running-the-server)
  - [Running the Client](#running-the-client)
- [Extending the Application](#extending-the-application)
- [Next Steps](#next-steps)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Introduction

This project is a simple implementation of a Job Queue Service using gRPC in Go. It demonstrates basic concepts of gRPC, including defining services, implementing server and client applications, and message exchange. The project serves as a stepping stone toward building a more complex Distributed Task Queue system.

## Features

- **Job Submission**: Clients can submit jobs with a unique ID and payload.
- **Job Status Retrieval**: Clients can query the status of their submitted jobs.
- **In-Memory Storage**: Uses an in-memory map to store job statuses.
- **Thread-Safe Access**: Implements mutex locks to prevent race conditions.
- **Extensible Design**: The codebase is structured to allow easy expansion and integration into larger projects.

## Prerequisites

- **Go**: Version 1.16 or higher. [Download Go](https://golang.org/dl/)
- **Protocol Buffers Compiler (`protoc`)**: [Installation Guide](https://grpc.io/docs/protoc-installation/)
- **Go Plugins for `protoc`**:
  - `protoc-gen-go`: Install via `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `protoc-gen-go-grpc`: Install via `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

Ensure that your `GOPATH/bin` is added to your `PATH` environment variable.

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/yourusername/job-queue.git
   cd job-queue
   ```
2. **Initialize Go Modules**

   ```bash
   go mod init github.com/yourusername/job-queue
   ```

   Replace 'github.com/yourusername/job-queue' with your actual module path if different

3. **Install Dependencies**

```bash
go mod tidy
```

4. **_Generate gRPCCode_**

```bash
protoc \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/jobqueue.proto
```

## Project Structure

```plaintext
job-queue/
├── client/
│   └── client.go
├── proto/
│   ├── jobqueue.proto
│   ├── jobqueue.pb.go
│   └── jobqueue_grpc.pb.go
├── server/
│   └── server.go
├── go.mod
└── README.md
```

- `client/`: Contains the client application code.
- `proto/`: Contains the `.proto` file and the generated Go code.
- `server/`: Contains the server application code.
- `go.mod`: Go module file.
- `README.md`: Project documentation.

## Usage

# Running the Server

1. **Navigate to the `server` directory:**

```bash
  cd server
```

2. **_Run the server application:_**

```bash
  go run server.go
```

The server will start and listen on port `50052`.

**Sample Output:**

```plaintext
2024/10/29 10:00:00 Server is running on port 50052...
```

# Running the Client

1. **_Open a new terminal window._**

2. **_Navigate to the `client` directory_**

```bash
cd client
```

3. **_Run the client application:_**

```bash
go run client.go
```

The client will submit a job and then request its status.

**_Sample Output:_**

```plaintext
2024/10/29 10:00:01 SubmitJob Response: Job submitted successfully
2024/10/29 10:00:01 GetJobStatus Response: Job job123 is completed
```

# Server Logs

The server terminal will display logs for job submissions and status requests:

```plaintext
2024/10/29 10:00:01 Job job123 submitted with payload: Process data X
2024/10/29 10:00:01 Job job123 status requested: completed
```
