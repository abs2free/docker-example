ARG STACK_VERSION

FROM docker.elastic.co/elasticsearch/elasticsearch:${STACK_VERSION}

MAINTAINER abs2free

# 在 FROM 之后重新声明 ARG
ARG STACK_VERSION
ENV STACK_VERSION=${STACK_VERSION}

RUN sh -c 'echo "y" | /usr/share/elasticsearch/bin/elasticsearch-plugin install http://get.infini.cloud/elasticsearch/analysis-ik/${STACK_VERSION}'
RUN sh '/usr/share/elasticsearch/bin/elasticsearch-plugin install mapper-size'
