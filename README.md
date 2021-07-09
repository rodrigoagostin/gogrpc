```
$ docker-compose up -d
$ docker exec -it gogrpc bash
$ protoc --proto_path=grpc grpc/protofile/*.proto --go_out=.
$ protoc --proto_path=grpc grpc/protofile/*.proto --go_out=. --go-grpc_out=.
$ evans -r repl
```
