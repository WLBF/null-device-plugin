build:
	go build ./cmd/null-device-plugin

image:
	docker build -t 10.27.44.1:5000/null-device-plugin .
	docker push 10.27.44.1:5000/null-device-plugin
