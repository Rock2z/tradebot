## build: Build source code for host platform.
.PHONY: build
gen:
	@protoc --proto_path=./proto --go_out=./ tb.proto
