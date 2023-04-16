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
	@docker compose run --rm gopacket-app go fmt
find-devices:
	@docker compose run --rm gopacket-app go run finds/find_devices.go
filter:
	@docker compose run --rm gopacket-app go run filters/setting_filters.go
capture:
	@docker compose run --rm gopacket-app go run livcap/live_capture.go
write:
	@docker compose run --rm gopacket-app go run writepcap/write_pcap_file.go