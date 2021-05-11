.PHONY: gen-proto gen-fbs gen-lexer test bench bench-report

default: test

gen-proto:
	@go install github.com/gogo/protobuf/protoc-gen-gogofaster
	@protoc --gogofaster_out=. ./ast/ast.proto

gen-fbs:
	@rm -f bytecode/*.go
	@flatc -g -o . bytecode/proto.fbs

gen-lexer: gen-proto
	@ragel -Z -G2 lexer/lexer.go.rl -o lexer/lexer.go
	@goimports -w lexer/lexer.go

test: gen-fbs gen-lexer
	@go test -v -cover ./...

bench: test
	@go test -v -bench=. -run=- -benchmem -cpuprofile cpu.out

bench-report: bench
	@go tool pprof -web cpu.out
