---
title: getWaveformPortion()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getWaveformPortion()

- [Home page](/)
- [@remotion/media-utils](/docs/media-utils/)
- getWaveformPortion()

On this page

# getWaveformPortion()

_Part of the `@remotion/media-utils` package of helper functions._

Takes bulky waveform data (for example fetched by [`getAudioData()`](/docs/get-audio-data)) and returns a trimmed and simplified version of it, for simpler visualization. This function is suitable if you only need volume data, if you need more detailed data about each frequency range, use [`visualizeAudio()`](/docs/visualize-audio).

## Arguments [​](\#arguments "Direct link to Arguments")

An object with the following arguments:

### `audioData` [​](\#audiodata "Direct link to audiodata")

_AudioData_

Information about the audio. Use [`getAudioData()`](/docs/get-audio-data) to fetch it.

### `startTimeInSeconds` [​](\#starttimeinseconds "Direct link to starttimeinseconds")

_number_

Trim the waveform to exclude all data before `startTimeInSeconds`.

### `durationInSeconds` [​](\#durationinseconds "Direct link to durationinseconds")

_number_

trim the waveform to exclude all data after `startTimeInSeconds + durationInSeconds`.

### `numberOfSamples` [​](\#numberofsamples "Direct link to numberofsamples")

_number_

How big you want the result array to be. The function will compress the waveform to fit in `numberOfSamples` data points.

## Return value [​](\#return-value "Direct link to Return value")

`Bar[]` \- An array of objects with the following properties:

### `index` [​](\#index "Direct link to index")

_number_

The index of the datapoint, starting at 0. Useful for specifying as React `key` attribute without getting a warning.

### `amplitude` [​](\#amplitude "Direct link to amplitude")

_number_

A value describing the amplitude / volume / loudness of the audio. Values range between 0 and 1.

## Example [​](\#example "Direct link to Example")

```

tsx

import { getAudioData, getWaveformPortion } from "@remotion/media-utils";
import { staticFile } from "remotion";

const audioData = await getAudioData(staticFile("music.mp3")); /* {
  channelWaveforms: [Float32Array(4410000), Float32Array(4410000)],
  sampleRate: 44100,
  durationInSeconds: 100.0000,
  numberOfChannels: 2,
  resultId: "0.432878981",
  isRemote: false
} */

const waveformPortion = await getWaveformPortion({
  audioData,
  // Will select time range of 20-40 seconds
  startTimeInSeconds: 20,
  durationInSeconds: 20,
  numberOfSamples: 10,
}); // [{index: 0, amplitude: 0.14}, ... {index: 9, amplitude: 0.79}]

console.log(waveformPortion.length); // 10
```

```

tsx

import { getAudioData, getWaveformPortion } from "@remotion/media-utils";
import { staticFile } from "remotion";

const audioData = await getAudioData(staticFile("music.mp3")); /* {
  channelWaveforms: [Float32Array(4410000), Float32Array(4410000)],
  sampleRate: 44100,
  durationInSeconds: 100.0000,
  numberOfChannels: 2,
  resultId: "0.432878981",
  isRemote: false
} */

const waveformPortion = await getWaveformPortion({
  audioData,
  // Will select time range of 20-40 seconds
  startTimeInSeconds: 20,
  durationInSeconds: 20,
  numberOfSamples: 10,
}); // [{index: 0, amplitude: 0.14}, ... {index: 9, amplitude: 0.79}]

console.log(waveformPortion.length); // 10
```

## Alternatives [​](\#alternatives "Direct link to Alternatives")

The [`visualizeAudio()`](/docs/visualize-audio) function is more suitable for visualizing audio based on frequency properties of the audio (bass, mids, highs, etc).

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/media-utils/src/get-waveform-portion.ts)
- [Using audio](/docs/using-audio)
- [`<Audio/>`](/docs/audio)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/get-waveform-portion.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getVideoMetadata()](/docs/get-video-metadata) [Next\
\
useAudioData()](/docs/use-audio-data)

- [Arguments](#arguments)
  - [`audioData`](#audiodata)
  - [`startTimeInSeconds`](#starttimeinseconds)
  - [`durationInSeconds`](#durationinseconds)
  - [`numberOfSamples`](#numberofsamples)
- [Return value](#return-value)
  - [`index`](#index)
  - [`amplitude`](#amplitude)
- [Example](#example)
- [Alternatives](#alternatives)
- [See also](#see-also)
