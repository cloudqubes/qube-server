`qube-server` is a tiny web server for testing and learning Docker and Kubernetes.

# Run qube-server on Docker

```shell
$ docker run -d -p 8080:8080 cloudqubes/qube-server:1.0.1
```

Test `qube-server` with `curl`.
```shell
$ curl http://127.0.0.1:8080
```

# Paths

`/`

Return 'Hello cloud' as the response.
Example:
```shell
$ curl http://127.0.0.1:8080
Hello cloud
```

`/count`

Starts a counter that increment by 1 at each request received to `/count`.



