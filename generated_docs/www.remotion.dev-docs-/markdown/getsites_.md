---
title: getSites()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getSites()

- [Home page](/)
- [@remotion/cloudrun](/docs/cloudrun/api)
- getSites()

On this page

# getSites()

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

Gets an array of Remotion projects in Cloud Storage, in your GCP project.

The projects are located in the `sites/` subdirectory of your Cloud Storage bucket. Remember - you should only have one bucket for Remotion Cloud Run per region, therefore you do not need to specify the name of the bucket for this function.

## Example [​](\#example "Direct link to Example")

Gets all sites and logs information about them.

```

ts

import { getSites } from "@remotion/cloudrun";

const { sites, buckets } = await getSites("europe-west4");

for (const site of sites) {
  console.log(site.id); // A unique ID for referring to that project
  console.log(site.bucketName); // In which bucket the site resides in.
  console.log(site.bucketRegion); // In which region the bucket resides in.
  console.log(site.serveUrl); // URL of the deployed site that you can pass to `renderMediaOnCloudRun()`
}

for (const bucket of buckets) {
  console.log(bucket.name); // The name of the Cloud Storage bucket.
  console.log(bucket.creationDate); // A unix timestamp of when the site was created.
  console.log(bucket.region); // 'europe-west4'
}
```

```

ts

import { getSites } from "@remotion/cloudrun";

const { sites, buckets } = await getSites("europe-west4");

for (const site of sites) {
  console.log(site.id); // A unique ID for referring to that project
  console.log(site.bucketName); // In which bucket the site resides in.
  console.log(site.bucketRegion); // In which region the bucket resides in.
  console.log(site.serveUrl); // URL of the deployed site that you can pass to `renderMediaOnCloudRun()`
}

for (const bucket of buckets) {
  console.log(bucket.name); // The name of the Cloud Storage bucket.
  console.log(bucket.creationDate); // A unix timestamp of when the site was created.
  console.log(bucket.region); // 'europe-west4'
}
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `region` [​](\#region "Direct link to region")

The [GCP region](/docs/cloudrun/region-selection) which you want to query. Alternatively, you can pass 'all regions' to return sites across all regions.

```

ts

import { getSites } from "@remotion/cloudrun";

const { sites, buckets } = await getSites("all regions");
```

```

ts

import { getSites } from "@remotion/cloudrun";

const { sites, buckets } = await getSites("all regions");
```

## Return value [​](\#return-value "Direct link to Return value")

A promise resolving to an object with the following properties:

### `sites` [​](\#sites "Direct link to sites")

An array of deployed Remotion projects that you can use for rendering.

Each item contains the following properties:

#### `id` [​](\#id "Direct link to id")

A unique identifier for that project.

#### `bucketName` [​](\#bucketname "Direct link to bucketname")

The bucket in which the project resides in.

#### `bucketRegion` [​](\#bucketregion "Direct link to bucketregion")

The region in which the bucket resides in.

#### `serveUrl` [​](\#serveurl "Direct link to serveurl")

URL of the deployed site. You can pass it into [`renderMediaOnCloudRun()`](/docs/cloudrun/rendermediaoncloudrun) to render a video or audio.

### `buckets` [​](\#buckets "Direct link to buckets")

An array of all buckets in the selected region in your account that start with `remotioncloudrun-`.

info

You should only have [1 bucket](/docs/cloudrun/multiple-buckets) per region for all your Remotion projects.

Each item contains the following properties:

#### `region` [​](\#region-1 "Direct link to region-1")

The region the bucket resides in.

#### `name` [​](\#name "Direct link to name")

The name of the bucket. Cloud Storage buckets have globally unique names.

#### `creationDate` [​](\#creationdate "Direct link to creationdate")

A UNIX timestamp of the point when the bucket was first created.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/cloudrun/src/api/get-sites.ts)
- [deleteSite()](/docs/cloudrun/deletesite)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/getsites.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
deleteSite()](/docs/cloudrun/deletesite) [Next\
\
getOrCreateBucket()](/docs/cloudrun/getorcreatebucket)

- [Example](#example)
- [Arguments](#arguments)
  - [`region`](#region)
- [Return value](#return-value)
  - [`sites`](#sites)
  - [`buckets`](#buckets)
- [See also](#see-also)
