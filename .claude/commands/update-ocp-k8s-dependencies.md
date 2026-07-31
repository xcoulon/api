---
name: update-ocp-k8s-dependencies
description: Update the OpenShift and Kubernetes dependencies
---

# Pre-check

The command requires the `major.minor` version of OpenShift and Kubernetes that this repository is now based on. Let's store the input as `OCP_VERSION` and `K8S_VERSION` for further use.

# Execution

Update the `k8s.io` modules to the corresponding `K8S_VERSION` and the `sigs.k8s.io` modules to the associated version in `go.mod` 

Update the github.com/openshift/api module to target the latest commit from the `release-$VERSION` branch on github.com/openshift/api

Also update the `CONTROLLER_TOOLS_VERSION` to the version `make/generate.mk` to the same version as `sigs.k8s.io/controller-tools` in `go.mod`

Ensure that the `toolchain` in `go.mod` refers to the latest version of Go based on the `go` directive 

Run `go mod tidy`

Run `make generate test` to verify that everything is good
