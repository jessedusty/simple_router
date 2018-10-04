# Simple Router

A simple reverse proxy for handling sub-domains written in golang.

The normal usage of this image is in a docker-compose setup.

Simply link this container to any other container you want to reverse-proxy to. The hostname of the other container will be used as the subdomain. 

The "home" hostname is reserved for the / path

## Example docker-compose setup

```yaml
version: '2'

services: 
  router: 
    image: jessedusty/simple_router
    links: 
      - home
    ports:
      - 80:80

  home:
    image: jessedusty/simple_hello
    environment:
       hello: home

  test:
    image: jessedusty/simple_hello
    environment:
       hello: test

```

The services will be accessable as http://localhost/ and http://test.localhost
