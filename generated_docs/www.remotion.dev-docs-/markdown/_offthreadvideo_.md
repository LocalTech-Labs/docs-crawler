---
title: <OffthreadVideo>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <OffthreadVideo>

- [Home page](/)
- [remotion](/docs/remotion)
- <OffthreadVideo>

On this page

# <OffthreadVideo>

_Available from Remotion 3.0.11_

This component imports and displays a video, similar to [`<Video/>`](/docs/video), but during rendering, extracts the exact frame from the video and displays it in a [`<Img>`](/docs/img) tag. This extraction process happens outside the browser using FFmpeg.

This component was designed to combat limitations of the default `<Video>` element. See: [`<Video>` vs `<OffthreadVideo>`](/docs/video-vs-offthreadvideo).

## Example [​](\#example "Direct link to Example")

```

tsx

import { AbsoluteFill, OffthreadVideo, staticFile } from "remotion";

export const MyVideo = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo src={staticFile("video.webm")} />
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, OffthreadVideo, staticFile } from "remotion";

export const MyVideo = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo src={staticFile("video.webm")} />
    </AbsoluteFill>
  );
};
```

You can load a video from an URL as well:

```

tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />
    </AbsoluteFill>
  );
};
```

```

tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />
    </AbsoluteFill>
  );
};
```

## Props [​](\#props "Direct link to Props")

### `src` [​](\#src "Direct link to src")

The URL of the video to be rendered. Can be a remote URL or a local file referenced with [`staticFile()`](/docs/staticfile).

### `startFrom?` [​](\#startfrom "Direct link to startfrom")

Will remove a portion of the video at the beginning.

