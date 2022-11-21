BIN=pirgb-web

all: clean build

clean:
	go clean -v
	@rm -rf frontend/dist/

build-frontend:
	@cd frontend && npm install && npm run build

build:
	go clean -v
	make build_frontend
	go build -v ./cmd/pirgb-web

build-dev: 
	make clean
	make build_frontend
	go build -v -tags dev ./cmd/pirgb-web

install:
ifeq ($(USER), root)
	@echo "err: don't run as root!"
	@exit 1
else
	@mkdir -p ${HOME}/.local/bin && cp ${BIN} ${HOME}/.local/bin/${BIN}
	@mkdir -p ${HOME}/.config/systemd/user/ && cp init/${BIN}.service ${HOME}/.config/systemd/user/${BIN}.service && systemctl --user daemon-reload
endif
