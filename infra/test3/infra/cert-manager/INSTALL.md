# cert-manager — Installation

## Add Helm repo

```bash
helm repo add jetstack https://charts.jetstack.io
helm repo update
```

## Apply namespace

```bash
kubectl apply -f infra/cert-manager/namespace.yaml
```

## Install chart

```bash
helm upgrade --install cert-manager jetstack/cert-manager \
  --version v1.14.4 \
  --namespace cert-manager \
  -f infra/cert-manager/helm/values.yaml
```

## Generate rendered YAML (for transparency/GitOps)

```bash
helm template cert-manager jetstack/cert-manager \
  --version v1.14.4 \
  --namespace cert-manager \
  -f infra/cert-manager/helm/values.yaml \
  > infra/cert-manager/rendered/all.yaml
```

Commit `rendered/all.yaml` to Git — enables `git diff` on upgrades.
