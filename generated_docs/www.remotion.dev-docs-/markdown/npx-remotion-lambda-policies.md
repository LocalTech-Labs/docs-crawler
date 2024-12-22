---
title: npx remotion lambda policies
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion lambda policies

- [Home page](/)
- [Command line](/docs/cli/)
- [lambda](/docs/lambda/cli)
- policies

On this page

# npx remotion lambda policies

Prints the necessary permissions to be inserted into the AWS console during [setup](/docs/lambda/setup).

- [`role`](#role)
- [`user`](#user)
- [`validate`](#validate)

tip

On macOS, add `| pbcopy` to the end of the command to copy the output.

## role [​](\#role "Direct link to role")

```

npx remotion lambda policies role
```

```

npx remotion lambda policies role
```

Show output

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

## user [​](\#user "Direct link to user")

```

npx remotion lambda policies user
```

```

npx remotion lambda policies user
```

Show output

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

## validate [​](\#validate "Direct link to validate")

Goes through all user permissions and validates them against the AWS Policy simulator. (Role permissions cannot be validated)

## See also [​](\#see-also "Direct link to See also")

- [Setup guide](/docs/lambda/setup)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/cli/policies.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
still](/docs/lambda/cli/still) [Next\
\
compositions](/docs/lambda/cli/compositions)

- [role](#role)
- [user](#user)
- [validate](#validate)
- [See also](#see-also)
