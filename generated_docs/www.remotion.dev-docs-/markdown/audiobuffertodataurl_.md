---
title: audioBufferToDataUrl()
source: Remotion Documentation
last_updated: 2024-12-22
---

# audioBufferToDataUrl()

- [Home page](/)
- [@remotion/media-utils](/docs/media-utils/)
- audioBufferToDataUrl()

On this page

# audioBufferToDataUrl()

_Part of the `@remotion/media-utils` package of helper functions. Available from v2.5.7._

This API takes an [`AudioBuffer`](https://developer.mozilla.org/en-US/docs/Web/API/AudioBuffer) instance and converts it to a Base 64 Data URL so it can be passed to an [`<Audio />`](/docs/audio) tag.

## API Usage [​](\#api-usage "Direct link to API Usage")

```

tsx

import { audioBufferToDataUrl } from "@remotion/media-utils";

const str = audioBufferToDataUrl(audioBuffer);
```

```

tsx

import { audioBufferToDataUrl } from "@remotion/media-utils";

const str = audioBufferToDataUrl(audioBuffer);
```

## Example: Rendering a sine tone [​](\#example-rendering-a-sine-tone "Direct link to Example: Rendering a sine tone")

The following composition will render a sine tone with a C4 pitch.

```

tsx

import { audioBufferToDataUrl } from "@remotion/media-utils";
import { useCallback, useEffect, useState } from "react";
import {
  Audio,
  cancelRender,
  continueRender,
  delayRender,
  interpolate,
  useVideoConfig,
} from "remotion";

const C4_FREQUENCY = 261.63;
const sampleRate = 44100;

export const OfflineAudioBufferExample: React.FC = () => {
  const [handle] = useState(() => delayRender());
  const [audioBuffer, setAudioBuffer] = useState<string | null>(null);
  const { fps, durationInFrames } = useVideoConfig();
  const lengthInSeconds = durationInFrames / fps;

  const renderAudio = useCallback(async () => {
    const offlineContext = new OfflineAudioContext({
      numberOfChannels: 2,
      length: sampleRate * lengthInSeconds,
      sampleRate,
    });
    const oscillatorNode = offlineContext.createOscillator();
    const gainNode = offlineContext.createGain();
    oscillatorNode.connect(gainNode);
    gainNode.connect(offlineContext.destination);
    gainNode.gain.setValueAtTime(0.5, offlineContext.currentTime);

    oscillatorNode.type = "sine";
    oscillatorNode.frequency.value = C4_FREQUENCY;

    const { currentTime } = offlineContext;
    oscillatorNode.start(currentTime);
    oscillatorNode.stop(currentTime + lengthInSeconds);

    const buffer = await offlineContext.startRendering();
    setAudioBuffer(audioBufferToDataUrl(buffer));

    continueRender(handle);
  }, [handle, lengthInSeconds]);

  useEffect(() => {
    renderAudio().catch((err) => cancelRender(err));
  }, [renderAudio]);

  return audioBuffer ? (
    <Audio
      src={audioBuffer}
      startFrom={0}
      endAt={100}
      volume={(f) =>
        interpolate(f, [0, 50, 100], [0, 1, 0], {
          extrapolateLeft: "clamp",
          extrapolateRight: "clamp",
        })
      }
    />
  ) : null;
};
```

```

tsx

import { audioBufferToDataUrl } from "@remotion/media-utils";
import { useCallback, useEffect, useState } from "react";
import {
  Audio,
  cancelRender,
  continueRender,
  delayRender,
  interpolate,
  useVideoConfig,
} from "remotion";

const C4_FREQUENCY = 261.63;
const sampleRate = 44100;

export const OfflineAudioBufferExample: React.FC = () => {
  const [handle] = useState(() => delayRender());
  const [audioBuffer, setAudioBuffer] = useState<string | null>(null);
  const { fps, durationInFrames } = useVideoConfig();
  const lengthInSeconds = durationInFrames / fps;

  const renderAudio = useCallback(async () => {
    const offlineContext = new OfflineAudioContext({
      numberOfChannels: 2,
      length: sampleRate * lengthInSeconds,
      sampleRate,
    });
    const oscillatorNode = offlineContext.createOscillator();
    const gainNode = offlineContext.createGain();
    oscillatorNode.connect(gainNode);
    gainNode.connect(offlineContext.destination);
    gainNode.gain.setValueAtTime(0.5, offlineContext.currentTime);

    oscillatorNode.type = "sine";
    oscillatorNode.frequency.value = C4_FREQUENCY;

    const { currentTime } = offlineContext;
    oscillatorNode.start(currentTime);
    oscillatorNode.stop(currentTime + lengthInSeconds);

    const buffer = await offlineContext.startRendering();
    setAudioBuffer(audioBufferToDataUrl(buffer));

    continueRender(handle);
  }, [handle, lengthInSeconds]);

  useEffect(() => {
    renderAudio().catch((err) => cancelRender(err));
  }, [renderAudio]);

  return audioBuffer ? (
    <Audio
      src={audioBuffer}
      startFrom={0}
      endAt={100}
      volume={(f) =>
        interpolate(f, [0, 50, 100], [0, 1, 0], {
          extrapolateLeft: "clamp",
          extrapolateRight: "clamp",
        })
      }
    />
  ) : null;
};
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/media-utils/src/audio-buffer/audio-url-helpers.ts)
- [Rendering audio only](/docs/using-audio/#rendering-audio-only)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/audiobuffertodataurl.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
@remotion/media-utils](/docs/media-utils/) [Next\
\
getAudioData()](/docs/get-audio-data)

- [API Usage](#api-usage)
- [Example: Rendering a sine tone](#example-rendering-a-sine-tone)
- [See also](#see-also)
