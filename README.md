# Supercluster Client

The desktop client for Supercluster. It comprises of a Go executable that serves a React frontend.

## Running
- Install React dependencies and build frontend:
``` shell
cd ui
yarn
yarn build
```
- From the project root, you can start the server by running `go run cmd/server/main.go`
- This will start the server, which can be accessed at `http://localhost:8080`
