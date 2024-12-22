---
title: The bucket does not allow ACLs
source: Remotion Documentation
last_updated: 2024-12-22
---

# The bucket does not allow ACLs

- [Home page](/)
- [Lambda](/docs/lambda)
- Troubleshooting
- The bucket does not allow ACLs

On this page

# The bucket does not allow ACLs

If you encounter the following error while rendering a video on Remotion Lambda:

```

AccessControlListNotSupported: The bucket does not allow ACLs
```

```

AccessControlListNotSupported: The bucket does not allow ACLs
```

You are trying to render into a bucket that has the ACL feature disabled and handles its permissions in a different way, for example through bucket policies.

## Solution [​](\#solution "Direct link to Solution")

Pass `no-acl` to the `privacy` option of [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda#privacy), [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda#privacy) or [`npx remotion lambda render`](/docs/lambda/cli/render#--privacy).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/troubleshooting/bucket-disallows-acl.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
UnrecognizedClientException](/docs/lambda/troubleshooting/unrecognizedclientexception) [Next\
\
Security token is invalid](/docs/lambda/troubleshooting/security-token)

- [Solution](#solution)
