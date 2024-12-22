---
title: "The security token included in the request is invalid"
source: Remotion Documentation
last_updated: 2024-12-22
---

# "The security token included in the request is invalid"

- [Home page](/)
- [Lambda](/docs/lambda)
- Troubleshooting
- Security token is invalid

On this page

# "The security token included in the request is invalid"

Some regions that are supported by Remotion are not enabled by default in an AWS account.

If you get a message:

```

The security token included in the request is invalid
```

```

The security token included in the request is invalid
```

it means the region is not enabled in your AWS account.

## Option 1: Enable the regions [​](\#option-1-enable-the-regions "Direct link to Option 1: Enable the regions")

Go to the AWS console, then:

- [1](#1) Click on your name at the top right
- [2](#2) Click `Account`
- [3](#3) Enable the regions that you would like to use.

## Option 2: Make `getRegions()` return only regions by default [​](\#option-2-make-getregions-return-only-regions-by-default "Direct link to option-2-make-getregions-return-only-regions-by-default")

Use the [`enabledByDefaultOnly`](/docs/lambda/getregions#enabledbydefaultonly) option.

## See also [​](\#see-also "Direct link to See also")

- [Region selection](/docs/lambda/region-selection)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/troubleshooting/security-token.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
The bucket does not allow ACLs](/docs/lambda/troubleshooting/bucket-disallows-acl) [Next\
\
With IAM Roles](/docs/lambda/without-iam/)

- [Option 1: Enable the regions](#option-1-enable-the-regions)
- [Option 2: Make `getRegions()` return only regions by default](#option-2-make-getregions-return-only-regions-by-default)
- [See also](#see-also)
