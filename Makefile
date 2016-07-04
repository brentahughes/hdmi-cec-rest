all: deps fix

deps:
	go get github.com/chbmuc/cec
	go get github.com/gorilla/mux


fix:
	go fmt
	go fix
