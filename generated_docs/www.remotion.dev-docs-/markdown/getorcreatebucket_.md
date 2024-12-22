---
title: getOrCreateBucket()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getOrCreateBucket()

- [Home page](/)
- [@remotion/cloudrun](/docs/cloudrun/api)
- getOrCreateBucket()

On this page

# getOrCreateBucket()

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

Creates a Cloud Storage bucket for Remotion Cloud Run in your GCP project. If one already exists, it will get returned instead.

**Only 1 bucket per region** is necessary for Remotion Cloud Run to function.

```

ts

import { getOrCreateBucket } from "@remotion/cloudrun";

const { bucketName, alreadyExisted } = await getOrCreateBucket({
  region: "us-east1",
});

console.log(bucketName); // "remotioncloudrun-32df3p"
```

```

ts

import { getOrCreateBucket } from "@remotion/cloudrun";

const { bucketName, alreadyExisted } = await getOrCreateBucket({
  region: "us-east1",
});

console.log(bucketName); // "remotioncloudrun-32df3p"
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `region` [​](\#region "Direct link to region")

The [GCP region](/docs/cloudrun/region-selection) which you want to create a bucket in.

### `updateBucketState?` [​](\#updatebucketstate "Direct link to updatebucketstate")

_optional_

Callback function that returns a state ( _string_) of operation. Used by the CLI to provide a progress update. State will be one of the following;

- Checking for existing bucket
- Creating new bucket
- Created bucket
- Using existing bucket

## Return value [​](\#return-value "Direct link to Return value")

A promise resolving to an object with the following properties:

### `bucketName` [​](\#bucketname "Direct link to bucketname")

The name of your bucket that was found or created.

### `alreadyExisted` [​](\#alreadyexisted "Direct link to alreadyexisted")

A boolean indicating whether the bucket already existed or was newly created.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/cloudrun/src/api/get-or-create-bucket.ts)
- [getServices()](/docs/cloudrun/getservices)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/getorcreatebucket.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getSites()](/docs/cloudrun/getsites) [Next\
\
renderMediaOnCloudrun()](/docs/cloudrun/rendermediaoncloudrun)

- [Arguments](#arguments)
  - [`region`](#region)
  - [`updateBucketState?`](#updatebucketstate)
- [Return value](#return-value)
  - [`bucketName`](#bucketname)
  - [`alreadyExisted`](#alreadyexisted)
- [See also](#see-also)
