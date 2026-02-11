# KrakenD API gateway and plugins

A demonstration that a large number in a flexible configuration gets converted into exponential format in the json output file.

## Build

```bash
docker build -t krakend-example .
```

## Run

```bash
docker run --rm --name krakend-example krakend-example
```

In the container, the file `/tmp/krakend.json` will have the following content:

```json
{
  "version": 3,
  "debug_endpoint": true,
  "echo_endpoint": true,
  "host": ["http://localhost:8080/"],
  "endpoints": [
    {
      "endpoint": "/example",
      "method": "GET",
      "extra_config": {
        "proxy": {
          "max_payload": 5.36870912e+09
        }
      },
      "backend": [
        {
          "url_pattern": "/__debug/"
        }
      ]
    }
  ]
}
```
