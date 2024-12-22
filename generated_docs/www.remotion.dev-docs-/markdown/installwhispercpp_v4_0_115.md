---
title: installWhisperCpp()v4.0.115
source: Remotion Documentation
last_updated: 2024-12-22
---

# installWhisperCpp()v4.0.115

- [Home page](/)
- [@remotion/install-whisper-cpp](/docs/install-whisper-cpp/)
- installWhisperCpp()

On this page

# installWhisperCpp() [v4.0.115](https://github.com/remotion-dev/remotion/releases/v4.0.115)

Installs a Whisper.cpp version into a folder.

```

install-whisper.mjs
tsx

import path from "path";
import { installWhisperCpp } from "@remotion/install-whisper-cpp";

const { alreadyExisted } = await installWhisperCpp({
  to: path.join(process.cwd(), "whisper.cpp"),
  version: "1.5.5", // A Whisper.cpp semver or git tag
});
```

```

install-whisper.mjs
tsx

import path from "path";
import { installWhisperCpp } from "@remotion/install-whisper-cpp";

const { alreadyExisted } = await installWhisperCpp({
  to: path.join(process.cwd(), "whisper.cpp"),
  version: "1.5.5", // A Whisper.cpp semver or git tag
});
```

On Windows, if the provided version is not in semantic version format (e.g., `1.5.4`), the installation will throw an error. The function expects a precompiled semantic versioned release available for download.

On other platforms, the source is cloned from GitHub and built.

If you install Whisper.cpp into a folder which is in a Git repository, you should add the folder to the `.gitignore` file.

If the output folder already exists, the function will not do anything and return `{ alreadyExisted: true }`.

You also need to download a model for Whisper.cpp to work. You can do this with the [`downloadWhisperModel()`](/docs/install-whisper-cpp/download-whisper-model) function.

## Options [​](\#options "Direct link to Options")

### `to` [​](\#to "Direct link to to")

The folder to install Whisper.cpp into.

### `version` [​](\#version "Direct link to version")

The version of Whisper.cpp to install. Don't include the `v` prefix.
This can be either a hash of a Whisper.cpp commit or a semantic version of an official release.

info

On Windows, currently only release tags of Whisper.cpp are supported (e.g `1.5.4`).

### `printOutput?` [​](\#printoutput "Direct link to printoutput")

Whether to print the output of the installation process to the console. Defaults to `true`.

### `signal?` [v4.0.156](https://github.com/remotion-dev/remotion/releases/v4.0.156) [​](\#signal "Direct link to signal")

A signal from an [`AbortController`](https://developer.mozilla.org/en-US/docs/Web/API/AbortController) to cancel the installation process.

## Return value [​](\#return-value "Direct link to Return value")

Returns a `Promise` that resolves to an object with the following property:

### `alreadyExisted` [​](\#alreadyexisted "Direct link to alreadyexisted")

Whether the folder already existed. If it did, and contains the necessary executable, the function did not perform any installation and returned `true`. If the executable is missing in the existing folder, the function expects manual deletion of the folder before attempting another installation.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/install-whisper-cpp/src/install-whisper-cpp.ts)
- [`@remotion/install-whisper-cpp`](/docs/install-whisper-cpp)
- [`downloadWhisperModel()`](/docs/install-whisper-cpp/download-whisper-model)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/install-whisper-cpp/install-whisper-cpp.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
@remotion/install-whisper-cpp](/docs/install-whisper-cpp/) [Next\
\
downloadWhisperModel()](/docs/install-whisper-cpp/download-whisper-model)

- [Options](#options)
  - [`to`](#to)
  - [`version`](#version)
  - [`printOutput?`](#printoutput)
  - [`signal?`](#signal)
- [Return value](#return-value)
  - [`alreadyExisted`](#alreadyexisted)
- [See also](#see-also)
