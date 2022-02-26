.PHONY: all

check-install:
	which swagger || GO11MOBULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

docs: check-install
	GO11MOBULE=off swagger generate markdown -f ./swagger.yml --output ../docs/docs.md

swagger: check-install
	GO11MOBULE=off swagger generate spec -w . -o ./swagger.yml --scan-models

test:
	go test -v -cover ./...
sqlc:
	sqlc generate -f sqlc.yaml
