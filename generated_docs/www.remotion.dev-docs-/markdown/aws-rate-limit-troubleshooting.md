---
title: AWS Rate Limit Troubleshooting
source: Remotion Documentation
last_updated: 2024-12-22
---

# AWS Rate Limit Troubleshooting

- [Home page](/)
- [Lambda](/docs/lambda)
- Troubleshooting
- TooManyRequestsException

On this page

# AWS Rate Limit Troubleshooting

If you get an error message:

```

TooManyRequestsException: Rate Exceeded.
```

```

TooManyRequestsException: Rate Exceeded.
```

while calling a Lambda function, it means your concurrency limit has been reached.

- **Concurrency limit**: The maximum amount of Lambda functions that can run concurrently per region per account.
- **Burst limit**: The maximum amount of concurrency increase in 10 seconds (the burst limit is 1000)

## Default concurrency limits [​](\#default-concurrency-limits "Direct link to Default concurrency limits")

By default, the concurrency limit is 1000 functions per region, however in some regions the burst limit is only 500, somewhat limiting the scale you can use.

## New accounts using AWS Lambda [​](\#new-accounts-using-aws-lambda "Direct link to New accounts using AWS Lambda")

According to AWS, "some accounts" which are new to AWS Lambda might get a very low concurrency limit such as 10 when they first start with AWS Lambda. In that case, increase the limit via the AWS console or the Remotion CLI (see below).

### Workaround [​](\#workaround "Direct link to Workaround")

If you want to test Remotion Lambda while you are waiting for AWS to approve a higher concurrency limit, you can pass a higher number to [`framesPerLambda`](/docs/lambda/rendermediaonlambda#framesperlambda).

For example: If your account currently has a concurrency limit of 10 and you want to render a composition with 900 frames, you can set `framesPerLambda` to `100`.

9 Lambda renderer functions will be spawned: `900 / 100 = 9`.

An additional Lambda function will be spawned for orchestration, meaning a total of 10 Lambda functions will be used, keeping yourself within your limit.

## See your limits [​](\#see-your-limits "Direct link to See your limits")

To see your limits, run

```

npx remotion lambda quotas
```

```

npx remotion lambda quotas
```

note

If you get a permission error, repeat the user policy step in the [Setup Guide](/docs/lambda/setup) and update your user policy file in the AWS console.

## Request an increase [​](\#request-an-increase "Direct link to Request an increase")

You can request a quota increase under [`https://console.aws.amazon.com/servicequotas/home`](https://console.aws.amazon.com/servicequotas/home) or using the [Remotion CLI](/docs/lambda/cli/quotas):

```

npx remotion lambda quotas increase
```

```

npx remotion lambda quotas increase
```

note

This only works for AWS Root accounts, not the children of an organization. You can still request an increase via the console.

[See here](/docs/lambda/limits#if-aws-asks-you-for-the-reason) for a default answer if AWS asks why you requested the increase.

## Unhelpful? [​](\#unhelpful "Direct link to Unhelpful?")

Contact the Remotion team, preferrably via [Discord](https://remotion.dev/discord) and we will be happy to help you with your rate limit problem.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/troubleshooting/rate-limit.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Permissions](/docs/lambda/troubleshooting/permissions) [Next\
\
UnrecognizedClientException](/docs/lambda/troubleshooting/unrecognizedclientexception)

- [Default concurrency limits](#default-concurrency-limits)
- [New accounts using AWS Lambda](#new-accounts-using-aws-lambda)
  - [Workaround](#workaround)
- [See your limits](#see-your-limits)
- [Request an increase](#request-an-increase)
- [Unhelpful?](#unhelpful)