usage="Run linters"

go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.18.0

golangci-lint run -v
