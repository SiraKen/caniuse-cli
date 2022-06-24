NAME = cani
BIN := bin/$(NAME)

LDFLAGS := -w \
		   -s

.PHONY: build
build:
	go build -ldflags "$(LDFLAGS)" -o $(BIN)

.PHONY: install
install:
	@make build
	@sudo cp $(BIN) /usr/local/bin
	@echo "Installed $(BIN) to /usr/local/bin"

.PHONY: uninstall
uninstall:
	@sudo rm /usr/local/bin/$(NAME)
	@echo "Uninstalled $(NAME) from /usr/local/bin"
