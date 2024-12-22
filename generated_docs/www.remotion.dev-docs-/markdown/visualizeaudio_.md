---
title: visualizeAudio()
source: Remotion Documentation
last_updated: 2024-12-22
---

# visualizeAudio()

- [Home page](/)
- [@remotion/media-utils](/docs/media-utils/)
- visualizeAudio()

On this page

# visualizeAudio()

_Part of the `@remotion/media-utils` package of helper functions._

This function takes in `AudioData` (preferably fetched by the [`useAudioData()`](/docs/use-audio-data) hook) and processes it in a way that makes visualizing the audio that is playing at the current frame easy.

## Arguments [​](\#arguments "Direct link to Arguments")

Takes an object containing the following values:

### `audioData` [​](\#audiodata "Direct link to audiodata")

_`AudioData`_

An object containing audio data. You can fetch this object using [`useAudioData()`](/docs/use-audio-data) or [`getAudioData()`](/docs/get-audio-data).

### `frame` [​](\#frame "Direct link to frame")

_`number`_

The time of the track that you want to get the audio information for. The `frame` always refers to the position in the audio track - if you have shifted or trimmed the audio in your timeline, the frame returned by `useCurrentFrame` must also be tweaked before you pass it into this function.

### `fps` [​](\#fps "Direct link to fps")

_`number`_

The frame rate of the composition. This helps the function understand the meaning of the `frame` input.

### `numberOfSamples` [​](\#numberofsamples "Direct link to numberofsamples")

`number`

Must be a power of two, such as `32`, `64`, `128`, etc. This parameter controls the length of the output array. A lower number will simplify the spectrum and is useful if you want to animate elements roughly based on the level of lows, mids and highs. A higher number will give the spectrum in more detail, which is useful for displaying a bar chart or waveform-style visualization of the audio.

### `smoothing` [​](\#smoothing "Direct link to smoothing")

`boolean`

When set to `true` the returned values will be an average of the current, previous and next frames. The result is a smoother transition for quickly changing values. Default value is `true`.

### `optimizeFor?` [v4.0.83](https://github.com/remotion-dev/remotion/releases/v4.0.83) [​](\#optimizefor "Direct link to optimizefor")

_`"accuracy" | "speed"`_

