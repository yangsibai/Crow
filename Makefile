PORT := 9093

less:
	lessc static/css/style.less static/css/style.css

build:
	go build -o Crow

test: less build
	go run

run: stop build
	nohup ./Newtonia>/dev/null 2>&1 &

stop:
	-lsof -t -i:${PORT} | xargs kill

.PHONY: build, run, test
