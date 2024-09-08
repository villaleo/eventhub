# EventHub

![EventHub Logo](/doc/assets/images/color-lg.svg)

EventHub is a Go-based event management system that leverages Protocol Buffers (Protobuf), gRPC, and MongoDB. It is designed with high-performance logging using Uber Zap and modularized dependency injection through Uber FX.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- **Go 1.20+**: The language this project is written in. [Installation instructions](https://go.dev/doc/install).
- **MongoDB**: The database management system used. Create a free MongoDB Atlas cluster [here](https://www.mongodb.com/docs/atlas/getting-started/) and follow the installation steps.

- **Mage**: A build tool CLI, that automatically uses Go functions as Makefile-like runnable targets. [Installation instructions](https://magefile.org/).

## Setup

1. **Clone the Repository:**

```bash
git clone https://github.com/villaleo/eventhub.git
cd eventhub
```

2. **Install Dependencies:**

After installing the prerequisites, run this command to install all other required tools.

```bash
mage
```

3. **Generate Protobuf Files:**
 After setup, generate the necessary protobuf files.

 ```bash
 mage gen
 ```

### Starting the Server

1. **Default Port:**

Starts the server with the default port 50051.

```bash
mage server:run
```

2. **Using a Custom Port:**

Starts the server using a specific port.

```bash
mage server:start 50051
```

### Starting the Client

1. **Default Address:**

Runs the client using the default address `localhost:50051`.

```bash
mage client:run
```

2. **Using a Custom Address:**

Starts the client using a specific address.

```bash
mage client:start localhost:50051
```

### Linting

Linting is ran by default before starting the server. It is required and cannot be disabled.

```bash
mage lint
```

### All Commands

To see all the commands available for this project, run this command.

```bash
mage -l
```

Or checkout a specific target's usage.

```bash
mage -h server:start
```
