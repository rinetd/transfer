PROJECT = transfer
VERSION ="$(echo $TRAVIS_TAG)"
ARCNAME = $(PROJECT)-$(VERSION)-$(GOOS)-$(GOARCH)

all: release-all

release:
	env CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(FLAGS) -ldflags='-s -w -X github.com/rinetd/transfer/version.Version=$(VERSION)' -o build/$(GOOS)-$(GOARCH)/$(PROJECT)$(EXT)
	tar czf build/$(ARCNAME).tar.gz -C build/$(GOOS)-$(GOARCH)/ $(PROJECT)$(EXT)

.PHONY: release-all
release-all:
	-@$(MAKE) release GOOS=darwin GOARCH=amd64
	-@$(MAKE) release GOOS=linux GOARCH=386
	-@$(MAKE) release GOOS=linux GOARCH=amd64
	-@$(MAKE) release GOOS=windows GOARCH=386 EXT=.exe
	-@$(MAKE) release GOOS=windows GOARCH=amd64 EXT=.exe
platform:
	# @$(MAKE) releaseGOOS=js GOARCH=wasm
	# @$(MAKE) release GOOS=windows GOARCH=386 FLAGS='-ldflags="-H=windowsgui"' EXE=.exe
	# @$(MAKE) release GOOS=windows GOARCH=amd64 FLAGS='-ldflags="-H=windowsgui"' EXE=.exe
	# @$(MAKE) release GOOS=linux GOARCH=arm
	# @$(MAKE) release GOOS=linux GOARCH=arm64
	# @$(MAKE) release GOOS=linux GOARCH=mips
	# @$(MAKE) release GOOS=linux GOARCH=mips64
	# @$(MAKE) release GOOS=linux GOARCH=mips64le
	# @$(MAKE) release GOOS=linux GOARCH=mipsle
	# @$(MAKE) release GOOS=linux GOARCH=ppc64
	# @$(MAKE) release GOOS=linux GOARCH=ppc64le
	# @$(MAKE) release GOOS=linux GOARCH=s390x
	# @$(MAKE) release GOOS=android GOARCH=386
	# @$(MAKE) release GOOS=android GOARCH=amd64
	# @$(MAKE) release GOOS=android GOARCH=arm
	# @$(MAKE) release GOOS=android GOARCH=arm64
	# @$(MAKE) release GOOS=darwin GOARCH=386
	# @$(MAKE) release GOOS=darwin GOARCH=arm
	# @$(MAKE) release GOOS=darwin GOARCH=arm64
	# @$(MAKE) release GOOS=dragonfly GOARCH=amd64
	# @$(MAKE) release GOOS=freebsd GOARCH=386
	# @$(MAKE) release GOOS=freebsd GOARCH=amd64
	# @$(MAKE) release GOOS=freebsd GOARCH=arm
	# @$(MAKE) release GOOS=nacl GOARCH=386
	# @$(MAKE) release GOOS=nacl GOARCH=amd64p32
	# @$(MAKE) release GOOS=nacl GOARCH=arm
	# @$(MAKE) release GOOS=netbsd GOARCH=386
	# @$(MAKE) release GOOS=netbsd GOARCH=amd64
	# @$(MAKE) release GOOS=netbsd GOARCH=arm
	# @$(MAKE) release GOOS=openbsd GOARCH=386
	# @$(MAKE) release GOOS=openbsd GOARCH=amd64
	# @$(MAKE) release GOOS=openbsd GOARCH=arm
	# @$(MAKE) release GOOS=plan9 GOARCH=386
	# @$(MAKE) release GOOS=plan9 GOARCH=amd64
	# @$(MAKE) release GOOS=plan9 GOARCH=arm
	# @$(MAKE) release GOOS=solaris GOARCH=amd64


.PHONY: build
build:
	go get ./...

.PHONY: test
test:
	go get -t ./...
	go test -v ...

build-image:
	docker build -t rientd/transfer .

clean:
	rm -rf build dist
