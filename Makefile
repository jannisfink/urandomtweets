rpi:
	env GOOS=linux GOARCH=arm go build urandomtweets.go

build:
	go build

install:
	go install
