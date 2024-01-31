# =========================================================================== #
# HELPERS
# =========================================================================== #

## help: print this help message
.PHONY: help 
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: help 
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# =========================================================================== #
# DEVELOPMENT
# =========================================================================== #

## run/api: run the cmd/api application
.PHONY: help 
run/api:
	go run ./cmd/api

## db/psql: connect to the database using psql
.PHONY: help 
db/psql:
	psql ${GREENLIGHT_DB_DSN}

## db/migrations/new name=$1: create a new db migration
.PHONY: help 
db/migrations/new:
	@echo 'Creating migration files for ${name}'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: help 
db/migrations/up: confirm
	@echo 'Running up migrations'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up

# =========================================================================== #
# QUALITY CONTROL
# =========================================================================== #

## audit: tidy dependencies, format, vet and detect race conditions in code
.PHONY: audit 
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...

	@echo 'Vetting code'
	go vet ./...

	@echo 'Detect race conditions'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies'
	go mod tidy
	go mod verify

	@echo 'Vendoring dependencies'
	go mod vendor 

# =========================================================================== #
# BUILD
# =========================================================================== #

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api'
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api