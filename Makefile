create-db:
	docker run --name go-chat-db -e POSTGRES_USER=gochat -e POSTGRES_PASSWORD=gochat -e POSTGRES_DB=gochat -p 5432:5432 -d postgres

start-db:
	docker start go-chat-db