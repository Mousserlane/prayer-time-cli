BINARY_NAME = prayer-time-cli

MAIN_PACKAGE = ./cmd/cli
GO_BUILD_FLAGS = 
INSTALL_DIR = /usr/local/bin

build:
	@echo "Compiling ${BINARY_NAME}..."
	go build ${GO_BUILD_FLAGS} -o ${BINARY_NAME} ${MAIN_PACKAGE}
	@echo "Build complete: ${BINARY_NAME}"

install: build
	@echo "Installing ${BINARY_NAME} to ${INSTALL_DIR}"
	@mkdir -p ${INSTALL_DIR}
	mv ${BINARY_NAME} ${INSTALL_DIR}/${BINARY_NAME}
	@echo "${BINARY_NAME} installed to ${INSTALL_DIR}"

run: build
	@echo "Running ${BINARY_NAME}"
	./${BINARY_NAME}
