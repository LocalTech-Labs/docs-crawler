---
title: downloadMedia()
source: Remotion Documentation
last_updated: 2024-12-22
---

# downloadMedia()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- downloadMedia()

On this page

# downloadMedia()

Downloads a rendered video, audio or still to the disk of the machine this API is called from.

If you want to let the user download a result to their machine, use [`renderMediaOnLambda()` -\> `downloadBehavior`](/docs/lambda/rendermediaonlambda#downloadbehavior) instead.

```

ts

import { downloadMedia } from "@remotion/lambda";

const { outputPath, sizeInBytes } = await downloadMedia({
  bucketName: "remotionlambda-r42fs9fk",
  region: "us-east-1",
  renderId: "8hfxlw",
  outPath: "out.mp4",
  onProgress: ({ totalSize, downloaded, percent }) => {
    console.log(
      `Download progress: ${totalSize}/${downloaded} bytes (${(
        percent * 100
      ).toFixed(0)}%)`
    );
  },
});

console.log(outputPath); // "/Users/yourname/remotion-project/out.mp4"
console.log(sizeInBytes); // 21249541
```

```

ts

import { downloadMedia } from "@remotion/lambda";

const { outputPath, sizeInBytes } = await downloadMedia({
  bucketName: "remotionlambda-r42fs9fk",
  region: "us-east-1",
  renderId: "8hfxlw",
  outPath: "out.mp4",
  onProgress: ({ totalSize, downloaded, percent }) => {
    console.log(
      `Download progress: ${totalSize}/${downloaded} bytes (${(
        percent * 100
      ).toFixed(0)}%)`
    );
  },
});

console.log(outputPath); // "/Users/yourname/remotion-project/out.mp4"
console.log(sizeInBytes); // 21249541
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `region` [​](\#region "Direct link to region")

The [AWS region](/docs/lambda/region-selection) in which the render has performed.

### `bucketName` [​](\#bucketname "Direct link to bucketname")

The bucket name in which the render was stored. This should be the same variable you used for [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) or [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda).

### `renderId` [​](\#renderid "Direct link to renderid")

The ID of the render. You can retrieve this ID by calling [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) or [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda).

### `outPath` [​](\#outpath "Direct link to outpath")

Where the video should be saved. Pass an absolute path, or it will be resolved relative to your current working directory.

### `onProgress` [​](\#onprogress "Direct link to onprogress")

_optional_

Callback function that gets called with the following properties:

- `totalSize` in bytes
- `downloaded` number of bytes downloaded
- `percent` relative progress between 0 and 1

### `customCredentials` [​](\#customcredentials "Direct link to customcredentials")

_optional, available from v3.2.23_

If the render was saved to a [different cloud](/docs/lambda/custom-destination#saving-to-another-cloud), pass an object with the same `endpoint`, `accessKeyId` and `secretAccessKey` as you passed to [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda#outname) or [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda#outname).

## Return value [​](\#return-value "Direct link to Return value")

Returns a promise resolving to an object with the following properties:

### `outputPath` [​](\#outputpath "Direct link to outputpath")

The absolute path of where the file got saved.

### `sizeInBytes` [​](\#sizeinbytes "Direct link to sizeinbytes")

The size of the file in bytes.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/download-media.ts)
- [renderMediaOnLambda()](/docs/lambda/rendermediaonlambda)
- [renderStillOnLambda()](/docs/lambda/renderstillonlambda)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/downloadmedia.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getSites()](/docs/lambda/getsites) [Next\
\
deleteRender()](/docs/lambda/deleterender)

- [Arguments](#arguments)
  - [`region`](#region)
  - [`bucketName`](#bucketname)
  - [`renderId`](#renderid)
  - [`outPath`](#outpath)
  - [`onProgress`](#onprogress)
  - [`customCredentials`](#customcredentials)
- [Return value](#return-value)
  - [`outputPath`](#outputpath)
  - [`sizeInBytes`](#sizeinbytes)
- [See also](#see-also)
