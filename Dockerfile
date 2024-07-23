FROM krakend/builder-ee:2.6.4 as builder

WORKDIR /app
COPY --chown=krakend:nogroup plugins /app

# build plugins
RUN ./build.sh


FROM krakend/krakend-ee:2.6.4

WORKDIR /opt/krakend

COPY . .

RUN krakend check --debug 3 -t --config krakend.json

# Disable telemetry: https://www.krakend.io/docs/configuration/environment-vars/#usage-reporting-env-var
ENV USAGE_DISABLE=1

COPY LICENSE LICENSE

# Copy built plugins with other plugins
COPY --from=builder --chown=krakend:nogroup /app/*.so /opt/krakend/plugins/

CMD [ "run", "-c", "/opt/krakend/krakend.json" ]
