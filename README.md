# HarborMaster

<img align="right" width="180px" src="logo.png">

#### HarborMaster is a simple frontend for a private docker registry.

---

![](screenshots/dashboard.png)

## Installation

### Docker

Docker images are available on [GHCR](https://github.com/users/mheidinger/packages/container/package/harbor-master).
The usage is described below.

### Compiling and running locally

Requires Go 1.26 or newer. The Docker builder uses Go 1.27.3.

For compiling the source, first download all dependencies:
```
go mod download
```

After that the compilation and execution can be done via the Makefile:
```
make
make run
```

A cleanup is also possible:
```
make clean
```

## Usage

```
./HarborMaster -help

Usage of ./HarborMaster:
  -password_file string
        Password of the docker registry (default "/run/secrets/hm_registry_password")
  -port int
        Port on which the application will run (default 4181)
  -url string
        URL of the docker registry (default "localhost:8080")
  -username_file string
        Username of the docker registry (default "/run/secrets/hm_registry_username")
```

HarborMaster does not perform any authentication itself — all routes are public. Put it behind a reverse proxy (e.g. Traefik, nginx) that handles authentication if you need access control.

## Dependency maintenance

Go modules and checksums are tracked in `go.mod` and `go.sum`. Keep
`github.com/distribution/reference` at v0.5.0 while using the current registry
client: v0.6.0 removes `SplitHostname`, which `github.com/docker/distribution` v2
still requires.

Frontend libraries are vendored under `static/`: Materialize 1.0.0 (the latest
stable original release) and SweetAlert2 11.26.25. The SweetAlert2 file uses the
upstream `sweetalert2.all.min.js` bundle, including CSS, under the existing local
filename `sweetalert2.min.js`.
