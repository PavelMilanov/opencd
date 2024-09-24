# make compile version=

version=

compile:
	GOOS=linux GOARCH=amd64 go install -ldflags="-X 'main.VERSION=${version}'"
	cp ~/go/bin/opencd opcd/usr/bin/
	size=$(du -k opcd/usr/bin/opencd | awk '{print $1}')
	sed -i "s/Version: [0-9].[0-9].[0-9]/Version: ${version}/g" opcd/DEBIAN/control
	sed -i "s/Installed-Size:\s\d+/Installed-Size: ${size}/g" opcd/DEBIAN/control
	dpkg-deb --build opcd/

build:
	go build -ldflags="-X 'main.VERSION=${version}'"

install:
	GOOS=linux GOARCH=amd64 go install -ldflags="-X 'main.VERSION=${version}'"