
# use ca
docker cp es01:/usr/share/elasticsearch/config/certs/http_ca.crt .
curl --cacert ca.crt -u "elastic:123456"  -X GET "https://localhost:9200/_cat/nodes?v&pretty"


# use api_key
export ES_URL="https://localhost:9200"
export API_KEY="eV9pTmZwSUJMa2NYSUM2UHBoNGU6NWFGRjIteHdRNi1icHpnRXQ1Y3Vrdw=="

curl -k  "${ES_URL}/poem"  -H "Authorization: ApiKey "${API_KEY}""  -H "Content-Type: application/json"
