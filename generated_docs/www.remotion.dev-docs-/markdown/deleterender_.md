---
title: deleteRender()
source: Remotion Documentation
last_updated: 2024-12-22
---

# deleteRender()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- deleteRender()

On this page

# deleteRender()

Deletes a rendered video, audio or still and its associated metada.

```

ts

import { deleteRender } from "@remotion/lambda";

const { freedBytes } = await deleteRender({
  bucketName: "remotionlambda-r42fs9fk",
  region: "us-east-1",
  renderId: "8hfxlw",
});

console.log(freedBytes); // 21249541
```

```

ts

import { deleteRender } from "@remotion/lambda";

const { freedBytes } = await deleteRender({
  bucketName: "remotionlambda-r42fs9fk",
  region: "us-east-1",
  renderId: "8hfxlw",
});

console.log(freedBytes); // 21249541
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `region` [​](\#region "Direct link to region")

The [AWS region](/docs/lambda/region-selection) in which the render has performed.

### `bucketName` [​](\#bucketname "Direct link to bucketname")

The bucket name in which the render was stored. This should be the same variable you used for [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) or [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda).

### `renderId` [​](\#renderid "Direct link to renderid")

The ID of the render. You can retrieve this ID by calling [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) or [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda).

### `customCredentials` [​](\#customcredentials "Direct link to customcredentials")

_optional, available from v3.2.23_

If the render was saved to a [different cloud](/docs/lambda/custom-destination#saving-to-another-cloud), pass an object with the same `endpoint`, `accessKeyId` and `secretAccessKey` as you passed to [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda#outname) or [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda#outname).

### `forcePathStyle?` [v4.0.202](https://github.com/remotion-dev/remotion/releases/v4.0.202) [​](\#forcepathstyle "Direct link to forcepathstyle")

Passes `forcePathStyle` to the AWS S3 client. If you don't know what this is, you probably don't need it.

## Return value [​](\#return-value "Direct link to Return value")

Returns a promise resolving to an object with the following properties:

### `freedBytes` [​](\#freedbytes "Direct link to freedbytes")

The amount of bytes that were removed from the S3 bucket.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/delete-render.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/deleterender.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
downloadMedia()](/docs/lambda/downloadmedia) [Next\
\
getUserPolicy()](/docs/lambda/getuserpolicy)

- [Arguments](#arguments)
  - [`region`](#region)
  - [`bucketName`](#bucketname)
  - [`renderId`](#renderid)
  - [`customCredentials`](#customcredentials)
  - [`forcePathStyle?`](#forcepathstyle)
- [Return value](#return-value)
  - [`freedBytes`](#freedbytes)
- [See also](#see-also)
