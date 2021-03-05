build:
	go build -o bin/translator cmd/main.go

deploy: build
	sudo cp bin/translator /usr/local/bin/translator
	sudo chmod +x /usr/local/bin/translator

run:
	go run cmd/main.go