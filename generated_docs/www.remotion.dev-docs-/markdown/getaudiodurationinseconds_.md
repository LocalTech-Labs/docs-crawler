---
title: getAudioDurationInSeconds()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getAudioDurationInSeconds()

- [Home page](/)
- [@remotion/media-utils](/docs/media-utils/)
- getAudioDurationInSeconds()

On this page

# getAudioDurationInSeconds()

_Part of the `@remotion/media-utils` package of helper functions._

_Previously called `getAudioDuration()`._

Gets the duration in seconds of an audio source. Remotion will create an invisible `<audio>` tag, load the audio and return the duration.

## Arguments [​](\#arguments "Direct link to Arguments")

### `src` [​](\#src "Direct link to src")

A string pointing to an audio asset

## Return value [​](\#return-value "Direct link to Return value")

`Promise<number>` \- the duration of the audio file.

## Example [​](\#example "Direct link to Example")

```

tsx

import { getAudioDurationInSeconds } from "@remotion/media-utils";
import music from "./music.mp3";

const MyComp: React.FC = () => {
  const getDuration = useCallback(async () => {
    const publicFile = await getAudioDurationInSeconds(
      staticFile("voiceover.wav"),
    ); // 33.221
    const imported = await getAudioDurationInSeconds(music); // 127.452
    const remote = await getAudioDurationInSeconds(
      "https://example.com/remote-audio.aac",
    ); // 50.24
  }, []);

  useEffect(() => {
    getDuration();
  }, []);

  return null;
};
```

```

tsx

import { getAudioDurationInSeconds } from "@remotion/media-utils";
import music from "./music.mp3";

const MyComp: React.FC = () => {
  const getDuration = useCallback(async () => {
    const publicFile = await getAudioDurationInSeconds(
      staticFile("voiceover.wav"),
    ); // 33.221
    const imported = await getAudioDurationInSeconds(music); // 127.452
    const remote = await getAudioDurationInSeconds(
      "https://example.com/remote-audio.aac",
    ); // 50.24
  }, []);

  useEffect(() => {
    getDuration();
  }, []);

  return null;
};
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/media-utils/src/get-audio-duration-in-seconds.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/get-audio-duration-in-seconds.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getAudioData()](/docs/get-audio-data) [Next\
\
getVideoMetadata()](/docs/get-video-metadata)

- [Arguments](#arguments)
  - [`src`](#src)
- [Return value](#return-value)
- [Example](#example)
- [See also](#see-also)
