ifneq (,$(wildcard ./.env))
    include .env
    export
endif


env:
	export ES_URL="https://localhost:9200"
	export API_KEY="eV9pTmZwSUJMa2NYSUM2UHBoNGU6NWFGRjIteHdRNi1icHpnRXQ1Y3Vrdw=="
cl:
	curl -k  "$(ES_URL)/poem"  -H "Authorization: ApiKey "$(API_KEY)""  -H "Content-Type: application/json"



ik-build:
	docker build --build-arg STACK_VERSION=${STACK_VERSION}  -t elasticsearch-ik:$(STACK_VERSION) .

ik-compose:
	docker compose --file docker-compose-ik.yml up -d
