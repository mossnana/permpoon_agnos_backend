compose:
	docker-compose up -d
test:
	go test ./... -v
coverage:
	go test ./... -v -coverprofile=./testing/coverage.out -coverpkg=./... \
	&& go tool cover -html=./testing/coverage.out -o=./testing/coverage.html
