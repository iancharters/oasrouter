# OASRouter Demo

## Goals

1. ~~Generate models from spec.~~
2. ~~Validate OpenAPI spec on service start~~
   1. ~~Parsable~~
   2. ~~Valid example data~~
3. Validate request
   1. On error
      1. Build a slice containing validation errors
      2. Provide a hook to allow the consumer to modify the errors into their own shape (Google BadRequest, JSON:API, etc).
   2. On success
      1. Unmarshal body into the generated struct
      2. Unmarshal query params into struct
      3. Provide a hook to allow the consumer to modify the error on failure of 1 or 2
4. Validate Response
   1. The response should match the schema
      1. On error
         1. Build a slice containing validation errors
         2. Provide a hook to allow the consumer to log and alert.

## Stretch Goals

1. SDK Generation

## Regenerating Structs from Open API spec

First, make sure you have oapi-codegen installed:

```zsh
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
```

Then run:

```zsh
go generate ./...
```
