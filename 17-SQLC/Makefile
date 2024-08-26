create-migration:
	migrate create -ext=sql/migrations -dir=sql/migrations -seq $(name)

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

rollback:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down

.PHONY: migrate rollback create-migration