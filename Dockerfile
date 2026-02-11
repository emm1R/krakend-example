FROM krakend/krakend-ee:2.12.3

WORKDIR /opt/krakend

ENV FC_ENABLE=1
ENV FC_OUT=/tmp/krakend.json
ENV FC_SETTINGS=/opt/krakend/config/settings

COPY krakend.tmpl /opt/krakend/config/
COPY settings $FC_SETTINGS

ENV USAGE_DISABLE=1

COPY LICENSE LICENSE

CMD [ "run", "-c", "/opt/krakend/config/krakend.tmpl" ]
