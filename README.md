# [Gardener Extension for Registry Cache](https://gardener.cloud)

[![REUSE status](https://api.reuse.software/badge/github.com/gardener/gardener-extension-registry-cache)](https://api.reuse.software/info/github.com/gardener/gardener-extension-registry-cache)
[![CI Build status](https://concourse.ci.gardener.cloud/api/v1/teams/gardener-tests/pipelines/gardener-extension-registry-cache-main/jobs/main-head-update-job/badge)](https://concourse.ci.gardener.cloud/teams/gardener-tests/pipelines/gardener-extension-registry-cache-main/jobs/main-head-update-job)
[![Go Report Card](https://goreportcard.com/badge/github.com/gardener/gardener-extension-registry-cache)](https://goreportcard.com/report/github.com/gardener/gardener-extension-registry-cache)

Gardener extension controller which deploys pull-through caches for container registries.

## Usage

- [Configuring the Registry Cache Extension](docs/usage/registry-cache/configuration.md) - learn what is the use-case for a pull-through cache, how to enable it and configure it
- [How to provide credentials for upstream repository?](docs/usage/registry-cache/upstream-credentials.md)
- [Migration from `v1alpha2` to `v1alpha3`](docs/usage/registry-cache/migration-from-v1alpha2-to-v1alpha3.md) - learn how to migrate from the `v1alpha2` API version of the `RegistryConfig` to `v1alpha3`
- [Configuring the Registry Mirror Extension](docs/usage/registry-mirror/configuration.md) - learn what is the use-case for a registry mirror, how to enable and configure it

## Local setup and development

- [Deploying Registry Cache Extension Locally](docs/development/getting-started-locally.md) - learn how to set up a local development environment
- [Developer Docs for Gardener Extension Registry Cache](docs/development/extension-registry-cache.md) - learn about the inner workings
