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

kibana: prekb
	docker run -p 127.0.0.1:5601:5601 -d --name kibana --network elastic-net \
  -e ELASTICSEARCH_URL=http://elasticsearch:9200 \
  -e ELASTICSEARCH_HOSTS=http://elasticsearch:9200 \
  -e ELASTICSEARCH_USERNAME=kibana_system \
  -e ELASTICSEARCH_PASSWORD=${KIBANA_PASSWORD} \
  -e "xpack.security.enabled=false" \
  -e "xpack.license.self_generated.type=trial" \
  docker.elastic.co/kibana/kibana:${STACK_VERSION}


metricbeat:
	docker run -d \
	  --name=metricbeat \
	--network elastic-net \
	  --user=root \
	  --volume="./metricbeat.docker.yml:/usr/share/metricbeat/metricbeat.yml:ro" \
	  --volume="/var/run/docker.sock:/var/run/docker.sock:ro" \
	  --volume="/sys/fs/cgroup:/hostfs/sys/fs/cgroup:ro" \
	  --volume="/proc:/hostfs/proc:ro" \
	  --volume="/:/hostfs:ro" \
	  --volume="./elasticsearch-xpack.yml:/usr/share/metricbeat/modules.d/elasticsearch-xpack.yml" \
     docker.elastic.co/beats/metricbeat:8.15.2


filebeat:
	 docker run -d \
  --name=filebeat \
  --network elastic-net \
  --user=root \
  --volume="./filebeat.docker.yml:/usr/share/filebeat/filebeat.yml:ro" \
  --volume="/var/lib/docker/containers:/var/lib/docker/containers:ro" \
  --volume="/var/run/docker.sock:/var/run/docker.sock:ro" \
  --volume="registry:/usr/share/filebeat/data:rw" \
  docker.elastic.co/beats/filebeat:8.15.2 filebeat -e --strict.perms=false


logst::
	docker run -d \
	--name logstash \
	-m 1000M --network elastic-net \
	-p 5044:5044/udp -p 9600:9600 --privileged=true \
	-e ES_JAVA_OPTS="-Duser.timezone=Asia/Shanghai" \
	-v ./logstash.conf:/usr/share/logstash/pipeline/logstash.conf \
	-v ./logstash.yml:/usr/share/logstash/config/logstash.yml \
	docker.elastic.co/logstash/logstash:8.15.2
