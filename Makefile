go-clean:
	rm -fr ./pkg

go-generate:
	go generate ./cli/gpa/main.go

go-build:
	go build  -o ./pkg/gpa ./cli/gpa

build:
	$(MAKE) go-clean
	$(MAKE) go-generate
	$(MAKE) go-build

install:
	$(MAKE) build
	cp -pr ./pkg/gpa /usr/local/bin

uninstall:
	rm -fr /usr/local/bin/gpa
