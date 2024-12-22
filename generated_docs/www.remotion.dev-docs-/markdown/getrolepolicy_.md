---
title: getRolePolicy()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getRolePolicy()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- getRolePolicy()

On this page

# getRolePolicy()

Returns an inline JSON policy to be assigned to the 'remotion-lambda-role' role that needs to be created in your AWS account.

These permissions will be given to the Lambda function itself.

See [Setup tutorial](/docs/lambda/setup) for setting up Lambda from scratch or [Role permissions](/docs/lambda/permissions#role-permissions) to see a copy of the current policy file with explanations.

## Example [​](\#example "Direct link to Example")

```

ts

import { getRolePolicy } from "@remotion/lambda";

console.log(getRolePolicy()); /* `
{
  "Version": "2012-10-17",
  "Statements": [
    // ...
  ]
}
` */
```

```

ts

import { getRolePolicy } from "@remotion/lambda";

console.log(getRolePolicy()); /* `
{
  "Version": "2012-10-17",
  "Statements": [
    // ...
  ]
}
` */
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/iam-validation/role-permissions.ts)
- [getUserPolicy()](/docs/lambda/getuserpolicy)
- [Permissions](/docs/lambda/permissions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/getrolepolicy.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getUserPolicy()](/docs/lambda/getuserpolicy) [Next\
\
getCompositionsOnLambda()](/docs/lambda/getcompositionsonlambda)

- [Example](#example)
- [See also](#see-also)
