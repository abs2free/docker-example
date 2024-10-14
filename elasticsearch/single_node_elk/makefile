ifneq (,$(wildcard ./.env))
    include .env
    export
endif


network:
	docker network create elastic-net

es:
	docker run -p 127.0.0.1:9200:9200 -d --name elasticsearch --network elastic-net \
  -e ELASTIC_PASSWORD=${ELASTIC_PASSWORD} \
  -e "discovery.type=single-node" \
  -e "xpack.security.http.ssl.enabled=false" \
  -e "xpack.license.self_generated.type=trial" \
  docker.elastic.co/elasticsearch/elasticsearch:${STACK_VERSION}

ik:
	docker exec -it elasticsearch bash -c 'echo "y" | /usr/share/elasticsearch/bin/elasticsearch-plugin install http://get.infini.cloud/elasticsearch/analysis-ik/${STACK_VERSION}'
	#docker exec -it elasticsearch bash /usr/share/elasticsearch/bin/elasticsearch-plugin install http://get.infini.cloud/elasticsearch/analysis-ik/${STACK_VERSION}

prekb:
	curl -u elastic:${ELASTIC_PASSWORD} \
  -X POST \
  http://localhost:9200/_security/user/kibana_system/_password \
  -d '{"password":"'"${KIBANA_PASSWORD}"'"}' \
  -H 'Content-Type: application/json'

kibana:
	docker run -p 127.0.0.1:5601:5601 -d --name kibana --network elastic-net \
  -e ELASTICSEARCH_URL=http://elasticsearch:9200 \
  -e ELASTICSEARCH_HOSTS=http://elasticsearch:9200 \
  -e ELASTICSEARCH_USERNAME=kibana_system \
  -e ELASTICSEARCH_PASSWORD=${KIBANA_PASSWORD} \
  -e "xpack.security.enabled=false" \
  -e "xpack.license.self_generated.type=trial" \
  docker.elastic.co/kibana/kibana:${STACK_VERSION}




