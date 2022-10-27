build:
	docker buildx build --platform=linux/amd64 --push -t docker.io/imroc/topology-server:latest .