Default `"accuracy"`. When set to `"speed"`, a faster Fast Fourier transform is used. Recommended for Remotion Lambda and when using a high number of samples. Read [user](https://discord.com/channels/809501355504959528/1189048518988550264/1190228606287360030) [experiences](https://discord.com/channels/809501355504959528/1155110845488046111/1155111360481480725) [here](https://github.com/remotion-dev/remotion/issues/2925).

## Return value [​](\#return-value "Direct link to Return value")

`number[]`

An array of values describing the amplitude of each frequency range. Each value is between 0 and 1. The array is of length defined by the `numberOfSamples` parameter.

The values on the left of the array are low frequencies (for example bass) and as we move towards the right, we go through the mid and high frequencies like drums and vocals.

Usually the values on left side of the array can become much larger than the values on the right. This is not a mistake, but for some visualizations you might have to apply some postprocessing to it, you can flatten the curve by for example mapping each value to a logarithm of the original function.

## Example [​](\#example "Direct link to Example")

In this example, we render a bar chart visualizing the audio spectrum of an audio file we imported using [`useAudioData()`](/docs/use-audio-data) and `visualizeAudio()`.

```

tsx

import { useAudioData, visualizeAudio } from "@remotion/media-utils";
import { Audio, staticFile, useCurrentFrame, useVideoConfig } from "remotion";

export const MyComponent: React.FC = () => {
  const frame = useCurrentFrame();
  const { width, height, fps } = useVideoConfig();
  const audioData = useAudioData(staticFile("music.mp3"));

  if (!audioData) {
    return null;
  }

  const visualization = visualizeAudio({
    fps,
    frame,
    audioData,
    numberOfSamples: 16,
  }); // [0.22, 0.1, 0.01, 0.01, 0.01, 0.02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

  // Render a bar chart for each frequency, the higher the amplitude,
  // the longer the bar
  return (
    <div>
      <Audio src={staticFile("music.mp3")} />
      {visualization.map((v) => {
        return (
          <div
            style={{ width: 1000 * v, height: 15, backgroundColor: "blue" }}
          />
        );
      })}
    </div>
  );
};
```

```

tsx

import { useAudioData, visualizeAudio } from "@remotion/media-utils";
import { Audio, staticFile, useCurrentFrame, useVideoConfig } from "remotion";

export const MyComponent: React.FC = () => {
  const frame = useCurrentFrame();
  const { width, height, fps } = useVideoConfig();
  const audioData = useAudioData(staticFile("music.mp3"));

  if (!audioData) {
    return null;
  }

  const visualization = visualizeAudio({
    fps,
    frame,
    audioData,
    numberOfSamples: 16,
  }); // [0.22, 0.1, 0.01, 0.01, 0.01, 0.02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

  // Render a bar chart for each frequency, the higher the amplitude,
  // the longer the bar
  return (
    <div>
      <Audio src={staticFile("music.mp3")} />
      {visualization.map((v) => {
        return (
          <div
            style={{ width: 1000 * v, height: 15, backgroundColor: "blue" }}
          />
        );
      })}
    </div>
  );
};
```

## Postprocessing example [​](\#postprocessing-example "Direct link to Postprocessing example")

A logarithmic representation of the audio will look more appealing than a linear one. Below is an example of a postprocessing step that looks prettier than the default one.

```

tsx

/**
 * This postprocessing step will match the values with what you'd
 * get from WebAudio's `AnalyserNode.getByteFrequencyData()`.
 *
 * MDN: https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/getByteFrequencyData
 * W3C Spec: https://www.w3.org/TR/webaudio/#AnalyserNode-methods
 */

// get the frequency data
const frequencyData = visualizeAudio(params);

// default scaling factors from the W3C spec for getByteFrequencyData
const minDb = -100;
const maxDb = -30;

const amplitudes = frequencyData.map((value) => {
  // convert to decibels (will be in the range `-Infinity` to `0`)
  const db = 20 * Math.log10(value);

  // scale to fit between min and max
  const scaled = (db - minDb) / (maxDb - minDb);

  return scaled;
});
```

```

tsx

/**
 * This postprocessing step will match the values with what you'd
 * get from WebAudio's `AnalyserNode.getByteFrequencyData()`.
 *
 * MDN: https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/getByteFrequencyData
 * W3C Spec: https://www.w3.org/TR/webaudio/#AnalyserNode-methods
 */

// get the frequency data
const frequencyData = visualizeAudio(params);

// default scaling factors from the W3C spec for getByteFrequencyData
const minDb = -100;
const maxDb = -30;

const amplitudes = frequencyData.map((value) => {
  // convert to decibels (will be in the range `-Infinity` to `0`)
  const db = 20 * Math.log10(value);

  // scale to fit between min and max
  const scaled = (db - minDb) / (maxDb - minDb);

  return scaled;
});
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/media-utils/src/visualize-audio.ts)
- [Audio visualization](/docs/audio-visualization)
- [`useAudioData()`](/docs/use-audio-data)
- [`getAudioData()`](/docs/get-audio-data)
- [`<Audio/>`](/docs/audio)
- [Using audio](/docs/using-audio)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/visualize-audio.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
useWindowedAudioData()](/docs/use-windowed-audio-data) [Next\
\
getImageDimensions()](/docs/get-image-dimensions)

- [Arguments](#arguments)
  - [`audioData`](#audiodata)
  - [`frame`](#frame)
  - [`fps`](#fps)
  - [`numberOfSamples`](#numberofsamples)
  - [`smoothing`](#smoothing)
  - [`optimizeFor?`](#optimizefor)
- [Return value](#return-value)
- [Example](#example)
- [Postprocessing example](#postprocessing-example)
- [See also](#see-also)