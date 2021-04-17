test:
	go test ./...

bench:
	go test -bench=. -cover ./...

godoc:
	godoc -http=:6060 

.PHONY: test bench
