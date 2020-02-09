# grpc-server-fahrenheit
gRPC server for celsuis/fahrenheit conversion

## API
### Health
displays service health status
#### input
NoParamsMsg
#### output
HealthMsg

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
