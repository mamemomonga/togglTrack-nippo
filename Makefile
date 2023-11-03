

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
