---
title: deleteSite()
source: Remotion Documentation
last_updated: 2024-12-22
---

# deleteSite()

- [Home page](/)
- [@remotion/cloudrun](/docs/cloudrun/api)
- deleteSite()

On this page

# deleteSite()

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

Removes a Remotion project from your Cloud Storage bucket.

Each project is located in the `sites/` subdirectory of your Cloud Storage bucket. Calling this function is equivalent of deleting all files inside a subfolder of your `sites/` subdirectory.

## Example [​](\#example "Direct link to Example")

Gets all sites and deletes them.

```

ts

import { GcpRegion, deleteSite, getSites } from "@remotion/cloudrun";

const region: GcpRegion = "australia-southeast1";

const { sites } = await getSites(region);

for (const site of sites) {
  await deleteSite({
    bucketName: site.bucketName,
    siteName: site.id,
  });
  console.log(`Site ${site.id} deleted.`);
}
```

```

ts

import { GcpRegion, deleteSite, getSites } from "@remotion/cloudrun";

const region: GcpRegion = "australia-southeast1";

const { sites } = await getSites(region);

for (const site of sites) {
  await deleteSite({
    bucketName: site.bucketName,
    siteName: site.id,
  });
  console.log(`Site ${site.id} deleted.`);
}
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `bucketName` [​](\#bucketname "Direct link to bucketname")

_string_

The name of the Cloud Storage bucket in which your site resides in.

### `siteName` [​](\#sitename "Direct link to sitename")

_string_

The unique ID of the project you want to delete.

## Return value [​](\#return-value "Direct link to Return value")

A promise resolving nothing.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/cloudrun/src/api/delete-site.ts)
- [getSites()](/docs/cloudrun/getsites)
- [deploySite()](/docs/cloudrun/deploysite)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/deletesite.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
deploySite()](/docs/cloudrun/deploysite) [Next\
\
getSites()](/docs/cloudrun/getsites)

- [Example](#example)
- [Arguments](#arguments)
  - [`bucketName`](#bucketname)
  - [`siteName`](#sitename)
- [Return value](#return-value)
- [See also](#see-also)
