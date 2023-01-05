.PHONY: test
test:
	go test ./...

.PHONY: test-and-update-snaps
test-and-update-snaps:
	UPDATE_SNAPS=true go test ./...
