all: install

build/pro: main.go commands.go utils.go
	go build -o build/pro

install: build/pro
	cp ./build/pro $(HOME)/bin/pro


#setup:
#	mkdir $(HOME)/.pro
#	echo "{}" > $(HOME)/.pro-config