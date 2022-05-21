.PHONY: swagger
swagger:
	mkdir -p gen
	swagger generate server --exclude-main -A shortly -t gen -f ./swagger.yml