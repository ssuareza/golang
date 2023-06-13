dev:
	docker compose exec filesplit sh

start:
	docker build -t filesplit .
	docker run -it --rm -v ./:/app --env-file env.dev filesplit sh
