PHONY:
run:
		go run cmd/main.go

swag:
	swag init -g cmd/main.go