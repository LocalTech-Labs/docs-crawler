---
title: getAudioData()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getAudioData()

- [Home page](/)
- [@remotion/media-utils](/docs/media-utils/)
- getAudioData()

On this page

# getAudioData()

_Part of the `@remotion/media-utils` package of helper functions._

Takes an audio `src`, loads it and returns data and metadata for the specified source.

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

## Arguments [​](\#arguments "Direct link to Arguments")

### `src` [​](\#src "Direct link to src")

A string pointing to an audio asset.

### `options?` [v4.0.121](https://github.com/remotion-dev/remotion/releases/v4.0.121) [​](\#options "Direct link to options")

#### `sampleRate?` [v4.0.121](https://github.com/remotion-dev/remotion/releases/v4.0.121) [​](\#samplerate "Direct link to samplerate")

The `sampleRate` that should be passed into the [`AudioContext`](https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/AudioContext) constructor. If not provided, the default value is `48000`.

In versions before 4.0.121, the default value was `undefined`, leading to undeterministic behavior across devices rendering.

## Return value [​](\#return-value "Direct link to Return value")

_`Promise<AudioData>`_

An object with information about the audio data:

### `channelWaveforms` [​](\#channelwaveforms "Direct link to channelwaveforms")

_Float32Array\[\]_

An array with waveform information for each channel.

### `sampleRate` [​](\#samplerate-1 "Direct link to samplerate-1")

_number_

The sample rate of the generated `AudioContext`. This will be the same as the `sampleRate` input option if passed, `48000` otherwise, and in version previous to 4.0.121, the sample rate of the device's preferred output device.

note

Previously, this documentation stated that this is the sample rate of the audio file. This is incorrect. [The sample rate of the audio file is not exposed to the browser's JavaScript environment.](https://github.com/WebAudio/web-audio-api/issues/30)

### `durationInSeconds` [​](\#durationinseconds "Direct link to durationinseconds")

_number_

The duration of the audio in seconds.

### `numberOfChannels` [​](\#numberofchannels "Direct link to numberofchannels")

_number_

The number of channels contained in the audio file. This corresponds to the length of the `channelWaveforms` array.

### `resultId` [​](\#resultid "Direct link to resultid")

_string_ Unique identifier of this audio data fetching call. Other functions can cache expensive operations if they get called with the same resultId multiple times.

### `isRemote` [​](\#isremote "Direct link to isremote")

_boolean_

Whether the audio was imported locally or from a different origin.

## Example [​](\#example "Direct link to Example")

```

ts

import {getAudioData} from '@remotion/media-utils';
import music from './music.mp3';

await getAudioData(music); /* {
  channelWaveforms: [Float32Array(4410000), Float32Array(4410000)],
  sampleRate: 44100,
  durationInSeconds: 100.0000,
  numberOfChannels: 2,
  resultId: "0.432878981",
  isRemote: false
} */
await getAudioData('https://example.com/remote-audio.aac'); /* {
  channelWaveforms: [Float32Array(4800000)],
  sampleRate: 48000,
  durationInSeconds: 100.0000,
  numberOfChannels: 1,
  resultId: "0.432324444",
  isRemote: true
} */
await getAudioData(staticFile('my-file.wav')); /* {
  channelWaveforms: [Float32Array(4800000)],
  sampleRate: 48000,
  durationInSeconds: 100.0000,
  numberOfChannels: 1,
  resultId: "0.6891332223",
  isRemote: false
} */
```

```

ts

import {getAudioData} from '@remotion/media-utils';
import music from './music.mp3';

await getAudioData(music); /* {
  channelWaveforms: [Float32Array(4410000), Float32Array(4410000)],
  sampleRate: 44100,
  durationInSeconds: 100.0000,
  numberOfChannels: 2,
  resultId: "0.432878981",
  isRemote: false
} */
await getAudioData('https://example.com/remote-audio.aac'); /* {
  channelWaveforms: [Float32Array(4800000)],
  sampleRate: 48000,
  durationInSeconds: 100.0000,
  numberOfChannels: 1,
  resultId: "0.432324444",
  isRemote: true
} */
await getAudioData(staticFile('my-file.wav')); /* {
  channelWaveforms: [Float32Array(4800000)],
  sampleRate: 48000,
  durationInSeconds: 100.0000,
  numberOfChannels: 1,
  resultId: "0.6891332223",
  isRemote: false
} */
```

## Errors [​](\#errors "Direct link to Errors")

If you pass in a file that has no audio track, this function will throw an error you need to handle.

To determine if a file has an audio track, you may use the [`getVideoMetadata()`](/docs/renderer/get-video-metadata#audiocodec) function on the server to reject a file if it has no audio track. To do so, check if the `audioCodec` field is `null`.

## Caching behavior [​](\#caching-behavior "Direct link to Caching behavior")

This function is memoizing the results it returns.

If you pass in the same argument to `src` multiple times, it will return a cached version from the second time on, regardless of if the file has changed.

To clear the cache, you have to reload the page.

## Alternatives [​](\#alternatives "Direct link to Alternatives")

If you need only the duration, prefer [`getAudioDurationInSeconds()`](/docs/get-audio-duration-in-seconds) which is faster because it doesn't need to read waveform data.

Use the [`useAudioData()`](/docs/use-audio-data) helper hook to not have to do state management yourself and to wrap the call in [`delayRender()`](/docs/delay-render).

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/media-utils/src/get-audio-data.ts)
- [Using audio](/docs/using-audio)
- [Audio visualization](/docs/audio-visualization)
- [`<Audio/>`](/docs/audio)
- [`visualizeAudio()`](/docs/visualize-audio)
- [`useAudioData()`](/docs/use-audio-data)
- [`useWindowedAudioData()`](/docs/use-windowed-audio-data)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/get-audio-data.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
audioBufferToDataUrl()](/docs/audio-buffer-to-data-url) [Next\
\
getAudioDurationInSeconds()](/docs/get-audio-duration-in-seconds)

- [Arguments](#arguments)
  - [`src`](#src)
  - [`options?`](#options)
- [Return value](#return-value)
  - [`channelWaveforms`](#channelwaveforms)
  - [`sampleRate`](#samplerate-1)
  - [`durationInSeconds`](#durationinseconds)
  - [`numberOfChannels`](#numberofchannels)
  - [`resultId`](#resultid)
  - [`isRemote`](#isremote)
- [Example](#example)
- [Errors](#errors)
- [Caching behavior](#caching-behavior)
- [Alternatives](#alternatives)
- [See also](#see-also)
