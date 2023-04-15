USER:=$(shell id -u)
GROUP:=$(shell id -g)

build:
	@docker compose build
up:
	@docker compose up -d
exec-user:
	@docker compose exec -u $(USER):$(GROUP) -it gopacket-app bash
exec:
	@docker compose exec -it gopacket-app bash
down:
	@docker compose down

fmt:
	@docker compose run --rm go-app go fmt