---
title: Lambda Limits
source: Remotion Documentation
last_updated: 2024-12-22
---

# Lambda Limits

- [Home page](/)
- [Lambda](/docs/lambda)
- Limits

On this page

# Lambda Limits

The standard AWS Lambda quotas apply ( [see here](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html)), most notably:

- [**Concurrency**](/docs/lambda/concurrency): By default, the maximum concurrent executions per region per account is 1000 executions. This limit might be lower for new accounts and users within an enterprise.
- [**Storage**](/docs/lambda/disk-size): Configurable, limited to 10GB at most
- [**RAM**](/docs/lambda/runtime#memory-size): Configurable, limited to 10GB at most
- [**Execution limit**](/docs/lambda/runtime#timeout): Configurable, at most 15 minutes

## Upgrading your concurrency limit [​](\#upgrading-your-concurrency-limit "Direct link to Upgrading your concurrency limit")

For scaling your renders, you should request a quota increase under [`https://console.aws.amazon.com/servicequotas/home`](https://console.aws.amazon.com/servicequotas/home) or using the [Remotion CLI](/docs/lambda/cli/quotas):

```

npx remotion lambda quotas increase
```

```

npx remotion lambda quotas increase
```

This only works for AWS Root accounts, not the children of an organization. You can still request an increase via the console.

## If AWS asks you for the reason [​](\#if-aws-asks-you-for-the-reason "Direct link to If AWS asks you for the reason")

AWS might ask you why you want to increase your concurrency limit. Commonly they send you this questionnaire.

> If you would still like a higher limit, please provide the following:
>
> 1. The main lambda functions ARNs of this application
> 2. Expected transactions per second per function
> 3. Expected duration per function
> 4. How is each function invoked
> 5. What services or resources do primary functions interact with
> 6. Time periods of high demand that demonstrate approaching the provided concurrency limit.

You can answer it as follows:

> 1. The functions are \[function-names\]

You can get a list of Remotion Lambda functions using `npx remotion lambda functions ls -q`.

> 1. The transactions per second are \[Provide a rough estimation\]
> 2. The duration is usually around 30 seconds to 1 minute.
> 3. In development phase it will be invoked via the Remotion CLI, and in production via a NodeJS app.
> 4. FFmpeg and Puppeteer. The npm package is called @remotion/lambda.
> 5. The process has a rendering strategy of splitting the rendering in multiple threads, to speed up the result. That is the reason why the quota increase is needed, as explained here: [https://www.remotion.dev/lambda](https://www.remotion.dev/lambda)

This response has previously been successfully used to obtain an increase of the concurrency limit.

## See also [​](\#see-also "Direct link to See also")

- [Lambda Rate Limit Troubleshooting](/docs/lambda/troubleshooting/rate-limit)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/limits.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Optimizing for speed](/docs/lambda/optimizing-speed) [Next\
\
Prerelease Changelog](/docs/lambda/changelog)

- [Upgrading your concurrency limit](#upgrading-your-concurrency-limit)
- [If AWS asks you for the reason](#if-aws-asks-you-for-the-reason)
- [See also](#see-also)
