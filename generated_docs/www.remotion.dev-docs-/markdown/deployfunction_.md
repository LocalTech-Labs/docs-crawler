---
title: deployFunction()
source: Remotion Documentation
last_updated: 2024-12-22
---

# deployFunction()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- deployFunction()

On this page

# deployFunction()

Creates an [AWS Lambda](https://aws.amazon.com/lambda/) function in your AWS account that will be able to render a video in the cloud.

If a function with the same version, memory size and timeout already existed, it will be returned instead without a new one being created. This means this function can be treated as idempotent.

## Example [​](\#example "Direct link to Example")

```

ts

import { deployFunction } from "@remotion/lambda";

const { functionName } = await deployFunction({
  region: "us-east-1",
  timeoutInSeconds: 120,
  memorySizeInMb: 2048,
  createCloudWatchLogGroup: true,
  diskSizeInMb: 2048,
});
console.log(functionName);
```

```

ts

import { deployFunction } from "@remotion/lambda";

const { functionName } = await deployFunction({
  region: "us-east-1",
  timeoutInSeconds: 120,
  memorySizeInMb: 2048,
  createCloudWatchLogGroup: true,
  diskSizeInMb: 2048,
});
console.log(functionName);
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `region` [​](\#region "Direct link to region")

The [AWS region](/docs/lambda/region-selection) which you want to deploy the Lambda function too. It must be the same region that your Lambda Layer resides in.

### `timeoutInSeconds` [​](\#timeoutinseconds "Direct link to timeoutinseconds")

How long the Lambda function may run before it gets killed. Must be below 900 seconds.
We recommend a timeout of 120 seconds or lower - remember, Remotion Lambda is the fastest if you render with a high concurrency. If your video takes longer to render, the concurrency should be increased rather than the timeout.

### `memorySizeInMb` [​](\#memorysizeinmb "Direct link to memorysizeinmb")

How many megabytes of RAM the Lambda function should have. By default we recommend a value of 2048MB. You may increase or decrease it depending on how memory-consuming your video is. The minimum allowed number is `512`, the maximum allowed number is `10240`. Since the costs of Remotion Lambda is directly proportional to the amount of RAM, we recommend to keep this amount as low as possible.

### `createCloudWatchLogGroup?` [​](\#createcloudwatchloggroup "Direct link to createcloudwatchloggroup")

Whether logs should be saved into CloudWatch. We recommend enabling this option.

### `cloudWatchLogRetentionPeriodInDays?` [​](\#cloudwatchlogretentionperiodindays "Direct link to cloudwatchlogretentionperiodindays")

Retention period for the CloudWatch Logs. Default: 14 days.

### `diskSizeInMb?` [​](\#disksizeinmb "Direct link to disksizeinmb")

Sets the amount of disk storage that is available in the Lambda function. Must be between 512MB and 10240MB (10GB). Set this higher if you want to render longer videos. See also: [Disk size](/docs/lambda/disk-size)

Remotion VersionDefault<5.0.02048MB>=5.0.010240MB

### `customRoleArn?` [​](\#customrolearn "Direct link to customrolearn")

Use a custom role for the function instead of the default ( `arn:aws:iam::[aws-account-id]:role/remotion-lambda-role`)

### `enableLambdaInsights?` [v4.0.61](https://github.com/remotion-dev/remotion/releases/v4.0.61) [​](\#enablelambdainsights "Direct link to enablelambdainsights")

Enable [Lambda Insights in AWS CloudWatch](https://remotion.dev/docs/lambda/insights). For this to work, you may have to update your role permission.

### `enableV5Runtime?` [v4.0.148](https://github.com/remotion-dev/remotion/releases/v4.0.148) [​](\#enablev5runtime "Direct link to enablev5runtime")

Enable the [upcoming v5 runtime](/docs/lambda/runtime#runtime-changes-in-remotion-50) with newer Chrome and Node versions early.

### `runtimePreference?` [v4.0.205](https://github.com/remotion-dev/remotion/releases/v4.0.205) [​](\#runtimepreference "Direct link to runtimepreference")

One of:

- `default`: Currently resolving to `prefer-cjk`
- `prefer-apple-emojis`: Use Apple Emojis instead of Google Emojis. CJK characters will be removed.
- `prefer-cjk`: Include CJK (Chinese, Japanese, Korean) characters and Google Emojis. Apple Emojis will be removed.

note

Apple Emojis are intellectual property of Apple Inc.

You are responsible for the use of Apple Emojis in your project.

## Return value [​](\#return-value "Direct link to Return value")

An object with the following values:

- `functionName` ( _string_): The name of the function just created.
- `alreadyExisted`: ( _boolean_): Whether the creation was skipped because the function already existed.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/deploy-function.ts)
- [CLI equivalent: `npx remotion lambda functions deploy`](/docs/lambda/cli/functions#deploy)
- [`deleteFunction()`](/docs/lambda/deletefunction)
- [`getFunctions()`](/docs/lambda/getfunctions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/deployfunction.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
estimatePrice()](/docs/lambda/estimateprice) [Next\
\
deleteFunction()](/docs/lambda/deletefunction)

- [Example](#example)
- [Arguments](#arguments)
  - [`region`](#region)
  - [`timeoutInSeconds`](#timeoutinseconds)
  - [`memorySizeInMb`](#memorysizeinmb)
  - [`createCloudWatchLogGroup?`](#createcloudwatchloggroup)
  - [`cloudWatchLogRetentionPeriodInDays?`](#cloudwatchlogretentionperiodindays)
  - [`diskSizeInMb?`](#disksizeinmb)
  - [`customRoleArn?`](#customrolearn)
  - [`enableLambdaInsights?`](#enablelambdainsights)
  - [`enableV5Runtime?`](#enablev5runtime)
  - [`runtimePreference?`](#runtimepreference)
- [Return value](#return-value)
- [See also](#see-also)
