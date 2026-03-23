# KrakenD API gateway and plugins

A demonstration that `xml` encoding is a lot slower in returning a response than `no-op` encoding.

## Build

```bash
docker build -t krakend-example .
```

## Run

```bash
docker run --rm --network host krakend-example
```

After KrakenD is up and running you can call `time curl http://localhost:8080/xml > /dev/null` and see that command is a lot slower than calling `time curl http://localhost:8082/no-op > /dev/null`. For me, the request with `xml` encoding is roughly 10x slower.
