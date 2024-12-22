---
title: AWS Permissions Troubleshooting
source: Remotion Documentation
last_updated: 2024-12-22
---

# AWS Permissions Troubleshooting

- [Home page](/)
- [Lambda](/docs/lambda)
- Troubleshooting
- Permissions

On this page

# AWS Permissions Troubleshooting

If you get an error message saying that something is wrong with your AWS permissions, read this page for troubleshooting ideas.

## Node.JS APIs don't read the `.env` file [​](\#nodejs-apis-dont-read-the-env-file "Direct link to nodejs-apis-dont-read-the-env-file")

If you use the Node.JS APIs, be aware that they don't read the .env file automatically. This means you need to either load the .env file or set the environment variables manually.

## User and role permissions have been mixed up [​](\#user-and-role-permissions-have-been-mixed-up "Direct link to User and role permissions have been mixed up")

There are two policy files, one for an AWS user, one for an AWS role. The policies are different, make sure you have correctly assigned both.

## Not enough time has passed for settings to apply [​](\#not-enough-time-has-passed-for-settings-to-apply "Direct link to Not enough time has passed for settings to apply")

It can take several minutes until the policies you entered into AWS propagate. Sometimes waiting 2-3 minutes can solve the problem.

## Required permissions have changed [​](\#required-permissions-have-changed "Direct link to Required permissions have changed")

A newer Remotion Lambda version might have required additional permissions. We will note this in the [changelog](https://github.com/remotion-dev/remotion/releases). Make sure you update the policies in AWS.

## Other AWS credentials might have been applied [​](\#other-aws-credentials-might-have-been-applied "Direct link to Other AWS credentials might have been applied")

If AWS environment variables failed to load, other credentials might have been loaded from other places such as the AWS CLI. Log the environment variables to ensure you loaded the correct ones.

## Use the validate command [​](\#use-the-validate-command "Direct link to Use the validate command")

Use the [`npx remotion lambda policies validate`](/docs/lambda/cli/policies) command to validate the user policy. Note that this can still mean the role policy is set wrongly or the environment variables don't get loaded when using the Node.JS APIs.

## See also [​](\#see-also "Direct link to See also")

- [Lambda - Permissions](/docs/lambda/permissions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/troubleshooting/permissions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Debugging failures](/docs/lambda/troubleshooting/debug) [Next\
\
TooManyRequestsException](/docs/lambda/troubleshooting/rate-limit)

- [Node.JS APIs don't read the `.env` file](#nodejs-apis-dont-read-the-env-file)
- [User and role permissions have been mixed up](#user-and-role-permissions-have-been-mixed-up)
- [Not enough time has passed for settings to apply](#not-enough-time-has-passed-for-settings-to-apply)
- [Required permissions have changed](#required-permissions-have-changed)
- [Other AWS credentials might have been applied](#other-aws-credentials-might-have-been-applied)
- [Use the validate command](#use-the-validate-command)
- [See also](#see-also)
