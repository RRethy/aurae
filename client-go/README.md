# client-go

Go gRPC client for Aurae

## Installation

```shell
go get github.com/aurae-runtime/aurae/client-go
```

## Usage

### Connecting with TLS

```go
tlsConfig := &tls.Config{} // Provide your TLS configuration

client, err := aurae.NewClient(context.TODO(), "localhost:8080", grpc.WithTransportCredentials(credentials))
if err != nil {
    panic(err)
}
```

### Interacting with auraed

```go
exec := &cellv0.Executable {
    Command: "sleep 2",
}
request := &cellv0.CellServiceStartRequest {
    CellName: "foobar",
    Executable: exec,
}
response, err := client.Runtime().Cell().V0().Start(context.TODO(), request)
if err != nil {
    panic(err)
}

fmt.Printf("Executable running with pid %d\n", response.Pid)
```
