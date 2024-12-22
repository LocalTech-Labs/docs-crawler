---
title: simulatePermissions()
source: Remotion Documentation
last_updated: 2024-12-22
---

# simulatePermissions()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- simulatePermissions()

On this page

# simulatePermissions()

Runs tests through the AWS Simulator ensuring that all the necessary permissions are set for the authenticated user.

The CLI equivalent is `npx remotion lambda policies validate`.

The function does not reject with an error if a permission is missing, rather the missing permission is indicated in the return value.

This function does only validate the validity of the **user policy**, not the **role policy**.

## Example [​](\#example "Direct link to Example")

```

ts

import { simulatePermissions } from "@remotion/lambda";

const { results } = await simulatePermissions({
  region: "us-east-1",
});

for (const result of results) {
  console.log(result.decision); // "allowed"
  console.log(result.name); // "iam:SimulatePrincipalPolicy"
}
```

```

ts

import { simulatePermissions } from "@remotion/lambda";

const { results } = await simulatePermissions({
  region: "us-east-1",
});

for (const result of results) {
  console.log(result.decision); // "allowed"
  console.log(result.name); // "iam:SimulatePrincipalPolicy"
}
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following properties:

### `region` [​](\#region "Direct link to region")

The [AWS region](/docs/lambda/region-selection) that you would like to query.

### `onSimulation` [​](\#onsimulation "Direct link to onsimulation")

_optional_

A callback function that gets called every time a new simulation has been executed. This allows you to react to new simulation results coming in much faster than waiting for the return value of the function. Example:

```

ts

import { simulatePermissions } from "@remotion/lambda";

const { results } = await simulatePermissions({
  region: "us-east-1",
  onSimulation: (result) => {
    console.log(result.decision); // "allowed"
    console.log(result.name); // "iam:SimulatePrincipalPolicy"
  },
});
```

```

ts

import { simulatePermissions } from "@remotion/lambda";

const { results } = await simulatePermissions({
  region: "us-east-1",
  onSimulation: (result) => {
    console.log(result.decision); // "allowed"
    console.log(result.name); // "iam:SimulatePrincipalPolicy"
  },
});
```

## Return value [​](\#return-value "Direct link to Return value")

**An array of objects** containing simulation results of each necessary permission. The objects contain the following keys:

### `decision` [​](\#decision "Direct link to decision")

Either `"allowed"`, `"implicitDeny"` or `"explicitDeny"`.

### `name` [​](\#name "Direct link to name")

The identifier of the required permission. See the [Permissions page](/docs/lambda/permissions) to see a list of required permissions.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/iam-validation/simulate.ts)
- [getUserPolicy()](/docs/lambda/getuserpolicy)
- [getRolePolicy()](/docs/lambda/getrolepolicy)
- [Permissions](/docs/lambda/permissions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/simulatepermissions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
renderStillOnLambda()](/docs/lambda/renderstillonlambda) [Next\
\
speculateFunctionName()](/docs/lambda/speculatefunctionname)

- [Example](#example)
- [Arguments](#arguments)
  - [`region`](#region)
  - [`onSimulation`](#onsimulation)
- [Return value](#return-value)
  - [`decision`](#decision)
  - [`name`](#name)
- [See also](#see-also)
