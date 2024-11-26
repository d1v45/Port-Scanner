# Port Scanner

A fast and efficient port scanner built using Go, designed to scan a range of ports on a given host concurrently. It supports concurrency control, progress reporting, and customizable timeouts.

## Features:
- Scans a range of ports concurrently using goroutines.
- Limits the number of concurrent goroutines to avoid overwhelming the system.
- Reports open ports in real-time.
- Supports customizable timeout settings.
- Handles error conditions gracefully and logs results.

## Prerequisites:
- Go 1.18+ installed.
- Basic understanding of Go, networking, and port scanning.

## Installation

1. Clone the repository or download the Go source code.

   ```bash
   git clone https://github.com/yourusername/go-port-scanner.git
   cd go-port-scanner
   ```

2. Build the Go project.

   ```bash
   go build -o portscanner main.go
   ```

## Usage

To run the port scanner, use the following command:

```bash
go run main.go <IP> <start_port> <end_port> <max_concurrency>
```

- `<IP>`: The target host's IP address or domain name.
- `<start_port>`: The starting port number to scan.
- `<end_port>`: The ending port number to scan.
- `<max_concurrency>`: The maximum number of concurrent goroutines (threads) to scan ports.

### Example

Scan ports 80-100 on `192.168.1.1`, using a maximum of 20 concurrent goroutines:

```bash
go run main.go 192.168.1.1 80 100 20
```

## Output

The program will display the open ports in real-time:

```bash
Scanning ports 80-100 on host 192.168.1.1 with max concurrency of 20...
Port 80 is open
Port 81 is open
Port 85 is open
...
Port scan completed.
```

### Options:
- If no ports are open, the program will simply complete without displaying any open ports.
- If an invalid port range is provided, the program will output an error message and exit.

## Code Structure

- **scanPort**: Scans a single port on the given host.
- **scanPorts**: Scans a range of ports concurrently, controlling the number of goroutines used.
- **main**: Takes user input from command-line arguments, validates it, and starts the scanning process.

## Contributing

Feel free to fork this repository and submit pull requests for improvements or bug fixes. Contributions are welcome!

### Improvements and features to consider:
- Add support for UDP scanning.
- Implement more detailed error handling and logging.
- Add customizable timeout durations per port.
- Implement retries with exponential backoff for slow or unreachable hosts.
