create-db:
	docker run --name go-chat-db -d -p 9042:9042 scylladb/scylla --cpus=1

create-keyspace:
	docker exec -it go-chat-db cqlsh -e "CREATE KEYSPACE IF NOT EXISTS gochat WITH replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };"

start-db:
	docker start go-chat-db

swagger-init:
	swag init --output ./static && cp ./static/swagger.json ./static/openapi.json

build-app:
	docker build --tag gochat .

create-app:
	docker run --name go-chat-core -d -p 8000:8000 gochat --cpus=1

run:
	docker-compose up -d
