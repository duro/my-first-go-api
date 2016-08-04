
db_url := postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:5432/${DB_NAME}?sslmode=disable

migrate-create:
	migrate -url ${db_url} -path ${MIGRATION_PATH} create ${NAME}

migrate-up:
	migrate -url ${db_url} -path ${MIGRATION_PATH} up

migrate-undo:
	migrate -url ${db_url} -path ${MIGRATION_PATH} migrate -1
