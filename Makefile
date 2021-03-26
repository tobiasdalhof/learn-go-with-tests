test:
	go test -v -cover ./...

bench:
	go test -bench=. -cover ./...

.PHONY: test bench
