run:
	$(info ****************** make run ******************)
	go run main.go
build:
	$(info ****************** make build ******************)
	go build -o ./bin/gobang main.go
clean:
	$(info ****************** make clean ******************)
	rm ./bin/*
e2eTest:
	$(info ****************** make e2eTest ******************)
	go test -v ./test