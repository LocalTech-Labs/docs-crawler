---
title: npx remotion lambda functions
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion lambda functions

- [Home page](/)
- [Command line](/docs/cli/)
- [lambda](/docs/lambda/cli)
- functions

On this page

# npx remotion lambda functions

The `npx remotion lambda functions` command allows you to deploy, view and delete AWS lambda functions that can render videos.

- [`deploy`](#deploy)
- [`ls`](#ls)
- [`rm`](#rm)
- [`rmall`](#rmall)

You only need one function per AWS region and Remotion version. Suggested reading: [Do I need to deploy a function for each render?](/docs/lambda/faq#do-i-need-to-deploy-a-function-for-each-render)

## deploy [​](\#deploy "Direct link to deploy")

```

npx remotion lambda functions deploy
```

```

npx remotion lambda functions deploy
```

Creates a new function in your AWS account. If a function in the same region, with the same Remotion version, with the same amount of memory, disk space and timeout already exists, the name of the already deployed function will be returned instead.

By default, a CloudWatch Log Group will be created that will log debug information to CloudWatch that you can consult in the case something is going wrong. The default retention period for these logs is 14 days, which can be changed.

Example output

```
Region = eu-central-1,
Memory = 2048MB,
Disk = 2048MB,
Timeout = 120sec,
Version = 2021-12-17,
CloudWatch Enabled = true,
CloudWatch Retention Period = 14 days
VPC Subnet IDs = subnet-0f6a0f6a0f6a0f6a0, subnet-0f6a0f6a0f6a0f6a1
VPC Security Group IDs = sg-0f6a0f6a0f6a0f6a0, sg-0f6a0f6a0f6a0f6a1
Deployed as remotion-render-2021-12-17-2048mb-120sec

```

### `--region` [​](\#--region "Direct link to --region")

The [AWS region](/docs/lambda/region-selection) to select.

### `--memory` [​](\#--memory "Direct link to --memory")

Memory size in megabytes. Default: 2048 MB.

### `--disk` [​](\#--disk "Direct link to --disk")

Disk size in megabytes. See also: [Disk size](/docs/lambda/disk-size).

Remotion VersionDefault<5.0.02048MB>=5.0.010240MB

### `--timeout` [​](\#--timeout "Direct link to --timeout")

Timeout of the Lambda function in seconds. Default: 120 seconds.

note

Not to be confused with the [`--timeout` flag for `npx remotion lambda render` which defines the timeout for `delayRender()`](/docs/cli/render#--timeout).

### `--disable-cloudwatch` [​](\#--disable-cloudwatch "Direct link to --disable-cloudwatch")

Does not create a CloudWatch log group.

### `--retention-period` [​](\#--retention-period "Direct link to --retention-period")

Retention period for the CloudWatch Logs in days. Default: 14 days.

### `--enable-lambda-insights` [v4.0.61](https://github.com/remotion-dev/remotion/releases/v4.0.61) [​](\#--enable-lambda-insights "Direct link to --enable-lambda-insights")

Enable [Lambda Insights in AWS CloudWatch](https://remotion.dev/docs/lambda/insights). For this to work, you may have to update your role permission.

### `--custom-role-arn` [​](\#--custom-role-arn "Direct link to --custom-role-arn")

Use a custom role for the function instead of the default ( `arn:aws:iam::[aws-account-id]:role/remotion-lambda-role`)

### `--quiet`, `-q` [​](\#--quiet--q "Direct link to --quiet--q")

Only logs the function name.

### `--enable-v5-runtime` [v4.0.148](https://github.com/remotion-dev/remotion/releases/v4.0.148) [​](\#--enable-v5-runtime "Direct link to --enable-v5-runtime")

Enable the [upcoming v5 runtime](/docs/lambda/runtime#runtime-changes-in-remotion-50) with newer Chrome and Node versions early.

### `--vpc-subnet-ids` [v4.0.160](https://github.com/remotion-dev/remotion/releases/v4.0.160) [​](\#--vpc-subnet-ids "Direct link to --vpc-subnet-ids")

Comma separated list of VPC subnet IDs to use for the Lambda function VPC configuration.

### `--vpc-security-group-ids` [v4.0.160](https://github.com/remotion-dev/remotion/releases/v4.0.160) [​](\#--vpc-security-group-ids "Direct link to --vpc-security-group-ids")

Comma separated list of VPC security group IDs to use for the Lambda function VPC configuration.

### `--runtime-preference` [v4.0.205](https://github.com/remotion-dev/remotion/releases/v4.0.205) [​](\#--runtime-preference "Direct link to --runtime-preference")

One of:

- `default`: Currently resolving to `prefer-cjk`
- `prefer-apple-emojis`: Use Apple Emojis instead of Google Emojis. CJK characters will be removed.
- `prefer-cjk`: Include CJK (Chinese, Japanese, Korean) characters and Google Emojis. Apple Emojis will be removed.

note

Apple Emojis are intellectual property of Apple Inc.

You are responsible for the use of Apple Emojis in your project.

## ls [​](\#ls "Direct link to ls")

```

npx remotion lambda functions ls
```

```

npx remotion lambda functions ls
```

Lists the functions that you have deployed to AWS in the [selected region](/docs/lambda/region-selection).

Example output

```
6 functions in the eu-central-1 region

Name                                              Version        Memory (MB)    Timeout (sec)

remotion-render-2021-12-16-2048mb-240sec          2021-12-16     2048           240

remotion-render-2021-12-17-2048mb-120sec          2021-12-17     2048           120

remotion-render-2021-12-15-2048mb-240sec          2021-12-15     2048           240

```

### `--compatible-only` [v4.0.164](https://github.com/remotion-dev/remotion/releases/v4.0.164) [​](\#--compatible-only "Direct link to --compatible-only")

Only lists functions that have the same version as the CLI you are calling the command from.

### `--region` [​](\#--region-1 "Direct link to --region-1")

The [AWS region](/docs/lambda/region-selection) to select.

### `--quiet`, `-q` [​](\#--quiet--q-1 "Direct link to --quiet--q-1")

Prints only the function names in a space-separated list. If no functions exist, prints `()`

## rm [​](\#rm "Direct link to rm")

```

npx remotion lambda functions rm remotion-render-2021-12-16-2048mb-240sec
```

```

npx remotion lambda functions rm remotion-render-2021-12-16-2048mb-240sec
```

Removes one or more functions from your AWS infrastructure. Pass a space-separated list of functions you'd like to delete.

Example output

```

Function name:   remotion-render-2021-12-16-2048mb-240sec

Memory:          2048MB

Timeout:         120sec

Version:         2021-12-16

Delete? (Y/n):  Y

Deleted!

```

### `--region` [​](\#--region-2 "Direct link to --region-2")

The [AWS region](/docs/lambda/region-selection) to select.

### `--yes`, `-y` [​](\#--yes--y "Direct link to --yes--y")

Skips confirmation.

## rmall [​](\#rmall "Direct link to rmall")

```

npx remotion lambda functions rmall
```

```

npx remotion lambda functions rmall
```

Removes all functions in a region from your AWS infrastructure.

Example output

```

Function name:   remotion-render-2021-12-16-2048mb-240sec

Memory:          2048MB

Timeout:         120sec

Version:         2021-12-16

Delete? (Y/n):  Y

Deleted!
Function name:   remotion-render-2021-12-18-2048mb-240sec

Memory:          2048MB

Timeout:         120sec

Version:         2021-12-16

Delete? (Y/n):  Y

Deleted!

```

### `--region` [​](\#--region-3 "Direct link to --region-3")

The [AWS region](/docs/lambda/region-selection) to select.

### `--yes`, `-y` [​](\#--yes--y-1 "Direct link to --yes--y-1")

Skips confirmation.

## See also [​](\#see-also "Direct link to See also")

- [Do I need to deploy a function for each render?](/docs/lambda/faq#do-i-need-to-deploy-a-function-for-each-render)
- [Setup guide](/docs/lambda/setup)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/cli/functions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
sites](/docs/lambda/cli/sites) [Next\
\
render](/docs/lambda/cli/render)

- [deploy](#deploy)
  - [`--region`](#--region)
  - [`--memory`](#--memory)
  - [`--disk`](#--disk)
  - [`--timeout`](#--timeout)
  - [`--disable-cloudwatch`](#--disable-cloudwatch)
  - [`--retention-period`](#--retention-period)
  - [`--enable-lambda-insights`](#--enable-lambda-insights)
  - [`--custom-role-arn`](#--custom-role-arn)
  - [`--quiet`, `-q`](#--quiet--q)
  - [`--enable-v5-runtime`](#--enable-v5-runtime)
  - [`--vpc-subnet-ids`](#--vpc-subnet-ids)
  - [`--vpc-security-group-ids`](#--vpc-security-group-ids)
  - [`--runtime-preference`](#--runtime-preference)
- [ls](#ls)
  - [`--compatible-only`](#--compatible-only)
  - [`--region`](#--region-1)
  - [`--quiet`, `-q`](#--quiet--q-1)
- [rm](#rm)
  - [`--region`](#--region-2)
  - [`--yes`, `-y`](#--yes--y)
- [rmall](#rmall)
  - [`--region`](#--region-3)
  - [`--yes`, `-y`](#--yes--y-1)
- [See also](#see-also)
