1. go mod init
2. go get google.golang.org/protobuf/cmd/protoc-gen-go

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

3. https://stackoverflow.com/a/71407805/2845389
4. export GOBIN=$(go env GOPATH)/bin
5. protoc \
   --go_out=invoicer \
   --go_opt=paths=source_relative \
   --go-grpc_out=invoicer \
   --go-grpc_opt=paths=source_relative \
   invoicer.proto

6.
