# make compile version=

version=

build:
	go build -ldflags="-X 'main.VERSION=${version}'"

install:
	GOOS=linux GOARCH=amd64 go install -ldflags="-X 'main.VERSION=${version}'"