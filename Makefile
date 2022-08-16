
all: clean build

clean:
	go clean -v

build_frontend:
	@cd frontend && yarn && yarn run build

build:
	go build -v ./cmd/pirgb-web

dev: 
	cd frontend; yarn; cd ..
	go build -v -tags dev ./cmd/pirgb-web
