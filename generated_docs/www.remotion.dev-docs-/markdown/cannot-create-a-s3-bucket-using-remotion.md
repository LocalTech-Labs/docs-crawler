---
title: Cannot create a S3 bucket using Remotion
source: Remotion Documentation
last_updated: 2024-12-22
---

# Cannot create a S3 bucket using Remotion

- [Home page](/)
- [Lambda](/docs/lambda)
- Cannot create public bucket

On this page

# Cannot create a S3 bucket using Remotion

Since approximately April 25th 2023, AWS blocks the creation of public buckets without extra configuration, making it impossible to create new S3 buckets with Remotion version v3.3.86 and older.

To make bucket creation work again, you need to upgrade to a newer Remotion version and update your user policy.

## Problem [​](\#problem "Direct link to Problem")

Users might see an error:

```

sh

InvalidBucketAclWithObjectOwnership: Bucket cannot have ACLs set with ObjectOwnership's BucketOwnerEnforced setting.
```

```

sh

InvalidBucketAclWithObjectOwnership: Bucket cannot have ACLs set with ObjectOwnership's BucketOwnerEnforced setting.
```

or

```

shell

AccessDenied: Access Denied
```

```

shell

AccessDenied: Access Denied
```

coming from the AWS SDK when trying to create a new site or bucket.

## Cause [​](\#cause "Direct link to Cause")

AWS does make all buckets private by default and Remotion tries to create a public bucket.

## Resolution [​](\#resolution "Direct link to Resolution")

[1](#1)

Upgrade to Remotion `v3.3.87` or later.

```

bash

npx remotion upgrade
```

```

bash

npx remotion upgrade
```

note

Note: When you upgrade Remotion, you will need to deploy new functions as well.

[2](#2)

change your policies to give your user the `s3:PutBucketOwnershipControls` and `s3:PutBucketPublicAccessBlock` permission. The easiest way
is to copy the following user policy:

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

or type in `npx remotion lambda policies user` after upgrading Remotion Lambda.

Go to the [Users](https://us-east-1.console.aws.amazon.com/iamv2/home?region=us-east-1#/users) section in the AWS console and overwrite the JSON policy of your Remotion Lambda user with the above copied JSON.

You can verify that it worked by running `npx remotion lambda policies validate`.

[3](#3)

Redeploy your functions. You don't need to delete your old functions
as it might cause downtime for your application.

[4](#4)

As a reminder, you also need to redeploy your site when you upgrade
to a higher Remotion version.

[5](#5)

If any values are hardcoded, update the function and site name in
your application code.

## Questions? [​](\#questions "Direct link to Questions?")

Join our [Discord](https://remotion.dev/discord) community to get help from the Remotion team and other users.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/s3-public-access.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Uninstall Lambda](/docs/lambda/uninstall) [Next\
\
Function naming convention](/docs/lambda/naming-convention)

- [Problem](#problem)
- [Cause](#cause)
- [Resolution](#resolution)
- [Questions?](#questions)
