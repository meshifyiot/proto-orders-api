.PHONY: all build go swagger typescript clean

all: go swagger typescript

build:
	docker-compose build
	
go:
	docker-compose run orders-api-go

swagger:
	docker-compose run orders-api-swagger

typescript:
	@echo "\n[Generate Typescript (based on swagger documentation)...]"
	docker-compose run orders-api-typescript
	./typescript/cleanup.sh

clean:
	rm -rf go/orders/*.go
	rm -rf swagger/orders/*.json
	rm -rf typescript/orders/*