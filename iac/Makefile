all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## init: project initialization - install tools and register git hook
.PHONY: init
init:
	asdf install
	pre-commit install

## checks: run all pre-commit checks
.PHONY: checks
checks:
	pre-commit run --all-files
