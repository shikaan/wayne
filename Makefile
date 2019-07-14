PROJECT_NAME := $(shell basename "$(PWD)")
ENTRY_POINT := "${PWD}/cmd/main.go"

# TODO: grab me from environemtn
PORT := 8080

MAKEFLAGS += --silent
MAKEFLAGS += --debug=n

run: stop
	@sleep 2
	clear
	@echo "  >  Starting ${PROJECT_NAME}..."
	go run ${ENTRY_POINT}

build:
	@echo "  >  Building binary for ${PROJECT_NAME}"
	go build -o ${PROJECT_NAME} ${ENTRY_POINT}

clean:
	rm "${PWD}/${PROJECT_NAME}"

stop:
	@echo "Killing process `lsof -i:$${PORT} -t`"
	@-kill -9 `lsof -i:$${PORT} -t` 2> /dev/null || true

start:
	make run &
	while true; do \
	  inotifywait -q -r . -e modify,move,delete,create 1> /dev/null; \
	  make run & echo '\n'; \
	done;