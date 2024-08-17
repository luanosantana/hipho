## Configure

```
export INGRESS_NAME=istio-ingressgateway                                       
export INGRESS_NS=istio-system
export INGRESS_HOST=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export SECURE_INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="https")].port}')
export TCP_INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="tcp")].port}')
```

```
minikube tunnel
```

### Configure DNS

IP From `minikube tunnel`

```
10.110.175.60 httpbin.example.com
10.110.175.60 {{ .Release.Name }}.example.com
```


curl --location 'http://httpbin.example.com/health' \
--header 'header: acessp' \
--header 'istio_ola: agora'


kubectl -n istio-system create token kiali

kubectl port-forward svc/kiali -n istio-system 8080:20001


# Application

## Build and Push

docker build -t docker.io/lsantana/alpine-mod:1.0.4 .

docker push lsantana/alpine-mod:1.0.4