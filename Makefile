.PHONY: all build run test

all: clean build run

build:
	cd src && go build -o redis-go

run:
	cd src && ./redis-go

test:
	redis-cli ping

clean:
	rm -f src/redis-go
	rm -f persist.aof