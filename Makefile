# vi:syntax=make

BUILD_DIR ?= build

.PHONY: *

.envrc:
	[[ -f .envrc ]] || cp .envrc.example .envrc
env: .envrc

test:
	@./bin/test

build:
	@./bin/build

all:
	@./bin/build --all
