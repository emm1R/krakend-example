FROM krakend/krakend-ee:2.13.0

WORKDIR /opt/krakend

COPY . .

RUN krakend check --debug 3 -t --config krakend.json

ENV USAGE_DISABLE=1

COPY LICENSE LICENSE

CMD [ "run", "-c", "/opt/krakend/krakend.json" ]
