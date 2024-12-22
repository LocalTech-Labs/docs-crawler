---
title: openAiWhisperApiToCaptions()v4.0.217
source: Remotion Documentation
last_updated: 2024-12-22
---

# openAiWhisperApiToCaptions()v4.0.217

- [Home page](/)
- [@remotion/openai-whisper](/docs/openai-whisper/)
- openAiWhisperApiToCaptions()

On this page

# openAiWhisperApiToCaptions() [v4.0.217](https://github.com/remotion-dev/remotion/releases/v4.0.217)

Turns the output from [`openai.audio.transcriptions.create`](https://platform.openai.com/docs/guides/speech-to-text/transcriptions) from the [openai](https://npm.im/openai) package into an array of [`Caption`](/docs/captions/caption) objects.

This package performs processing on the captions in order to retain the punctuation in the words, which is not by default included in the OpenAI response.

This function can be used in any JavaScript environment, but you should not use the OpenAI API in the browser because your API key will be exposed to the browser.

```

Example usage
tsx

import fs from 'fs';
import {OpenAI} from 'openai';
import {openAiWhisperApiToCaptions} from '@remotion/openai-whisper';

const openai = new OpenAI();

const transcription = await openai.audio.transcriptions.create({
  file: fs.createReadStream('audio.mp3'),
  model: 'whisper-1',
  response_format: 'verbose_json',
  prompt: 'Hello, welcome to my lecture.',
  timestamp_granularities: ['word'],
});

const {captions} = openAiWhisperApiToCaptions({transcription});
```

```

Example usage
tsx

import fs from 'fs';
import {OpenAI} from 'openai';
import {openAiWhisperApiToCaptions} from '@remotion/openai-whisper';

const openai = new OpenAI();

const transcription = await openai.audio.transcriptions.create({
  file: fs.createReadStream('audio.mp3'),
  model: 'whisper-1',
  response_format: 'verbose_json',
  prompt: 'Hello, welcome to my lecture.',
  timestamp_granularities: ['word'],
});

const {captions} = openAiWhisperApiToCaptions({transcription});
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/openai-whisper/src/openai-whisper-api-to-captions.ts)
- [`@remotion/openai-whisper`](/docs/openai-whisper)
- [`@remotion/install-whisper-cpp`](/docs/install-whisper-cpp)
- [`@remotion/captions`](/docs/captions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/openai-whisper/openai-whisper-api-to-captions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
@remotion/openai-whisper](/docs/openai-whisper/) [Next\
\
Overview](/docs/fonts-api/)

- [See also](#see-also)
