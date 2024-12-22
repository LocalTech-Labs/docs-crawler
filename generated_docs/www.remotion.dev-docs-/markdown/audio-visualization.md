---
title: Audio visualization
source: Remotion Documentation
last_updated: 2024-12-22
---

# Audio visualization

- [Home page](/)
- Designing visuals
- Audio visualization

On this page

# Audio visualization

Remotion has APIs for visualizing audio, for example for creating audiograms or music visualizers.

The `@remotion/media-utils` package provides helper functions for reading and processing audio. Using the [`getAudioData()`](/docs/get-audio-data) API you can read audio, and using the [`useAudioData()`](/docs/use-audio-data) helper hook you can load this audio data directly into your component.

Using the [`visualizeAudio()`](/docs/visualize-audio) API, you can get an audio spectrum for the current frame, with the `numberOfSamples` parameter being an option to control the amount of detail you need.

Refer to the documentation of the above mentioned functions to learn more.

```

tsx

import {useAudioData, visualizeAudio} from '@remotion/media-utils';
import {Audio, staticFile, useCurrentFrame, useVideoConfig} from 'remotion';

const music = staticFile('music.mp3');

export const MyComponent: React.FC = () => {
  const frame = useCurrentFrame();
  const {width, height, fps} = useVideoConfig();
  const audioData = useAudioData(music);

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
      <Audio src={music} />
      {visualization.map((v) => {
        return (
          <div style={{width: 1000 * v, height: 15, backgroundColor: 'blue'}} />
        );
      })}
    </div>
  );
};
```

```

tsx

import {useAudioData, visualizeAudio} from '@remotion/media-utils';
import {Audio, staticFile, useCurrentFrame, useVideoConfig} from 'remotion';

const music = staticFile('music.mp3');

export const MyComponent: React.FC = () => {
  const frame = useCurrentFrame();
  const {width, height, fps} = useVideoConfig();
  const audioData = useAudioData(music);

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
      <Audio src={music} />
      {visualization.map((v) => {
        return (
          <div style={{width: 1000 * v, height: 15, backgroundColor: 'blue'}} />
        );
      })}
    </div>
  );
};
```

## Working with large files [​](\#working-with-large-files "Direct link to Working with large files")

[`useAudioData()`](/docs/use-audio-data) loads the entire audio file into memory.

This is fine for small files, but for large files, it can be slow and consume a lot of memory.

Use [`useWindowedAudioData()`](/docs/use-windowed-audio-data) to only load a portion of the audio around the current frame.

The tradeoff is that this API only works with `.wav` files.

## See also [​](\#see-also "Direct link to See also")

- [Using audio](/docs/using-audio)
- [`useAudioData()`](/docs/use-audio-data)
- [`useWindowedAudioData()`](/docs/use-windowed-audio-data)
- [`visualizeAudio()`](/docs/visualize-audio)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/audio-visualization.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Randomness](/docs/using-randomness) [Next\
\
Noise visualization](/docs/noise-visualization)

- [Working with large files](#working-with-large-files)
- [See also](#see-also)