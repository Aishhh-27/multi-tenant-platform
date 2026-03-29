# GitLab Multi-Tenant Environment Automation Platform

## Overview

This project implements a multi-tenant environment automation platform inspired by GitLab-style infrastructure workflows. It provisions, manages, monitors, and heals isolated tenant environments using a combination of Go, Terraform, Kubernetes, and Helm.

The system is designed to demonstrate practical Site Reliability Engineering (SRE) concepts including infrastructure as code, workload isolation, observability, and automated remediation.

---

## Features

* Command-line interface written in Go for tenant lifecycle management
* Tenant isolation using Kubernetes namespaces
* Infrastructure provisioning using Terraform workspaces
* Application deployment using Helm charts
* Observability using Prometheus and Grafana
* Automated detection and recovery of failing workloads

---

## Architecture

The platform follows a simple execution flow:

CLI → Terraform → Kubernetes → Helm
↓
Prometheus + Grafana
↓
Auto-healing loop

Each tenant is provisioned with:

* A dedicated Terraform workspace
* A Kubernetes namespace
* An isolated Helm release

---

## Usage

### Create a tenant

```bash
./platform create-tenant --name tenant-a
```

### Delete a tenant

```bash
./platform delete-tenant --name tenant-a
```

### Run auto-healing

```bash
./platform heal-tenant --name tenant-a
```

---

## Observability

Grafana dashboards are used to monitor tenant workloads.

### CPU Usage

![CPU](screenshots/cpu.png)

### Memory Usage

![Memory](screenshots/memory.png)

### Pod Restarts

![Restarts](screenshots/restart.png)

---

## Auto-healing

The platform continuously monitors pod states and detects failure conditions such as:

* CrashLoopBackOff
* Error
* ImagePullBackOff

When a failure is detected, the system automatically deletes the affected pods, allowing Kubernetes to recreate them.

![Healing](screenshots/healing.png)

The system detects failing pods and deletes only the affected workloads instead of restarting the entire namespace.

---


## Failure Simulation

To simulate a failure scenario:

```bash
kubectl run crash-test \
  --image=busybox \
  -n tenant-a \
  --restart=Always \
  -- sh -c "exit 1"
```

This triggers repeated restarts which are detected and handled by the auto-healing loop.

---

## Project Structure

```
cmd/                CLI commands
internal/           Core logic (terraform, kubernetes, helm, healing)
terraform/          Infrastructure configuration
helm/               Helm charts
screenshots/        Dashboard and execution output
```

---

## Notes

* The auto-healing logic currently operates at the namespace level and deletes all pods when a failure is detected. This can be refined to target only failing pods.
* Metrics visibility depends on Prometheus scrape intervals, so short-lived pods may not always appear in dashboards.

---

## Future Improvements

* Fine-grained healing (target specific pods instead of all pods)
* Support for multiple cloud providers
* Role-based access control per tenant
* GitOps integration for deployment workflows
* Operator-based implementation for reconciliation

---

## Summary

This project demonstrates a complete workflow for managing multi-tenant environments with an emphasis on reliability and observability. It reflects practical patterns used in real-world SRE and platform engineering systems.


## Author
Aishwarya Ganesh
