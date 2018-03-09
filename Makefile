all: compile install

compile: main.go
	go build -o pro

install: ./pro
	cp ./pro $(HOME)/bin/pro


#setup:
#	mkdir $(HOME)/.pro
#	echo "{}" > $(HOME)/.pro-config