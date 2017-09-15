Prometheus: New Relic Exporter
==============================

[![CircleCI](https://circleci.com/gh/previousnext/newrelic-exporter.svg?style=svg)](https://circleci.com/gh/previousnext/newrelic-exporter)

**Maintainer**: Nick Schuch

Prometheus exporter for a single New Relic application.

Exposes the following metrics:

* Response Time
* Throughput
* Error Rate
* Apdex Target
* Apdex Score
* Host Count
* Instance Count

## Usage

**CLI**

```bash
newrelic-exporter --application="My Application" \
                  --api-key="xxxxxxxxxxxxxxxxx"
```

**Kubernetes**

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: amaysim-newrelic-exporter
  namespace: kube-system
spec:
  replicas: 1
  template:
    metadata:
      annotations:
        prometheus.io/port: "9000"
        prometheus.io/scrape: "true"
      labels:
        app: amaysim-newrelic-exporter
        task: monitoring
    spec:
      containers:
      - name: exporter
        image: previousnext/newrelic-exporter:0.0.2
        imagePullPolicy: Always
        resources:
          limits:
            cpu: 5m
            memory: 25Mi
          requests:
            cpu: 5m
            memory: 25Mi
        env:
          - name: NEW_RELIC_APPLICATION 
            value: "My Application"
          - name: NEW_RELIC_API_KEY
            value: "xxxxxxxxxxxxxxxxx"
```

## Resources

* [Dave Cheney - Reproducible Builds](https://www.youtube.com/watch?v=c3dW80eO88I)

## Development

### Tools

* **Dependency management** - https://getgb.io
* **Build** - https://github.com/mitchellh/gox
* **Linting** - https://github.com/golang/lint

### Workflow

(While in the `workspace` directory)

**Installing a new dependency**

```bash
gb vendor fetch github.com/foo/bar
```

**Running quality checks**

```bash
make lint test
```

**Building binaries**

```bash
make build
```
