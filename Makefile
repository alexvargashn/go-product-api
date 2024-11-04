check_install: 
	go get -u github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models