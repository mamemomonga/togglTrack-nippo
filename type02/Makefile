APPNAME=toggl-nippo-type02

bin/$(APPNAME): vendor
	go build -o $@ ./

vendor:
	go mod vendor

release-build:
	go build -o release/$(APPNAME)-$(GOOS)-$(GOARCH)

release:
	GOOS=darwin GOARCH=amd64 $(MAKE) release-build
	GOOS=darwin GOARCH=arm64 $(MAKE) release-build

clean:
	rm -rf bin release

.PHONY: release release-build clean
