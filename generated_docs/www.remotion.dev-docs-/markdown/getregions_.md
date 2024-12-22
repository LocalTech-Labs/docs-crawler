---
title: getRegions()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getRegions()

- [Home page](/)
- [@remotion/lambda](/docs/lambda/api)
- getRegions()

On this page

# getRegions()

Gets an array of all supported AWS regions of this release of Remotion Lambda.

## Example [​](\#example "Direct link to Example")

```

tsx

const regions = getRegions();
// ["eu-central-1", "us-east-1"]
```

```

tsx

const regions = getRegions();
// ["eu-central-1", "us-east-1"]
```

## API [​](\#api "Direct link to API")

The function takes an optional object, with the following options:

### `enabledByDefaultOnly` [​](\#enabledbydefaultonly "Direct link to enabledbydefaultonly")

_available from v3.3.11_

Only return [the regions which are enabled by default](https://docs.aws.amazon.com/general/latest/gr/rande-manage.html) in a new AWS account.

## Return value [​](\#return-value "Direct link to Return value")

An array of supported regions by this release of Remotion Lambda.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/lambda/src/api/get-regions.ts)
- [Region selection](/docs/lambda/region-selection)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/getregions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getAwsClient()](/docs/lambda/getawsclient) [Next\
\
getSites()](/docs/lambda/getsites)

- [Example](#example)
- [API](#api)
  - [`enabledByDefaultOnly`](#enabledbydefaultonly)
- [Return value](#return-value)
- [See also](#see-also)
