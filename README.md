# JWT CLI Tool

A lightweight, command-line utility written in Go to decode JSON Web Tokens (JWT), format their payload into human-readable JSON, and validate their expiration claims.

## Prerequisites
- Go 1.21 or higher installed locally.

## Installation

1. Clone the repository:
   ```bash
   git clone [https://github.com/Isshaan-Dhar/jwt-cli.git](https://github.com/Isshaan-Dhar/jwt-cli.git)
   cd jwt-cli
   ```

2. Build the binary (optional, but recommended):
   ```bash
   go build -o jwt-cli
   ```

## Usage

Pass a JWT token as the first argument to the program:

```bash
# If using the built binary
./jwt-cli <your_token_here>

# If running directly via Go
go run main.go <your_token_here>
```

## Example Output

```text
$ go run main.go eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0LXVzZXIiLCJleHAiOjE4OTM0NTYwMDB9.dummy_signature

=== HEADER ===
{
  "alg": "HS256",
  "typ": "JWT"
}
=== PAYLOAD ===
{
  "exp": 1893456000,
  "sub": "test-user"
}
=== STATUS ===
TOKEN VALID (expires at: Tue, 01 Jan 2030 00:00:00 UTC)
```

## Testing

This tool includes automated tests covering valid, expired, and malformed token scenarios. To run the test suite:

```bash
go test ./... -v
```