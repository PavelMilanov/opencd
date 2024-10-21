# make compile version= logfile=

version=
logfile=

build:
	go build -ldflags="-X 'main.VERSION=${version}' -X 'main.LOGFILE=${logfile}'"

install:
	GOOS=linux GOARCH=amd64 go install -ldflags="-X 'main.VERSION=${version}' -X 'main.LOGFILE=${logfile}'"
	cp ~/go/bin/linux_amd64/opencd .