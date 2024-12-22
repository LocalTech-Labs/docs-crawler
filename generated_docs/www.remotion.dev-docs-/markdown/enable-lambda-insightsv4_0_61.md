---
title: Enable Lambda Insightsv4.0.61
source: Remotion Documentation
last_updated: 2024-12-22
---

# Enable Lambda Insightsv4.0.61

- [Home page](/)
- [Lambda](/docs/lambda)
- Insights

On this page

# Enable Lambda Insights [v4.0.61](https://github.com/remotion-dev/remotion/releases/v4.0.61)

You may enable [AWS Lambda Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights.html) for your Remotion Lambda function.

## Prerequisites [​](\#prerequisites "Direct link to Prerequisites")

[1](#1)

Ensure you are on at least Remotion v4.0.61.

[2](#2) If you started using Remotion before v4.0.61, update both your [AWS user permission](/docs/lambda/permissions#user-permissions) and [AWS role permission](/docs/lambda/permissions#role-permissions), since now more permissions are needed.

## Enable Lambda Insights [​](\#enable-lambda-insights-1 "Direct link to Enable Lambda Insights")

**Via CLI**:

```

npx remotion lambda functions deploy --enable-lambda-insights
```

```

npx remotion lambda functions deploy --enable-lambda-insights
```

If the function already existed before, you need to delete it beforehand.

**Via Node.js APIs:**

```

deploy.ts
tsx

import { deployFunction } from "@remotion/lambda";

const { alreadyExisted } = await deployFunction({
  createCloudWatchLogGroup: true,
  region: "us-east-1",
  timeoutInSeconds: 120,
  memorySizeInMb: 3009,
  enableLambdaInsights: true,
});

// Note: If the function previously already existed, Lambda insights are not applied.
// Delete the old function and deploy again.
assert(!alreadyExisted);
```

```

deploy.ts
tsx

import { deployFunction } from "@remotion/lambda";

const { alreadyExisted } = await deployFunction({
  createCloudWatchLogGroup: true,
  region: "us-east-1",
  timeoutInSeconds: 120,
  memorySizeInMb: 3009,
  enableLambdaInsights: true,
});

// Note: If the function previously already existed, Lambda insights are not applied.
// Delete the old function and deploy again.
assert(!alreadyExisted);
```

## Add a role to the Lambda function [​](\#add-a-role-to-the-lambda-function "Direct link to Add a role to the Lambda function")

In order to actually allow Lambda to send data to CloudWatch, you need to do this once:

- Go to the Lambda Dashboard and select any Remotion Lambda function.
- In the "Configuration" tab, scroll down to "Monitoring and operations tools" and in the "Additional monitoring tools" section, click "Edit".
- Toggle the switch to "Enable AWS Lambda Insights". If it is already turned on, disable it, save, and then enable it again.

This will add the necessary permissions to the role of the Lambda function.

All Lambda functions share the same role with the default setup, so you only need to do this once.

## View Lambda insights [​](\#view-lambda-insights "Direct link to View Lambda insights")

In your CloudWatch dashboard ( [link for `us-east-1`](https://eu-central-1.console.aws.amazon.com/cloudwatch/home)) under Insights ➞ Lambda Insights ➞ Single function, you can view the metrics of the Remotion Lambda function.

A link to the Lambda insights is also included in the return value of [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda).

If you render via the CLI with the `--log=verbose` flag, a link to the Lambda insights is also printed, regardless of if Lambda insights are enabled or not.

## See also [​](\#see-also "Direct link to See also")

- [Debugging failed Lambda renders](/docs/lambda/troubleshooting/debug)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/insights.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Serverless.com integration](/docs/lambda/serverless-framework-integration) [Next\
\
Multiple buckets](/docs/lambda/multiple-buckets)

- [Prerequisites](#prerequisites)
- [Enable Lambda Insights](#enable-lambda-insights-1)
- [Add a role to the Lambda function](#add-a-role-to-the-lambda-function)
- [View Lambda insights](#view-lambda-insights)
- [See also](#see-also)
