---
title: testPermissions()
source: Remotion Documentation
last_updated: 2024-12-22
---

# testPermissions()

- [Home page](/)
- [@remotion/cloudrun](/docs/cloudrun/api)
- testPermissions()

On this page

# testPermissions()

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

Makes a call to the [Test Iam Permissions](https://cloud.google.com/resource-manager/reference/rest/v3/projects/testIamPermissions) method of the Resource Manager API in GCP, which returns the list of permissions the Service Account has on the GCP Project. This is then validated against the list of permissions required for the version of Remotion.

The CLI equivalent is `npx remotion cloudrun permissions`.

The function does not reject with an error if a permission is missing, rather the missing permission is indicated in the return value.

## Example [​](\#example "Direct link to Example")

```

ts

import { testPermissions } from "@remotion/cloudrun";

const { results } = await testPermissions();

for (const result of results) {
  console.log(result.decision); // "allowed"
  console.log(result.permissionName); // "iam.serviceAccounts.actAs"
}
```

```

ts

import { testPermissions } from "@remotion/cloudrun";

const { results } = await testPermissions();

for (const result of results) {
  console.log(result.decision); // "allowed"
  console.log(result.permissionName); // "iam.serviceAccounts.actAs"
}
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following property:

### `onTest?` [​](\#ontest "Direct link to ontest")

_optional_

A callback function that gets called every time a new test has been executed. This allows you to react to new test results coming in much faster than waiting for the return value of the function. Example:

```

ts

import { testPermissions } from "@remotion/cloudrun";

const { results } = await testPermissions({
  onTest: (result) => {
    console.log(result.decision); // "allowed"
    console.log(result.permissionName); // "iam.serviceAccounts.actAs"
  },
});
```

```

ts

import { testPermissions } from "@remotion/cloudrun";

const { results } = await testPermissions({
  onTest: (result) => {
    console.log(result.decision); // "allowed"
    console.log(result.permissionName); // "iam.serviceAccounts.actAs"
  },
});
```

## Return value [​](\#return-value "Direct link to Return value")

**An array of objects** containing simulation results of each necessary permission. The objects contain the following keys:

### `decision` [​](\#decision "Direct link to decision")

Either `true` or `false`.

### `permissionName` [​](\#permissionname "Direct link to permissionname")

The identifier of the required permission. See the [Permissions page](/docs/cloudrun/permissions) to see a list of required permissions.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/cloudrun/src/api/iam-validation/testPermissions.ts)
- [Permissions](/docs/cloudrun/permissions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/testpermissions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
renderStillOnCloudrun()](/docs/cloudrun/renderstilloncloudrun) [Next\
\
Overview](/docs/tailwind/tailwind)

- [Example](#example)
- [Arguments](#arguments)
  - [`onTest?`](#ontest)
- [Return value](#return-value)
  - [`decision`](#decision)
  - [`permissionName`](#permissionname)
- [See also](#see-also)
