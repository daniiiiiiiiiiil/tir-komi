include .env
export

migrate-up:
	migrate -path migrations -database ${DATABASE_URL} up

migrate-down:
	migrate -path migrations -database ${DATABASE_URL} down

migrate-force:
	migrate -path migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" force 1

migrate-create:
	migrate create -ext sql -dir migrations -seq $(NAME)