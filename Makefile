OPENAPI_GENERATOR_VERSION=v7.3.0
OPENAPI_GENERATOR=go
.PHONY: all clean imports fetch-openapi

all: clean pkg/vimeoapi
fetch-openapi:
	curl -o api/openapi.yaml -X GET \
      'https://api.vimeo.com/?openapi=1' \
      -H 'Accept: application/json;version=3.4' \
      -H 'Authorization: bearer ${VIMEO_BEARER_TOKEN}'
clean:
	rm -rf pkg/vimeoapi
imports:
	goimports -w pkg/*
pkg/vimeoapi:
	mkdir -p pkg/vimeoapi
	cp api/.openapi-generator-ignore pkg/vimeoapi/.openapi-generator-ignore
	docker run --rm -v "$(PWD):/local" openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_VERSION) generate \
		-g $(OPENAPI_GENERATOR) \
		-c /local/api/config.json \
		-i /local/api/openapi.yaml \
		-o /local/pkg/vimeoapi