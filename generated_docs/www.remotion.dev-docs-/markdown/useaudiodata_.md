---
title: useAudioData()
source: Remotion Documentation
last_updated: 2024-12-22
---

# useAudioData()

- [Home page](/)
- [@remotion/media-utils](/docs/media-utils/)
- useAudioData()

On this page

# useAudioData()

_Part of the `@remotion/media-utils` package of helper functions._

This convenience function wraps the [`getAudioData()`](/docs/get-audio-data) function into a hook and does 3 things:

- Keeps the audio data in a state
- Wraps the function in a [`delayRender()` / `continueRender()`](/docs/data-fetching) pattern.
- Handles the case where the component gets unmounted while the fetching is in progress and a React error is thrown.

Using this function, you can elegantly render a component based on audio properties, for example together with the [`visualizeAudio()`](/docs/visualize-audio) function.

info

Remote audio files need to support [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS).

More info

- Remotion's origin is usually `http://localhost:3000`, but it may be different if rendering on Lambda or the port is busy.

- You can use [`getAudioDurationInSeconds()`](/docs/get-audio-duration-in-seconds) without the audio needing CORS.

- You can [disable CORS](/docs/chromium-flags#--disable-web-security) during renders.

## Arguments [​](\#arguments "Direct link to Arguments")

### `src` [​](\#src "Direct link to src")

A string pointing to an audio asset.

## Return value [​](\#return-value "Direct link to Return value")

`AudioData | null` \- An object containing audio data (see documentation of [`getAudioData()`](/docs/get-audio-data)) or `null` if the data has not been loaded.

## Example [​](\#example "Direct link to Example")

```

tsx

import { useAudioData } from "@remotion/media-utils";
import { staticFile } from "remotion";

export const MyComponent: React.FC = () => {
  const audioData = useAudioData(staticFile("music.mp3"));

  if (!audioData) {
    return null;
  }

  return <div>This file has a {audioData.sampleRate} sampleRate.</div>;
};
```

```

tsx

import { useAudioData } from "@remotion/media-utils";
import { staticFile } from "remotion";

export const MyComponent: React.FC = () => {
  const audioData = useAudioData(staticFile("music.mp3"));

  if (!audioData) {
    return null;
  }

  return <div>This file has a {audioData.sampleRate} sampleRate.</div>;
};
```

## Errors [​](\#errors "Direct link to Errors")

If you pass in a file that has no audio track, this hook will throw an error ( _from v4.0.75_) or lead to an unhandled rejection ( _until v4.0.74_).

To determine if a file has an audio track, you may use the [`getVideoMetadata()`](/docs/renderer/get-video-metadata#audiocodec) function on the server to reject a file if it has no audio track. To do so, check if the `audioCodec` field is `null`.

If you want to catch the error in the component, you need to make your own inline hook by taking the source code from the bottom of this page.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/media-utils/src/use-audio-data.ts)
- [`getAudioData()`](/docs/get-audio-data)
- [`visualizeAudio()`](/docs/visualize-audio)
- [Audio visualization](/docs/audio-visualization)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/use-audio-data.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getWaveformPortion()](/docs/get-waveform-portion) [Next\
\
useWindowedAudioData()](/docs/use-windowed-audio-data)

- [Arguments](#arguments)
  - [`src`](#src)
- [Return value](#return-value)
- [Example](#example)
- [Errors](#errors)
- [See also](#see-also)
