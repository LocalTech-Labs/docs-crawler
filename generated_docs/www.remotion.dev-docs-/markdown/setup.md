---
title: Setup
source: Remotion Documentation
last_updated: 2024-12-22
---

# Setup

- [Home page](/)
- [Lambda](/docs/lambda)
- Setup

On this page

# Setup

[![](https://i.ytimg.com/vi/kFVd3KnfwYY/hqdefault.jpg?sqp=-oaymwEbCKgBEF5IVfKriqkDDggBFQAAiEIYAXABwAEG&rs=AOn4CLCJOz2mu_A24NvjFQGm6GjUVssnUQ)\
\
Also available as a 18min video\
\
**How to set up Remotion Lambda**](https://www.youtube.com/watch?v=kFVd3KnfwYY)

## 1\. Install `@remotion/lambda` [​](\#1-install-remotionlambda "Direct link to 1-install-remotionlambda")

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/lambda@4.0.242Copy
```

```

npm i --save-exact @remotion/lambda@4.0.242Copy
```

```

pnpm i @remotion/lambda@4.0.242Copy
```

```

pnpm i @remotion/lambda@4.0.242Copy
```

```

bun i @remotion/lambda@4.0.242Copy
```

```

bun i @remotion/lambda@4.0.242Copy
```

```

yarn --exact add @remotion/lambda@4.0.242Copy
```

```

yarn --exact add @remotion/lambda@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

## 2\. Create role policy [​](\#2-create-role-policy "Direct link to 2. Create role policy")

- Go to [AWS account IAM Policies section](https://console.aws.amazon.com/iamv2/home?#/policies)
- Click on "Create policy"
- Click on JSON
- In your project, type `npx remotion lambda policies role` in the command line and copy it into the "JSON" field on AWS.
- Click next. On the tags page, you don't need to fill in anything. Click next again.
- Give the policy **exactly** the name `remotion-lambda-policy`. The other fields can be left as they are.

## 3\. Create a role [​](\#3-create-a-role "Direct link to 3. Create a role")

- Go to [AWS account IAM Roles section](https://console.aws.amazon.com/iamv2/home#/roles)
- Click "Create role".
- Under "Use cases", select "Lambda". Click next.
- Under "Permissions policies", filter for `remotion-lambda-policy` and click the checkbox to assign this policy. Click next.
- In the final step, name the role `remotion-lambda-role` **exactly**. You can leave the other fields as is.
- Click "Create role" to confirm.

## 4\. Create a user [​](\#4-create-a-user "Direct link to 4. Create a user")

- Go to [AWS account IAM Users section](https://console.aws.amazon.com/iamv2/home#/users)
- Click `Add users`
- Enter any username, such as `remotion-user`.
- **Don't check** the "Enable console access" option. You don't need it.
- Click "Next".
- Click "Next" again without changing any settings. You should now be on the "Review and Create" step.
- Click "Create user".

## 5\. Create an access key for the user [​](\#5-create-an-access-key-for-the-user "Direct link to 5. Create an access key for the user")

- Go to [AWS account IAM Users section](https://console.aws.amazon.com/iamv2/home#/users)
- Click on the name of the user that was created in step 4.
- Navigate to the "Security Credentials" tab, and scroll down to the "Access Keys" section.
- Click the "Create access key" button.
- Select "Application running on an AWS compute service".
- Ignore warnings that might appear and check the "I understand the recommendations..." checkbox.
- Click "Next".
- Click "Create access key".
- Add a `.env` file to your project's root and add the credentials you just copied in the following format:

```

.env
txt

REMOTION_AWS_ACCESS_KEY_ID=<Access key ID>
REMOTION_AWS_SECRET_ACCESS_KEY=<Secret access key>
```

```

.env
txt

REMOTION_AWS_ACCESS_KEY_ID=<Access key ID>
REMOTION_AWS_SECRET_ACCESS_KEY=<Secret access key>
```

## 6\. Add permissions to your user [​](\#6-add-permissions-to-your-user "Direct link to 6. Add permissions to your user")

- Go to [AWS account IAM Users section](https://console.aws.amazon.com/iamv2/home#/users)
- Select the user you just created.
- Click "Add inline policy" under the "Add Permissions" dropdown in the "Permissions policies" panel.
- Click the tab "JSON".
- Enter in your terminal: `npx remotion lambda policies user` and copy into the AWS text field what gets printed.
- Click "Review policy".
- Give the policy a name. For example `remotion-user-policy`, but it can be anything.
- Click "Create policy" to confirm.

## 7\. Optional: Validate the permission setup [​](\#7-optional-validate-the-permission-setup "Direct link to 7. Optional: Validate the permission setup")

Check all user permissions and validate them against the AWS Policy simulator by executing the following command:

```

bash

npx remotion lambda policies validate
```

```

bash

npx remotion lambda policies validate
```

* * *

For the following steps, you may execute them on the CLI, or programmatically using the Node.JS APIs.

## 8\. Deploy a function [​](\#8-deploy-a-function "Direct link to 8. Deploy a function")

- CLI
- Node.JS

Deploy a function that can render videos into your AWS account by executing the following command:

```

bash

npx remotion lambda functions deploy
```

```

bash

npx remotion lambda functions deploy
```

You can deploy a function that can render videos into your AWS account using [`deployFunction()`](/docs/lambda/deployfunction).

```

ts

const {functionName} = await deployFunction({
  region: 'us-east-1',
  timeoutInSeconds: 120,
  memorySizeInMb: 2048,
  createCloudWatchLogGroup: true,
});
```

```

ts

const {functionName} = await deployFunction({
  region: 'us-east-1',
  timeoutInSeconds: 120,
  memorySizeInMb: 2048,
  createCloudWatchLogGroup: true,
});
```

The function name is returned which you'll need for rendering.

The function consists of necessary binaries and JavaScript code that can take a [serve URL](/docs/terminology/serve-url) and make renders from it. A function is bound to the Remotion version, if you upgrade Remotion, you [need to deploy a new function](/docs/lambda/upgrading). A function does not include your Remotion code, it will be deployed in the next step instead.

## 9\. Deploy a site [​](\#9-deploy-a-site "Direct link to 9. Deploy a site")

- CLI
- Node.JS

Run the following command to deploy your Remotion project to an S3 bucket. Pass as the last argument the [entry point](/docs/terminology/entry-point) of the project.

```

bash

npx remotion lambda sites create src/index.ts --site-name=my-video
```

```

bash

npx remotion lambda sites create src/index.ts --site-name=my-video
```

A [Serve URL](/docs/terminology/serve-url) will be printed pointing to the deployed project.

When you update your Remotion code in the future, redeploy your site. Pass the same [`--site-name`](/docs/lambda/cli/sites#--site-name) to overwrite the previous deploy. If you don't pass [`--site-name`](/docs/lambda/cli/sites#--site-name), a unique URL will be generated on every deploy.

First, you need to create an S3 bucket in your preferred region. If one already exists, it will be used instead:

```

ts

import path from 'path';
import {deploySite, getOrCreateBucket} from '@remotion/lambda';

const {bucketName} = await getOrCreateBucket({
  region: 'us-east-1',
});
```

```

ts

import path from 'path';
import {deploySite, getOrCreateBucket} from '@remotion/lambda';

const {bucketName} = await getOrCreateBucket({
  region: 'us-east-1',
});
```

Next, upload your Remotion project to an S3 bucket. Specify the entry point of your Remotion project, this is the file where [`registerRoot()`](/docs/register-root) is called.

```

ts

const {serveUrl} = await deploySite({
  bucketName,
  entryPoint: path.resolve(process.cwd(), 'src/index.ts'),
  region: 'us-east-1',
  siteName: 'my-video',
});
```

```

ts

const {serveUrl} = await deploySite({
  bucketName,
  entryPoint: path.resolve(process.cwd(), 'src/index.ts'),
  region: 'us-east-1',
  siteName: 'my-video',
});
```

When you update your Remotion code in the future, redeploy your site. Pass the same [`siteName`](/docs/lambda/deploysite#sitename) to overwrite the previous deploy. If you don't pass [`siteName`](/docs/lambda/deploysite#sitename), a unique URL will be generated on every deploy.

## 10\. Check AWS concurrency limit [​](\#10-check-aws-concurrency-limit "Direct link to 10. Check AWS concurrency limit")

Check the concurrency limit that AWS has given to your account:

```

npx remotion lambda quotas
```

```

npx remotion lambda quotas
```

By default, it is `1000` concurrent invocations per region. However, new accounts might have a limit [as low as `10`](/docs/lambda/troubleshooting/rate-limit#new-accounts-using-aws-lambda). Each Remotion render may use as much as 200 functions per render concurrently, so if your assigned limit is very low, [you might want to request an increase right away](/docs/lambda/troubleshooting/rate-limit#request-an-increase).

## 11\. Render a video [​](\#11-render-a-video "Direct link to 11. Render a video")

- CLI
- Node.JS

Take the URL you received from the step 9 - your "serve URL" - and run the following command. Also pass in the [ID of the composition](/docs/composition) you'd like to render.

```

bash

npx remotion lambda render <serve-url> <composition-id>
```

```

bash

npx remotion lambda render <serve-url> <composition-id>
```

Progress will be printed until the video finished rendering. Congrats! You rendered your first video using Remotion Lambda 🚀

You already have the function name from a previous step. But since you only need to deploy a function once, it's useful to retrieve the name of your deployed function programmatically before rendering a video in case your Node.JS program restarts. We can call [`getFunctions()`](/docs/lambda/getfunctions) with the `compatibleOnly` flag to get only functions with a matching version.

```

ts

import {
  getFunctions,
  renderMediaOnLambda,
  getRenderProgress,
} from '@remotion/lambda';

const functions = await getFunctions({
  region: 'us-east-1',
  compatibleOnly: true,
});

const functionName = functions[0].functionName;
```

```

ts

import {
  getFunctions,
  renderMediaOnLambda,
  getRenderProgress,
} from '@remotion/lambda';

const functions = await getFunctions({
  region: 'us-east-1',
  compatibleOnly: true,
});

const functionName = functions[0].functionName;
```

We can now trigger a render using the [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) function.

```

ts

const {renderId, bucketName} = await renderMediaOnLambda({
  region: 'us-east-1',
  functionName,
  serveUrl: url,
  composition: 'HelloWorld',
  inputProps: {},
  codec: 'h264',
  imageFormat: 'jpeg',
  maxRetries: 1,
  framesPerLambda: 20,
  privacy: 'public',
});
```

```

ts

const {renderId, bucketName} = await renderMediaOnLambda({
  region: 'us-east-1',
  functionName,
  serveUrl: url,
  composition: 'HelloWorld',
  inputProps: {},
  codec: 'h264',
  imageFormat: 'jpeg',
  maxRetries: 1,
  framesPerLambda: 20,
  privacy: 'public',
});
```

The render will now run and after a while the video will be available in your S3 bucket. You can at any time get the status of the video render by calling [`getRenderProgress()`](/docs/lambda/getrenderprogress).

```

ts

while (true) {
  await new Promise((resolve) => setTimeout(resolve, 1000));
  const progress = await getRenderProgress({
    renderId,
    bucketName,
    functionName,
    region: 'us-east-1',
  });
  if (progress.done) {
    console.log('Render finished!', progress.outputFile);
    process.exit(0);
  }
  if (progress.fatalErrorEncountered) {
    console.error('Error enountered', progress.errors);
    process.exit(1);
  }
}
```

```

ts

while (true) {
  await new Promise((resolve) => setTimeout(resolve, 1000));
  const progress = await getRenderProgress({
    renderId,
    bucketName,
    functionName,
    region: 'us-east-1',
  });
  if (progress.done) {
    console.log('Render finished!', progress.outputFile);
    process.exit(0);
  }
  if (progress.fatalErrorEncountered) {
    console.error('Error enountered', progress.errors);
    process.exit(1);
  }
}
```

This code will poll every second to check the progress of the video and exit the script if the render is done. Congrats! [Check your S3 Bucket](https://s3.console.aws.amazon.com/s3/) \- you just rendered your first video using Remotion Lambda 🚀

## Next steps [​](\#next-steps "Direct link to Next steps")

- Select [which region(s)](/docs/lambda/region-selection) you want to run Remotion Lambda in.
- Familiarize yourself with the CLI and the Node.JS APIs (list in sidebar).
- Learn how to [upgrade Remotion Lambda](/docs/lambda/upgrading).
- Before going live, go through the [Production checklist](/docs/lambda/checklist).
- If you have any questions, go through the [FAQ](/docs/lambda/faq) or ask in our [Discord channel](https://remotion.dev/discord)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/setup.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/lambda) [Next\
\
Authentication](/docs/lambda/authentication)

- [1\. Install `@remotion/lambda`](#1-install-remotionlambda)
- [2\. Create role policy](#2-create-role-policy)
- [3\. Create a role](#3-create-a-role)
- [4\. Create a user](#4-create-a-user)
- [5\. Create an access key for the user](#5-create-an-access-key-for-the-user)
- [6\. Add permissions to your user](#6-add-permissions-to-your-user)
- [7\. Optional: Validate the permission setup](#7-optional-validate-the-permission-setup)
- [8\. Deploy a function](#8-deploy-a-function)
- [9\. Deploy a site](#9-deploy-a-site)
- [10\. Check AWS concurrency limit](#10-check-aws-concurrency-limit)
- [11\. Render a video](#11-render-a-video)
- [Next steps](#next-steps)
