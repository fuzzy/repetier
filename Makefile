all: test
	go build -v -x .

test:
	make -C tests
