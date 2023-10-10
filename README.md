# Sybline Go Client

This is the go client for talking to [Sybline]().

### How to install

Go install command:
```
go get ....
```

### Documentation

All of the go documentation can be found at the [Sybline Documentation]() website.


### Proto Command

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    mq.proto
```