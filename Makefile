
test.all:
	docker run --rm -d --name test_db -p 6380:6379  redis:7.0.11-alpine
	go test -v ./...
	docker stop test_db
	