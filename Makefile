BINARY_NAME=scraper
IMAGE_NAME=jwilder/scraper

all: build build-image

build: deps
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) .

build-image: build
	docker build -t $(IMAGE_NAME) .

clean:
	rm -f $(BINARY_NAME)

run:
	kubectl run scraper -it --image $(IMAGE_NAME) --rm=true --restart=Never

deps:
	go get ./...

