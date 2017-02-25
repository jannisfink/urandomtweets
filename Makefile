rpi:
	env GOOS=linux GOARCH=arm go build urandomtweets.go

build:
	go build

install:
	go install

upload:
	scp ./urandomtweets pi3:/home/pi/urandomtweets/

publish:
	make rpi
	make upload
