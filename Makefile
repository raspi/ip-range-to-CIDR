LAST_TAG := $(shell git describe --abbrev=0 --always --tags)
BUILD := $(shell git rev-parse $(LAST_TAG))

BINARY := ip-range-to-cidr
UNIXBINARY := $(BINARY)
WINBINARY := $(UNIXBINARY).exe
BUILDDIR := build

LINUXRELEASE := $(BINARY)-$(LAST_TAG)-linux-x64.tar.gz
WINRELEASE := $(BINARY)-$(LAST_TAG)-windows-x64.zip

LDFLAGS := -ldflags "-s -w -X=main.VERSION=$(LAST_TAG) -X=main.BUILD=$(BUILD)"

bin:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -v -o $(BUILDDIR)/$(UNIXBINARY)
	upx -v -9 $(BUILDDIR)/$(UNIXBINARY)

bin-windows:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -v -o $(BUILDDIR)/$(WINBINARY)
	upx -v -9 $(BUILDDIR)/$(WINBINARY)

release:
	cd $(BUILDDIR); tar cvzf $(LINUXRELEASE) $(UNIXBINARY)

release-windows:
	cd $(BUILDDIR); zip -v -9 $(WINRELEASE) $(WINBINARY)

.PHONY: all clean test