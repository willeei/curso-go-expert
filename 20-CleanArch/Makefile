migrate:
	migrate -path=internal/infra/database/migrations -database "mysql://root:root@tcp(mysql:3306)/orders" -verbose up

migratedown:
	migrate -path=internal/infra/database/migrations -database "mysql://root:root@tcp(mysql:3306)/orders" -verbose down

.PHONY: migrate migratedown