all: deps build

deps:
	@echo "Installing go dependencies"
	go get

build:
	@echo "Compiling"
	go install

fix:
	@echo "Runnong go tools to clean up source code"
	go fmt
	go fix
