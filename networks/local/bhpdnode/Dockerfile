#FROM alpine:3.7
#MAINTAINER Greg Szabo <greg@tendermint.com>
#
#RUN apk update && \
#    apk upgrade && \
#    apk --no-cache add curl jq file
#
#VOLUME [ /bhpd ]
#WORKDIR /bhpd
#EXPOSE 26656 26657
#ENTRYPOINT ["/usr/bin/wrapper.sh"]
#CMD ["start"]
#STOPSIGNAL SIGTERM
#
#COPY wrapper.sh /usr/bin/wrapper.sh
#
#RUN chmod 777 -R /usr/bin/wrapper.sh

FROM ubuntu:18.04

RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get -y install curl jq file

VOLUME [ /bhpd ]
WORKDIR /bhpd
EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start"]
STOPSIGNAL SIGTERM

COPY wrapper.sh /usr/bin/wrapper.sh

RUN chmod 777 -R /usr/bin/wrapper.sh

