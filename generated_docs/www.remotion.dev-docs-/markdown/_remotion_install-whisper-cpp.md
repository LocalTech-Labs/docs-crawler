---
title: @remotion/install-whisper-cpp
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/install-whisper-cpp

- [Home page](/)
- @remotion/install-whisper-cpp

On this page

# @remotion/install-whisper-cpp

_Available from v4.0.115_

With [Whisper.cpp](https://github.com/ggerganov/whisper.cpp), you can transcribe audio locally on your machine.

This package provides easy to use cross-platform functions to install Whisper.cpp and a model.

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/install-whisper-cpp@4.0.242Copy
```

```

npm i --save-exact @remotion/install-whisper-cpp@4.0.242Copy
```

```

pnpm i @remotion/install-whisper-cpp@4.0.242Copy
```

```

pnpm i @remotion/install-whisper-cpp@4.0.242Copy
```

```

bun i @remotion/install-whisper-cpp@4.0.242Copy
```

```

bun i @remotion/install-whisper-cpp@4.0.242Copy
```

```

yarn --exact add @remotion/install-whisper-cpp@4.0.242Copy
```

```

yarn --exact add @remotion/install-whisper-cpp@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

## Example usage [​](\#example-usage "Direct link to Example usage")

Install Whisper `1.5.5` (the latest version at the time of writing that we find works well and supports token-level timestamps) and the `medium.en` model to the `whisper.cpp` folder.

```

install-whisper.cpp
tsx

import path from 'path';
import {
  downloadWhisperModel,
  installWhisperCpp,
  transcribe,
  convertToCaptions,
} from '@remotion/install-whisper-cpp';

const to = path.join(process.cwd(), 'whisper.cpp');

await installWhisperCpp({
  to,
  version: '1.5.5',
});

await downloadWhisperModel({
  model: 'medium.en',
  folder: to,
});

const {transcription} = await transcribe({
  model: 'medium.en',
  whisperPath: to,
  inputPath: '/path/to/audio.wav',
  tokenLevelTimestamps: true,
});

for (const token of transcription) {
  console.log(token.timestamps.from, token.timestamps.to, token.text);
}

// Optional: Apply our recommended postprocessing
const {captions} = convertToCaptions({
  transcription,
  combineTokensWithinMilliseconds: 200,
});

for (const line of captions) {
  console.log(line.text, line.startInSeconds);
}
```

```

install-whisper.cpp
tsx

import path from 'path';
import {
  downloadWhisperModel,
  installWhisperCpp,
  transcribe,
  convertToCaptions,
} from '@remotion/install-whisper-cpp';

const to = path.join(process.cwd(), 'whisper.cpp');

await installWhisperCpp({
  to,
  version: '1.5.5',
});

await downloadWhisperModel({
  model: 'medium.en',
  folder: to,
});

const {transcription} = await transcribe({
  model: 'medium.en',
  whisperPath: to,
  inputPath: '/path/to/audio.wav',
  tokenLevelTimestamps: true,
});

for (const token of transcription) {
  console.log(token.timestamps.from, token.timestamps.to, token.text);
}

// Optional: Apply our recommended postprocessing
const {captions} = convertToCaptions({
  transcription,
  combineTokensWithinMilliseconds: 200,
});

for (const line of captions) {
  console.log(line.text, line.startInSeconds);
}
```

## Functions [​](\#functions "Direct link to Functions")

[**installWhisperCpp()** \
\
Install the whisper.cpp software](/docs/install-whisper-cpp/install-whisper-cpp) [**downloadWhisperModel()** \
\
Download a Whisper model](/docs/install-whisper-cpp/download-whisper-model) [**transcribe()** \
\
Transcribe an audio file](/docs/install-whisper-cpp/transcribe) [**convertToCaptions()** \
\
Postprocessing for TikTok-style captions](/docs/install-whisper-cpp/convert-to-captions)

## License [​](\#license "Direct link to License")

MIT

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/install-whisper-cpp/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
createTikTokStyleCaptions()](/docs/captions/create-tiktok-style-captions) [Next\
\
installWhisperCpp()](/docs/install-whisper-cpp/install-whisper-cpp)

- [Example usage](#example-usage)
- [Functions](#functions)
- [License](#license)
