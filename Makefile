# make compile version=

version=

compile:
	GOOS=linux GOARCH=amd64 go install -ldflags="-X 'main.VERSION=${version}'"

build:
	go build -ldflags="-X 'main.VERSION=${version}'"