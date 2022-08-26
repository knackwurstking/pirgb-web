SYSTEMD_DIR=${HOME}/.config/systemd/user
SYSTEMD_FILE_FROM=init
SYSTEMD_FILE=pirgb-web.service

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
	@if [[ ! -e pirgb-web ]]; then echo -e "\nRun \`make\` first!"; exit 1; fi
	@mkdir -pv ${HOME}/.local/bin && cp -v pirgb-web ${HOME}/.local/bin

	@if [[ "${USER}" == "root" ]]; then echo -e "\nerr: don't run as root"; exit 1; fi

	@if [[ ! -e "${SYSTEMD_DIR}/${SYSTEMD_FILE}" ]]; then mkdir -pv ${SYSTEMD_DIR} && cp -v ${SYSTEMD_FILE_FROM}/${SYSTEMD_FILE}	${SYSTEMD_DIR} && systemctl --user daemon-reload && echo -e "\n\033[32mstart/enable ${SYSTEMD_FILE}\033[0m"; fi
