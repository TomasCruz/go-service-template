# go-service-template
Go web service template

## API
### /health/
displays service health status

#### Return codes:
- OK (200) if service is working as expected
- InternalServerError (500) in case of general errors

### /hello/
saying hello to the user, using a name passed as input. If name is "*", request fails with InternalServerError. If there is no input, response message is "hello world"

#### input
path variable containing the username

#### Return codes:
- OK (200) for successfuly saying hello
- InternalServerError (500) in case of general errors

## Build
From terminal, run 'source ./env' then 'make'

## Run
run build steps, then 'make run'

## Unit tests
run 'make test' for unit tests

## End-to-end tests
If not ran already, run steps for building, then 'make run'
In another terminal, navigate to directory containing Makefile, run 'source ./env' then 'make end2end'
