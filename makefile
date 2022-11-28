default: test

test:
	cd	calculator
	go test	-v	./...

build:
	cd	calculator && go build -o calculator

