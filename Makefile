go-clean:
	rm -fr ./pkg

go-build:
	go build  -o ./pkg/pag ./cli/pag

build:
	$(MAKE) go-clean
	$(MAKE) go-build

install:
	$(MAKE) build
	cp -pr ./pkg/pag /usr/local/bin

uninstall:
	rm -fr /usr/local/bin/pag
