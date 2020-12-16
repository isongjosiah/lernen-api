build:
	go build -o bin/lernen-api main.go

run: build
	./bin/lernen-api

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' make run