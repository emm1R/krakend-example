FROM krakend/builder-ee:2.11.0 AS builder

WORKDIR /app
COPY --chown=krakend:nogroup plugins /app

# build plugins
RUN cd mw-example && go build '-buildmode=plugin' -o ../mw-example.so middleware.go
RUN cd modifier-example && go build '-buildmode=plugin' -o ../modifier-example.so modifier.go
RUN cd client-example && go build '-buildmode=plugin' -o ../client-example.so client.go

FROM krakend/krakend-ee:2.11.0

WORKDIR /opt/krakend

COPY . .

RUN krakend check --debug 3 -t --config krakend.json

ENV USAGE_DISABLE=1

COPY LICENSE LICENSE

# Copy built plugins with other plugins
COPY --from=builder --chown=krakend:nogroup /app/*.so /opt/krakend/plugins/

CMD [ "run", "-c", "/opt/krakend/krakend.json" ]
