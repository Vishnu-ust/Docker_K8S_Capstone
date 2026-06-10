# Docker_K8S_Capstone

Simple Go web app that serves an `index.html` template with application info. This repo contains a small HTTP server, a Dockerfile, and Kubernetes manifests to deploy the app.

**Files of interest**
- `app/main.go` — Go HTTP server and template rendering.
- `app/index.html` — HTML template; currently shows a version and template fields (`{{.AppName}}`, `{{.Hostname}}`, `{{.CurrentTime}}`, `{{.Environment}}`).
- `app/Dockerfile` — multi-stage Dockerfile that builds the binary and packages `index.html`.
- `kubernetes/` — Kubernetes manifests: `deployment.yaml`, `service.yaml`, `ingress.yaml`, `namespace.yaml`.

## Prerequisites
- Go (1.18+ recommended)
- Docker
- kubectl (if deploying to Kubernetes)

## Run locally (development)
Run from the `app` directory so `index.html` is available to the running process:

```powershell
Set-Location 'C:\Users\Administrator\Documents\Docker_K8S_Capstone\app'
# Run directly
go run main.go
# or build and run
go build -o myapp .
.\myapp
```

Open: `http://localhost:8080`

When run locally, logs (startup and each request) are printed to the console.

## Build and run with Docker
From the `app` directory:

```powershell
Set-Location 'C:\Users\Administrator\Documents\Docker_K8S_Capstone\app'
# Build
docker build -t shadowxlab-app:latest .
# Run
docker run --rm -p 8080:8080 shadowxlab-app:latest
```

Open: `http://localhost:8080`

To see container logs:

```powershell
docker ps
docker logs <container-id>
# or stream
docker logs -f <container-id>
```

## Deploy to Kubernetes
1. Build & push an image to your registry (or load it into your local cluster):

```powershell
# Example (replace <your-repo> and tag):
docker build -t <your-repo>/shadowxlab-app:feature-app-v2 .
docker push <your-repo>/shadowxlab-app:feature-app-v2
```

2. Update `kubernetes/deployment.yaml` to use the new image tag.
3. Apply manifests:

```powershell
kubectl apply -f .\kubernetes\namespace.yaml
kubectl apply -f .\kubernetes\deployment.yaml
kubectl apply -f .\kubernetes\service.yaml
kubectl apply -f .\kubernetes\ingress.yaml
```

4. Check pods and logs:

```powershell
kubectl get pods -n <namespace>
kubectl logs -f <pod-name> -n <namespace>
```

## Version handling (automatic)
Current state in this workspace: `app/index.html` currently displays `v2.0.0` hardcoded in the badge and Version row.

To automatically inject the version at build time (recommended), do the following changes:

1. Make a package-level `Version` variable in `main.go` and use it when building the `AppInfo` (so the template receives `{{.Version}}`). Example:

```go
// package main
var Version = "v1.0.0" // overridden at link time

// use Version in AppInfo:
info := AppInfo{ Version: Version, ... }
```

2. Update `index.html` to use `{{.Version}}` for the badge and Version row (remove hardcoded `v2.0.0`).

3. Build with `go build` and inject the version via `-ldflags` or use the Docker build-arg in the `Dockerfile`:

Direct `go build` example:

```powershell
# inject version at link time
go build -ldflags "-X main.Version=v2.1.0" -o myapp main.go
```

Docker `Dockerfile` pattern (example) — allow a build-arg `VERSION` and pass it with `-ldflags`:

```dockerfile
ARG VERSION=v1.0.0
RUN go build -ldflags "-X main.Version=${VERSION}" -o myapp main.go
```

Build passing the arg:

```powershell
docker build --build-arg VERSION=v2.1.0 -t shadowxlab-app:2.1.0 .
```

Now the binary will set `main.Version` to the injected value and the template will render it automatically.

## Logs verification
- When running locally with `go run`, logs are printed to stdout.
- In Docker, container stdout is visible via `docker logs`.
- In Kubernetes, `kubectl logs <pod>` shows stdout/stderr from the container; `kubectl logs -f` to stream.

## Notes
- Ensure you run the binary from the `app` directory (or package `index.html` into the container) since the template is loaded from `index.html` at runtime.
- If you see stale content in the browser, clear cache or open an incognito window.

---
If you want, I can: 
- Update `main.go` to add the `Version` variable and wire it through (I already added logging earlier).
- Modify `index.html` to use `{{.Version}}` instead of a hardcoded value.
- Update the `Dockerfile` to accept a `VERSION` build-arg and inject it at build time.

Tell me which of these you'd like me to apply and I will make the changes and rebuild for you.
