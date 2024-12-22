---
title: getServices()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getServices()

- [Home page](/)
- [@remotion/cloudrun](/docs/cloudrun/api)
- getServices()

On this page

# getServices()

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

Retrieves a list of Remotion services deployed to GCP Cloud Run.

The parameter `compatibleOnly` determines whether only services that are compatible with the installed version of Remotion Cloud Run should be returned.

note

The Cloud Run service is versioned and the version of the service must match the version of the `@remotion/cloudrun` package. So if you upgrade Remotion, you should deploy a new service or otherwise you might get an empty array from this function.

To get information about only a single service, use [`getServiceInfo()`](/docs/cloudrun/getserviceinfo).

If you are sure that a service exists, you can also guess the name of it using [`speculateServiceName()`](/docs/cloudrun/speculateservicename) and save an API call to Cloud Run.

## Example [​](\#example "Direct link to Example")

```

ts

import { getServices } from "@remotion/cloudrun/client";

const info = await getServices({
  region: "us-east1",
  compatibleOnly: true,
});

for (const service of info) {
  console.log(service.serviceName); // "remotion--3-3-82--mem512mi--cpu1-0"
  console.log(service.timeoutInSeconds); // 300
  console.log(service.memoryLimit); // 512Mi
  console.log(service.cpuLimit); // 1.0
  console.log(service.remotionVersion); // "4.0.1"
  console.log(service.uri); // "https://remotion--3-3-82--mem512mi--cpu1-0--t-300-1a2b3c4d5e-ue.a.run.app"
  console.log(service.region); // "us-east1"
  console.log(service.consoleUrl); // "https://console.cloud.google.com/run/detail/us-east1/remotion--3-3-82--mem512mi--cpu1-0--t-300/logs"
}
```

```

ts

import { getServices } from "@remotion/cloudrun/client";

const info = await getServices({
  region: "us-east1",
  compatibleOnly: true,
});

for (const service of info) {
  console.log(service.serviceName); // "remotion--3-3-82--mem512mi--cpu1-0"
  console.log(service.timeoutInSeconds); // 300
  console.log(service.memoryLimit); // 512Mi
  console.log(service.cpuLimit); // 1.0
  console.log(service.remotionVersion); // "4.0.1"
  console.log(service.uri); // "https://remotion--3-3-82--mem512mi--cpu1-0--t-300-1a2b3c4d5e-ue.a.run.app"
  console.log(service.region); // "us-east1"
  console.log(service.consoleUrl); // "https://console.cloud.google.com/run/detail/us-east1/remotion--3-3-82--mem512mi--cpu1-0--t-300/logs"
}
```

note

Import from [`@remotion/cloudrun/client`](/docs/cloudrun/light-client) to not import the whole renderer, which cannot be bundled.

## Argument [​](\#argument "Direct link to Argument")

An object containing the following properties:

### `region` [​](\#region "Direct link to region")

The [GCP region](/docs/cloudrun/region-selection) that you would like to query.

```

ts

import { getServices } from "@remotion/cloudrun";

const info = await getServices({
  region: "us-west1",
  compatibleOnly: true,
});

for (const service of info) {
  console.log(service.serviceName); // "remotion--3-3-82--mem2gi--cpu2--t-1100"
  console.log(service.timeoutInSeconds); // 1100
  console.log(service.memoryLimit); // 2Gi
  console.log(service.cpuLimit); // 2
  console.log(service.remotionVersion); // "3.3.82"
  console.log(service.uri); // "https://remotion--3-3-82--mem2gi--cpu2--t-1100-1a2b3c4d5e-uw.a.run.app"
  console.log(service.region); // "us-west1"
  console.log(service.consoleUrl); // "https://console.cloud.google.com/run/detail/us-west1/remotion--3-3-82--mem2gi--cpu2--t-1100/logs"
}
```

```

ts

import { getServices } from "@remotion/cloudrun";

const info = await getServices({
  region: "us-west1",
  compatibleOnly: true,
});

for (const service of info) {
  console.log(service.serviceName); // "remotion--3-3-82--mem2gi--cpu2--t-1100"
  console.log(service.timeoutInSeconds); // 1100
  console.log(service.memoryLimit); // 2Gi
  console.log(service.cpuLimit); // 2
  console.log(service.remotionVersion); // "3.3.82"
  console.log(service.uri); // "https://remotion--3-3-82--mem2gi--cpu2--t-1100-1a2b3c4d5e-uw.a.run.app"
  console.log(service.region); // "us-west1"
  console.log(service.consoleUrl); // "https://console.cloud.google.com/run/detail/us-west1/remotion--3-3-82--mem2gi--cpu2--t-1100/logs"
}
```

### `compatibleOnly` [​](\#compatibleonly "Direct link to compatibleonly")

If `true`, only services that match the version of the current Remotion Cloud run package are returned. If `false`, all Remotion services are returned.

## Return value [​](\#return-value "Direct link to Return value")

A promise resolving to an **array of objects** with the following properties:

### `serviceName` [​](\#servicename "Direct link to servicename")

The name of the service.

### `memoryLimit` [​](\#memorylimit "Direct link to memorylimit")

The upper bound on the amount of RAM that the Cloud Run service can consume.

### `cpuLimit` [​](\#cpulimit "Direct link to cpulimit")

The maximum number of CPU cores that the Cloud Run service can use to process requests.

### `remotionVersion` [​](\#remotionversion "Direct link to remotionversion")

The Remotion version of the service. Remotion is versioning the Cloud Run service and a render can only be triggered from a version of `@remotion/cloudrun` that is matching the one of the service.

### `timeoutInSeconds` [​](\#timeoutinseconds "Direct link to timeoutinseconds")

The timeout that has been assigned to the Cloud Run service.

### `uri` [​](\#uri "Direct link to uri")

The endpoint of the service.

### `region` [​](\#region-1 "Direct link to region-1")

The region of the deployed service. Useful if passing 'all regions' to the region input.

### `consoleUrl` [​](\#consoleurl "Direct link to consoleurl")

A link to the GCP console page for this service. Specifically, a link to logs display.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/cloudrun/src/api/get-services.ts)
- [`getServiceInfo()`](/docs/cloudrun/getserviceinfo)
- [`deployService()`](/docs/cloudrun/deployservice)
- [`deleteService()`](/docs/cloudrun/deleteservice)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/getservices.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
deleteService()](/docs/cloudrun/deleteservice) [Next\
\
speculateServiceName()](/docs/cloudrun/speculateservicename)

- [Example](#example)
- [Argument](#argument)
  - [`region`](#region)
  - [`compatibleOnly`](#compatibleonly)
- [Return value](#return-value)
  - [`serviceName`](#servicename)
  - [`memoryLimit`](#memorylimit)
  - [`cpuLimit`](#cpulimit)
  - [`remotionVersion`](#remotionversion)
  - [`timeoutInSeconds`](#timeoutinseconds)
  - [`uri`](#uri)
  - [`region`](#region-1)
  - [`consoleUrl`](#consoleurl)
- [See also](#see-also)
