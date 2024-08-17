# Hip, Ho!

## Configure Environment

The assembled architecture depends on some tools for operation, which are:

- Minikube (k8s)
- Istio and Kiali
- Prometheus
- Jaeger

Below is a step-by-step guide for configuring each one.

### 1. Install minikube

In In the official minkube documentation with istio it is recommended that minikube be started with the following configuration. [[Documnetation](https://minikube.sigs.k8s.io/docs/handbook/addons/istio/)]

```
minikube start --memory=8192mb --cpus=4
```

### 2. Install Istio

Install istio according to the official configuration and then initialize istio in the kubernetes cluster using the CDR, using the command below. [[doc](https://istio.io/latest/docs/setup/getting-started/#download)]

```
$ istioctl install -f https://raw.githubusercontent.com/istio/istio/release-1.23/samples/bookinfo/demo-profile-no-gateways.yamldemo-profile-no-gateways.yaml -y
✔ Istio core installed
✔ Istiod installed
✔ Installation complete
Made this installation the default for injection and validation.
```

### 3. Prometheus to Istio

The official istio documentation shows how to configure Prometheus with istio [[Doc](https://istio.io/latest/docs/ops/integrations/prometheus/)]

```
$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.23/samples/addons/prometheus.yaml
```

### 4. Jaeger to Istio

The official istio documentation shows how to configure Jaeger with istio in a simple and direct way [[Doc](https://istio.io/latest/docs/ops/integrations/jaeger/#installation)]

```
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.23/samples/addons/jaeger.yaml
```
