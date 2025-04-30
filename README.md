# Stress Test Tool

A Docker-based tool for performing HTTP stress tests on a specified URL.

## Building the Project
To build the Docker image, run:
```bash
make build
```

## Running the Stress Test
Execute the stress test using the Docker image with the following command:
```bash
make run
```

`make run` will run the following command `docker run --rm pimentafm/stress-test:1.0 --url=https://stackoverflow.com --requests=100 --concurrency=5`

## Help information
```bash
make help
```