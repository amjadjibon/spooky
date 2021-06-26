VERSION = "0.0.1"
change-version:
	@echo $(VERSION)>VERSION
	@echo "package constant\n\n//Version constant of the service\nconst Version = \"$(VERSION)\"">pkg/constant/version.go

update-module:
	go get -u github.com/urfave/cli/v2
	go get -u github.com/sirupsen/logrus
	go get -u github.com/alexandrevicenzi/unchained
	go get -u github.com/oklog/ulid
	go get github.com/google/uuid
	go get -u github.com/go-resty/resty/v2


run:
	go run cmd/spooky.go

build:
	go build -v -o bin/spooky cmd/spooky.go

clean:
	@rm -rf bin/spooky

push:
	git push origin master

install:
	go install cmd/spooky.go
