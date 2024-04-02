
# run main.go in cmd/run
.PHONY: run
run:
	@templ generate
	@go run ./cmd/main.go

# run main.go in cmd/create
.PHONY: create
create:
	@go run ./cmd/create

# run main.go in cmd/seed
.PHONY: seed
seed:
	@go run ./cmd/seed

# run main.go in cmd/drop
.PHONY: drop
drop:
	@go run ./cmd/drop

# generate templ files
.PHONY: templ
templ:
	@templ generate


.PHONY: watch
watch:
	air

# recreate database
.PHONY: recreate
recreate:
	@go run ./cmd/drop
	@go run ./cmd/create
	@go run ./cmd/seed