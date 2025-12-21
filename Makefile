redis-cli:
	docker exec -it notification-service-redis-1 redis-cli

migrate-up:
	migrate -path ./migrations -database "postgres://notify:notify@localhost:5432/notify?sslmode=disable" up

migrate-test-up:
	migrate -path ./migrations -database "postgres://notify:notify@localhost:5433/notify_test?sslmode=disable" up

migrate-test-down:
	migrate -path ./migrations -database "postgres://notify:notify@localhost:5433/notify_test?sslmode=disable" down