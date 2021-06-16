.PHONY: gen-proto gen-fbs gen-tokenz test bench bench-report escape-analysis

default: test

gen-proto:
	@go install github.com/gogo/protobuf/protoc-gen-gogofaster
	@protoc --gogofaster_out=. ./ast/ast.proto

gen-fbs:
	@rm -f bytecode/*.go
	@flatc -g -o . bytecode/proto.fbs

gen-tokenz: gen-proto
	@ragel -Z -G2 tokenz/tokenz.go.rl -o tokenz/tokenz.go
	@goimports -w tokenz/tokenz.go

test: gen-fbs gen-tokenz
	@go test -v -cover ./...

bench: test
	@go test -v -bench=. -run=- -benchmem -cpuprofile cpu.out

bench-report: bench
	@go tool pprof -web cpu.out

escape-analysis: ; @go test -gcflags="-m" ./...
