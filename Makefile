.DEFAULT_GOAL := install

build: tfac.go
	go mod tidy
	go build tfac.go

install: tfac.go
	go install tfac.go
