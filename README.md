# k8s

**A Kubernetes learning sandbox**: one small Go HTTP server, deployed the same
way six times, so the difference between a Pod, a ReplicaSet and a Deployment is
something you watch happen rather than something you read about.

> ## ⚠️ Read this before copying anything
>
> `k8s/secret.yaml` contains a committed Kubernetes Secret. **A Kubernetes
> Secret is base64, not encryption** — `echo <value> | base64 -d` reveals it in
> one command, and this repository is public. The values here are throwaway
> study credentials that were never real.
>
> The demonstration is the point: this is exactly what committing a Secret to
> git looks like, and why `Secret` manifests belong in Sealed Secrets, SOPS,
> External Secrets Operator, or your cloud provider's secret manager instead.
>
> `server.go` also exposes `/secret`, which prints those values back over plain
> HTTP. That endpoint exists to prove the ConfigMap and Secret reached the
> container. **Never do either of these in real code.**

## The server

Two endpoints, a handful of lines, deliberately boring — the point is the
manifests, not the app:

| Route | Returns |
| --- | --- |
| `/hello` | A greeting built from `NAME` and `AGE`, injected by the ConfigMap |
| `/secret` | `USER` and `PASSWORD`, injected by the Secret (see the warning above) |

Reading configuration from the environment is what makes the ConfigMap and
Secret visible: change the manifest, restart the pod, watch the response change.

## The manifests

| File | Object | What it teaches |
| --- | --- | --- |
| `kind.yaml` | Cluster | A local cluster: one control plane, three workers |
| `pod.yaml` | Pod | The smallest unit — and why you never deploy one directly |
| `replicaset.yaml` | ReplicaSet | Replica count maintained; delete a pod and watch it return |
| `deployment.yaml` | Deployment | A ReplicaSet with rollout and rollback on top |
| `service.yaml` | Service | Stable networking; port `8080` → container port `3000` |
| `configmap-env.yaml` | ConfigMap | Non-secret configuration as environment variables |
| `secret.yaml` | Secret | Sensitive configuration — and its limits |

## Running it

Requires [kind](https://kind.sigs.k8s.io/), `kubectl` and Docker.

```bash
kind create cluster --config k8s/kind.yaml

kubectl apply -f k8s/configmap-env.yaml
kubectl apply -f k8s/secret.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

kubectl get pods -w
kubectl port-forward service/goserver-service 8080:8080
curl localhost:8080/hello
```

The experiment worth running, once it is up:

```bash
kubectl delete pod <one-of-the-pods>   # the Deployment replaces it
kubectl scale deployment goserver --replicas=5
```

### Note on the image

The manifests reference `jonasborgeslm/hello-go:v4` from Docker Hub. To use your
own build:

```bash
docker build -t <your-user>/hello-go:v1 .
docker push <your-user>/hello-go:v1
# then update the image field in the manifests
```

`kind load docker-image <your-user>/hello-go:v1` also works, and skips the
registry entirely.

## Related

[`gitops`](https://github.com/JonasBorgesLM/gitops) takes the next step: the
same kind of Go server, but with the manifests updated by a CD pipeline instead
of by hand.
