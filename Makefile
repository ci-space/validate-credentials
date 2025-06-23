build:
	docker build -t ci-space/validate-github-credentials:local .

run:
	INPUT_GITHUB_TOKEN=${TOKEN} go run main.go

lint:
	golangci-lint run --fix
