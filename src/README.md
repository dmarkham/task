# Task - implementation

The server requests name and joke to 2 backing public APIs. Then it returns the result to its client per GET request received.

To run the server, 
- Go SDK should be installed.
- Build and run the project (go build -race -tags debug -o bin/task && ./bin/task).
- Open another terminal and query with curl (curl $SERVER_IP:5000). Watch command can be helpful to repeat.

