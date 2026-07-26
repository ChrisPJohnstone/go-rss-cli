.PHONY: build run clean

build:
	go build -o rss .

run: build
	./rss

clean:
	rm -f rss