In the following example, we assume that the [`fps`](/docs/composition#fps) of the composition is `30`.

By passing `startFrom={60}`, the playback starts immediately, but with the first 2 seconds of the video trimmed away.

By passing `endAt={120}`, any video after the 4 second mark in the file will be trimmed away.

The video will play the range from `00:02:00` to `00:04:00`, meaning the video will play for 2 seconds.

```

tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo
        src={staticFile("video.webm")}
        startFrom={60}
        endAt={120}
      />
    </AbsoluteFill>
  );
};
```

```

tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo
        src={staticFile("video.webm")}
        startFrom={60}
        endAt={120}
      />
    </AbsoluteFill>
  );
};
```

### `endAt?` [​](\#endat "Direct link to endat")

Removes a portion of the video at the end. See [`startAt`](/docs/video#startfrom) for an explanation.

### `transparent?` [v4.0.0](https://github.com/remotion-dev/remotion/releases/v4.0.0) [​](\#transparent "Direct link to transparent")

If set to `true`, frames will be extracted as PNG, enabling transparency but also slowing down your render.

If set to `false` ( _default_), frames will be extracted as bitmap (BMP), which is faster.

### `volume?` [​](\#volume "Direct link to volume")

Allows you to control the volume for the whole track or change it on a per-frame basis. Refer to the [using audio](/docs/using-audio#controlling-volume) guide to learn how to use it.

```

Example using static volume
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo volume={0.5} src={staticFile("video.webm")} />
    </AbsoluteFill>
  );
};
```

```

Example using static volume
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo volume={0.5} src={staticFile("video.webm")} />
    </AbsoluteFill>
  );
};
```

```

Example of a ramp up over 100 frames
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo
        volume={(f) =>
          interpolate(f, [0, 100], [0, 1], { extrapolateLeft: "clamp" })
        }
        src={staticFile("video.webm")}
      />
    </AbsoluteFill>
  );
};
```

```

Example of a ramp up over 100 frames
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo
        volume={(f) =>
          interpolate(f, [0, 100], [0, 1], { extrapolateLeft: "clamp" })
        }
        src={staticFile("video.webm")}
      />
    </AbsoluteFill>
  );
};
```

note

On iOS Safari, it's not possible to granularly change the volume of a media tag.

Only values `0` and `1` will be respected by the browser.

### `loopVolumeCurveBehavior?` [v4.0.142](https://github.com/remotion-dev/remotion/releases/v4.0.142) [​](\#loopvolumecurvebehavior "Direct link to loopvolumecurvebehavior")

Controls the `frame` which is returned when using the [`volume`](#volume) callback function and wrapping `OffthreadVideo` in a [`<Loop>`](/docs/loop).

Can be either `"repeat"` (default, start from 0 on each iteration) or `"extend"` (keep increasing frames).

### `style?` [​](\#style "Direct link to style")

You can pass any style you can pass to a native HTML element. Keep in mind that during rendering, `<OffthreadVideo>` renders an [`<Img>`](/docs/img) tag, but a `<video>` tag is used during preview.

```

tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Img
        src={staticFile("video.webm")}
        style={{ height: 720, width: 1280 }}
      />
    </AbsoluteFill>
  );
};
```

```

tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Img
        src={staticFile("video.webm")}
        style={{ height: 720, width: 1280 }}
      />
    </AbsoluteFill>
  );
};
```

### `allowAmplificationDuringRender?` [v3.3.17](https://github.com/remotion-dev/remotion/releases/v3.3.17) [​](\#allowamplificationduringrender "Direct link to allowamplificationduringrender")

Make values for [`volume`](/docs/video#volume) greater than `1` result in amplification during renders.

During Preview, the volume will be limited to `1`, since the browser cannot amplify audio.

### `name?` [v4.0.71](https://github.com/remotion-dev/remotion/releases/v4.0.71) [​](\#name "Direct link to name")

A name and that will be shown as the label of the sequence in the timeline of the Remotion Studio. This property is purely for helping you keep track of items in the timeline.

### `toneFrequency?` [v4.0.47](https://github.com/remotion-dev/remotion/releases/v4.0.47) [​](\#tonefrequency "Direct link to tonefrequency")

Adjust the pitch of the audio - will only be applied during rendering.

Accepts a number between `0.01` and `2`, where `1` represents the original pitch. Values less than `1` will decrease the pitch, while values greater than `1` will increase it.

A `toneFrequency` of 0.5 would lower the pitch by half, and a `toneFrequency` of `1.5` would increase the pitch by 50%.

### `onError?` [​](\#onerror "Direct link to onerror")

Handle an error playing the video. From v3.3.89, if you pass an `onError` callback, then no exception will be thrown. Previously, the error could not be caught.

### `playbackRate?` [v2.2.0](https://github.com/remotion-dev/remotion/releases/v2.2.0) [​](\#playbackrate "Direct link to playbackrate")

Controls the speed of the video. `1` is the default and means regular speed, `0.5` slows down the video so it's twice as long and `2` speeds up the video so it's twice as fast.

While Remotion doesn't limit the range of possible playback speeds, in development mode the [`HTMLMediaElement.playbackRate`](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/playbackRate) API is used which throws errors on extreme values. At the time of writing, Google Chrome throws an exception if the playback rate is below `0.0625` or above `16`.

```

Example of a video playing twice as fast
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo playbackRate={2} src={staticFile("video.webm")} />
    </AbsoluteFill>
  );
};
```

```

Example of a video playing twice as fast
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo playbackRate={2} src={staticFile("video.webm")} />
    </AbsoluteFill>
  );
};
```

### `muted?` [​](\#muted "Direct link to muted")

You can drop the audio of the video by adding a `muted` prop:

```

Example of a muted video
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo
        muted
        src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"
      />
    </AbsoluteFill>
  );
};
```

```

Example of a muted video
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <OffthreadVideo
        muted
        src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"
      />
    </AbsoluteFill>
  );
};
```

### `acceptableTimeShiftInSeconds?` [v3.2.42](https://github.com/remotion-dev/remotion/releases/v3.2.42) [​](\#acceptabletimeshiftinseconds "Direct link to acceptabletimeshiftinseconds")

In the [Studio](/docs/terminology/studio) or in the [Remotion Player](/docs/player), Remotion will seek the video if it gets too much out of sync with Remotion's internal time - be it due to the video loading or the page being too slow to keep up in real-time. By default, a seek is triggered if `0.45` seconds of time shift is encountered. Using this prop, you can customize the threshold.

### `allowAmplificationDuringRender?` [v3.3.17](https://github.com/remotion-dev/remotion/releases/v3.3.17) [​](\#allowamplificationduringrender-1 "Direct link to allowamplificationduringrender-1")

Make values for [`volume`](#volume) greater than `1` result in amplification during renders.

During Preview, the volume will be limited to `1`, since the browser cannot amplify audio.

### `toneFrequency?` [v4.0.47](https://github.com/remotion-dev/remotion/releases/v4.0.47) [​](\#tonefrequency-1 "Direct link to tonefrequency-1")

Adjust the pitch of the audio - will only be applied during rendering.

Accepts a number between `0.01` and `2`, where `1` represents the original pitch. Values less than `1` will decrease the pitch, while values greater than `1` will increase it.

A `toneFrequency` of 0.5 would lower the pitch by half, and a `toneFrequency` of `1.5` would increase the pitch by 50%.

### `pauseWhenBuffering?` [v4.0.111](https://github.com/remotion-dev/remotion/releases/v4.0.111) [​](\#pausewhenbuffering "Direct link to pausewhenbuffering")

If set to `true` and the video is loading, the Player will enter into the [native buffering state](/docs/player/buffer-state). The default is `false`, but will become `true` in Remotion 5.0.

### `toneMapped?` [v4.0.117](https://github.com/remotion-dev/remotion/releases/v4.0.117) [​](\#tonemapped "Direct link to tonemapped")

Since Remotion 4.0.117, Remotion will adjust the colors of videos in different color spaces (such as HDR) when converting to RGB, to counteract color shifts.

Since the browser is painting in sRGB, this is necessary to ensure that the colors are displayed correctly.

This behavior is by default `true` and can be disabled by setting `toneMapped` to `false`.

Disabling tone mapping will speed up rendering, but may result in less vibrant colors.

Prior to Remotion 4.0.117, tone mapping was always disabled, and the `toneMapped` prop was not available.

### `showInTimeline?` [v4.0.122](https://github.com/remotion-dev/remotion/releases/v4.0.122) [​](\#showintimeline "Direct link to showintimeline")

If set to `false`, no layer will be shown in the timeline of the Remotion Studio. The default is `true`.

### `delayRenderTimeoutInMilliseconds?` [v4.0.150](https://github.com/remotion-dev/remotion/releases/v4.0.150) [​](\#delayrendertimeoutinmilliseconds "Direct link to delayrendertimeoutinmilliseconds")

Customize the [timeout](/docs/delay-render#modifying-the-timeout) of the [`delayRender()`](/docs/delay-render) call that this component makes.

### `delayRenderRetries?` [v4.0.178](https://github.com/remotion-dev/remotion/releases/v4.0.178) [​](\#delayrenderretries "Direct link to delayrenderretries")

Customize the [number of retries](/docs/delay-render#retrying) of the [`delayRender()`](/docs/delay-render) call that this component makes.

### `onAutoPlayError?` [v4.0.187](https://github.com/remotion-dev/remotion/releases/v4.0.187) [​](\#onautoplayerror "Direct link to onautoplayerror")

A callback function that gets called when the video fails to play due to autoplay restrictions.

If you don't pass a callback, the video will be muted and be retried once.

This prop is useful if you want to handle the error yourself, e.g. for pausing the Player.

Read more here about [autoplay restrictions](/docs/player/autoplay).

### `onVideoFrame?` [v4.0.190](https://github.com/remotion-dev/remotion/releases/v4.0.190) [​](\#onvideoframe "Direct link to onvideoframe")

A callback function that gets called when a frame is extracted from the video.

Useful for [video manipulation](/docs/video-manipulation).

The callback is called with a [`CanvasImageSource`](https://developer.mozilla.org/en-US/docs/Web/API/CanvasImageSource) object.

During preview, this is a `HTMLVideoElement` object, during rendering, it is an `HTMLImageElement`.

### `crossOrigin?` [v4.0.190](https://github.com/remotion-dev/remotion/releases/v4.0.190) [​](\#crossorigin "Direct link to crossorigin")

Corresponds to the [`crossOrigin`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video#attr-crossorigin) attribute of the `<video>` element.

One of `"anonymous"`, `"use-credentials"` or `undefined`.

Default: `"anonymous"` if `onVideoFrame` is specified, `undefined`, otherwise.

### `imageFormat` [v3.0.22](https://github.com/remotion-dev/remotion/releases/v3.0.22) [​](\#imageformat- "Direct link to imageformat-")

_removed in v4.0.0_

Either `jpeg` or `png`. Default `jpeg`.

With `png`, transparent videos (VP8, VP9, ProRes) can be displayed, however it is around 40% slower, with VP8 videos being [much slower](/docs/slow-method-to-extract-frame).

### Other props [​](\#other-props "Direct link to Other props")

The props [`onError`](/docs/img#onerror), `className` and `style` are supported and get passed to the underlying HTML element. Remember that during render, this is a `<img>` element, and during Preview, this is a `<video>` element.

## Performance tips [​](\#performance-tips "Direct link to Performance tips")

Only set `transparent` to `true` if you need transparency. It is slower than non-transparent frame extraction.

If you don't care about color accuracy, you can set `toneMapped` to `false` as well to save time on color conversion.

## Looping a video [​](\#looping-a-video "Direct link to Looping a video")

Unlike [`<Video>`](/docs/video), `OffthreadVideo` does not currently implement the `loop` property.

You can use the following snippet that uses [`@remotion/media-utils`](/docs/media-utils/) to loop a video.

Note that this will mount a `<video>` tag in the browser, meaning only codecs supported by the browser can be used.

```

LoopedOffthreadVideo.tsx
tsx

import {
  Loop,
  OffthreadVideo,
  useVideoConfig,
} from "remotion";

export const LoopedOffthreadVideo: React.FC<{
  durationInSeconds: number | null;
  src: string;
}> = ({durationInSeconds, src}) => {
  const { fps } = useVideoConfig();

  if (durationInSeconds === null) {
    return null;
  }

  return (
    <Loop durationInFrames={Math.floor(fps * durationInSeconds)}>
      <OffthreadVideo src={src} />
    </Loop>
  );
};
```

```

LoopedOffthreadVideo.tsx
tsx

import {
  Loop,
  OffthreadVideo,
  useVideoConfig,
} from "remotion";

export const LoopedOffthreadVideo: React.FC<{
  durationInSeconds: number | null;
  src: string;
}> = ({durationInSeconds, src}) => {
  const { fps } = useVideoConfig();

  if (durationInSeconds === null) {
    return null;
  }

  return (
    <Loop durationInFrames={Math.floor(fps * durationInSeconds)}>
      <OffthreadVideo src={src} />
    </Loop>
  );
};
```

```

Root.tsx
tsx

import React, { useState } from "react";
import { Composition, staticFile } from "remotion";
import { getVideoMetadata } from "@remotion/media-utils";

export const RemotionRoot: React.FC = () => {
  return (
    <Composition
      id="MyComp"
      component={LoopedOffthreadVideo}
      defaultProps={{
        durationInSeconds: null,
        src: staticFile("myvideo.mp4"),
      }}
      calculateMetadata={async ({props}) => {
        const { durationInSeconds, width, height } = await getVideoMetadata(props.src);
        const fps = 30;

        return {
          // Set any duration, here as an example, we loop the video 5 times
          durationInFrames: Math.floor(durationInSeconds * fps * 5),
          fps,
          width,
          height,
          props: {
            ...props,
            // Pass the durationInSeconds as prop to the React component
            durationInSeconds,
          }
        };
      }}
    />
  )
}
```

```

Root.tsx
tsx

import React, { useState } from "react";
import { Composition, staticFile } from "remotion";
import { getVideoMetadata } from "@remotion/media-utils";

export const RemotionRoot: React.FC = () => {
  return (
    <Composition
      id="MyComp"
      component={LoopedOffthreadVideo}
      defaultProps={{
        durationInSeconds: null,
        src: staticFile("myvideo.mp4"),
      }}
      calculateMetadata={async ({props}) => {
        const { durationInSeconds, width, height } = await getVideoMetadata(props.src);
        const fps = 30;

        return {
          // Set any duration, here as an example, we loop the video 5 times
          durationInFrames: Math.floor(durationInSeconds * fps * 5),
          fps,
          width,
          height,
          props: {
            ...props,
            // Pass the durationInSeconds as prop to the React component
            durationInSeconds,
          }
        };
      }}
    />
  )
}
```

## Supported codecs by `<OffthreadVideo>` [​](\#supported-codecs-by-offthreadvideo "Direct link to supported-codecs-by-offthreadvideo")

The following codecs can be read by `<OffthreadVideo>`:

- H.264 ("MP4")
- H.265 ("HEVC")
- VP8 and VP9 ("WebM")
- AV1 (from _v4.0.6_)
- ProRes

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/video/OffthreadVideo.tsx)
- [`<Video />`](/docs/video)
- [`<Video>` vs `<OffthreadVideo>`](/docs/video-vs-offthreadvideo)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/offthreadvideo.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
measureSpring()](/docs/measure-spring) [Next\
\
prefetch()](/docs/prefetch)

- [Example](#example)
- [Props](#props)
  - [`src`](#src)
  - [`startFrom?`](#startfrom)
  - [`endAt?`](#endat)
  - [`transparent?`](#transparent)
  - [`volume?`](#volume)
  - [`loopVolumeCurveBehavior?`](#loopvolumecurvebehavior)
  - [`style?`](#style)
  - [`allowAmplificationDuringRender?`](#allowamplificationduringrender)
  - [`name?`](#name)
  - [`toneFrequency?`](#tonefrequency)
  - [`onError?`](#onerror)
  - [`playbackRate?`](#playbackrate)
  - [`muted?`](#muted)
  - [`acceptableTimeShiftInSeconds?`](#acceptabletimeshiftinseconds)
  - [`allowAmplificationDuringRender?`](#allowamplificationduringrender-1)
  - [`toneFrequency?`](#tonefrequency-1)
  - [`pauseWhenBuffering?`](#pausewhenbuffering)
  - [`toneMapped?`](#tonemapped)
  - [`showInTimeline?`](#showintimeline)
  - [`delayRenderTimeoutInMilliseconds?`](#delayrendertimeoutinmilliseconds)
  - [`delayRenderRetries?`](#delayrenderretries)
  - [`onAutoPlayError?`](#onautoplayerror)
  - [`onVideoFrame?`](#onvideoframe)
  - [`crossOrigin?`](#crossorigin)
  - [`imageFormat`](#imageformat-)
  - [Other props](#other-props)
- [Performance tips](#performance-tips)
- [Looping a video](#looping-a-video)
- [Supported codecs by `<OffthreadVideo>`](#supported-codecs-by-offthreadvideo)
- [See also](#see-also)
