DB_HOST=mysql
DB_PORT=3306
DB_USER=root
DB_PASS=root
DB_NAME=orders

createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate: 
	migrate -path=sql/migrations \
	-database "mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" \
	-verbose up

migratedown:
	migrate -path=sql/migrations \
	-database "mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" \
	-verbose down

.PHONY: migrate migratedown createmigration