.PHONY: build
build:
	go build -o ./bin/micorkart

.PHONY: run
run: build
	./bin/micorkart	

.PHONY: clean
clean:
	rm -f ./bin/micorkart

.PHONY: help
help:
	@echo "*** make commands ***"
	@echo "build: create microkart app using go build"
	@echo "run: start microkart app"
	@echo "clean: remove microkart app"
