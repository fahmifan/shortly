.PHONY: swagger
swagger:
	mkdir -p gen
	swagger generate server --exclude-main -A shortly -t gen -f ./swagger.yml

.PHONY: migrate
# migrate up
migrate:
	mkdir -p sqlitedb
	migrate -path repository/sqlite/migrations -database "sqlite3://sqlitedb/shortly.db?cache=shared&mode=rwc&_journal_mode=WAL" up