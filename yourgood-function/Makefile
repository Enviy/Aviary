build:
	GOOS=windows GOARCH=amd64 go build -o main.exe cmd/main.go
	zip -r function.zip .

deploy:
	az functionapp deployment source config-zip -g yourgood-rg -n yourgood-function --src function.zip

