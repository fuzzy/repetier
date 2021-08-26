all: test
	go get -v -x .

test:
	make -C tests
