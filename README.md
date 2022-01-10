# demoapp
Go demoapp

## /readyz
``` bash
curl -X POST "status: notReady" http://127.0.0.1:8080/readyz
# status code is 503 and NotReady
curl -X POST "status: ready" http://127.0.0.1:8080/readyz
# status code is 200 and Ready
```

## /livez
``` bash
curl -X POST "status: notLive" http://127.0.0.1:8080/livez
# status code is 503 and NotReady
curl -X POST "status: Live" http://127.0.0.1:8080/livez
# status code is 200 and Ready
```

## Docker
``` bash
docker run -d -p 80:8080 demoapp:v1.0
# docker run -d -e PORT:80 -p 80:80 demoapp:v1.0
```

## Kubernetes

``` bash
kubectl apply -f deploy.yaml
```