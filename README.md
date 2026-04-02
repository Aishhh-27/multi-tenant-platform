![CI](https://github.com/Aishhh-27/multi-tenant-platform/actions/workflows/ci.yml/badge.svg)
# Multi-Tenant Environment Automation Platform
Inspired by GitLab Dedicated multi-tenant environment automation systems.

## Overview

This project implements a multi-tenant environment automation platform inspired by GitLab-style infrastructure workflows. It provisions, manages, monitors, and heals isolated tenant environments using a combination of Go, Terraform, Kubernetes, and Helm.

The system is designed to demonstrate practical Site Reliability Engineering (SRE) concepts including infrastructure as code, workload isolation, observability, and automated remediation.Designed to simulate GitLab-style environment automation workflows.

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

User
  ↓
Go CLI (tenant manager)
  ↓
Terraform (provisions infra per tenant)
  ↓
Kubernetes Cluster
  ↓
Namespace per Tenant
  ↓
Helm → GitLab / App Deployment
  ↓
Prometheus → Monitoring

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

## System Design Decisions

### Why Namespace-per-Tenant?
Each tenant is isolated using Kubernetes namespaces to ensure:
- Resource separation
- Fault isolation
- Independent scaling

### Why Terraform?
Ensures reproducible infrastructure and safe state management.

### Tradeoffs
- Namespace isolation vs cluster-per-tenant
- Shared cluster reduces cost but requires strict isolation

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


##  Failure Simulation

Scenario:
- Pod deleted manually

Result:
- Kubernetes automatically recreates pod
- Ensures high availability

This demonstrates self-healing infrastructure.

This triggers repeated restarts which are detected and handled by the auto-healing loop.

---

## Tenant Lifecycle

1. Create tenant → CLI triggers Terraform
2. Namespace created in Kubernetes
3. Helm deploys application
4. Prometheus starts monitoring
5. Failures auto-recovered
6. Tenant can be deleted cleanly

---

## Scaling Capability

This system is designed to support multiple tenants using:
- Automated provisioning
- Namespace isolation
- Repeatable Terraform modules

Tested with 5–10 tenants (can scale further with cluster capacity).

## Project Structure

```
cmd/                CLI commands
internal/           Core logic (terraform, kubernetes, helm, healing)
terraform/          Infrastructure configuration
helm/               Helm charts
screenshots/        Dashboard and execution output
```

---

## Observability

- Prometheus collects cluster and pod metrics
- Enables monitoring across multiple tenants
- Alerts can be configured for:
  - Pod failures
  - Resource exhaustion

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
