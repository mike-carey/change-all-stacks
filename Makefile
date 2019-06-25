# vi:syntax=make

BUILD_DIR ?= build

BATS ?= ./.lib/bats/bin/bats

.PHONY: *

.envrc:
	[[ -f .envrc ]] || cp .envrc.example .envrc
env: .envrc

test: build
	@echo "Running ginkgo suite"
	@./bin/test
	@echo "Running integration suite"
	@$(BATS) test/*.bats

build:
	@./bin/build

all:
	@./bin/build --all
