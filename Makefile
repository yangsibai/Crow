PORT := 9093

build:
	go build -o Crow

stop:
	-lsof -t -i:${PORT} | xargs kill

dev:
	go run

run: stop build
	nohup ./Crow>/dev/null 2>&1 &

.PHONY: build, stop, dev
