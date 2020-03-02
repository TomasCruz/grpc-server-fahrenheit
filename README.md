# grpc-server-fahrenheit
gRPC server for celsuis/fahrenheit conversion

## API
See [API messages](#api-messages) for message formats referenced here

### Health
[Health proto](https://github.com/grpc/grpc/blob/master/src/proto/grpc/health/v1/health.proto "gRPC Health v1 proto")
[grpc_health_v1 library](https://godoc.org/google.golang.org/grpc/health/grpc_health_v1)

grpc_health_v1.HealthClient's Check method is used for healthcheck
HealthClient's Watch is not currently used as server doesn't expose streaming APIs

### C2F
converts celsius to fahrenheit
#### input
ConversionMsg
#### output
ConversionMsg

### F2C
converts fahrenheit to celsius
#### input
ConversionMsg
#### output
ConversionMsg

## API messages
- NoParamsMsg is an empty message
- HealthMsg contains bool health
- ConversionMsg contains double number

## Build
### Prerequisites:
- standard Docker installation
- standard gRPC installation

### building
- run 'source ./env'
- either create a db with 'make create_db', or start it if it's been created with 'make start_db'
- run 'make'

## Run
'make run' from terminal

## Unit tests
'make test' from terminal for unit tests

## Integration tests
If not ran already, run steps for build and run.
In another terminal, navigate to directory containing Makefile, then 'source ./env', then 'make integration'
