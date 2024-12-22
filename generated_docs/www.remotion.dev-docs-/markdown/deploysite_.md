---
title: deploySite()
source: Remotion Documentation
last_updated: 2024-12-22
---

# deploySite()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- deploySite()

On this page

# deploySite()

Takes a Remotion project, bundles it and uploads it to an S3 bucket. Once uploaded, a Lambda function can render any composition in the Remotion project by specifying the URL.

- If you make changes locally, you need to redeploy the site. You can use [`siteName`](#sitename) to overwrite the previous site.
- Note that the Remotion project will be deployed to a subdirectory, not the root of the domain. Therefore you must ensure that if you have specified paths in your Remotion project, they are able to handle this scenario.
- Before calling this function, you should create a bucket, see [`getOrCreateBucket()`](/docs/lambda/getorcreatebucket).

## Example [​](\#example "Direct link to Example")

```

ts

import { deploySite } from "@remotion/lambda";
import path from "path";

const { serveUrl } = await deploySite({
  entryPoint: path.resolve(process.cwd(), "src/index.ts"),
  bucketName: "remotionlambda-c7fsl3d",
  region: "us-east-1",
  options: {
    onBundleProgress: (progress) => {
      // Progress is between 0 and 100
      console.log(`Bundle progress: ${progress}%`);
    },
    onUploadProgress: ({
      totalFiles,
      filesUploaded,
      totalSize,
      sizeUploaded,
    }) => {
      console.log(
        `Upload progress: Total files ${totalFiles}, Files uploaded ${filesUploaded}, Total size ${totalSize}, Size uploaded ${sizeUploaded}`,
      );
    },
  },
});
console.log(serveUrl);
```

```

ts

import { deploySite } from "@remotion/lambda";
import path from "path";

const { serveUrl } = await deploySite({
  entryPoint: path.resolve(process.cwd(), "src/index.ts"),
  bucketName: "remotionlambda-c7fsl3d",
  region: "us-east-1",
  options: {
    onBundleProgress: (progress) => {
      // Progress is between 0 and 100
      console.log(`Bundle progress: ${progress}%`);
    },
    onUploadProgress: ({
      totalFiles,
      filesUploaded,
      totalSize,
      sizeUploaded,
    }) => {
      console.log(
        `Upload progress: Total files ${totalFiles}, Files uploaded ${filesUploaded}, Total size ${totalSize}, Size uploaded ${sizeUploaded}`,
      );
    },
  },
});
console.log(serveUrl);
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `entryPoint` [​](\#entrypoint "Direct link to entrypoint")

An absolute path pointing to the entry point of your Remotion project. [Usually the entry point in your Remotion project is stored at `src/entry.tsx`](/docs/terminology/entry-point).

### `bucketName` [​](\#bucketname "Direct link to bucketname")

The bucket to where the website will be deployed. The bucket must have been created by Remotion Lambda.

### `region` [​](\#region "Direct link to region")

The AWS region in which the bucket resides.

### `siteName?` [​](\#sitename "Direct link to sitename")

_optional_

Specify the subfolder in your S3 bucket that you want the site to deploy to. If you omit this property, a new subfolder with a random name will be created. If a site already exists with the name you passed, it will be overwritten. Can only contain the following characters: `0-9`, `a-z`, `A-Z`, `-`, `!`, `_`, `.`, `*`, `'`, `(`, `)`

### `logLevel?` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#loglevel "Direct link to loglevel")

_optional_

One of `verbose`, `info`, `warn`, `error`.

Determines how much is being logged to the console.

`verbose` will also log `console.log`'s from the browser.

Default `info`.

### `options?` [​](\#options "Direct link to options")

_optional_

An object with the following properties:

#### `onBundleProgress` [​](\#onbundleprogress "Direct link to onbundleprogress")

Callback from Webpack when the bundling has progressed. Passes a number between 0 and 100 to the callback, see example at the top of the page.

#### `onUploadProgress` [​](\#onuploadprogress "Direct link to onuploadprogress")

Callback function that gets called when uploading of the assets has progressed. Passes an object with the following properties to the callback:

- `totalFiles` ( _number_): Total number of files in the bundle.
- `filesUploaded` ( _number_): Number of files that have been fully uploaded so far.
- `totalSize` ( _number_): Total size in bytes of all the files in the bundle.
- `sizeUploaded` ( _number_): Amount of bytes uploaded so far.

#### `webpackOverride` [​](\#webpackoverride "Direct link to webpackoverride")

Allows to pass a custom webpack override. See [`bundle()` -\> webpackOverride](/docs/bundle#webpackoverride) for more information.

#### `enableCaching` [​](\#enablecaching "Direct link to enablecaching")

Whether webpack caching should be enabled. Default true. See [`bundle()` -\> enableCaching](/docs/bundle#enablecaching) for more information.

#### `publicDir` [​](\#publicdir "Direct link to publicdir")

_available from v3.2.17_

Set the directory in which the files that can be loaded using [`staticFile()`](/docs/staticfile) are located. By default it is the folder `public/` located in the [Remotion Root](/docs/terminology/remotion-root) folder. If you pass a relative path, it will be resolved against the [Remotion Root](/docs/terminology/remotion-root).

#### `rootDir` [​](\#rootdir "Direct link to rootdir")

_available from v3.2.17_

The directory in which the Remotion project is rooted in. This should be set to the directory that contains the `package.json` which installs Remotion. By default, it is the current working directory.

note

The current working directory is the directory from which your program gets executed from. It is not the same as the file where bundle() gets called.

#### `ignoreRegisterRootWarning` [​](\#ignoreregisterrootwarning "Direct link to ignoreregisterrootwarning")

_available from v3.3.55_

Ignore an error that gets thrown if you pass an entry point file which does not contain `registerRoot`.

### `privacy` [​](\#privacy "Direct link to privacy")

_available from v3.3.97_

Either `public` (default) or `no-acl` if you are not using ACL. Sites must have a public URL to be able to be rendered on Lambda, since the headless browser opens that URL.

### `throwIfSiteExists?` [v4.0.141](https://github.com/remotion-dev/remotion/releases/v4.0.141) [​](\#throwifsiteexists "Direct link to throwifsiteexists")

Prevents accidential update of an existing site. If there are any files in the subfolder where the site should be placed, the function will throw.

### `forcePathStyle?` [v4.0.202](https://github.com/remotion-dev/remotion/releases/v4.0.202) [​](\#forcepathstyle "Direct link to forcepathstyle")

Passes `forcePathStyle` to the AWS S3 client. If you don't know what this is, you probably don't need it.

## Return value [​](\#return-value "Direct link to Return value")

An object with the following values:

- `serveUrl` _(string)_: An URL such as `https://remotionlambda-12345.s3.eu-central-1.amazonaws.com/sites/abcdef/index.html`.

You can use this "Serve URL" to render a video on Remotion Lambda using:

  - The [`npx remotion lambda render`](/docs/lambda/cli/render) and [`npx remotion lambda still`](/docs/lambda/cli/still) commands
  - The [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) and [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda) functions.
  - Locally using the [`renderMedia()`](/docs/renderer/render-media) and [`renderStill()`](/docs/renderer/render-still) functions.
  - Locally using the [`npx remotion render`](/docs/cli) and [`npx remotion still`](/docs/cli) commands

If you are rendering on Lambda, you can also pass the site name (in this case `abcdef`) as an abbreviation.

- `siteName` _(string)_: The identifier of the site that was given. Is either the site name that you have passed into this function, or a random string that was generated if you didn't pass a site name.

- `stats`: ( _available from v3.3.7_) An object with 3 entries: `uploadedFiles`, `deletedFiles` and `untouchedFiles`. Each one is a `number`.

## Changelog [​](\#changelog "Direct link to Changelog")

From `v3.3.7`, this function is incremental: It only compares the contents of the local files and the files on S3 and only executes the necessary operations.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/deploy-site.ts)
- [CLI equivalent: `npx remotion lambda sites create`](/docs/lambda/cli/sites#create)
- [getSites()](/docs/lambda/getsites)
- [deleteSite()](/docs/lambda/deletesite)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/deploysite.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
deleteSite()](/docs/lambda/deletesite) [Next\
\
getAwsClient()](/docs/lambda/getawsclient)

- [Example](#example)
- [Arguments](#arguments)
  - [`entryPoint`](#entrypoint)
  - [`bucketName`](#bucketname)
  - [`region`](#region)
  - [`siteName?`](#sitename)
  - [`logLevel?`](#loglevel)
  - [`options?`](#options)
  - [`privacy`](#privacy)
  - [`throwIfSiteExists?`](#throwifsiteexists)
  - [`forcePathStyle?`](#forcepathstyle)
- [Return value](#return-value)
- [Changelog](#changelog)
- [See also](#see-also)
