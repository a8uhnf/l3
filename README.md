# DNS:: Drone Navigation System

L3, named after L3-37 is a drone navigation service for Atlas Corp, which
helps navigating drone fleet to locate appropriate data bank for a specified sector.
It exposes JSON REST API's and excepts HTTP requests from drones deployed in the specified sector.

## Instruction to Build
The Build commands are abstracted using make, following make commnads are available to use with the project.
```bash
make compile
```
### Build Docker
```bash
 make docker-build image=<IMAGE-TAG>
```


## Instruction to Run

### Run Docker

```bash
make docker-run
```

### Run Webserver

```bash
l3 serve --sector-id 2
```

## Running Tests

```bash
go test ./...
```

## Web Server Documentation


## API Documentation
---

- [Locations API](docs/api/location.md)
- [Metrics API](docs/api/metrics.md)
- [Health check APIs](docs/api/health.md)

## Deployment
Deployment manifests with kuberentes is available at [hack/k8s](hack/k8s).  Following instructions can be used with any k8s cluster
including minikube or kind.

## Monitoring
Promethues monitoring pull endpoint is available at /metrics. We can configure a promethues to pull data by adding the following config

```text
global:
  evaluation_interval: 1m
  scrape_interval: 10s
  scrape_timeout: 10s

rule_files:
- /etc/config/rules
- /etc/config/alerts
scrape_configs:
- job_name: prometheus
  static_configs:
  - targets:
    - localhost:9090
- job_name: kubernetes-service-endpoints
  kubernetes_sd_configs:
  - role: endpoints
  relabel_configs:
  - action: replace
    regex: (https?)
    source_labels:
    - __meta_kubernetes_service_annotation_prometheus_io_scheme
    target_label: __scheme__
  - action: replace
    regex: (.+)
    source_labels:
    - __meta_kubernetes_service_annotation_prometheus_io_path
    target_label: __metrics_path__
  - action: replace
    regex: ([^:]+)(?::\d+)?;(\d+)
    replacement: $1:$2
    source_labels:
    - __address__
    - __meta_kubernetes_service_annotation_prometheus_io_port
    target_label: __address__
  - action: labelmap
    regex: __meta_kubernetes_service_label_(.+)
  - action: replace
    source_labels:
    - __meta_kubernetes_namespace
    target_label: kubernetes_namespace
  - action: replace
    source_labels:
    - __meta_kubernetes_service_name
    target_label: kubernetes_name

```
