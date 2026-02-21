# ingress-nginx — Installation

## Add Helm repo

```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
```

## Apply namespace

```bash
kubectl apply -f infra/ingress-nginx/namespace.yaml
```

## Install chart

```bash
helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx \
  --version 4.10.1 \
  --namespace ingress-nginx \
  -f infra/ingress-nginx/helm/values.yaml
```

## Generate rendered YAML (for transparency/GitOps)

```bash
helm template ingress-nginx ingress-nginx/ingress-nginx \
  --version 4.10.1 \
  --namespace ingress-nginx \
  -f infra/ingress-nginx/helm/values.yaml \
  > infra/ingress-nginx/rendered/all.yaml
```

Commit `rendered/all.yaml` to Git — enables `git diff` on upgrades.
