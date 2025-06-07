# syntax=docker/dockerfile:latest
FROM --platform=$BUILDPLATFORM golang:latest AS build

# Build aura-core
WORKDIR /src
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -o aura -trimpath -ldflags "-s -w -buildid=" ./main

# Download geodat into a staging directory
ADD https://raw.githubusercontent.com/Loyalsoldier/v2ray-rules-dat/release/geoip.dat /tmp/geodat/geoip.dat
ADD https://raw.githubusercontent.com/Loyalsoldier/v2ray-rules-dat/release/geosite.dat /tmp/geodat/geosite.dat

RUN mkdir -p /tmp/empty

# Create config files with empty JSON content
RUN mkdir -p /tmp/usr/local/etc/aura
RUN cat <<EOF >/tmp/usr/local/etc/aura/00_log.json
{
  "log": {
    "error": "/var/log/aura/error.log",
    "loglevel": "warning",
    "access": "none",
    "dnsLog": false
  }
}
EOF
RUN echo '{}' >/tmp/usr/local/etc/aura/01_api.json
RUN echo '{}' >/tmp/usr/local/etc/aura/02_dns.json
RUN echo '{}' >/tmp/usr/local/etc/aura/03_routing.json
RUN echo '{}' >/tmp/usr/local/etc/aura/04_policy.json
RUN echo '{}' >/tmp/usr/local/etc/aura/05_inbounds.json
RUN echo '{}' >/tmp/usr/local/etc/aura/06_outbounds.json
RUN echo '{}' >/tmp/usr/local/etc/aura/07_transport.json
RUN echo '{}' >/tmp/usr/local/etc/aura/08_stats.json
RUN echo '{}' >/tmp/usr/local/etc/aura/09_reverse.json

# Create log files
RUN mkdir -p /tmp/var/log/aura && touch \
  /tmp/var/log/aura/access.log \
  /tmp/var/log/aura/error.log

# Build finally image
FROM gcr.io/distroless/static:nonroot

COPY --from=build --chown=0:0 --chmod=755 /src/aura /usr/local/bin/aura
COPY --from=build --chown=0:0 --chmod=644 /tmp/geodat/*.dat /usr/local/share/aura/
COPY --from=build --chown=0:0 --chmod=755 /tmp/empty /usr/local/etc/aura
COPY --from=build --chown=0:0 --chmod=644 /tmp/usr/local/etc/aura/*.json /usr/local/etc/aura/
COPY --from=build --chown=0:0 --chmod=755 /tmp/empty /var/log/aura
COPY --from=build --chown=65532:65532 --chmod=600 /tmp/var/log/aura/*.log /var/log/aura/

VOLUME /usr/local/etc/aura
VOLUME /var/log/aura

ARG TZ=Etc/UTC
ENV TZ=$TZ

ENTRYPOINT [ "/usr/local/bin/aura" ]
CMD [ "-confdir", "/usr/local/etc/aura/" ]
