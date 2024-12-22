---
title: getUserPolicy()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getUserPolicy()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- getUserPolicy()

On this page

# getUserPolicy()

Returns an inline JSON policy to be assigned to the AWS user whose credentials are being used for excuting CLI commands or calling Node.JS functions.

See [Setup tutorial](/docs/lambda/setup) for setting up Lambda from scratch or [User permissions](/docs/lambda/permissions#user-permissions) to see a copy of the current policy file with explanations.

## Example [​](\#example "Direct link to Example")

```

ts

import { getUserPolicy } from "@remotion/lambda";

console.log(getUserPolicy()); /* `
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

import { getUserPolicy } from "@remotion/lambda";

console.log(getUserPolicy()); /* `
{
  "Version": "2012-10-17",
  "Statements": [
    // ...
  ]
}
` */
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/iam-validation/user-permissions.ts)
- [getRolePolicy()](/docs/lambda/getrolepolicy)
- [Permissions](/docs/lambda/permissions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/getuserpolicy.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
deleteRender()](/docs/lambda/deleterender) [Next\
\
getRolePolicy()](/docs/lambda/getrolepolicy)

- [Example](#example)
- [See also](#see-also)
