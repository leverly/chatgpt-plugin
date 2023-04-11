build:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o plugin
	zip plugin.zip plugin logo.png openai.yaml .well-known/ai-plugin.json
	GOOS=windows GOARCH=amd64 go build -o plugin.ext
	zip wechat.zip plugin.exe logo.png openai.yaml .well-known/ai-plugin.json
	GOOS=darwin GOARCH=amd64 go build -o plugin
	zip wechat.zip plugin logo.png openai.yaml .well-known/ai-plugin.json
	GOOS=darwin GOARCH=arm64 go build -o plugin
	zip wechat.zip plugin logo.png openai.yaml .well-known/ai-plugin.json
