---
title: npx remotion cloudrun services
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion cloudrun services

- [Home page](/)
- [Command line](/docs/cli/)
- [cloudrun](/docs/cloudrun/cli)
- services

On this page

# npx remotion cloudrun services

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs
may change in any version and documentation is not yet finished. See the

[changelog to stay up to date with breaking changes](https://remotion.dev/changelog)

.

The `npx remotion cloudrun services` command allows you to deploy, view and delete GCP Cloud Run services that can render videos and stills.

- [`deploy`](#deploy)
- [`ls`](#ls)
- [`rm`](#rm)
- [`rmall`](#rmall)

You only need one service per GCP region and Remotion version.

## deploy [​](\#deploy "Direct link to deploy")

```

npx remotion cloudrun services deploy
```

```

npx remotion cloudrun services deploy
```

Creates a new service in your GCP project. If a service exists in the same region, with the same Remotion version, with the same amount of memory, disk space and timeout duration, the name of the already deployed service will be returned instead.

Example output

```
Validating Deployment of Cloud Run Service:

Remotion Version: 3.3.95
Memory Limit: 2Gi
CPU Limit: 1.0
Timeout: 300
Project Name: remotion-example
Region: us-east1

Deploying Cloud Run Service...

Cloud Run Deployed!

Service name: remotion--3-3-95--mem512mi--cpu2--t-1200
Version: 3.3.95
CPU Limit: 2

Memory Limit: 512Mi
Timeout: 1200sec
Region: us-east1
Service URL:
https://remotion--3-3-95--mem512mi--cpu2--t-1200-1a2b3c4d5e-ue.a.run.app
GCP Console URL:
https://console.cloud.google.com/run/detail/us-east1/remotion--3-3-95--mem512mi--cpu2--t-1200/logs

```

### `--region` [​](\#--region "Direct link to --region")

The [GCP region](/docs/cloudrun/region-selection) to select. The site that the service will be accessing should also be in this same region to minimise latency.

### `--memoryLimit` [​](\#--memorylimit "Direct link to --memorylimit")

The upper bound on the amount of RAM that the Cloud Run service can consume. Default: 2 GB.

### `--cpuLimit` [​](\#--cpulimit "Direct link to --cpulimit")

The maximum number of CPU cores that the Cloud Run service can use to process requests. Default: 1.0.

### `--minInstances` [​](\#--mininstances "Direct link to --mininstances")

The minimum number of service instances to have available, regardless of requests. Default: 0.

note

Any running instances, even if they are not performing a render, will be billable in GCP. The default minimum number of instances is zero, which means that when no requests are made to your service, you are not billed.

### `--maxInstances` [​](\#--maxinstances "Direct link to --maxinstances")

The maximum number of service instances that can be create by GCP in response to incoming requests. Default: 100.

### `--timeoutSeconds` [​](\#--timeoutseconds "Direct link to --timeoutseconds")

Timeout of the Cloud Run service. Default: 300 seconds.

info

Not to be confused with the [`--timeout` flag when rendering which defines the timeout for `delayRender()`](/docs/cli/render#--timeout).

### `--onlyAllocateCpuDuringRequestProcessing` [v4.0.221](https://github.com/remotion-dev/remotion/releases/v4.0.221) [​](\#--onlyallocatecpuduringrequestprocessing "Direct link to --onlyallocatecpuduringrequestprocessing")

If this is set to true, `cpu_idle` will be set to `true` in the service manifest.

CPU alloction will be disabled while no request is being processed, which can lead to significant cost savings.

### `--quiet`, `-q` [​](\#--quiet--q "Direct link to --quiet--q")

Only logs the service name, and 'Authenticated access granted'.

## ls [​](\#ls "Direct link to ls")

```

npx remotion cloudrun services ls
```

```

npx remotion cloudrun services ls
```

Lists the services that you have deployed to GCP in the [selected region](/docs/cloudrun/region-selection).

Example output

```
2 services in us-east1

Service name: remotion--3-3-95--mem512mi--cpu2--t-1200
Version: 3.3.95
CPU Limit: 2

Memory Limit: 512Mi
Timeout: 1200sec
Region: us-east1
Service URL:
https://remotion--3-3-95--mem512mi--cpu2--t-1200-1a2b3c4d5e-ue.a.run.app
GCP Console URL:
https://console.cloud.google.com/run/detail/us-east1/remotion--3-3-95--mem512mi--cpu2--t-1200/logs

Service name: remotion--3-3-82--mem512mi--cpu1-0--t-800
Version: 3.3.82
CPU Limit: 1.0
Memory Limit: 512Mi
Timeout: 800sec
Region: us-east1
Service URL:
https://remotion--3-3-82--mem512mi--cpu1-0--t-800-1a2b3c4d5e-ue.a.run.app
GCP Console URL:
https://console.cloud.google.com/run/detail/us-east1/remotion--3-3-82--mem512mi--cpu1-0--t-800/logs

```

### `--region` [​](\#--region-1 "Direct link to --region-1")

The [GCP region](/docs/cloudrun/region-selection) to list services from.

### `--quiet`, `-q` [​](\#--quiet--q-1 "Direct link to --quiet--q-1")

Prints only the service names in a space-separated list. If no services exist, prints `()`

## rm [​](\#rm "Direct link to rm")

```

npx remotion cloudrun services rm remotion--3-3-82--mem512mi--cpu1-0--t-800
```

```

npx remotion cloudrun services rm remotion--3-3-82--mem512mi--cpu1-0--t-800
```

Removes one or more services from your GCP project. Pass a space-separated list of services you'd like to delete.

Example output

```

Service name:      remotion--3-3-82--mem2gi--cpu1-0--t-800

Version:           3.3.82

CPU Limit:         1.0

Memory Limit:      2Gi

Timeout:           300sec

Region:            us-east1

Service URL:       https://remotion--3-3-82--mem2gi--cpu1-0--t-800-1a2b3c4d5e-ue.a.run.app

GCP Console URL:   https://console.cloud.google.com/run/detail/us-east1/remotion--3-3-82--mem2gi--cpu1-0--t-800/logs
Delete? (Y/n):  Y

Deleted!

```

### `--region` [​](\#--region-2 "Direct link to --region-2")

The [GCP region](/docs/cloudrun/region-selection) to select.

### `--yes`, `-y` [​](\#--yes--y "Direct link to --yes--y")

Skips confirmation.

## rmall [​](\#rmall "Direct link to rmall")

```

npx remotion cloudrun services rmall
```

```

npx remotion cloudrun services rmall
```

Removes all services from your GCP project for a certain region.

Example output

```
2 services in us-east1

Service name: remotion--3-3-95--mem512mi--cpu2--t-1200
Version: 3.3.95
CPU Limit: 2

Memory Limit: 512Mi
Timeout: 1200sec
Region: us-east1
Service URL:
https://remotion--3-3-95--mem512mi--cpu2--t-1200-1a2b3c4d5e-ue.a.run.app
GCP Console URL:
https://console.cloud.google.com/run/detail/us-east1/remotion--3-3-95--mem512mi--cpu2--t-1200/logs

Delete? (Y/n) n

Skipping service - remotion--3-3-95--mem512mi--cpu2--t-1200.

Service name: remotion--3-3-82--mem512mi--cpu1-0--t-800
Version: 3.3.82
CPU Limit: 1.0
Memory Limit: 512Mi
Timeout: 800sec
Region: us-east1
Service URL:
https://remotion--3-3-82--mem512mi--cpu1-0--t-800-1a2b3c4d5e-ue.a.run.app
GCP Console URL:
https://console.cloud.google.com/run/detail/us-east1/remotion--3-3-82--mem512mi--cpu1-0--t-800/logs

Delete? (Y/n) n

Skipping service - remotion--3-3-82--mem512mi--cpu1-0--t-800.

```

### `--region` [​](\#--region-3 "Direct link to --region-3")

The [GCP region](/docs/cloudrun/region-selection) to remove services from.

### `--yes`, `-y` [​](\#--yes--y-1 "Direct link to --yes--y-1")

Skips confirmation.

## See also [​](\#see-also "Direct link to See also")

- [Setup guide](/docs/cloudrun/setup)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/cli/services.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
sites](/docs/cloudrun/cli/sites) [Next\
\
render](/docs/cloudrun/cli/render)

- [deploy](#deploy)
  - [`--region`](#--region)
  - [`--memoryLimit`](#--memorylimit)
  - [`--cpuLimit`](#--cpulimit)
  - [`--minInstances`](#--mininstances)
  - [`--maxInstances`](#--maxinstances)
  - [`--timeoutSeconds`](#--timeoutseconds)
  - [`--onlyAllocateCpuDuringRequestProcessing`](#--onlyallocatecpuduringrequestprocessing)
  - [`--quiet`, `-q`](#--quiet--q)
- [ls](#ls)
  - [`--region`](#--region-1)
  - [`--quiet`, `-q`](#--quiet--q-1)
- [rm](#rm)
  - [`--region`](#--region-2)
  - [`--yes`, `-y`](#--yes--y)
- [rmall](#rmall)
  - [`--region`](#--region-3)
  - [`--yes`, `-y`](#--yes--y-1)
- [See also](#see-also)
