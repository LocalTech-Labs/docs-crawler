---
title: Auto-delete renders
source: Remotion Documentation
last_updated: 2024-12-22
---

# Auto-delete renders

- [Home page](/)
- [Lambda](/docs/lambda)
- Auto-delete renders

On this page

# Auto-delete renders

_available from v4.0.32_

To automatically delete renders and associated files after some time, you need to:

[1](#1)

Enable the lifecycle rules on the AWS bucket.

[2](#2) Render a video with the `deleteAfter` option.

## Apply the lifecycle rules [​](\#apply-the-lifecycle-rules "Direct link to Apply the lifecycle rules")

[1](#1) **Ensure the right permission**

To put a lifecycle rule on the bucket, you need to have the `s3:PutLifecycleConfiguration` permission for your user.

- If you set up Remotion Lambda after 4.0.32, you have it automatically.
- If you set up Remotion Lambda previously, execute the following:
1. Upgrade to at least Remotion 4.0.32.
2. [Click here](https://us-east-1.console.aws.amazon.com/iamv2/home) to go to the users section in the AWS console.
3. Select your user.
4. Under "Permissions policies", click your policy (underlined in blue)
5. Choose the "JSON" mode for editing.
6. Under the `"Sid": "Storage"` section, add `"s3:PutLifecycleConfiguration"` to the `"Action"` array.
7. Save.

[2](#2) **Enable the lifecycle rules**

Redeploy the site with the `--enable-folder-expiry` option. This operation will modify the S3 bucket to apply [AWS Lifecycle Rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html).

- CLI
- Node.JS

```

In the command line
bash

npx remotion lambda sites create --site-name=my-site-name --enable-folder-expiry
```

```

In the command line
bash

npx remotion lambda sites create --site-name=my-site-name --enable-folder-expiry
```

```

deploy.mjs
ts

import { deploySite, getOrCreateBucket } from "@remotion/lambda";
import path from "path";

const { bucketName } = await getOrCreateBucket({
  region: "us-east-1",
  enableFolderExpiry: true,
});

const { serveUrl } = await deploySite({
  entryPoint: path.resolve(process.cwd(), "src/index.ts"),
  bucketName, // use the bucket with lifecyle rules
  region: "us-east-1",
});
console.log(serveUrl);
```

```

deploy.mjs
ts

import { deploySite, getOrCreateBucket } from "@remotion/lambda";
import path from "path";

const { bucketName } = await getOrCreateBucket({
  region: "us-east-1",
  enableFolderExpiry: true,
});

const { serveUrl } = await deploySite({
  entryPoint: path.resolve(process.cwd(), "src/index.ts"),
  bucketName, // use the bucket with lifecyle rules
  region: "us-east-1",
});
console.log(serveUrl);
```

Verify that it worked

In your S3 bucket, under "Management", you should see 4 lifecycle rules.

![](/img/lambda/applied-lc-rules.png)

## Trigger a render with expiration [​](\#trigger-a-render-with-expiration "Direct link to Trigger a render with expiration")

Valid values are `"1-day"`, `"3-days"`, `"7-days"` and `"30-days"`.

- CLI
- renderMediaOnLambda()
- renderStillOnLambda()

```

bash

npx remotion lambda render testbed-v6 react-svg --delete-after="1-day"
```

```

bash

npx remotion lambda render testbed-v6 react-svg --delete-after="1-day"
```

```

render.ts
tsx

import { renderMediaOnLambda } from "@remotion/lambda/client";

const { bucketName, renderId } = await renderMediaOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "MyVideo",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  codec: "h264",
  deleteAfter: "1-day",
});
```

```

render.ts
tsx

import { renderMediaOnLambda } from "@remotion/lambda/client";

const { bucketName, renderId } = await renderMediaOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "MyVideo",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  codec: "h264",
  deleteAfter: "1-day",
});
```

```

render.ts
tsx

import { renderStillOnLambda } from "@remotion/lambda/client";

const { bucketName, renderId } = await renderStillOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "MyVideo",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  inputProps: {},
  privacy: "public",
  imageFormat: "png",
  deleteAfter: "1-day",
});
```

```

render.ts
tsx

import { renderStillOnLambda } from "@remotion/lambda/client";

const { bucketName, renderId } = await renderStillOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "MyVideo",
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  inputProps: {},
  privacy: "public",
  imageFormat: "png",
  deleteAfter: "1-day",
});
```

Verify that it worked

All files should have an expiration rule and expiration date set when viewing the object properties.

![](/img/lambda/rendered-file-path.png)![](/img/lambda/rendered-file-management.png)

AWS does not delete the file at the exact time, but the deletion will happen.

## How it works [​](\#how-it-works "Direct link to How it works")

By applying the AWS Lifecycle rules, we are instructing AWS S3 to delete files based on their prefixes. When `deleteAfter` is defined with a value of `"1-day"`, the `renderId` will be prefixed with `1-day`, and the S3 key will start with `renders/1-day-*`, to which the deletion rule will be applied.

The basis of the deletion is based on the `Last modified date` of the file/folder.

`deleteAfter` value

Render Prefix

`1-day``renders/1-day``3-days``renders/3-days``7-days``renders/7-days``30-days``renders/30-days`

## See also [​](\#see-also "Direct link to See also")

- [AWS Expiring objects](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-expire-general-considerations.html)
- [`deploySite()`](/docs/lambda/deploysite)
- [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/autodelete.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Cost example](/docs/lambda/cost-example) [Next\
\
Debugging failures](/docs/lambda/troubleshooting/debug)

- [Apply the lifecycle rules](#apply-the-lifecycle-rules)
- [Trigger a render with expiration](#trigger-a-render-with-expiration)
- [How it works](#how-it-works)
- [See also](#see-also)
