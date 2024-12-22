---
title: useOffthreadVideoTexture()
source: Remotion Documentation
last_updated: 2024-12-22
---

# useOffthreadVideoTexture()

- [Home page](/)
- [@remotion/three](/docs/three)
- useOffthreadVideoTexture()

On this page

# useOffthreadVideoTexture()

_available from v4.0.83_

Allows you to use a video in React Three Fiber that is synchronized with Remotion's `useCurrentFrame()`, similar to [`useVideoTexture()`](/docs/use-video-texture), but uses [`<OffthreadVideo>`](/docs/offthreadvideo) instead.

This hook only works during rendering. In the Player and the Remotion Studio, use [`useVideoTexture()`](/docs/use-video-texture) instead. See below for an example.

This hook was designed to combat limitations of the default `<Video>` element that is used with `useVideoTexture()` hook.
See: [`<OffthreadVideo> vs <Video>`](/docs/video-vs-offthreadvideo)

The return type of it is a `THREE.Texture | null` which you can assign as a `map` to for example `meshBasicMaterial`. We recommend to only render the material when the texture is not `null` to prevent bugs.

## Example [​](\#example "Direct link to Example")

```

Simple usage (only works during rendering)
tsx

import {ThreeCanvas, useOffthreadVideoTexture} from '@remotion/three';
import {staticFile, useVideoConfig} from 'remotion';

const videoSrc = staticFile('/vid.mp4');

const My3DVideo = () => {
  const {width, height} = useVideoConfig();

  const videoTexture = useOffthreadVideoTexture({src: videoSrc});

  return (
    <ThreeCanvas width={width} height={height}>
      <mesh>
        {videoTexture ? <meshBasicMaterial map={videoTexture} /> : null}
      </mesh>
    </ThreeCanvas>
  );
};
```

```

Simple usage (only works during rendering)
tsx

import {ThreeCanvas, useOffthreadVideoTexture} from '@remotion/three';
import {staticFile, useVideoConfig} from 'remotion';

const videoSrc = staticFile('/vid.mp4');

const My3DVideo = () => {
  const {width, height} = useVideoConfig();

  const videoTexture = useOffthreadVideoTexture({src: videoSrc});

  return (
    <ThreeCanvas width={width} height={height}>
      <mesh>
        {videoTexture ? <meshBasicMaterial map={videoTexture} /> : null}
      </mesh>
    </ThreeCanvas>
  );
};
```

```

Use useVideoTexture() only during rendering
tsx

import {
  ThreeCanvas,
  useOffthreadVideoTexture,
  useVideoTexture,
} from '@remotion/three';
import {useRef} from 'react';
import {
  getRemotionEnvironment,
  staticFile,
  useVideoConfig,
  Video,
} from 'remotion';

const videoSrc = staticFile('/vid.mp4');

const useVideoOrOffthreadVideoTexture = (
  videoSrc: string,
  videoRef: React.RefObject<HTMLVideoElement | null>,
) => {
  if (getRemotionEnvironment().isRendering) {
    // eslint-disable-next-line react-hooks/rules-of-hooks
    return useOffthreadVideoTexture({src: videoSrc});
  }

  // eslint-disable-next-line react-hooks/rules-of-hooks
  return useVideoTexture(videoRef);
};

const My3DVideo = () => {
  const {width, height} = useVideoConfig();

  const videoRef = useRef<HTMLVideoElement | null>(null);
  const videoTexture = useVideoOrOffthreadVideoTexture(videoSrc, videoRef);

  return (
    <>
      {getRemotionEnvironment().isRendering ? null : (
        <Video
          ref={videoRef}
          src={videoSrc}
          style={{position: 'absolute', opacity: 0}}
        />
      )}
      <ThreeCanvas width={width} height={height}>
        <mesh>
          {videoTexture ? <meshBasicMaterial map={videoTexture} /> : null}
        </mesh>
      </ThreeCanvas>
    </>
  );
};
```

```

Use useVideoTexture() only during rendering
tsx

import {
  ThreeCanvas,
  useOffthreadVideoTexture,
  useVideoTexture,
} from '@remotion/three';
import {useRef} from 'react';
import {
  getRemotionEnvironment,
  staticFile,
  useVideoConfig,
  Video,
} from 'remotion';

const videoSrc = staticFile('/vid.mp4');

const useVideoOrOffthreadVideoTexture = (
  videoSrc: string,
  videoRef: React.RefObject<HTMLVideoElement | null>,
) => {
  if (getRemotionEnvironment().isRendering) {
    // eslint-disable-next-line react-hooks/rules-of-hooks
    return useOffthreadVideoTexture({src: videoSrc});
  }

  // eslint-disable-next-line react-hooks/rules-of-hooks
  return useVideoTexture(videoRef);
};

const My3DVideo = () => {
  const {width, height} = useVideoConfig();

  const videoRef = useRef<HTMLVideoElement | null>(null);
  const videoTexture = useVideoOrOffthreadVideoTexture(videoSrc, videoRef);

  return (
    <>
      {getRemotionEnvironment().isRendering ? null : (
        <Video
          ref={videoRef}
          src={videoSrc}
          style={{position: 'absolute', opacity: 0}}
        />
      )}
      <ThreeCanvas width={width} height={height}>
        <mesh>
          {videoTexture ? <meshBasicMaterial map={videoTexture} /> : null}
        </mesh>
      </ThreeCanvas>
    </>
  );
};
```

## API [​](\#api "Direct link to API")

Takes an object with the following properties:

### `src` [​](\#src "Direct link to src")

The video source. Can be a URL or a [`staticFile()`](/docs/staticfile).

### `playbackRate` [​](\#playbackrate "Direct link to playbackrate")

The playback rate of the video. Defaults to `1`.

### `transparent` [​](\#transparent "Direct link to transparent")

_optional, boolean_

If set to `true`, frames will be extracted as PNG, enabling transparency but also slowing down your render.

If set to `false` ( _default_), frames will be extracted as bitmap (BMP), which is faster.

### `toneMapped` [v4.0.117](https://github.com/remotion-dev/remotion/releases/v4.0.117) [​](\#tonemapped "Direct link to tonemapped")

Since Remotion 4.0.117, Remotion will adjust the colors of videos in different color spaces (such as HDR) when converting to RGB, to counteract color shifts.

Since the browser is painting in sRGB, this is necessary to ensure that the colors are displayed correctly.

This behavior is by default `true` and can be disabled by setting `toneMapped` to `false`.

Disabling tone mapping will speed up rendering, but may result in less vibrant colors.

Prior to Remotion 4.0.117, tone mapping was always disabled, and the `toneMapped` prop was not available.

## Looping a texture [​](\#looping-a-texture "Direct link to Looping a texture")

Like `<OffthreadVideo>`, looping a video [is not supported out of the box](/docs/offthreadvideo#looping-a-video) but can be achieved with the `<Loop>` component.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/three/src/use-offthread-video-texture.ts)
- [`useVideoTexture()`](/docs/use-video-texture)
- [`<ThreeCanvas />`](/docs/three-canvas)
- [`<OffthreadVideo> vs <Video>`](/docs/video-vs-offthreadvideo)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/use-offthread-video-texture.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
useVideoTexture()](/docs/use-video-texture) [Next\
\
Overview](/docs/skia/)

- [Example](#example)
- [API](#api)
  - [`src`](#src)
  - [`playbackRate`](#playbackrate)
  - [`transparent`](#transparent)
  - [`toneMapped`](#tonemapped)
- [Looping a texture](#looping-a-texture)
- [See also](#see-also)
