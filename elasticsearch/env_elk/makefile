crt:
	 docker cp env_elk-es01-1:/usr/share/elasticsearch/config/certs/http_ca.crt .

cel:
	curl --cacert es01.crt -u elastic:123456 ${url}

