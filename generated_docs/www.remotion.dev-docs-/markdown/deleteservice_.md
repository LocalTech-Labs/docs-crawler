---
title: deleteService()
source: Remotion Documentation
last_updated: 2024-12-22
---

# deleteService()

- [Home page](/)
- [@remotion/cloudrun](/docs/cloudrun/api)
- deleteService()

On this page

# deleteService()

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

Deletes a deployed Cloud Run service based on its name.

To retrieve a list of services, call [`getServices()`](/docs/cloudrun/getservices) first.

## Example [​](\#example "Direct link to Example")

```

ts

import { deleteService, getServices } from "@remotion/cloudrun";

const services = await getServices({
  region: "us-east1",
  compatibleOnly: false,
});
for (const service of services) {
  await deleteService({
    region: "us-east1",
    serviceName: service.serviceName,
  });
}
```

```

ts

import { deleteService, getServices } from "@remotion/cloudrun";

const services = await getServices({
  region: "us-east1",
  compatibleOnly: false,
});
for (const service of services) {
  await deleteService({
    region: "us-east1",
    serviceName: service.serviceName,
  });
}
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `region` [​](\#region "Direct link to region")

The [GCP region](/docs/cloudrun/region-selection) to which the service was deployed to.

### `serviceName` [​](\#servicename "Direct link to servicename")

The name of the service to be deleted.

## Return value [​](\#return-value "Direct link to Return value")

Nothing. If the deletion failed, the service rejects with an error.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this service](https://github.com/remotion-dev/remotion/blob/main/packages/cloudrun/src/api/delete-service.ts)
- [deployService()](/docs/cloudrun/deployservice)
- [getServices()](/docs/cloudrun/getservices)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/deleteservice.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
deployService()](/docs/cloudrun/deployservice) [Next\
\
getServices()](/docs/cloudrun/getservices)

- [Example](#example)
- [Arguments](#arguments)
  - [`region`](#region)
  - [`serviceName`](#servicename)
- [Return value](#return-value)
- [See also](#see-also)
