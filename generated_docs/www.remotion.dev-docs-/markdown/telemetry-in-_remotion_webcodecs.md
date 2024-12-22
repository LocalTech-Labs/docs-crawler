---
title: Telemetry in @remotion/webcodecs
source: Remotion Documentation
last_updated: 2024-12-22
---

# Telemetry in @remotion/webcodecs

- [Home page](/)
- [WebCodecs](/docs/webcodecs/)
- Telemetry in @remotion/webcodecs

On this page

# Telemetry in @remotion/webcodecs

Upon finishing a conversion with [`convertMedia()`](/docs/webcodecs/convert-media), a HTTP request is sent to an endpoint on [`remotion.pro`](https://remotion.pro), registering the domain and whether the conversion was successful.

Telemetry will never cause any conversions to fail.

💼 Important License Disclaimer

This package is licensed under the [Remotion License](/docs/license).

We consider a team of 4 or more people a "company".

**For "companies"**: A Remotion Company license needs to be obtained to use this package.

In a future version of`@remotion/webcodecs`, this package will also require the purchase of a newly created "WebCodecs Conversion Seat". [Get in touch](/contact) with us if you are planning to use this package.

**For individuals and teams up to 3:** You can use this package for free.

This is a short, non-binding explanation of our license. See the [License](/docs/license) itself for more details.

## Data collection [​](\#data-collection "Direct link to Data collection")

No user data is collected, however the IP address may be used to reject invalid telemetry.

## Use of data [​](\#use-of-data "Direct link to Use of data")

Telemetry data helps Remotion understand which sites use the [`@remotion/webcodecs`](/docs/webcodecs) package.

Historically, it's been difficult to understand which sites use Remotion and various sizable companies have been using it without getting the appropriate license.

Hence we have decided to add telemetry to this new package.

Remotion will not share any telemetry data with others.

## Track conversions with an API key [​](\#track-conversions-with-an-api-key "Direct link to Track conversions with an API key")

You may add an API key that you obtained from [`remotion.pro`](https://remotion.pro) to the [`convertMedia()`](/docs/webcodecs/convert-media) function.

```

Adding an API key
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  apiKey: 'rm_pub_abcdefghijklo',
});
```

```

Adding an API key
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  apiKey: 'rm_pub_abcdefghijklo',
});
```

You will then be able to see the amount of conversions you did on your dashboard.

For now, you still need to adjust the license manually to cover your WebCodecs conversions.

## Disabling telemetry [​](\#disabling-telemetry "Direct link to Disabling telemetry")

A version of [`@remotion/webcodecs`](/docs/webcodecs) without telemetry can be obtained with an Remotion Enterprise License.

Contact [hi@remotion.dev](mailto:hi@remotion.dev) to request an Enterprise License.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/licensing`](/docs/licensing)
- [`convertMedia()`](/docs/webcodecs/convert-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/telemetry.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Fixing a MediaRecorder video](/docs/webcodecs/fix-mediarecorder-video) [Next\
\
Installation in existing project](/docs/brownfield)

- [Data collection](#data-collection)
- [Use of data](#use-of-data)
- [Track conversions with an API key](#track-conversions-with-an-api-key)
- [Disabling telemetry](#disabling-telemetry)
- [See also](#see-also)
