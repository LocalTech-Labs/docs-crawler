---
title: registerUsagePoint()
source: Remotion Documentation
last_updated: 2024-12-22
---

# registerUsagePoint()

- [Home page](/)
- [@remotion/licensing](/docs/licensing/)
- registerUsagePoint()

On this page

# registerUsagePoint()

Registers a usage point for your Remotion license.

Allows for accurate and up to date reporting and to track your usage on the Remotion dashboard.

note

[`@remotion/webcodecs` automatically calls this function for you.](/docs/webcodecs/telemetry)

You do not need to call this function if you are using `@remotion/webcodecs`.

```

Register a usage point
tsx

import {registerUsageEvent} from '@remotion/licensing';

await registerUsageEvent({
  apiKey: 'rm_pub_xxxxx',
  event: 'webcodec-conversion',
  host: 'https://myapp.com',
  succeeded: true,
});
```

```

Register a usage point
tsx

import {registerUsageEvent} from '@remotion/licensing';

await registerUsageEvent({
  apiKey: 'rm_pub_xxxxx',
  event: 'webcodec-conversion',
  host: 'https://myapp.com',
  succeeded: true,
});
```

## API [​](\#api "Direct link to API")

An object with the following properties:

### `apiKey` [​](\#apikey "Direct link to apikey")

Type: `string`

Your Remotion public API key. You can get it from your Remotion.pro dashboard.

### `event` [​](\#event "Direct link to event")

Type: `string`

The event you want to register. This can be one of the following:

- `webcodec-conversion`
- `cloud-render`

### `host` [​](\#host "Direct link to host")

The domain at here you host your app.

This should be the value of what `window.location.origin` evaluates to on your frontend.

If the host is `localhost` or similar, it will be registered as non-billable.

### `succeeded` [​](\#succeeded "Direct link to succeeded")

Whether the event was successful or not.

If the event was not successful, it will be registered as a non-billable event.

## Return value [​](\#return-value "Direct link to Return value")

A promise that resolves when the event was successfully registered.

The resolved object contains two properties:

### `billable` [​](\#billable "Direct link to billable")

Whether this was an event that should be billed.

### `classification` [​](\#classification "Direct link to classification")

Either, `"billable"`, `"development"` or `"failed"`.

You do not have to pay for failed renders or renders doing development.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/licensing`](/docs/licensing)
- [`getUsage()`](/docs/licensing/get-usage)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/licensing/register-usage-point.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/licensing/) [Next\
\
getUsage()](/docs/licensing/get-usage)

- [API](#api)
  - [`apiKey`](#apikey)
  - [`event`](#event)
  - [`host`](#host)
  - [`succeeded`](#succeeded)
- [Return value](#return-value)
  - [`billable`](#billable)
  - [`classification`](#classification)
- [See also](#see-also)
