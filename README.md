# go-service-template
Go web service template

## API
### /health/
displays service health status

#### Response codes and content:
- NoContent (204) if service is working as expected
- InternalServerError (500) in case of general errors, JSON body containing errorMessage

### /hello/
saying hello to the user, using a name passed as input. If name is "*", request fails with NotAcceptable. If name is not a valid UTF-8 string, request fails with UnprocessableEntity. If there is no input, response message is "hello world"

#### input
path variable containing the username

#### Response codes and content:
- OK (200) for successfuly saying hello, JSON body containing message
- NotAcceptable (406) for unacceptable input, JSON body containing errorMessage
- UnprocessableEntity (422) for invalid input, JSON body containing errorMessage
- InternalServerError (500) in case of general errors, JSON body containing errorMessage

## Build
From terminal, run 'source ./env' then 'make'

## Run
run build steps, then 'make run'

## Unit tests
run 'make test' for unit tests

## End-to-end tests
If not ran already, run steps for building, then 'make run'
In another terminal, navigate to directory containing Makefile, run 'source ./env' then 'make end2end'
