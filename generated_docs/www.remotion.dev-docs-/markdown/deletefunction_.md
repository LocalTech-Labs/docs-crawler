---
title: deleteFunction()
source: Remotion Documentation
last_updated: 2024-12-22
---

# deleteFunction()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- deleteFunction()

On this page

# deleteFunction()

Deletes a deployed Lambda function based on its name.

To retrieve a list of functions, call [`getFunctions()`](/docs/lambda/getfunctions) first.

## Example [​](\#example "Direct link to Example")

```

ts

import { deleteFunction, getFunctions } from "@remotion/lambda";

const functions = await getFunctions({
  region: "us-east-1",
  compatibleOnly: false,
});
for (const fn of functions) {
  await deleteFunction({
    region: "us-east-1",
    functionName: fn.functionName,
  });
}
```

```

ts

import { deleteFunction, getFunctions } from "@remotion/lambda";

const functions = await getFunctions({
  region: "us-east-1",
  compatibleOnly: false,
});
for (const fn of functions) {
  await deleteFunction({
    region: "us-east-1",
    functionName: fn.functionName,
  });
}
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `region` [​](\#region "Direct link to region")

The [AWS region](/docs/lambda/region-selection) to which the function was deployed to.

### `functionName` [​](\#functionname "Direct link to functionname")

The name of the function to be deleted.

## Return value [​](\#return-value "Direct link to Return value")

Nothing. If the deletion failed, the function rejects with an error.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/delete-function.ts)
- [deployFunction()](/docs/lambda/deployfunction)
- [getFunctions()](/docs/lambda/getfunctions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/deletefunction.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
deployFunction()](/docs/lambda/deployfunction) [Next\
\
getFunctionInfo()](/docs/lambda/getfunctioninfo)

- [Example](#example)
- [Arguments](#arguments)
  - [`region`](#region)
  - [`functionName`](#functionname)
- [Return value](#return-value)
- [See also](#see-also)
