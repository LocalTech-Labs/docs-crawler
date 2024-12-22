---
title: downloadWhisperModel()v4.0.115
source: Remotion Documentation
last_updated: 2024-12-22
---

# downloadWhisperModel()v4.0.115

- [Home page](/)
- [@remotion/install-whisper-cpp](/docs/install-whisper-cpp/)
- downloadWhisperModel()

On this page

# downloadWhisperModel() [v4.0.115](https://github.com/remotion-dev/remotion/releases/v4.0.115)

Downloads a Whisper.cpp model to a folder.

You should first install Whisper.cpp, for example through [`installWhisperCpp()`](/docs/install-whisper-cpp/install-whisper-cpp).

```

install-whisper.mjs
tsx

import path from 'path';
import {downloadWhisperModel} from '@remotion/install-whisper-cpp';

const {alreadyExisted} = await downloadWhisperModel({
  model: 'medium.en',
  folder: path.join(process.cwd(), 'whisper.cpp'),
});
```

```

install-whisper.mjs
tsx

import path from 'path';
import {downloadWhisperModel} from '@remotion/install-whisper-cpp';

const {alreadyExisted} = await downloadWhisperModel({
  model: 'medium.en',
  folder: path.join(process.cwd(), 'whisper.cpp'),
});
```

## Options [​](\#options "Direct link to Options")

### `folder` [​](\#folder "Direct link to folder")

The folder to download the model to. The model will be downloaded into this folder with the filename `ggml-${model}.bin`.

### `model` [​](\#model "Direct link to model")

The model to download. Possible values: `tiny`, `tiny.en`, `base`, `base.en`, `small`, `small.en`, `medium`, `medium.en`, `large-v1`, `large-v2`, `large-v3`, `large-v3-turbo`.

### `onProgress?` [​](\#onprogress "Direct link to onprogress")

Act upon download progress. This is the function signature:

```

tsx

import type {OnProgress} from '@remotion/install-whisper-cpp';

const onProgress: OnProgress = (
  downloadedBytes: number,
  totalBytes: number,
) => {
  const progress = downloadedBytes / totalBytes;
};
```

```

tsx

import type {OnProgress} from '@remotion/install-whisper-cpp';

const onProgress: OnProgress = (
  downloadedBytes: number,
  totalBytes: number,
) => {
  const progress = downloadedBytes / totalBytes;
};
```

### `printOutput?` [​](\#printoutput "Direct link to printoutput")

Print human-readable progress to the console. Default: `true`.

### `signal?` [v4.0.156](https://github.com/remotion-dev/remotion/releases/v4.0.156) [​](\#signal "Direct link to signal")

A signal from an [`AbortController`](https://developer.mozilla.org/en-US/docs/Web/API/AbortController) to cancel the download process.

## Return Value [​](\#return-value "Direct link to Return Value")

Returns an object with the following property:

### `alreadyExisted` [​](\#alreadyexisted "Direct link to alreadyexisted")

Indicates whether a file at the output path already existed. If it did, the function did not do anything and this property is set to `true`.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/install-whisper-cpp/src/download-whisper-model.ts)
- [`@remotion/install-whisper-cpp`](/docs/install-whisper-cpp)
- [`installWhisperCpp()`](/docs/install-whisper-cpp/install-whisper-cpp)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/install-whisper-cpp/download-whisper-model.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
installWhisperCpp()](/docs/install-whisper-cpp/install-whisper-cpp) [Next\
\
transcribe()](/docs/install-whisper-cpp/transcribe)

- [Options](#options)
  - [`folder`](#folder)
  - [`model`](#model)
  - [`onProgress?`](#onprogress)
  - [`printOutput?`](#printoutput)
  - [`signal?`](#signal)
- [Return Value](#return-value)
  - [`alreadyExisted`](#alreadyexisted)
- [See also](#see-also)
