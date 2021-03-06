.DEFAULT: install

install:
	go install ./cmd/protoc-gen-go-patch

tools: internal/tools/*.go
	go generate --tags tools ./internal/tools

vet:
	go vet ./...

test: test-go test-cgo-disabled

test-go:
	go test -i -mod=readonly -race ./...
	go test -mod=readonly -v -race ./...

test-cgo-disabled:
	CGO_ENABLED=0 go test -i -mod=readonly ./...
	CGO_ENABLED=0 go test -mod=readonly -v ./...

go_module = $(shell go list -m)
proto_files = $(sort $(shell find . -name '*.proto'))
proto_includes = \
	-I . \
	-I $(shell go list -m -f {{.Dir}} google.golang.org/protobuf) \
	-I $(shell go list -m -f {{.Dir}} github.com/envoyproxy/protoc-gen-validate) \

protos: $(proto_files)

.PHONY: $(proto_files)
$(proto_files): tools Makefile
	# protoc-gen-go
	protoc \
		$(proto_includes) \
		--go-patch_out=plugin=go-grpc,paths=import,module=$(go_module):. \
		$@

	# protoc-gen-go-grpc
	protoc \
		$(proto_includes) \
		--go-patch_out=plugin=go,paths=import,module=$(go_module):. \
		$@

	# protoc-gen-validate
	if grep -q validate/validate\.proto $@; then protoc \
		$(proto_includes) \
		--go-patch_out=plugin=validate,paths=source_relative,lang=go:. \
		$@ ; \
	fi
