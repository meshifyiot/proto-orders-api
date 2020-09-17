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

You can see Gateway API documentation by downloading this repo and then run:

    make swagger-ui
