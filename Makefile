.PHONY: all
all: build

.PHONY: build
build:
	go build -o pciids

.PHONY: clean
clean:
	rm -f pciids coverage.out
	rm -rf dist/ site/

.PHONY: docs
docs:
	mkdocs build

.PHONY: docs-api
docs-api:
	echo "View docs at: http://localhost:6060/pkg/pciids/"
	godoc -http=localhost:6060

.PHONY: lint
lint:
	golangci-lint run

.PHONY: release
release: clean
	goreleaser

.PHONY: release-snapshot
release-snapshot: clean
	goreleaser --rm-dist --skip-publish --snapshot

.PHONY: test
test:
	go test -cover -coverprofile=coverage.out  ./...

.PHONY: test-coverage
test-coverage: test
	go tool cover -html=coverage.out
