.PHONY: run
run: build image
	./scripts/run.sh

.PHONY: build
build:
	./scripts/build_proto.sh
	./scripts/build.sh

.PHONY: check
check:
	./scripts/check.sh

.PHONY: image
image:
	./scripts/image.sh

.PHONY: test
test:
	./scripts/test.sh
