# Vimeo API Client

Library with vimeo API client.

# Install

Before make sure you have configured `GOPRIVATE`, you can do it like this

`go env -w GOPRIVATE='github.com/Odrabiamy-pl/*'`
or
`export GOPRIVATE=github.com/Odrabiamy-pl/*`

and
`git config --global url."git@github.com:".insteadOf "https://github.com/"`

Now you are ready to go

`go get github.com/Odrabiamy-pl/go-vimeo`

# Regenerate

If you want to regenerate API client, first fetch new version of OpenAPI, to do so you'll need Vimeo bearer token (sic!).

```
VIMEO_BEARER_TOKEN=... make fetch-openapi
```

then just run

```
make all
```

...not so fast, now you need to fix broken API documentation and if you are ready - run `make all`.

But it is better to add missing things manually rather than downloading new openapi schema and fixing everything again.

![cry.jpg](cry.jpg)

# Examples

Go to `./test`