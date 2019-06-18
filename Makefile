# vi:syntax=make

BUILD_DIR ?= build

.PHONY: *

.envrc:
	[[ -f .envrc ]] || cp .envrc.example .envrc
env: .envrc

test: build
	@echo "Running ginkgo suite"
	@./bin/test
	@echo "Running integration suite"
	@./vendor/github.com/sstephenson/bats/bin/bats test/*.bats

build:
	@./bin/build

all:
	@./bin/build --all
