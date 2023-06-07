start:
	docker build -t api .
	docker run -it --rm -v ./:/app --env-file env.dev -p 127.0.0.1:8888:8888/tcp api
