check_install:
	which swagger || GO111MODULE=off go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger: check_install
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
