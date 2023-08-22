default: run

run:
	@docker-compose up --build

stop:
	@docker-compose kill

rm:
	@docker-compose rm -svf
