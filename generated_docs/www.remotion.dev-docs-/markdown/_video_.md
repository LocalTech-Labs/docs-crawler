---
title: <Video>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Video>

- [Home page](/)
- [remotion](/docs/remotion)
- <Video>

On this page

# <Video>

Wraps the native [`<video>`](https://developer.mozilla.org/en-US/docs/Web/API/HTMLVideoElement) element to include video in your component that is synchronized with Remotion's time.

info

Prefer [`<OffthreadVideo>`](/docs/offthreadvideo) which during render is faster and supports more codecs.

## API [​](\#api "Direct link to API")

[Put a video file into the `public/` folder](/docs/assets) and use [`staticFile()`](/docs/staticfile) to reference it.

All the props that the native `<video>` element accepts (except `autoplay` and `controls`) will be forwarded (but of course not all are useful for Remotion). This means you can use all CSS to style the video.

```

tsx

import { AbsoluteFill, staticFile, Video } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Video src={staticFile("video.webm")} />
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, staticFile, Video } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Video src={staticFile("video.webm")} />
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
      <Video src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />
    </AbsoluteFill>
  );
};
```

```

tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Video src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />
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
      <Video src={staticFile("video.webm")} startFrom={60} endAt={120} />
    </AbsoluteFill>
  );
};
```

```

tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Video src={staticFile("video.webm")} startFrom={60} endAt={120} />
    </AbsoluteFill>
  );
};
```

### `endAt?` [​](\#endat "Direct link to endat")

Removes a portion of the video at the end. See [`startFrom`](/docs/video#startfrom) for an explanation.

### `style?` [​](\#style "Direct link to style")

You can pass any style you can pass to a native `<video>` element. For example, set its size:

```

tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Video
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
      <Video
        src={staticFile("video.webm")}
        style={{ height: 720, width: 1280 }}
      />
    </AbsoluteFill>
  );
};
```

### `volume?` [​](\#volume "Direct link to volume")

Allows you to control the volume for the whole track or change it on a per-frame basis. Refer to the [using audio](/docs/using-audio#controlling-volume) guide to learn how to use it.

```

Example using static volume
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Video volume={0.5} src={staticFile("video.webm")} />
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
      <Video volume={0.5} src={staticFile("video.webm")} />
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
      <Video
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
      <Video
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

Controls the `frame` which is returned when using the [`volume`](#volume) callback function and adding the [`loop`](#loop) attribute.

Can be either `"repeat"` (default, start from 0 on each iteration) or `"extend"` (keep increasing frames).

### `name?` [v4.0.71](https://github.com/remotion-dev/remotion/releases/v4.0.71) [​](\#name "Direct link to name")

A name and that will be shown as the label of the sequence in the timeline of the Remotion Studio. This property is purely for helping you keep track of items in the timeline.

### `playbackRate?` [v2.2.0](https://github.com/remotion-dev/remotion/releases/v2.2.0) [​](\#playbackrate "Direct link to playbackrate")

Controls the speed of the video. `1` is the default and means regular speed, `0.5` slows down the video so it's twice as long and `2` speeds up the video so it's twice as fast.

While Remotion doesn't limit the range of possible playback speeds, in development mode the [`HTMLMediaElement.playbackRate`](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/playbackRate) API is used which throws errors on extreme values. At the time of writing, Google Chrome throws an exception if the playback rate is below `0.0625` or above `16`.

```

Example of a video playing twice as fast
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Video playbackRate={2} src={staticFile("video.webm")} />
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
      <Video playbackRate={2} src={staticFile("video.webm")} />
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
      <Video
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
      <Video
        muted
        src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"
      />
    </AbsoluteFill>
  );
};
```

This has the benefit that Remotion will not have to download the video file during rendering in order to extract the audio from it.

### `loop?` [v3.2.29](https://github.com/remotion-dev/remotion/releases/v3.2.29) [​](\#loop "Direct link to loop")

Makes the video loop indefinitely.

```

Example of a looped video
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Video
        loop
        src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"
      />
    </AbsoluteFill>
  );
};
```

```

