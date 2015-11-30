PORT := 9093

less:
	lessc static/css/style.less static/css/style.css

build:
	go build -o Crow

test: less build
	./Crow

run: stop build
	nohup ./Crow>/dev/null 2>&1 &

stop:
	-lsof -t -i:${PORT} | xargs kill

pull:
	git pull

update: pull run

.PHONY: build, run, test
