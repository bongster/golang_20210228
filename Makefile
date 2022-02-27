.PHONY: all

check-sqlc:
	which sqlc || brew install sqlc
check-install:
	which swagger || GO11MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

docs: check-install
	GO11MODULE=off swagger generate markdown -f ./src/swagger.yml --output ./docs/docs.md

swagger: check-install
	GO11MODULE=off swagger generate spec -w ./src -o ./src/swagger.yml --scan-models

test:
	go test -v -cover ./src/...
sqlc: check-sqlc
	sqlc generate -f src/sqlc.yaml
