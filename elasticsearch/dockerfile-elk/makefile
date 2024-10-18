ifneq (,$(wildcard ./.env))
    include .env
    export
endif


env:
	export ES_URL="https://localhost:9200"
	export API_KEY="eV9pTmZwSUJMa2NYSUM2UHBoNGU6NWFGRjIteHdRNi1icHpnRXQ1Y3Vrdw=="
cl:
	curl -k  "$(ES_URL)/poem"  -H "Authorization: ApiKey "$(API_KEY)""  -H "Content-Type: application/json"



build:
	docker build --build-arg STACK_VERSION=${STACK_VERSION}  -t elasticsearch-ik:$(STACK_VERSION) .

compose:
	docker compose --file docker-compose.yml up -d

basicup:
	docker compose --profile basic  up -d

moniup:
	docker compose --profile monitoring  up -d

basicdown:
	docker compose --profile basic  down

monidown:
	docker compose --profile monitoring  down
