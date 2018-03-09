build: build/pro

build/pro: main.go commands.go utils.go
	mkdir build
	go build -o build/pro

install: setup build/pro
	cp ./build/pro $(HOME)/bin/pro


setup:
	mkdir $(HOME)/.pro