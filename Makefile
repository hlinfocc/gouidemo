export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on

LDFLAGS := -s -w

all: fmt build

build: gouidemo


fmt:
	go fmt ./...

fmt-more:
	gofumpt -l -w .

vet:
	go vet ./...


gouidemo:
	env CGO_ENABLED=1 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/gouidemo ./
	

gouidemoexe:
	fyne-cross windows -app-id net.hlinfo.gouidemo -image fyneio/fyne-cross-images:1.1.0-windows -icon ./assets/static/logo.png

test: gotest

gotest:
	go test -v --cover ./...

	
clean:
	rm -f ./bin/gouidemo
	rm -f ./bin/gouidemoexe
