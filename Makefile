run:
	cd server; go run cmd/main.go;

build:
	cd server/cmd; go build -o ../bin/go-react;

exec:
	./server/bin/go-react

startapp: build exec