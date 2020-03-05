.PHONY: all cleanup go swagger typescript 

all:
	docker-compose up
	cleanup

go:
	docker-compose run orders-api-go

swagger:
	docker-compose run orders-api-swagger

typescript:
	@echo "Note: Typescript generated based on swagger documentation."
	docker-compose run orders-api-typescript
	./typescript/cleanup.sh

cleanup:
	./typescript/cleanup.sh
