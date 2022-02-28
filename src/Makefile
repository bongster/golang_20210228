.PHONY: all

check-install:
	which swagger || GO11MOBULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

docs: check-install
	GO11MODULE=off swagger generate markdown -f ./swagger.yml --output ../docs/docs.md

swagger: check-install
	GO11MODULE=off swagger generate spec -w . -o ./swagger.yml --scan-models

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/payments?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/payments?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/payments?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/payments?sslmode=disable" -verbose down 1

testdb:
	PGPASSWORD=postgres psql postgres -h localhost -p 5432 -U postgres -c "CREATE DATABASE testdb" || true

dropdb:
	PGPASSWORD=postgres psql postgres -h localhost -p 5432 -U postgres -c "DROP DATABASE testdb" || true

test_migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/testdb?sslmode=disable" -verbose up

test_migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/testdb?sslmode=disable" -verbose down

test:
	go test -v -cover -timeout 60s ./...
sqlc:
	sqlc generate -f sqlc.yaml
