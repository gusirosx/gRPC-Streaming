# gRPC Server Time Streaming Application

This demonstration presents:

- `./server`: a gRPC server application with an gRPC that streams the current time in the response
- `./client`: a small program to query the server and show the response messages

## How to run this example

1. run grpc server

```sh
$ make run_server
```
Make a request using the client:

2. run gin client

```sh
$ make run_client
```
or

```sh
   go run client/* -duration 5
```

3. Observe the output, it should be receiving and printing a message every second as the server sends them.

```sh
    gRPC established to timeserver, starting to stream
    Current TimeStamp: 2022-03-01T21:23:47Z
    Current TimeStamp: 2022-03-01T21:23:48Z
    Current TimeStamp: 2022-03-01T21:23:49Z
    Current TimeStamp: 2022-03-01T21:23:50Z
    Current TimeStamp: 2022-03-01T21:23:51Z
    end of stream
```
## How to create proto files

1. use the makefile

```sh
$ make generate