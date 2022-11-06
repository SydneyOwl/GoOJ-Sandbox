FROM criyle/executorserver:latest AS executorserver 

FROM ubuntu:18.04

ARG DEBIAN_FRONTEND=noninteractive

ENV TZ=Asia/Shanghai

RUN buildDeps='software-properties-common libtool wget unzip' && \
    apt-get update && apt-get install -y python python3.7 gcc g++ mono-devel $buildDeps curl bash && \
    add-apt-repository ppa:openjdk-r/ppa && add-apt-repository ppa:longsleep/golang-backports && apt-get update && apt-get install -y golang-go openjdk-8-jdk && \
	add-apt-repository ppa:pypy/ppa && apt-get update && apt install -y pypy pypy3 && \
	add-apt-repository ppa:ondrej/php && apt-get update && apt-get install -y php7.3-cli && \
	cd /tmp && wget -O jsv8.zip  https://storage.googleapis.com/chromium-v8/official/canary/v8-linux64-dbg-8.4.109.zip && \
	unzip -d /usr/bin/jsv8 jsv8.zip && rm -rf /tmp/jsv8.zip && \
	curl -fsSL https://deb.nodesource.com/setup_14.x | bash && \
	apt-get install -y nodejs && \
    apt-get purge -y --auto-remove $buildDeps && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /opt

COPY --from=executorserver /opt/executorserver /opt/mount.yaml /opt/

EXPOSE 5050/tcp 5051/tcp

ENTRYPOINT ["./executorserver"]