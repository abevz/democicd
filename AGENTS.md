# democicd - AI Agent Guide

This file provides repository guidance to AI coding agents working in
`/home/abevz/github/democicd`.

## Instruction Files

- `AGENTS.md` is tracked in git and lives on `main` so new worktrees inherit it.
- `CLAUDE.md` should point to `AGENTS.md`.

## Git Workflow

- This repository uses a bare repo plus git worktrees.
- Bare repo path: `/home/abevz/github/democicd/.bare`
- Main worktree path: `/home/abevz/github/democicd/main`
- Treat `main` as read, update, and status only.
- Keep repository-wide instruction files in `main` so every new worktree starts
  with the same baseline.
- Make changes on feature branches in sibling worktrees under
  `/home/abevz/github/democicd/`.
- One task means one branch, one sibling worktree, and one merge request.
- Create new worktrees from `main`, then do task changes inside that sibling
  worktree.
- Do not push directly to `main`.
- Remove completed worktrees with `git worktree remove`, not raw directory
  deletion.

## Attribution Rules

- Do not add AI assistant names, agent names, or bot identities to commit
  authors, committers, co-author trailers, commit messages, merge request
  descriptions, or repository documentation.
- Do not add `Co-authored-by`, `Generated-by`, `Assisted-by`, or similar
  attribution lines for AI tools.

## Project Context

- This repo is the workload side of the demo supply-chain path.
- It contains a small Go HTTP service, a Docker build, and basic Kubernetes
  manifests.
- Deployment contract changes here often need a matching change in
  `/home/abevz/github/platform-iac-gitops/main/k8s-lab-01/democicd`.

## Build And Verification

```bash
go test ./...
go build ./...
docker build -t democicd:dev .
```

## Code Style

- Go: keep code idiomatic and simple; prefer standard library unless there is a
  clear need for a dependency.
- Keep handlers and startup logic explicit; this repo is intentionally small and
  should stay easy to read.
- Preserve the container port and manifest wiring unless the deployment contract
  is intentionally changing.

## Kubernetes Manifest Rules

- Keep `Deployment.yaml` and `Service.yaml` aligned on namespace, labels,
  selector, and port wiring.
- Do not hardcode credentials or registry tokens in manifests.
- If the image name, pull secret, namespace, or ports change here, update the
  GitOps manifests in the sibling `platform-iac-gitops` repo in the same task.

## Commit Guidance

- Use Conventional Commits such as `feat:`, `fix:`, `docs:`, `refactor:`,
  `test:`, and `chore:`.
- Keep messages in English and focus on what changed and why.
