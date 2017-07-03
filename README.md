# dc-check
Simple Go Template syntax checker for Rancher's docker-compose files

## Introduction

Rancher allows to use Go's template engine for docker-compose files.
If you develop for a Rancher catalog item you might want to
check how your docker-compose files look like based on possible answers
given when you deploy a service from the catalog.

But how can you check how your docker-compose file looks like
without actually deploying it?

Rancher deploys your file regardless whether Go Template engine has
fully comiled your file or not. What's more, it's hard to interprete
Rancher deployment errors if your file has newline or indentation
errors and you can't actually see the resulting code.

This is where dc-check comes into play. It will just print the
resulting docker-compose file to the console. See examples below.


## Build

`go build -o dc-check`


## Examples

A common example is to make ports public available or to just
expose them internally.

mini-compose.yml:
```
{{- if eq .Values.PLAINTEXT_PORT_PUBLIC "true" }}
  ports:
  - 80:8080/tcp
{{- else }}
  expose:
  - 8080/tcp
{{- end }}
```


### Example 1:

Setup your environment:

`export PLAINTEXT_PORT_PUBLIC=true`

`./dc-check mini-compose.yml`

will result in:

```
ports:
- 80:8080/tcp
```



### Example 2:

Setup your environment:

`export PLAINTEXT_PORT_PUBLIC=true`

`./dc-check mini-compose.yml`

will result in:

```
expose:
- 8080/tcp
```


### More examples

More complex example with ports can be found in this repo which also
works if no port is public at all:

`. sample1.env && ./dc-check sample-compose.yml`

```
version: '2'
services:
  myservice:
    image: ubuntu:latest
    ports:
    - ${SSL_PORT}:8443/tcp
    expose:
    - 8080/tcp
    environment:
      LOGLEVEL: ${LOGLEVEL}
```


`. sample2.env && ./dc-check sample-compose.yml`

```
version: '2'
services:
  myservice:
    image: ubuntu:latest
    expose:
    - 8080/tcp
    - 8443/tcp
    environment:
      LOGLEVEL: ${LOGLEVEL}
```

`. sample3.env && ./dc-check sample-compose.yml`

```
version: '2'
services:
  myservice:
    image: ubuntu:latest
    ports:
    - ${PLAINTEXT_PORT}:8080/tcp
    - ${SSL_PORT}:8443/tcp
    environment:
      LOGLEVEL: ${LOGLEVEL}
```
