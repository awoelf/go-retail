include .env

init_graph:
	go run github.com/99designs/gqlgen init; \

generate_graph:
	go run github.com/99designs/gqlgen generate; \

create_migrations:
	sqlx migrate add -r init; \

migrate_up:
	sqlx migrate run --source ./migrations --database-url ${DB_URL}

migrate_down:
	sqlx migrate revert --source ./migrations --database-url ${DB_URL}

docker_create:
	docker run -d --name ${DB_CONTAINER_NAME} -p ${DB_PORT}:${DB_PORT} -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} postgres:16.2

docker_delete: docker_stop
	docker rm ${DB_CONTAINER_NAME}

docker_run:
	docker start ${DB_CONTAINER_NAME}

docker_stop:
	docker stop ${DB_CONTAINER_NAME}

create_db:
	docker exec -it ${DB_CONTAINER_NAME} createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

build:
	if [ -f "${BINARY}" ]; then \
		rm ${BINARY}; \
		echo "Deleted ${BINARY}"; \
	fi
	@echo "Building binary"; \
	go build -o ${BINARY} server.go; \

run: build
	./${BINARY}
	@echo "Server running..."; \

stop: 
	@echo "Stopping server"; \
	@-pkill -SIGTERM -f "./${BINARY}"; \
	@echo "Server stopped"; \

# Code for initializing a container with a .sql file.
# -v ${PWD}/mysql/init.sql:/docker-entrypoint-initdb.d/0_init.sql