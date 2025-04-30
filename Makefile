build:
	docker build -t pimentafm/stress-test:1.0 .

run:
	docker run --rm pimentafm/stress-test:1.0 --url=https://stackoverflow.com --requests=100 --concurrency=5

help:
	docker run pimentafm/stress-test:1.0 --help

.PHONY: build run help