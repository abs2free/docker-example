ifneq (,$(wildcard ./.env))
    include .env
    export
endif

crt:
	 docker cp config-elk-es01-1:/usr/share/elasticsearch/config/certs/es01/es01.crt .

cel:
	curl --cacert es01.crt -u elastic:123456 ${url}


env:
	export ES_URL="https://localhost:9200"
	export API_KEY="eV9pTmZwSUJMa2NYSUM2UHBoNGU6NWFGRjIteHdRNi1icHpnRXQ1Y3Vrdw=="
cl:
	curl -k  "$(ES_URL)/poem"  -H "Authorization: ApiKey "$(API_KEY)""  -H "Content-Type: application/json"

