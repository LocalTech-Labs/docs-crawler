---
title: UnrecognizedClientException
source: Remotion Documentation
last_updated: 2024-12-22
---

# UnrecognizedClientException

- [Home page](/)
- [Lambda](/docs/lambda)
- Troubleshooting
- UnrecognizedClientException

On this page

# UnrecognizedClientException

If you got a permission error while calling [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) or [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda):

```

txt

UnrecognizedClientException: The AWS credentials provided were probably mixed up.
```

```

txt

UnrecognizedClientException: The AWS credentials provided were probably mixed up.
```

it means that the AWS credentials were correct, but don't allow access to a certain resource.

## Most common cause: Calling a Remotion function inside a serverless function [​](\#most-common-cause-calling-a-remotion-function-inside-a-serverless-function "Direct link to Most common cause: Calling a Remotion function inside a serverless function")

When calling render inside an AWS Lambda function or a Vercel serverless function, that function already has it's own `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables set. However, these are from AWS itself and are different from the variables that Remotion uses to invoke serverless functions.

To avoid that conflict, you can change the name of the environment variables you set:

- Rename `AWS_ACCESS_KEY_ID` to `REMOTION_AWS_ACCESS_KEY_ID`
- Rename `AWS_SECRET_ACCESS_KEY` to `REMOTION_AWS_SECRET_ACCESS_KEY`

If both are set, Remotion will prefer the environment variables that are prefixed with `REMOTION_`, which should separate the two different credentials nicely.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/troubleshooting/unrecognizedclientexception.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
TooManyRequestsException](/docs/lambda/troubleshooting/rate-limit) [Next\
\
The bucket does not allow ACLs](/docs/lambda/troubleshooting/bucket-disallows-acl)

- [Most common cause: Calling a Remotion function inside a serverless function](#most-common-cause-calling-a-remotion-function-inside-a-serverless-function)
