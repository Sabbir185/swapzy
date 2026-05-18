include .env
export $(shell sed 's/=.*//' .env)

# Default migration path
MIGRATIONS_PATH=migrations

# Commands
migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

migrate-force:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force 1

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

migrate-version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version