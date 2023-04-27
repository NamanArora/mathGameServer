server:
	docker-compose up -d --build --remove-orphans

logs:
	docker logs -f mathgame_server

clean:
	docker compose down
