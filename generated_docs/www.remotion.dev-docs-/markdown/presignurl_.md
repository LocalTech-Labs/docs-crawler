---
title: presignUrl()
source: Remotion Documentation
last_updated: 2024-12-22
---

# presignUrl()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- presignUrl()

On this page

# presignUrl()

Takes a private S3 object and turns it into a public URL by signing with your AWS credentials.

## Example [​](\#example "Direct link to Example")

```

ts

import { presignUrl } from "@remotion/lambda/client";

const url = await presignUrl({
  region: "us-east-1",
  bucketName: "remotionlambda-c7fsl3d",
  objectKey: "assets/sample.png",
  expiresInSeconds: 900,
  checkIfObjectExists: true,
});

console.log(url); // `string` - or `null` if object doesn't exist

const url2 = await presignUrl({
  region: "us-east-1",
  bucketName: "remotionlambda-c7fsl3d",
  objectKey: "assets/sample.png",
  expiresInSeconds: 900,
  checkIfObjectExists: false,
});

console.log(url); // always a string, or exception if object doesn't exist
```

```

ts

import { presignUrl } from "@remotion/lambda/client";

const url = await presignUrl({
  region: "us-east-1",
  bucketName: "remotionlambda-c7fsl3d",
  objectKey: "assets/sample.png",
  expiresInSeconds: 900,
  checkIfObjectExists: true,
});

console.log(url); // `string` - or `null` if object doesn't exist

const url2 = await presignUrl({
  region: "us-east-1",
  bucketName: "remotionlambda-c7fsl3d",
  objectKey: "assets/sample.png",
  expiresInSeconds: 900,
  checkIfObjectExists: false,
});

console.log(url); // always a string, or exception if object doesn't exist
```

note

Preferrably import this function from `@remotion/lambda/client` (available from v3.3.42) to avoid problems [inside serverless functions](/docs/lambda/light-client).

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `region` [​](\#region "Direct link to region")

The AWS region in which the bucket resides.

### `bucketName` [​](\#bucketname "Direct link to bucketname")

The bucket where the asset exists. The bucket must have been created by Remotion Lambda.

### `objectKey` [​](\#objectkey "Direct link to objectkey")

They key that uniquely identifies the object stored in the bucket.

### `expiresInSeconds` [​](\#expiresinseconds "Direct link to expiresinseconds")

The number of seconds before the presigned URL expires.
Must be an integer and `>=1` and `<=604800` (maximum of 7 days as enforced by AWS)

### `checkIfObjectExists` [​](\#checkifobjectexists "Direct link to checkifobjectexists")

Whether the function should check if the object exists in the bucket before generating the presigned url. Default `false`.

If the object does not exist and `checkIfObjectExists` is:

- `true`, then `null` is returned
- `false`, then an exception is thrown

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/presign-url.ts)
- [`downloadMedia()`](/docs/lambda/downloadmedia)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/presignurl.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getRenderProgress()](/docs/lambda/getrenderprogress) [Next\
\
renderMediaOnLambda()](/docs/lambda/rendermediaonlambda)

- [Example](#example)
- [Arguments](#arguments)
  - [`region`](#region)
  - [`bucketName`](#bucketname)
  - [`objectKey`](#objectkey)
  - [`expiresInSeconds`](#expiresinseconds)
  - [`checkIfObjectExists`](#checkifobjectexists)
- [See also](#see-also)
