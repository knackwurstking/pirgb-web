
all: clean build

clean:
	go clean -v

build:
	cd frontend; yarn; cd ..
	go build -v ./cmd/pirgb-web

dev: 
	cd frontend; yarn; cd ..
	go build -v -tags dev ./cmd/pirgb-web
