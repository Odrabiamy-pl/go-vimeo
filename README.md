# Vimeo API Client

Library with vimeo API client.

# Regenerate

If you want to regenerate API client, first fetch new version of OpenAPI, to do so you'll need Vimeo bearer token (sic!).

```
VIMEO_BEARER_TOKEN=... make fetch-openapi
```

then just run

```
make all
```

...not so fast, now you need to fix broken API documentation and if you are ready run `make all`.

# Examples

Go to ./test