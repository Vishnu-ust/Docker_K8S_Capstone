**Project Report — Docker_K8S_Capstone**

**Summary**
- **Purpose:** : Deliver a small Go web app that serves an HTML template with runtime metadata and can be run locally, in Docker, or deployed to Kubernetes.
- **Repository:** `Docker_K8S_Capstone` (branch: `feature/app`).
- **Status:** Basic working app with logging added; frontend template updated for improved appearance. Version display is currently hardcoded in `app/index.html` as `v2.0.0` and the backend sets `info.Version` to `"1.0.0"` (these can be synchronized — see Versioning section).

**Objectives**
- Show application information (name, version, hostname, time, environment) via a simple HTTP server.
- Provide a Docker image and Kubernetes manifests for deployment.
- Ensure logs are visible (stdout) so `kubectl logs` and `docker logs` capture activity.
- Make the UI clean and responsive.

**Technology Stack**
- Language: Go (golang)
- Container: Docker (multi-stage build)
- Orchestration: Kubernetes manifests included in `kubernetes/`
- Web: HTML template (`app/index.html`), server in `app/main.go`

**Repository Layout**
- `app/`
  - `main.go` — Go HTTP server and template rendering
  - `index.html` — HTML template (currently modernized styling)
  - `Dockerfile` — multi-stage build
  - `go.mod` — Go module file
- `kubernetes/` — `deployment.yaml`, `service.yaml`, `ingress.yaml`, `namespace.yaml`
- `README.md` — usage + instructions (added)
- `PROJECT_REPORT.md` — this document

**Design / Implementation Notes**
- `main.go` reads host info and environment, renders `index.html` using Go templates, and now logs startup and each request to stdout (timestamped with microseconds).
- The UI was restyled to a centered card and responsive table. Template placeholders remain for `{{.AppName}}`, `{{.Hostname}}`, `{{.CurrentTime}}`, and `{{.Environment}}`.
- Version handling in the source is currently inconsistent (see Versioning). Logs include the version the server reports.


**Logging & Observability**
- `main.go` now sets `log.SetFlags(log.LstdFlags | log.Lmicroseconds)` and logs:
  - Server startup: `starting server on :8080`
  - Each HTTP request: remote addr, method, path, served version, environment, hostname.
- To view logs:
  - Local run (`go run main.go`) — logs on the console
  - Docker: `docker logs <container-id>` or `docker logs -f <container-id>`
  - Kubernetes: `kubectl logs <pod-name> -n <namespace>` or `kubectl logs -f <pod-name> -n <namespace>`

**Build & Run (commands)**
- Local (development):
```powershell
Set-Location 'C:\Users\Administrator\Documents\Docker_K8S_Capstone\app'
# run directly
go run main.go
# or build
go build -o myapp .
.\myapp
```
- Docker (build + run):
```powershell
Set-Location 'C:\Users\Administrator\Documents\Docker_K8S_Capstone\app'
docker build -t shadowxlab-app:latest .
docker run --rm -p 8080:8080 shadowxlab-app:latest
```
- Docker with version injection (example):
```powershell
docker build --build-arg VERSION=v2.1.0 -t shadowxlab-app:v2.1.0 .
```
- Kubernetes (apply manifests):
```powershell
kubectl apply -f .\kubernetes\namespace.yaml
kubectl apply -f .\kubernetes\deployment.yaml
kubectl apply -f .\kubernetes\service.yaml
kubectl apply -f .\kubernetes\ingress.yaml
```

**Testing & Verification**
- Manual verification steps:
  - Start the server and `curl http://localhost:8080/` to confirm response and check logs on the console.
  - Rebuild Docker image and run container, then `docker logs` to confirm logs and UI.
  - Deploy to Kubernetes and run `kubectl logs` and `kubectl port-forward`/`kubectl get svc` to access the app.

**Assumptions**
- The server expects `index.html` to be available in the working directory (the Dockerfile copies it into the image in the current layout).
- The user has Go and Docker installed locally for build/run steps.

**Next Steps / Recommendations**
- Implement build-time version injection and template binding (`{{.Version}}`) so displayed version matches the build tag.
- Update `Dockerfile` to accept `ARG VERSION` and pass it to `go build -ldflags`.
- Optionally add a small CI workflow to tag builds using `git describe --tags` and publish images.
- Add a health/readiness probe to `kubernetes/deployment.yaml` for production-quality rollout.
- Consider using a lightweight asset pipeline for templates and static files.

**Appendix — Useful commands**
- View pod logs (Kubernetes):
```powershell
kubectl get pods -n <namespace>
kubectl logs -f <pod-name> -n <namespace>
```

