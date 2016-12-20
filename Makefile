APP=subutai
CC=go
VERSION=$(shell git describe --abbrev=0 --tags | awk -F'.' '{print $$1"."$$2"."$$3+1}')
ifeq (${GIT_BRANCH}, )
	GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD | grep -iv head)
endif
ifneq (${GIT_BRANCH}, )
	VERSION:=${VERSION}-SNAPSHOT
endif
COMMIT=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags "-r /apps/subutai/current/lib -w -s -X main.version=${VERSION} -X main.commit=${COMMIT}"

all:
	$(CC) get
	$(CC) build ${LDFLAGS} -o $(APP)
snapcraft:
	GOPATH=$(shell pwd) GOBIN=$(shell pwd)/../go/bin $(CC) get
	GOPATH=$(shell pwd) GOBIN=$(shell pwd)/../go/bin $(CC) build ${LDFLAGS} -o $(APP)
install: 
	@mkdir -p $(DESTDIR)/bin
	@cp $(APP) $(DESTDIR)/bin
