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

After KrakenD is up and running you can call `curl http://localhost:8082/example` and see that the response contains the modifier error, not the expected middlware error.
