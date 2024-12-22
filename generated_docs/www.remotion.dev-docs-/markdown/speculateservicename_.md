---
title: speculateServiceName()
source: Remotion Documentation
last_updated: 2024-12-22
---

# speculateServiceName()

- [Home page](/)
- [@remotion/cloudrun](/docs/cloudrun/api)
- speculateServiceName()

On this page

# speculateServiceName()

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

Speculate the name of the Cloud Run service that will be created by [deployService()](/docs/cloudrun/deployservice) or its CLI equivalent, [`npx remotion cloudrun services deploy`](/docs/cloudrun/cli/services). This could be useful in cases when the configuration of the Cloud Run service is known in advance, and the name of the service is needed.

If you are not sure whether a service exists, use [`getServiceInfo()`](/docs/cloudrun/getserviceinfo) and catch the error that gets thrown if it does not exist.

If you want to get a list of deployed services, use [`getServices()`](/docs/cloudrun/getservices) instead.

## Service name pattern [​](\#service-name-pattern "Direct link to Service name pattern")

The service name depends on the following parameters:

- Remotion version
- Memory Limit
- CPU Limit
- Timeout in seconds

The name of the service resembles the following pattern:

```

remotion--3-3-96--mem2gi--cpu1-0--t-1900
          ^^^^^^   ^^^     ^^^      ^^^
            |       |       |        |-- Timeout in seconds
            |       |       |----------- Cpu limit
            |       |------------------- Memory limit
            |--------------------------- Remotion version with dots replaced by dashes
```

```

remotion--3-3-96--mem2gi--cpu1-0--t-1900
          ^^^^^^   ^^^     ^^^      ^^^
            |       |       |        |-- Timeout in seconds
            |       |       |----------- Cpu limit
            |       |------------------- Memory limit
            |--------------------------- Remotion version with dots replaced by dashes
```

## Example [​](\#example "Direct link to Example")

```

ts

import { speculateServiceName } from "@remotion/cloudrun";

const speculatedServiceName = speculateServiceName({
  memoryLimit: "2Gi",
  cpuLimit: "2",
  timeoutSeconds: 300,
});

console.log(speculatedServiceName); // remotion--3-3-96--mem2gi--cpu2-0--t-300
```

```

ts

import { speculateServiceName } from "@remotion/cloudrun";

const speculatedServiceName = speculateServiceName({
  memoryLimit: "2Gi",
  cpuLimit: "2",
  timeoutSeconds: 300,
});

console.log(speculatedServiceName); // remotion--3-3-96--mem2gi--cpu2-0--t-300
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `memoryLimit` [​](\#memorylimit "Direct link to memorylimit")

The upper bound on the amount of RAM that the Cloud Run service can consume.

### `cpuLimit` [​](\#cpulimit "Direct link to cpulimit")

The maximum number of CPU cores that the Cloud Run service can use to process requests.

### `timeoutSeconds` [​](\#timeoutseconds "Direct link to timeoutseconds")

The timeout that has been assigned to the Cloud Run service.

## Return value [​](\#return-value "Direct link to Return value")

A string with the speculated name of the service.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/cloudrun/src/api/speculate-service-name.ts)
- [deployService()](/docs/cloudrun/deployservice)
- CLI version of `deployService()`: [`npx remotion cloudrun services deploy`](/docs/cloudrun/cli/services#deploy)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/speculateservicename.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getServices()](/docs/cloudrun/getservices) [Next\
\
getRegions()](/docs/cloudrun/getregions)

- [Service name pattern](#service-name-pattern)
- [Example](#example)
- [Arguments](#arguments)
  - [`memoryLimit`](#memorylimit)
  - [`cpuLimit`](#cpulimit)
  - [`timeoutSeconds`](#timeoutseconds)
- [Return value](#return-value)
- [See also](#see-also)
