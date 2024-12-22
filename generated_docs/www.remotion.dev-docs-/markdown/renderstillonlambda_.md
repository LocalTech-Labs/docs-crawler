---
title: renderStillOnLambda()
source: Remotion Documentation
last_updated: 2024-12-22
---

# renderStillOnLambda()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- renderStillOnLambda()

On this page

# renderStillOnLambda()

Renders a still image inside a lambda function and writes it to the specified output location.

If you want to render a video or audio instead, use [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) instead.

If you want to render a still locally instead, use [`renderStill()`](/docs/renderer/render-still) instead.

## Example [​](\#example "Direct link to Example")

```

tsx

import { renderStillOnLambda } from "@remotion/lambda/client";

const { estimatedPrice, url, sizeInBytes } = await renderStillOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  composition: "MyVideo",
  inputProps: {},
  imageFormat: "png",
  maxRetries: 1,
  privacy: "public",
  envVariables: {},
  frame: 10,
});
```

```

tsx

import { renderStillOnLambda } from "@remotion/lambda/client";

const { estimatedPrice, url, sizeInBytes } = await renderStillOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  composition: "MyVideo",
  inputProps: {},
  imageFormat: "png",
  maxRetries: 1,
  privacy: "public",
  envVariables: {},
  frame: 10,
});
```

note

