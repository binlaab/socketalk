# SocketTalk

SockeTalk (thank you, ChatGPT) is a socket-based chat written in Go as a starter project.

**Disclaimer**: I do not recommend its use in production environments, as it was originally developed as a starter project, not for real-life use.

## Quick start
1. Build the project:
   ```bash
   go build . -ldflags "-s -w"
   ```
2. Run it:
    ```
    ./socketalk
    ```

And that's it!

## Configuration

You can change the default values for the app via command-line arguments or by editing the source code.

### Command-line arguments
```
Usage:
        -a, --address: IP address where the server will listen for connections. Defaults to all interfaces (0.0.0.0).
        -p, --port: Port where the server will listen for connection. Defaults to port 8000.
```

### Source code
Go to `main.go`, lines 21 and 22. Values should go between double quotes.

```go
var IP = "0.0.0.0" // Changes the listening interface
var PORT = "8000" // Changes the listening port
```

If you find any issues or want to change the code, feel free to submit a pull request or issue!

