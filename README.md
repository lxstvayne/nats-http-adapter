# Nats HTTP Adapter

## Usage

Requests should be sent POST HTTP Request to endpoint `localhost:8080/`

Example:

```
{
    "subject":"test",
    "type":"request",
    "data": "any data"
}
```

List of params:

* *subject* - is a subject to which message is sent
* *type* - is a type of nats publish. Available types:

  * *request* - stands for *request* . Expects response, fails after timeout
  * *publish* - stands for *publish* .
* *data* - raw text of data. Might be JSON or whatever you like.
* *timeout -* timeout in milliseconds, default is 5000.

Response:

```json
{
  "data": {
    "error": "any error from your service"
  },
  "headers": null,
  "reply": "",
  "subject": "_INBOX.3GY2aJwUqwpFNWKd8EJqv7.qz3BYyUS"
}
```

## Docker Compose

```yaml
services:

  nats-http-adapter:
    image: ghcr.io/lxstvayne/nats-http-adapter:latest
    restart: "unless-stopped"
    environment:
      - NATS_URL=nats://host.docker.internal:4222

```
