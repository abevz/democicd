# democicd

Minimal Go HTTP service used as the workload for a full signed CI/CD supply
chain, deployed into the [`platform-iac-gitops`](https://github.com/abevz/platform-iac-gitops)
GitOps repo.

## Pipeline

`.gitlab-ci.yml` runs the deploy path on every push to `main`:

1. **build** — Kaniko builds the image (no Docker daemon) and pushes it to
   a private Harbor registry.
2. **scan** — Trivy blocks the pipeline on any CRITICAL, fixable
   vulnerability in the image.
3. **sign** — Vault-issued short-lived credentials are used to cosign-sign
   the image by digest.
4. **deploy** — the pipeline commits the new image digest into the GitOps
   repo (`k8s-lab-01/democicd/deployment.yaml`), which ArgoCD then syncs to
   the cluster.

Every stage authenticates to Vault with a CI-issued JWT (OIDC-style) rather
than a long-lived secret; credentials are fetched per-job and never stored in
CI variables. Kyverno on the cluster enforces `require-signed-images`, so an
unsigned or unscanned image cannot be admitted even if the pipeline were
bypassed.

`.github/workflows/ci.yml` runs a separate, registry-agnostic check (test,
lint, build, smoke test, `docker build`) on GitHub for portability — it does
not push to the private Harbor registry or sign images.

## Repo contents

| Path | Purpose |
|---|---|
| `main.go` | The HTTP service itself |
| `Dockerfile` | Multi-stage build (Go builder → Alpine runtime) |
| `Deployment.yaml`, `Service.yaml` | Kubernetes manifests (mirrored into the GitOps repo for ArgoCD) |
| `.gitlab-ci.yml` | Build → scan → sign → deploy pipeline |
| `.github/workflows/ci.yml` | Test/lint/build check for GitHub |

## Related repos

- [`platform-iac-gitops`](https://github.com/abevz/platform-iac-gitops) —
  ArgoCD app-of-apps that deploys this service.
- [`platform-iac`](https://github.com/abevz/platform-iac) — cluster
  provisioning (OpenTofu + Ansible on Proxmox).
