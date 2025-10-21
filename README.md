# KrakenD API gateway and plugins

A demonstration that the modifier plugin is executed before the middleware plugin. Requires an Enterprise LICENSE.

## Build

```bash
docker build -t krakend .
```

## Run

```bash
docker run --rm --network host krakend
```

After KrakenD is up and running you can call `curl http://localhost:8082/xml-example` and see that the response modifier is not called. On the other hand, `curl http://localhost:8082/no-op-example` reaches the response plugin but does not return the body.
