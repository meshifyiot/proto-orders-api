# Interface definition for OrdersAPI

Be warned: this is en experimental repo.

## Workflow

- Make changes to \*.proto
- Build docker images, generate doc, clients and stubs: `make all`
- Tag a version
- Push to github
- Enjoy!

## Prerequisites

[Docker Compose](https://docs.docker.com/compose/install/)

## SwaggerUI for Gateway

If you are running mac then you can download this repo and then run:

    make swagger-ui

I'm sure that docker part will work on linux too, but PR's are welcome!

Another option is to use online swagger editor: https://editor.swagger.io/
