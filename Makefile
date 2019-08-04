default: docker-build

run: docker

docker:
	@docker-compose rm -fv
	@docker-compose up

docker-build:
	@docker-compose rm -fv
	@docker-compose up --build
