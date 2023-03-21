# kevin
GRPC golang application to be used as a simple, structured baseline for k8s cluster testing

# Generate from .proto 
```
protoc --go_out=pkg/ --go_opt=paths=source_relative \
    --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative \
    api/pong.proto
```
