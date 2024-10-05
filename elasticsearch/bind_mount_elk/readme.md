
docker cp es01:/usr/share/elasticsearch/config/certs/http_ca.crt .



curl --cacert ca.crt -u "elastic:123456"  -X GET "https://localhost:9200/_cat/nodes?v&pretty"
