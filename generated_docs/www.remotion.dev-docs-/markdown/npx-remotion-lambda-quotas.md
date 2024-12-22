---
title: npx remotion lambda quotas
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion lambda quotas

- [Home page](/)
- [Command line](/docs/cli/)
- [lambda](/docs/lambda/cli)
- quotas

On this page

# npx remotion lambda quotas

Displays the AWS Lambda concurrency and burst limits currently being applied to your account and region. You can use the `increase` command to send a request to AWS asking them to increase their quotas.

- [`quotas`](#quotas)
- [`quotas increase`](#quotas-increase)

## `quotas` [​](\#quotas "Direct link to quotas")

Print the current limits being applied to your AWS account and region.

```

npx remotion lambda quotas
```

```

npx remotion lambda quotas
```

Show example output

```
Region = us-east-1Concurrency limit: 1000 - Increase recommended!
A request to increase it to 5000 is pending:
https://us-east-1.console.aws.amazon.com/support/home#/case/?displayId=9742781451
The maximum amount of Lambda functions which can concurrently execute.
Run npx remotion lambda quotas increase to ask AWS to increase your limit.
```

### `--region` [​](\#--region "Direct link to --region")

For which region the quotas should be printed.

## `quotas increase` [​](\#quotas-increase "Direct link to quotas-increase")

Creates an AWS support request to increase the concurrency limit on your account as well as potential quota increase requests that might exist on your account.

```

npx remotion lambda quotas increase
```

```

npx remotion lambda quotas increase
```

note

This only works for AWS Root accounts, not the children of an organization. You can still request an increase via the console.

### `--region` [​](\#--region-1 "Direct link to --region-1")

For which region the quotas should be increased.

### `--yes` [​](\#--yes "Direct link to --yes")

Skips asking for confirmation.

### `--force` [​](\#--force "Direct link to --force")

Asks for an increase even if it is the second one. Without a message you specify, it is more unlikely you will get approved, prefer using console for further increases.

## See also [​](\#see-also "Direct link to See also")

- [Setup guide](/docs/lambda/setup)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/cli/quotas.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
regions](/docs/lambda/cli/regions) [Next\
\
Overview](/docs/cloudrun/cli)

- [`quotas`](#quotas)
  - [`--region`](#--region)
- [`quotas increase`](#quotas-increase)
  - [`--region`](#--region-1)
  - [`--yes`](#--yes)
  - [`--force`](#--force)
- [See also](#see-also)
