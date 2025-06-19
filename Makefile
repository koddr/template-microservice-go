POSTGRESQL_URL=postgres://postgres:password@localhost:5432/postgres

start-postgres-docker:
	@docker run -d \
	--name postgres_analytics \
	-e POSTGRES_PASSWORD=password \
	-p 5432:5432 \
	postgres:17

migrate-up:
	@migrate -database ${POSTGRESQL_URL}?sslmode=disable -path internal/attachments/migrations up

migrate-down:
	@migrate -database ${POSTGRESQL_URL}?sslmode=disable -path internal/attachments/migrations down

build-app-docker:
	@docker build -t app_analytics .

start-app-docker:
	@docker run -d \
	-e API_SERVER_PORT=8080 \
	-e API_SERVER_AUTH_USERNAME=user \
	-e API_SERVER_AUTH_PASSWORD=password \
	-e API_SERVER_DATABASE_URL=${POSTGRESQL_URL} \
	-p 8080:8080 \
	app_analytics

run:
	@API_SERVER_PORT=8080 \
	API_SERVER_AUTH_USERNAME=user \
	API_SERVER_AUTH_PASSWORD=password \
	API_SERVER_DATABASE_URL=${POSTGRESQL_URL} \
	go run .