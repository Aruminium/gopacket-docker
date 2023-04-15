FROM golang:1.20 as dev

ENV USER=gopacket-user
ENV GROUP=gopacket-group

RUN apt-get update && apt-get install -y --no-install-recommends \
  libpcap-dev=1.10.0-2  \
  && rm -rf /var/lib/apt/lists/*

# RUN groupadd -r ${GROUP} && useradd --no-log-init -r -g ${GROUP} ${USER}
# USER ${USER}
