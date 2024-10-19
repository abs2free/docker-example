ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ca:
	 docker cp dockerfile-elk-es01-1:/usr/share/elasticsearch/config/certs/ca/ca.crt http_ca.crt

footprint: ca
	openssl x509 -fingerprint -sha256 -noout -in ./http_ca.crt | awk --field-separator="=" '{print $2}' | sed 's/://g'

test:
	curl --cacert http_ca.crt -u "elastic:123456"  -X GET "https://localhost:9200/_cat/nodes?v&pretty"


env:
	export ES_URL="https://localhost:9200"
	export API_KEY="eV9pTmZwSUJMa2NYSUM2UHBoNGU6NWFGRjIteHdRNi1icHpnRXQ1Y3Vrdw=="
cl:
	curl -k  "$(ES_URL)/poem"  -H "Authorization: ApiKey "$(API_KEY)""  -H "Content-Type: application/json"



build:
	docker build --build-arg STACK_VERSION=${STACK_VERSION}  -t elasticsearch-ik:$(STACK_VERSION) .

compose:
	docker compose --file docker-compose.yml up -d

upbasic:
	docker compose --profile basic  up -d

upmonitor:
	docker compose --profile monitoring  up -d

downbasic:
	docker compose --profile basic  down

downmomitor:
	docker compose --profile monitoring  down