Example of a looped video
tsx

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Video
        loop
        src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"
      />
    </AbsoluteFill>
  );
};
```

### `acceptableTimeShiftInSeconds?` [v3.2.42](https://github.com/remotion-dev/remotion/releases/v3.2.42) [​](\#acceptabletimeshiftinseconds "Direct link to acceptabletimeshiftinseconds")

In the [Studio](/docs/terminology/studio) or in the [Remotion Player](/docs/player), Remotion will seek the video if it gets too much out of sync with Remotion's internal time - be it due to the video loading or the page being too slow to keep up in real-time. By default, a seek is triggered if `0.45` seconds of time shift is encountered. Using this prop, you can customize the threshold.

### `allowAmplificationDuringRender?` [v3.3.17](https://github.com/remotion-dev/remotion/releases/v3.3.17) [​](\#allowamplificationduringrender "Direct link to allowamplificationduringrender")

Make values for [`volume`](#volume) greater than `1` result in amplification during renders.

During Preview, the volume will be limited to `1`, since the browser cannot amplify audio.

### `toneFrequency?` [v4.0.47](https://github.com/remotion-dev/remotion/releases/v4.0.47) [​](\#tonefrequency "Direct link to tonefrequency")

Adjust the pitch of the audio - will only be applied during rendering.

Accepts a number between `0.01` and `2`, where `1` represents the original pitch. Values less than `1` will decrease the pitch, while values greater than `1` will increase it.

A `toneFrequency` of 0.5 would lower the pitch by half, and a `toneFrequency` of `1.5` would increase the pitch by 50%.

### `onError?` [​](\#onerror "Direct link to onerror")

Handle an error playing the video. From v3.3.89, if you pass an `onError` callback, then no exception will be thrown. Previously, the error could not be caught.

### `pauseWhenBuffering?` [v4.0.100](https://github.com/remotion-dev/remotion/releases/v4.0.100) [​](\#pausewhenbuffering "Direct link to pausewhenbuffering")

If set to `true` and the video is loading, the Player will enter into the [native buffering state](/docs/player/buffer-state). The default is `false`, but will become `true` in Remotion 5.0.

### `showInTimeline?` [v4.0.122](https://github.com/remotion-dev/remotion/releases/v4.0.122) [​](\#showintimeline "Direct link to showintimeline")

If set to `false`, no layer will be shown in the timeline of the Remotion Studio. The default is `true`.

### `delayRenderTimeoutInMilliseconds?` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#delayrendertimeoutinmilliseconds "Direct link to delayrendertimeoutinmilliseconds")

Customize the [timeout](/docs/delay-render#modifying-the-timeout) of the [`delayRender()`](/docs/delay-render) call that this component makes.

### `delayRenderRetries?` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#delayrenderretries "Direct link to delayrenderretries")

Customize the [number of retries](/docs/delay-render#retrying) of the [`delayRender()`](/docs/delay-render) call that this component makes.

### `onAutoPlayError?` [v4.0.187](https://github.com/remotion-dev/remotion/releases/v4.0.187) [​](\#onautoplayerror "Direct link to onautoplayerror")

A callback function that gets called when the video fails to play due to autoplay restrictions.

If you don't pass a callback, the video will be muted and be retried once.

This prop is useful if you want to handle the error yourself, e.g. for pausing the Player.

Read more here about [autoplay restrictions](/docs/player/autoplay).

### `crossOrigin?` [v4.0.190](https://github.com/remotion-dev/remotion/releases/v4.0.190) [​](\#crossorigin "Direct link to crossorigin")

Corresponds to the [`crossOrigin`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video#attr-crossorigin) attribute of the `<video>` element.

One of `"anonymous"`, `"use-credentials"` or `undefined`.

Default: `"anonymous"` if `onVideoFrame` is specified, `undefined`, otherwise.

## Speed up renders for video with silent audio [​](\#speed-up-renders-for-video-with-silent-audio "Direct link to Speed up renders for video with silent audio")

Remotion will download the whole video during render in order to mix its audio. If the video contains a silent audio track, you can add the muted property to signal to Remotion that it does not need to download the video and make the render more efficient.

## Codec support [​](\#codec-support "Direct link to Codec support")

See: [Which video formats does Remotion support?](/docs/miscellaneous/video-formats)

## Alternative: `<OffthreadVideo>` [​](\#alternative-offthreadvideo "Direct link to alternative-offthreadvideo")

[`<OffthreadVideo>`](/docs/offthreadvideo) is a drop-in alternative to `<Video>`. To decide which tag to use, see: [`<Video>` vs `<OffthreadVideo>`](/docs/video-vs-offthreadvideo)

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/video/Video.tsx)
- [`<Audio />`](/docs/audio)
- [`<OffthreadVideo />`](/docs/offthreadvideo)
- [`<Video>` vs `<OffthreadVideo>`](/docs/video-vs-offthreadvideo)
- [`Change the speed of a video over time`](/docs/miscellaneous/snippets/accelerated-video)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/video.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
VERSION](/docs/version) [Next\
\
watchStaticFile()](/docs/watchstaticfile)

- [API](#api)
- [Props](#props)
  - [`src`](#src)
  - [`startFrom?`](#startfrom)
  - [`endAt?`](#endat)
  - [`style?`](#style)
  - [`volume?`](#volume)
  - [`loopVolumeCurveBehavior?`](#loopvolumecurvebehavior)
  - [`name?`](#name)
  - [`playbackRate?`](#playbackrate)
  - [`muted?`](#muted)
  - [`loop?`](#loop)
  - [`acceptableTimeShiftInSeconds?`](#acceptabletimeshiftinseconds)
  - [`allowAmplificationDuringRender?`](#allowamplificationduringrender)
  - [`toneFrequency?`](#tonefrequency)
  - [`onError?`](#onerror)
  - [`pauseWhenBuffering?`](#pausewhenbuffering)
  - [`showInTimeline?`](#showintimeline)
  - [`delayRenderTimeoutInMilliseconds?`](#delayrendertimeoutinmilliseconds)
  - [`delayRenderRetries?`](#delayrenderretries)
  - [`onAutoPlayError?`](#onautoplayerror)
  - [`crossOrigin?`](#crossorigin)
- [Speed up renders for video with silent audio](#speed-up-renders-for-video-with-silent-audio)
- [Codec support](#codec-support)
- [Alternative: `<OffthreadVideo>`](#alternative-offthreadvideo)
- [See also](#see-also)
