---
title: useWindowedAudioData()v4.0.240
source: Remotion Documentation
last_updated: 2024-12-22
---

# useWindowedAudioData()v4.0.240

- [Home page](/)
- [@remotion/media-utils](/docs/media-utils/)
- useWindowedAudioData()

On this page

# useWindowedAudioData() [v4.0.240](https://github.com/remotion-dev/remotion/releases/v4.0.240)

_Part of the `@remotion/media-utils` package of helper functions._

This is an alternative to [`useAudioData()`](/docs/use-audio-data) that only loads a portion of the audio around the current frame.

It only works with `.wav` files.

Unlike [`useAudioData()`](/docs/use-audio-data), which keeps all of the audio data in memory, this function makes HTTP Range requests to only load the audio data around the current frame.

We recommend using this function for visualizing audio with a long duraiton.

info

Remote audio files need to support [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS).

More info

- Remotion's origin is usually `http://localhost:3000`, but it
may be different if rendering on Lambda or the port is busy.

- You can use
[`getAudioDurationInSeconds()`](/docs/get-audio-duration-in-seconds)
without the audio needing CORS.

- You can [disable CORS](/docs/chromium-flags#--disable-web-security)
during renders.

## Example [​](\#example "Direct link to Example")

```

tsx

import {useWindowedAudioData, visualizeAudio} from '@remotion/media-utils';
import {staticFile, useCurrentFrame, useVideoConfig} from 'remotion';

export const MyComponent: React.FC = () => {
  const {fps} = useVideoConfig();
  const frame = useCurrentFrame();
  const {audioData, dataOffsetInSeconds} = useWindowedAudioData({
    src: staticFile('podcast.wav'),
    frame,
    fps,
    windowInSeconds: 10,
  });

  if (!audioData) {
    return null;
  }

  const visualization = visualizeAudio({
    fps,
    frame,
    audioData,
    numberOfSamples: 16,
    dataOffsetInSeconds,
  });

  return null;
};
```

```

tsx

import {useWindowedAudioData, visualizeAudio} from '@remotion/media-utils';
import {staticFile, useCurrentFrame, useVideoConfig} from 'remotion';

export const MyComponent: React.FC = () => {
  const {fps} = useVideoConfig();
  const frame = useCurrentFrame();
  const {audioData, dataOffsetInSeconds} = useWindowedAudioData({
    src: staticFile('podcast.wav'),
    frame,
    fps,
    windowInSeconds: 10,
  });

  if (!audioData) {
    return null;
  }

  const visualization = visualizeAudio({
    fps,
    frame,
    audioData,
    numberOfSamples: 16,
    dataOffsetInSeconds,
  });

  return null;
};
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object with:

### `src` [​](\#src "Direct link to src")

A string pointing to an audio asset.

### `frame` [​](\#frame "Direct link to frame")

_`number`_

The current frame of the composition.

### `fps` [​](\#fps "Direct link to fps")

_`number`_

The frames per second of the composition. Should be taken from [`useVideoConfig()`](/docs/use-video-config).

### `windowInSeconds` [​](\#windowinseconds "Direct link to windowinseconds")

_`number`_

The audio will be segmented into windows of this length.

The function will load the audio data around the current frame and the windows before and after.

In this example, the window is 10 seconds long, so the function will load the current window as well as the previous and next one, leading to up to 30 seconds of audio being loaded at a time.

## Return value [​](\#return-value "Direct link to Return value")

An object with:

### `audioData` [​](\#audiodata "Direct link to audiodata")

_`AudioData | null`_

An object containing audio data (see documentation of [`getAudioData()`](/docs/get-audio-data)) or `null` if the data has not been loaded.

### `dataOffsetInSeconds` [​](\#dataoffsetinseconds "Direct link to dataoffsetinseconds")

_`number`_

The offset in seconds of the audio data that is currently loaded.

You should pass it through to [`visualizeAudio()`](/docs/visualize-audio).

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/media-utils/src/use-audio-data.ts)
- [`getAudioData()`](/docs/get-audio-data)
- [`visualizeAudio()`](/docs/visualize-audio)
- [Audio visualization](/docs/audio-visualization)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/use-windowed-audio-data.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
useAudioData()](/docs/use-audio-data) [Next\
\
visualizeAudio()](/docs/visualize-audio)

- [Example](#example)
- [Arguments](#arguments)
  - [`src`](#src)
  - [`frame`](#frame)
  - [`fps`](#fps)
  - [`windowInSeconds`](#windowinseconds)
- [Return value](#return-value)
  - [`audioData`](#audiodata)
  - [`dataOffsetInSeconds`](#dataoffsetinseconds)
- [See also](#see-also)
