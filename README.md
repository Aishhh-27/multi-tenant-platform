![CI](https://github.com/Aishhh-27/multi-tenant-platform/actions/workflows/ci.yml/badge.svg)
# Multi-Tenant Environment Automation Platform
Inspired by GitLab Dedicated multi-tenant environment automation systems.

## Overview

This project simulates a production-grade multi-tenant environment automation platform inspired by GitLab Dedicated, designed to provision, manage, and operate multiple isolated environments at scale using infrastructure-as-code and Kubernetes.

The system is designed to demonstrate practical Site Reliability Engineering (SRE) concepts including infrastructure as code, workload isolation, observability, and automated remediation.Designed to simulate GitLab-style environment automation workflows.

---

## Features

* Command-line interface written in Go for tenant lifecycle management
* Tenant isolation using Kubernetes namespaces
* Infrastructure provisioning using Terraform workspaces
* Application deployment using Helm charts
* Observability using Prometheus and Grafana
* Automated detection and recovery of failing workloads
* The Go CLI acts as a control plane for managing tenant lifecycle operations in a repeatable and automated manner.
* Designed with SRE principles such as automation, reliability, observability, and minimal manual intervention.

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
##  Why Multi-Tenant Architecture?

Multi-tenancy allows multiple environments (tenants) to run on shared infrastructure while maintaining isolation.

Benefits:
- Efficient resource utilization
- Easier scaling
- Centralized management
- Faster provisioning

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

##  Scaling Capability

The platform is designed to handle multiple tenants using:
- Terraform-based reproducible infrastructure
- Kubernetes namespace isolation
- Automated provisioning via CLI

Tested with multiple tenants and designed to scale further based on cluster capacity.

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

- Prometheus is used to monitor:
- Pod health across tenants
- Resource usage
- Cluster-level metrics

This enables visibility into multiple tenant environments and helps detect failures early.

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

This project demonstrates a complete workflow for managing multi-tenant environments with an emphasis on reliability and observability. It reflects practical patterns used in real-world SRE and platform engineering systems.It also demonstrates the transition from managing individual environments to operating a scalable platform capable of handling multiple tenants efficiently.


## Author
Aishwarya Ganesh
