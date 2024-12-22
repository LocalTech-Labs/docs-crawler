---
title: Uninstall Lambda
source: Remotion Documentation
last_updated: 2024-12-22
---

# Uninstall Lambda

- [Home page](/)
- [Lambda](/docs/lambda)
- Uninstall Lambda

On this page

# Uninstall Lambda

Of course we are bummed if you decided not to use Remotion Lambda anymore. We are very receptive to feedback if you'd like to share it with us.

If you would like to remove **all Remotion Lambda related objects** from your infrastructure, you can follow these steps.

warning

This will remove all videos already rendered, and break your programs that are using Remotion Lambda to render videos.

## Delete Lambda Functions [​](\#delete-lambda-functions "Direct link to Delete Lambda Functions")

You can delete all functions using the following command. The `yes` flag is already included, if you run this, it will delete all functions without confirmation.

```

npx remotion lambda functions rmall -y
```

```

npx remotion lambda functions rmall -y
```

## Delete projects [​](\#delete-projects "Direct link to Delete projects")

```

npx remotion lambda sites rmall -y
```

```

npx remotion lambda sites rmall -y
```

## Delete renders and artifacts [​](\#delete-renders-and-artifacts "Direct link to Delete renders and artifacts")

Delete all S3 buckets starting with `remotionlambda-` from your AWS account.

* * *

Make sure you have [applied the instructions to all regions](/docs/lambda/region-selection). You have now cleared up your AWS infrastructure from all Remotion Lambda-related resources.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/uninstall.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Upgrading Lambda](/docs/lambda/upgrading) [Next\
\
Cannot create public bucket](/docs/lambda/s3-public-access)

- [Delete Lambda Functions](#delete-lambda-functions)
- [Delete projects](#delete-projects)
- [Delete renders and artifacts](#delete-renders-and-artifacts)
