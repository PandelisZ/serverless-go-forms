# Serverless forms with go

The endpoint currently supports POSTS with JSON payloads in the following format:

```json
{
    "from": "sender@example.com",
    "message": "Hello there"
}
```

Or application/x-www-form-urlencoded content in the format 

```
from=sender@example.com&message=Hello%20there
```

If you need to change this format for your needs you can edit [lib/types/types.go](lib/types/types.go)

## Configuration

go-forms currently supports mailgun with plans to support slack and regular smtp or amazon sqs.

You can configure your mailgun keys by renaming `config.example.yml` -> `config.dev.yml` and entering your details.

If you want to create a production config use `config.prod.yml`

### CORS
If you want to lock down the requests to just your dowmain use the cors.origin setting in [serverless.yml](serverless.yml)

## Building and Deploying

This project can be deployed with the [serverless framework](https://github.com/serverless/serverless).

To build binaries for deployment run `make` before deploying.

You willl also need [go dep](https://github.com/golang/dep) to install the vendor dependencies.

```
npm install -g serverless
make
serverless deploy
```






