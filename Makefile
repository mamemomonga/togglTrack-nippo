RELEASE_VERSION := v0.0.4

release: type01/release type02/release
	mkdir -p release
	mv type01/release/* release/
	rmdir type01/release
	mv type02/release/* release/
	rmdir type02/release

type01/release:
	cd type01 && make release

type02/release:
	cd type02 && make release

clean:
	cd type01 && make clean
	cd type02 && make clean

gh-release:
	gh release create $(RELEASE_VERSION) --generate-notes --latest
	gh release upload $(RELEASE_VERSION) release/toggl-nippo-type01-darwin-amd64
	gh release upload $(RELEASE_VERSION) release/toggl-nippo-type01-darwin-arm64
	gh release upload $(RELEASE_VERSION) release/toggl-nippo-type02-darwin-amd64
	gh release upload $(RELEASE_VERSION) release/toggl-nippo-type02-darwin-arm64
