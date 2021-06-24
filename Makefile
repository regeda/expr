.PHONY: gen-fbs gen-lexer test bench bench-report escape-analysis

default: test

gen-fbs:
	@rm -f bytecode/*.go
	@flatc -g -o . bytecode/proto.fbs

gen-lexer:
	@ragel -Z -G2 lexer/lexer.go.rl -o lexer/lexer.go
	@goimports -w lexer/lexer.go

test: gen-fbs gen-lexer
	@go test -v -cover ./...

bench: test
	@go test -v -bench=. -run=- -benchmem -cpuprofile cpu.out

bench-report: bench
	@go tool pprof -web cpu.out

escape-analysis: ; @go test -gcflags="-m" ./...
