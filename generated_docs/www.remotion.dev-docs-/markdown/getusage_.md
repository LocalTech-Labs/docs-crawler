---
title: getUsage()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getUsage()

- [Home page](/)
- [@remotion/licensing](/docs/licensing/)
- getUsage()

On this page

# getUsage()

Allows you to get the current usage of your Remotion license.

This requires your secret key that you can obtain from the remotion.pro dashboard.

You should only call this API from the backend to avoid exposing your secret key to the client.

```

Get the current usage of your license
tsx

import {getUsage} from '@remotion/licensing';

const usage = await getUsage({
  apiKey: 'rm_sec_xxxxx',
  since: Date.now() - 1000 * 60 * 60 * 24 * 30, // 30 days ago
});

console.log(usage);
/*
{
  "webcodecConversions": {
    "billable": 10,
    "development": 5,
    "failed": 2
  },
  "cloudRenders": {
    "billable": 10,
    "development": 5,
    "failed": 2
  },
}
*/
```

```

Get the current usage of your license
tsx

import {getUsage} from '@remotion/licensing';

const usage = await getUsage({
  apiKey: 'rm_sec_xxxxx',
  since: Date.now() - 1000 * 60 * 60 * 24 * 30, // 30 days ago
});

console.log(usage);
/*
{
  "webcodecConversions": {
    "billable": 10,
    "development": 5,
    "failed": 2
  },
  "cloudRenders": {
    "billable": 10,
    "development": 5,
    "failed": 2
  },
}
*/
```

## API [​](\#api "Direct link to API")

An object with the following properties:

### `apiKey` [​](\#apikey "Direct link to apikey")

Type: `string`

Your Remotion secret API key. You can get it from your Remotion.pro dashboard.

### `since` [​](\#since "Direct link to since")

Type: `number`

The timestamp since when you want to get the usage.

The default is since the beginning of the current month in UTC.

The lowest timestamp you can use is 90 days ago ( `Date.now() - 90 * 24 * 60 * 60 * 1000`).

## Return value [​](\#return-value "Direct link to Return value")

An object with the following properties:

### `webcodecConversions` [​](\#webcodecconversions "Direct link to webcodecconversions")

An object with the following properties:

- `billable`: The number of billable webcodec conversions.
- `development`: The number of development webcodec conversions (on `localhost` or other local environments).
- `failed`: The number of failed webcodec conversions (you don't need to pay for them).

### `cloudRenders` [​](\#cloudrenders "Direct link to cloudrenders")

An object with the following properties:

- `billable`: The number of billable cloud renders.
- `development`: The number of development cloud renders.
- `failed`: The number of failed cloud renders.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/licensing`](/docs/licensing)
- [`registerUsagePoint()`](/docs/licensing/register-usage-point)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/licensing/get-usage.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
registerUsagePoint()](/docs/licensing/register-usage-point)

- [API](#api)
  - [`apiKey`](#apikey)
  - [`since`](#since)
- [Return value](#return-value)
  - [`webcodecConversions`](#webcodecconversions)
  - [`cloudRenders`](#cloudrenders)
- [See also](#see-also)