Preferrably import this function from `@remotion/lambda/client` to avoid problems [inside serverless functions](/docs/lambda/light-client).

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `region` [​](\#region "Direct link to region")

In which region your Lambda function is deployed. It's highly recommended that your Remotion site is also in the same region.

### `functionName` [​](\#functionname "Direct link to functionname")

The name of the deployed Lambda function.
Use [`deployFunction()`](/docs/lambda/deployfunction) to create a new function and [`getFunctions()`](/docs/lambda/getfunctions) to obtain currently deployed Lambdas.

### `serveUrl` [​](\#serveurl "Direct link to serveurl")

A URL pointing to a Remotion project. Use [`deploySite()`](/docs/lambda/deploysite) to deploy a Remotion project.

### `composition` [​](\#composition "Direct link to composition")

The `id` of the [composition](/docs/composition) you want to render..

### `inputProps` [​](\#inputprops "Direct link to inputprops")

[Input Props to pass to the selected composition of your video.](/docs/passing-props#passing-input-props-in-the-cli).

Must be a JSON object.

From the root component the props can be read using [`getInputProps()`](/docs/get-input-props).

You may transform input props using [`calculateMetadata()`](/docs/calculate-metadata).

### `privacy` [​](\#privacy "Direct link to privacy")

One of:

- `"public"` ( _default_): The rendered still is publicly accessible under the S3 URL.
- `"private"`: The rendered still is not publicly available, but signed links can be created using [presignUrl()](/docs/lambda/presignurl).
- `"no-acl"` ( _available from v.3.1.7_): The ACL option is not being set at all, this option is useful if you are writing to another bucket that does not support ACL using [`outName`](#outname).

### `frame?` [​](\#frame "Direct link to frame")

_optional - default `0`_

Which frame of the composition should be rendered. Frames are zero-indexed.

From v3.2.27, negative values are allowed, with `-1` being the last frame.

### `imageFormat?` [​](\#imageformat "Direct link to imageformat")

_optional - default `"png"`_

See [`renderStill() -> imageFormat`](/docs/renderer/render-still#imageformat).

### `onInit?` [v4.0.6](https://github.com/remotion-dev/remotion/releases/v4.0.6) [​](\#oninit "Direct link to oninit")

A callback function that gets called when the render starts, useful to obtain the link to the logs even if the render fails.

It receives an object with the following properties:

- `renderId`: The ID of the render.
- `cloudWatchLogs`: A link to the CloudWatch logs of the Lambda function, if you did not disable it.
- `lambdaInsightsUrl` [v4.0.61](https://github.com/remotion-dev/remotion/releases/v4.0.61): A link to the [Lambda insights](/docs/lambda/insights), if you enabled it.

Example usage:

```

tsx

import {
  renderStillOnLambda,
  RenderStillOnLambdaInput,
} from "@remotion/lambda/client";

const otherParameters: RenderStillOnLambdaInput = {
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  composition: "MyVideo",
  inputProps: {},
  imageFormat: "png",
  maxRetries: 1,
  privacy: "public",
  envVariables: {},
  frame: 10,
};
await renderStillOnLambda({
  ...otherParameters,
  onInit: ({ cloudWatchLogs, renderId, lambdaInsightsUrl }) => {
    console.log(console.log(`Render invoked with ID = ${renderId}`));
    console.log(`CloudWatch logs (if enabled): ${cloudWatchLogs}`);
    console.log(`Lambda Insights (if enabled): ${lambdaInsightsUrl}`);
  },
});
```

```

tsx

import {
  renderStillOnLambda,
  RenderStillOnLambdaInput,
} from "@remotion/lambda/client";

const otherParameters: RenderStillOnLambdaInput = {
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  composition: "MyVideo",
  inputProps: {},
  imageFormat: "png",
  maxRetries: 1,
  privacy: "public",
  envVariables: {},
  frame: 10,
};
await renderStillOnLambda({
  ...otherParameters,
  onInit: ({ cloudWatchLogs, renderId, lambdaInsightsUrl }) => {
    console.log(console.log(`Render invoked with ID = ${renderId}`));
    console.log(`CloudWatch logs (if enabled): ${cloudWatchLogs}`);
    console.log(`Lambda Insights (if enabled): ${lambdaInsightsUrl}`);
  },
});
```

### `jpegQuality?` [​](\#jpegquality "Direct link to jpegquality")

_optional_

Sets the quality of the generated JPEG images. Must be an integer between 0 and 100. Default is to leave it up to the browser, [current default is 80](https://github.com/chromium/chromium/blob/99314be8152e688bafbbf9a615536bdbb289ea87/headless/lib/browser/protocol/headless_handler.cc#L32).

Only applies if `imageFormat` is `"jpeg"`, otherwise this option is invalid.

### `quality?` [​](\#quality "Direct link to quality")

Renamed to `jpegQuality` in `v4.0.0`.

### `maxRetries?` [​](\#maxretries "Direct link to maxretries")

_optional - default `1`_

How often a frame render may be retried until it fails.

note

A retry only gets executed if a the error is in the [list of flaky errors](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/shared/is-flaky-error.ts).

### `envVariables?` [​](\#envvariables "Direct link to envvariables")

_optional - default `{}`_

See [`renderStill() -> envVariables`](/docs/renderer/render-still#envvariables).

### `forceHeight?` [​](\#forceheight "Direct link to forceheight")

_optional, available from v3.2.40_

Overrides the default composition height.

### `forceWidth?` [​](\#forcewidth "Direct link to forcewidth")

_optional, available from v3.2.40_

Overrides the default composition width.

### `scale?` [​](\#scale "Direct link to scale")

_optional_

Scales the output dimensions by a factor. See [Scaling](/docs/scaling) to learn more about this feature.

### `outName?` [​](\#outname "Direct link to outname")

_optional_

It can either be:

- `undefined` \- it will default to `out` plus the appropriate file extension, for example: `renders/${renderId}/out.mp4`.
- A `string` \- it will get saved to the same S3 bucket as your site under the key `renders/{renderId}/{outName}`. Make sure to include the file extension at the end of the string.
- An object if you want to render to a different bucket or cloud provider - [see here for detailed instructions](/docs/lambda/custom-destination).

### `timeoutInMilliseconds?` [​](\#timeoutinmilliseconds "Direct link to timeoutinmilliseconds")

_optional_

A number describing how long the render may take to resolve all [`delayRender()`](/docs/delay-render) calls [before it times out](/docs/timeout). Default: `30000`

### `downloadBehavior?` [​](\#downloadbehavior "Direct link to downloadbehavior")

_optional, available since v3.1.5_

How the output file should behave when accessed through the S3 output link in the browser.

Either:

- `{"type": "play-in-browser"}` \- the default. The video will play in the browser.
- `{"type": "download", fileName: null}` or `{"type": "download", fileName: "download.mp4"}` \- a `Content-Disposition` header will be added which makes the browser download the file. You can optionally override the filename.

### `offthreadVideoCacheSizeInBytes?` [v4.0.23](https://github.com/remotion-dev/remotion/releases/v4.0.23) [​](\#offthreadvideocachesizeinbytes "Direct link to offthreadvideocachesizeinbytes")

From v4.0, Remotion has a cache for [`<OffthreadVideo>`](https://remotion.dev/docs/offthreadvideo) frames. The default is `null`, corresponding to half of the system memory available when the render starts.

This option allows to override the size of the cache. The higher it is, the faster the render will be, but the more memory will be used.

The used value will be printed when running in verbose mode.

Default: `null`

### `deleteAfter?` [v4.0.32](https://github.com/remotion-dev/remotion/releases/v4.0.32) [​](\#deleteafter "Direct link to deleteafter")

Automatically delete the render after a certain period. Accepted values are `1-day`, `3-days`, `7-days` and `30-days`.

For this to work, your bucket needs to have [lifecycles enabled](/docs/lambda/autodelete).

### `chromiumOptions?` [​](\#chromiumoptions "Direct link to chromiumoptions")

_optional_

Allows you to set certain Chromium / Google Chrome flags. See: [Chromium flags](/docs/chromium-flags).

#### `disableWebSecurity` [​](\#disablewebsecurity "Direct link to disablewebsecurity")

_boolean - default `false`_

This will most notably disable CORS among other security features.

#### `ignoreCertificateErrors` [​](\#ignorecertificateerrors "Direct link to ignorecertificateerrors")

_boolean - default `false`_

Results in invalid SSL certificates, such as self-signed ones, being ignored.

#### `gl` [​](\#gl "Direct link to gl")

Changelog

- From Remotion v2.6.7 until v3.0.7, the default for Remotion Lambda was `swiftshader`, but from v3.0.8 the default is `swangle` (Swiftshader on Angle) since Chrome 101 added support for it.
- From Remotion v2.4.3 until v2.6.6, the default was `angle`, however it turns out to have a small memory leak that could crash long Remotion renders.

Select the OpenGL renderer backend for Chromium.

Accepted values:

- `"angle"`
- `"egl"`
- `"swiftshader"`
- `"swangle"`
- `"vulkan"` ( _from Remotion v4.0.41_)
- `"angle-egl"` ( _from Remotion v4.0.51_)

The default is `null`, letting Chrome decide, except on Lambda where the default is `"swangle"`

#### `userAgent` [v3.3.83](https://github.com/remotion-dev/remotion/releases/v3.3.83) [​](\#useragent "Direct link to useragent")

Lets you set a custom user agent that the headless Chrome browser assumes.

### `forceBucketName?` [​](\#forcebucketname "Direct link to forcebucketname")

_optional_

Specify a specific bucket name to be used. [This is not recommended](/docs/lambda/multiple-buckets), instead let Remotion discover the right bucket automatically.

### `logLevel?` [​](\#loglevel "Direct link to loglevel")

_optional_

One of `verbose`, `info`, `warn`, `error`.

Determines how much is being logged to the console.

`verbose` will also log `console.log`'s from the browser.

Default `info`.

Logs can be read through the CloudWatch URL that this function returns.

### `forcePathStyle?` [v4.0.202](https://github.com/remotion-dev/remotion/releases/v4.0.202) [​](\#forcepathstyle "Direct link to forcepathstyle")

Passes `forcePathStyle` to the AWS S3 client. If you don't know what this is, you probably don't need it.

### `dumpBrowserLogs?` [​](\#dumpbrowserlogs "Direct link to dumpbrowserlogs")

_optional - default `false`, deprecated in v4.0_

Deprecated in favor of [`logLevel`](#loglevel).

## Return value [​](\#return-value "Direct link to Return value")

Returns a promise resolving to an object with the following properties:

### `bucketName` [​](\#bucketname "Direct link to bucketname")

The S3 bucket in which the video was saved.

### `url` [​](\#url "Direct link to url")

An AWS S3 URL where the output is available.

### `outKey` [v4.0.141](https://github.com/remotion-dev/remotion/releases/v4.0.141) [​](\#outkey "Direct link to outkey")

The S3 key where the output is saved.

### `estimatedPrice` [​](\#estimatedprice "Direct link to estimatedprice")

Object containing roughly estimated information about how expensive this operation was.

### `sizeInBytes` [​](\#sizeinbytes "Direct link to sizeinbytes")

The size of the output image in bytes.

### `renderId` [​](\#renderid "Direct link to renderid")

A unique alphanumeric identifier for this render. Useful for obtaining status and finding the relevant files in the S3 bucket.

### `cloudWatchLogs` [​](\#cloudwatchlogs "Direct link to cloudwatchlogs")

_Available from v3.2.10_

A link to CloudWatch (if you haven't disabled it) that you can visit to see the logs for the render.

### `artifacts` [v4.0.176](https://github.com/remotion-dev/remotion/releases/v4.0.176) [​](\#artifacts "Direct link to artifacts")

Artifacts that were created so far during the render. [See here for an example of dealing with field.](/docs/artifacts#using-renderstillonlambda)

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/render-still-on-lambda.ts)
- [renderMediaOnLambda()](/docs/lambda/rendermediaonlambda)
- [renderStill()](/docs/renderer/render-still)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/renderstillonlambda.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
renderMediaOnLambda()](/docs/lambda/rendermediaonlambda) [Next\
\
simulatePermissions()](/docs/lambda/simulatepermissions)

- [Example](#example)
- [Arguments](#arguments)
  - [`region`](#region)
  - [`functionName`](#functionname)
  - [`serveUrl`](#serveurl)
  - [`composition`](#composition)
  - [`inputProps`](#inputprops)
  - [`privacy`](#privacy)
  - [`frame?`](#frame)
  - [`imageFormat?`](#imageformat)
  - [`onInit?`](#oninit)
  - [`jpegQuality?`](#jpegquality)
  - [`quality?`](#quality)
  - [`maxRetries?`](#maxretries)
  - [`envVariables?`](#envvariables)
  - [`forceHeight?`](#forceheight)
  - [`forceWidth?`](#forcewidth)
  - [`scale?`](#scale)
  - [`outName?`](#outname)
  - [`timeoutInMilliseconds?`](#timeoutinmilliseconds)
  - [`downloadBehavior?`](#downloadbehavior)
  - [`offthreadVideoCacheSizeInBytes?`](#offthreadvideocachesizeinbytes)
  - [`deleteAfter?`](#deleteafter)
  - [`chromiumOptions?`](#chromiumoptions)
  - [`forceBucketName?`](#forcebucketname)
  - [`logLevel?`](#loglevel)
  - [`forcePathStyle?`](#forcepathstyle)
  - [`dumpBrowserLogs?`](#dumpbrowserlogs)
- [Return value](#return-value)
  - [`bucketName`](#bucketname)
  - [`url`](#url)
  - [`outKey`](#outkey)
  - [`estimatedPrice`](#estimatedprice)
  - [`sizeInBytes`](#sizeinbytes)
  - [`renderId`](#renderid)
  - [`cloudWatchLogs`](#cloudwatchlogs)
  - [`artifacts`](#artifacts)
- [See also](#see-also)
