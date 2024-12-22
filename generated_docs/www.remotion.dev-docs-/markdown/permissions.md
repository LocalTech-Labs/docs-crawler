---
title: Permissions
source: Remotion Documentation
last_updated: 2024-12-22
---

# Permissions

- [Home page](/)
- [Lambda](/docs/lambda)
- Permissions

On this page

# Permissions

This document describes the necessary permissions for Remotion Lambda and explains to those interested why the permissions are necessary.

For a step by step guide on how to set up permissions, [follow the setup guide](/docs/lambda/setup).

## User permissions [​](\#user-permissions "Direct link to User permissions")

This policy should be assigned to the **AWS user**. To do so, go to the [AWS console](https://console.aws.amazon.com/console/home) ➞ [IAM](https://console.aws.amazon.com/iam/home) ➞ [Users](https://console.aws.amazon.com/iamv2/home#/users) ➞ Your created Remotion user ➞ Permissions tab ➞ Add inline policy ➞ JSON.

Show full user permissions JSON file for latest Remotion Lambda version

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "HandleQuotas",
      "Effect": "Allow",
      "Action": [
        "servicequotas:GetServiceQuota",
        "servicequotas:GetAWSDefaultServiceQuota",
        "servicequotas:RequestServiceQuotaIncrease",
        "servicequotas:ListRequestedServiceQuotaChangeHistoryByQuota"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "PermissionValidation",
      "Effect": "Allow",
      "Action": [
        "iam:SimulatePrincipalPolicy"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "LambdaInvokation",
      "Effect": "Allow",
      "Action": [
        "iam:PassRole"
      ],
      "Resource": [
        "arn:aws:iam::*:role/remotion-lambda-role"
      ]
    },
    {
      "Sid": "Storage",
      "Effect": "Allow",
      "Action": [
        "s3:GetObject",
        "s3:DeleteObject",
        "s3:PutObjectAcl",
        "s3:PutObject",
        "s3:CreateBucket",
        "s3:ListBucket",
        "s3:GetBucketLocation",
        "s3:PutBucketAcl",
        "s3:DeleteBucket",
        "s3:PutBucketOwnershipControls",
        "s3:PutBucketPublicAccessBlock",
        "s3:PutLifecycleConfiguration"
      ],
      "Resource": [
        "arn:aws:s3:::remotionlambda-*"
      ]
    },
    {
      "Sid": "BucketListing",
      "Effect": "Allow",
      "Action": [
        "s3:ListAllMyBuckets"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "FunctionListing",
      "Effect": "Allow",
      "Action": [
        "lambda:ListFunctions",
        "lambda:GetFunction"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "FunctionManagement",
      "Effect": "Allow",
      "Action": [
        "lambda:InvokeAsync",
        "lambda:InvokeFunction",
        "lambda:CreateFunction",
        "lambda:DeleteFunction",
        "lambda:PutFunctionEventInvokeConfig",
        "lambda:PutRuntimeManagementConfig",
        "lambda:TagResource"
      ],
      "Resource": [
        "arn:aws:lambda:*:*:function:remotion-render-*"
      ]
    },
    {
      "Sid": "LogsRetention",
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:PutRetentionPolicy"
      ],
      "Resource": [
        "arn:aws:logs:*:*:log-group:/aws/lambda/remotion-render-*"
      ]
    },
    {
      "Sid": "FetchBinaries",
      "Effect": "Allow",
      "Action": [
        "lambda:GetLayerVersion"
      ],
      "Resource": [
        "arn:aws:lambda:*:678892195805:layer:remotion-binaries-*",
        "arn:aws:lambda:*:580247275435:layer:LambdaInsightsExtension*"
      ]
    }
  ]
}
```

info

You can always get the suitable permission file for your Remotion Lambda version by typing `npx remotion lambda policies user`.

## Role permissions [​](\#role-permissions "Direct link to Role permissions")

This policy should be assigned to the **role `remotion-lambda-role`** in your AWS account. The permissions below are given to the Lambda function itself.

To assign, go to [AWS console](https://console.aws.amazon.com/console/home) ➞ [IAM](https://console.aws.amazon.com/iam/home) ➞ [Roles](https://console.aws.amazon.com/iamv2/home#/roles) ➞ [`remotion-lambda-role`](https://console.aws.amazon.com/iam/home#/roles/remotion-lambda-role) ➞ Permissions tab ➞ [Add inline policy](https://console.aws.amazon.com/iam/home#/roles/remotion-lambda-role$createPolicy?step=edit).

Show full role permissions JSON file for latest Remotion Lambda version

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "0",
      "Effect": "Allow",
      "Action": [
        "s3:ListAllMyBuckets"
      ],
      "Resource": [
        "*"
      ]
    },
    {
      "Sid": "1",
      "Effect": "Allow",
      "Action": [
        "s3:CreateBucket",
        "s3:ListBucket",
        "s3:PutBucketAcl",
        "s3:GetObject",
        "s3:DeleteObject",
        "s3:PutObjectAcl",
        "s3:PutObject",
        "s3:GetBucketLocation"
      ],
      "Resource": [
        "arn:aws:s3:::remotionlambda-*"
      ]
    },
    {
      "Sid": "2",
      "Effect": "Allow",
      "Action": [
        "lambda:InvokeFunction"
      ],
      "Resource": [
        "arn:aws:lambda:*:*:function:remotion-render-*"
      ]
    },
    {
      "Sid": "3",
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup"
      ],
      "Resource": [
        "arn:aws:logs:*:*:log-group:/aws/lambda-insights"
      ]
    },
    {
      "Sid": "4",
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": [
        "arn:aws:logs:*:*:log-group:/aws/lambda/remotion-render-*",
        "arn:aws:logs:*:*:log-group:/aws/lambda-insights:*"
      ]
    }
  ]
}
```

info

You can always get the suitable permission file for your Remotion Lambda version by typing `npx remotion lambda policies role`.

## Validation [​](\#validation "Direct link to Validation")

There are two ways in which you can test if the permissions for the user have been correctly set up. Either you execute the following command:

```

bash

npx remotion lambda policies validate
```

```

bash

npx remotion lambda policies validate
```

or if you want to validate it programmatically, using the [`simulatePermissions()`](/docs/lambda/simulatepermissions) function.

info

Policies for the role cannot be validated.

## Explanation [​](\#explanation "Direct link to Explanation")

The following table is a breakdown of why Remotion Lambda requires the permissions it does.

### User policies [​](\#user-policies "Direct link to User policies")

Permission

Scope

Reason

`iam:SimulatePrincipalPolicy``*`

Allows for `npx remotion lambda permissions validate`.

`iam:PassRole``arn:aws:iam::*:role/remotion-lambda-role`

Allows the Lambda function to assume a role with sufficient permissions.

`s3:GetObject`

`s3:DeleteObject`

`s3:PutObjectAcl`

`s3:PutObject`

`s3:CreateBucket`

`s3:ListBucket`

`s3:GetBucketLocation`

`s3:PutBucketAcl`

`s3:DeleteBucket`

`s3.PutBucketOwnershipControls`

`s3.PutBucketPublicAccessBlock`

`arn:aws:s3:::remotionlambda-*`

Allows to create and delete buckets and objects in your account, make objects public and configure them as websites. Only buckets that start with `remotionlambda-` can be accessed.

`s3:ListAllMyBuckets``arn:aws:s3:::*`

Allows listing the names of all buckets in your account, in order to detect an already existing Remotion bucket.

`lambda:GetLayerVersion``arn:aws:lambda:*:678892195805:layer:remotion-binaries-*`

Allows to read Chromium and FFMPEG binaries. These binaries are hosted in an account hosted by Remotion specifically dedicated to hosting those layers in all supported regions.

`lambda:ListFunctions`

`lambda:GetFunction`

`*`

Allows to read the functions in your AWS account in order to find the correct function to invoke. The loose `*` permission is because AWS doesn't allow this permission to be tightened.

`lambda:InvokeAsync`

`lambda:InvokeFunction`

`lambda:DeleteFunction`

`lambda:PutFunctionEventInvokeConfig`

`lambda:CreateFunction`

`lambda:PutRuntimeManagementConfig`

`lambda:TagResource`

`arn:aws:lambda:*:*:function:remotion-render-*`

Allows to create, delete, invoke and configure functions (such as disabling automatic retries). Used by the CLI and the Node.JS APIS to set up, execute and teardown the infrastructure. `lambda:TagResource` will optionally tag the functions

`logs:CreateLogGroup`

`logs:PutRetentionPolicy``arn:aws:logs:*:*:log-group:/aws/lambda/remotion-render-*`

Allows to create CloudWatch group, so logs can be saved in there later. Simplifies debugging.

`servicequotas:GetServiceQuota`

`servicequotas:GetAWSDefaultServiceQuota`

`servicequotas:RequestServiceQuotaIncrease`

`servicequotas:ListRequestedServiceQuotaChangeHistoryByQuota`

`*`

Powers the `lambda quotas` CLI command.

### Role policies [​](\#role-policies "Direct link to Role policies")

Permission

Scope

Reason

`s3.ListAllMyBuckets``*`

Get a list of Remotion buckets in order to find existing buckets that start with `remotionlambda-`.

`s3:CreateBucket`

`s3:ListBucket`

`s3:PutBucketAcl`

`s3:GetObject`

`s3:DeleteObject`

`s3:PutObjectAcl`

`s3:PutObject`

`s3:GetBucketLocation``arn:aws:s3:::remotionlambda-*`

Create and delete buckets and items, make them public or private and fetch their location. Since Remotion stores the videos in an S3 bucket, it needs basic CRUD capabilities over those buckets. The permission only applies to buckets that start with `remotionlambda-`

`lambda:InvokeFunction`

`arn:aws:lambda:*:*:function:remotion-render*`

Allow the function to recursively invoke itself. A render involves multiple function calls, which is to be orchestrated by the first function call.

`lambda:CreateLogStream`

`lambda:PutLogEvents`

`arn:aws:logs:*:*:log-group:/aws/lambda/remotion-render*`

Allows to write the function logs to CloudWatch for easier debugging.

## See also [​](\#see-also "Direct link to See also")

- [Set up guide](/docs/lambda/setup)
- [`simulatePermissions()`](/docs/lambda/simulatepermissions)
- [Permissions Troubleshooting](/docs/lambda/troubleshooting/permissions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/permissions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Authentication](/docs/lambda/authentication) [Next\
\
Region selection](/docs/lambda/region-selection)

- [User permissions](#user-permissions)
- [Role permissions](#role-permissions)
- [Validation](#validation)
- [Explanation](#explanation)
  - [User policies](#user-policies)
  - [Role policies](#role-policies)
- [See also](#see-also)
