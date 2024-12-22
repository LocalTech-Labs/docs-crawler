---
title: <Img>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Img>

- [Home page](/)
- [remotion](/docs/remotion)
- <Img>

On this page

# <Img>

The `<Img>` tag can be used like a regular `<img>` HTML tag.

If you use `<Img>`, Remotion will ensure that the image is loaded before rendering the frame. That way you can avoid flickers if the image does not load immediately during rendering.

## API [​](\#api "Direct link to API")

### `src` [​](\#src "Direct link to src")

[Put an image into the `public/` folder](/docs/assets) and use [`staticFile()`](/docs/staticfile) to reference it.

```

tsx

import { AbsoluteFill, Img, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Img src={staticFile("hi.png")} />
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, Img, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Img src={staticFile("hi.png")} />
    </AbsoluteFill>
  );
};
```

You can also load a remote image:

```

tsx

import { AbsoluteFill, Img } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Img src={"https://picsum.photos/200/300"} />
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, Img } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Img src={"https://picsum.photos/200/300"} />
    </AbsoluteFill>
  );
};
```

### `onError` [​](\#onerror "Direct link to onerror")

To use the `<Img>` tag in a resilient way, handle the error that occurs when an image can not be loaded:

```

tsx

import { AbsoluteFill, Img, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Img
        src={staticFile("hi.png")}
        onError={(event) => {
          // Handle image loading error here
        }}
      />
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, Img, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Img
        src={staticFile("hi.png")}
        onError={(event) => {
          // Handle image loading error here
        }}
      />
    </AbsoluteFill>
  );
};
```

If an error occurs, the component must be unmounted or the `src` must be replaced, otherwise the render will time out.

From `v3.3.82`, the image load will first be retried before `onError` is thrown.

### `maxRetries` [v3.3.82](https://github.com/remotion-dev/remotion/releases/v3.3.82) [​](\#maxretries "Direct link to maxretries")

If an image fails to load, it will be retried from `v3.3.82`. The default value is `2`.

An exponential backoff is being used, with 1000ms delay between the first and second attempt, then 2000ms, then 4000ms and so on.

### `pauseWhenLoading` [v4.0.111](https://github.com/remotion-dev/remotion/releases/v4.0.111) [​](\#pausewhenloading "Direct link to pausewhenloading")

_optional_

If set to `true`, pause the Player when the image is loading. See: [Buffer state](/docs/player/buffer-state).

### `delayRenderTimeoutInMilliseconds` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#delayrendertimeoutinmilliseconds "Direct link to delayrendertimeoutinmilliseconds")

Customize the [timeout](/docs/delay-render#modifying-the-timeout) of the [`delayRender()`](/docs/delay-render) call that this component makes.

### `delayRenderRetries` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#delayrenderretries "Direct link to delayrenderretries")

Customize the [number of retries](/docs/delay-render#retrying) of the [`delayRender()`](/docs/delay-render) call that this component makes.

Prefer the [`maxRetries`](#maxretries) prop over this.

## Other props [​](\#other-props "Direct link to Other props")

Remotion inherits the props of the regular `<img>` tag, like for example `style`.

## GIFs [​](\#gifs "Direct link to GIFs")

Don't use the `<Img>` tag for GIFs, use [`@remotion/gif`](/docs/gif) instead.

## Error behavior [​](\#error-behavior "Direct link to Error behavior")

From v4.0.0: If the image fails to load and no retries are left, then [`cancelRender`](/docs/cancel-render) will be called to throw an error, unless you handle the error using `onError()`.

From v3.3.82: If an image fails to load, it will be retried up to two times.

In earlier versions, failing to load an image would lead to an error message in the console and an eventual timeout.

## Restrictions [​](\#restrictions "Direct link to Restrictions")

- The maximum resolution that Chrome can display is `2^29` pixels (539 megapixels) [\[source\]](https://stackoverflow.com/questions/57223559/what-is-the-maximum-image-dimensions-supported-in-desktop-chrome#:~:text=than%202%5E29-,(539MP)). Remotion inherits this restriction.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/Img.tsx)
- [Use `<Img>` and `<IFrame>` tags](/docs/use-img-and-iframe)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/img.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<IFrame>](/docs/iframe) [Next\
\
interpolateColors()](/docs/interpolate-colors)

- [API](#api)
  - [`src`](#src)
  - [`onError`](#onerror)
  - [`maxRetries`](#maxretries)
  - [`pauseWhenLoading`](#pausewhenloading)
  - [`delayRenderTimeoutInMilliseconds`](#delayrendertimeoutinmilliseconds)
  - [`delayRenderRetries`](#delayrenderretries)
- [Other props](#other-props)
- [GIFs](#gifs)
- [Error behavior](#error-behavior)
- [Restrictions](#restrictions)
- [See also](#see-also)
