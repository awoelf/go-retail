include .env

init_graph:
	cd server; \
	go run github.com/99designs/gqlgen init; \

generate_graph:
	cd server; \
	go run github.com/99designs/gqlgen generate; \

create_migrations:
	cd server; \
	sqlx migrate add -r init; \

stop_containers:
	@echo "Stopping other docker containers"; \
	if [ $$(docker ps -q) ]; then \
		echo "Found and stopped containers"; \
		docker stop $$(docker ps -q); \
	else \
		echo "No containers running"; \
	fi

create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -v ${PWD}/mysql/init.sql:/docker-entrypoint-initdb.d/0_init.sql -e MYSQL_ROOT_PASSWORD=${MYSQL_PASSWORD} -d -p 3306:3306 mysql:8.1.0

run_container:
	docker start ${DB_DOCKER_CONTAINER}

mysql:
	docker run -it mysql bash -c 'mysql -h ${HOST} -u ${MYSQL_USER} -p"${MYSQL_PASSWORD}"'

build:
	if [ -f "${BINARY}" ]; then \
		rm ${BINARY}; \
		echo "Deleted ${BINARY}"; \
	fi
	@echo "Building binary"; \
	go build -o ${BINARY} server/server.go; \

run: build
	./${BINARY}
	@echo "Server running..."; \

stop: 
	@echo "Stopping server"; \
	@-pkill -SIGTERM -f "./${BINARY}"; \
	@echo "Server stopped"; \
