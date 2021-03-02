.PHONY: gen-proto gen-fbs gen-lexer test bench bench-report

default: test

gen-proto:
	@go install github.com/gogo/protobuf/protoc-gen-gogofaster
	@protoc --gogofaster_out=. ./internal/ast/ast.proto

gen-fbs:
	@rm -f internal/bytecode/*.go
	@flatc -g -o internal/ internal/bytecode/proto.fbs

gen-lexer: gen-proto
	@ragel -Z -G2 internal/lexer/lexer.go.rl -o internal/lexer/lexer.go

test: gen-fbs gen-lexer
	@go test -v -cover ./...

bench: test
	@go test -v -bench=. -run=- -benchmem -cpuprofile cpu.out

bench-report: bench
	@go tool pprof -web cpu.out
