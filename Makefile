
all: clean build_frontend build

clean:
	go clean -v

build_frontend:
	@cd frontend && yarn && yarn run build

build:
	go build -v ./cmd/pirgb-web

dev: 
	make clean
	make build_frontend
	go build -v -tags dev ./cmd/pirgb-web

install:
	# TODO: ...
