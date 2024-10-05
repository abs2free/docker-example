crt:
	 docker cp volume_elk-es01-1:/usr/share/elasticsearch/config/certs/es01/es01.crt .

cel:
	curl --cacert es01.crt -u elastic:123456 ${url}
