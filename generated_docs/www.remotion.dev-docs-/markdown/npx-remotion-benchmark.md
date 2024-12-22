---
title: npx remotion benchmark
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion benchmark

- [Home page](/)
- [Command line](/docs/cli/)
- benchmark

On this page

# npx remotion benchmark

_available from v3.2.28_

Measures render time by running a render multiple times, if desired with multiple compositions and concurrency values to compare against each other.

```

bash

npx remotion benchmark src/index.ts [composition-ids]
```

```

bash

npx remotion benchmark src/index.ts [composition-ids]
```

You can provide multiple composition IDs separated by comma, ex: `npx remotion benchmark src/index.ts --codec=h264 Main,Canvas,CSS`

If `composition-ids` is not passed, Remotion will let you select compositions from a list.

## Flags [​](\#flags "Direct link to Flags")

### `--runs` [​](\#--runs "Direct link to --runs")

_optional. Default is 3_

Specify how many times video must be rendered. Default value is 3.

### `--concurrencies` [​](\#--concurrencies "Direct link to --concurrencies")

_optional_

You can specify which concurrency value should be used while rendering the video. Multiple concurrency values can be passed separated by comma. Learn more about [concurrency](/docs/terminology/concurrency)

### `--codec` [​](\#--codec "Direct link to --codec")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--codec)

### `--audio-codec` [v3.3.42](https://github.com/remotion-dev/remotion/releases/v3.3.42) [​](\#--audio-codec "Direct link to --audio-codec")

_optional_

Set the format of the audio that is embedded in the video. Not all codec and audio codec combinations are supported and certain combinations require a certain file extension and container format. See the table in the docs to see possible combinations.

### `--crf` [​](\#--crf "Direct link to --crf")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--crf)

### `--frames` [​](\#--frames "Direct link to --frames")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--frames)

### `--image-format` [​](\#--image-format "Direct link to --image-format")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--image-format)

### `--pixel-format` [​](\#--pixel-format "Direct link to --pixel-format")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--pixel-format)

### `--props` [​](\#--props "Direct link to --props")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--props)

### `--prores-profile` [​](\#--prores-profile "Direct link to --prores-profile")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--prores-profile)

### `--jpeg-quality` [​](\#--jpeg-quality "Direct link to --jpeg-quality")

_optional, available from v4.0.0_

Inherited from [`npx remotion render`](/docs/cli/render#--jpeg-quality)

### `--quality` [​](\#--quality "Direct link to --quality")

_optional, removed in v4.0.0_

Renamed to `--jpeg-quality`.

### `--log` [​](\#--log "Direct link to --log")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--log)

### `--ignore-certificate-errors` [​](\#--ignore-certificate-errors "Direct link to --ignore-certificate-errors")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--ignore-certificate-errors)

### `--disable-web-security` [​](\#--disable-web-security "Direct link to --disable-web-security")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--disable-web-security)

### `--disable-headless?` [​](\#--disable-headless "Direct link to --disable-headless")

_optional_

Deprecated - will be removed in 5.0.0. With the migration to [Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell), this option is not functional anymore.

If disabled, the render will open an actual Chrome window where you can see the render happen. The default is headless mode.

### `--enable-multiprocess-on-linux` [v4.0.42](https://github.com/remotion-dev/remotion/releases/v4.0.42) [​](\#--enable-multiprocess-on-linux "Direct link to --enable-multiprocess-on-linux")

Removes the `--single-process` flag that gets passed to Chromium on Linux by default. This will make the render faster because multiple processes can be used, but may cause issues with some Linux distributions or if window server libraries are missing.

Default: `false` until v4.0.136, then `true` from v4.0.137 on because newer Chrome versions don't allow rendering with the `--single-process` flag.

This flag will be removed in Remotion v5.0.

### `--gl` [​](\#--gl "Direct link to --gl")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--gl)

### `--timeout` [​](\#--timeout "Direct link to --timeout")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--timeout)

### `--scale` [​](\#--scale "Direct link to --scale")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--scale)

### `--port` [​](\#--port "Direct link to --port")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--port)

### `--number-of-gif-loops` [​](\#--number-of-gif-loops "Direct link to --number-of-gif-loops")

_optional_

Allows you to set the number of loops as follows:

- `null` (or omitting in the CLI) plays the GIF indefinitely.
- `0` disables looping
- `1` loops the GIF once (plays twice in total)
- `2` loops the GIF twice (plays three times in total) and so on.

### `--every-nth-frame` [​](\#--every-nth-frame "Direct link to --every-nth-frame")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--every-nth-frame)

### `--log` [​](\#--log-1 "Direct link to --log-1")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--log)

### `--muted` [​](\#--muted "Direct link to --muted")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--muted)

### `--enforce-audio-track` [​](\#--enforce-audio-track "Direct link to --enforce-audio-track")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--enforce-audio-track)

### `--browser-executable` [​](\#--browser-executable "Direct link to --browser-executable")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--browser-executable)

### `--public-dir` [​](\#--public-dir "Direct link to --public-dir")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--public-dir)

### `--config` [​](\#--config "Direct link to --config")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--config)

### `--bundle-cache` [​](\#--bundle-cache "Direct link to --bundle-cache")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--bundle-cache)

### `--video-bitrate` [​](\#--video-bitrate "Direct link to --video-bitrate")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--video-bitrate)

### `--audio-bitrate` [​](\#--audio-bitrate "Direct link to --audio-bitrate")

_optional_

Inherited from [`npx remotion render`](/docs/cli/render#--audio-bitrate)

### `--color-space` [v4.0.28](https://github.com/remotion-dev/remotion/releases/v4.0.28) [​](\#--color-space "Direct link to --color-space")

Color space to use for the video. Acceptable values: `"default"`(default since 5.0), `"bt709"`(since v4.0.28), `"bt2020-ncl"`(since v4.0.88), `"bt2020-cl"`(since v4.0.88), .

For best color accuracy, it is recommended to also use `"png"` as the image format to have accurate color transformations throughout.

Only since v4.0.83, colorspace conversion is actually performed, previously it would only tag the metadata of the video.

### `--hardware-acceleration` [v4.0.228](https://github.com/remotion-dev/remotion/releases/v4.0.228) [​](\#--hardware-acceleration "Direct link to --hardware-acceleration")

One of
"disable", "if-possible", or "required"
. Default "disable". Encode using a hardware-accelerated encoder if
available. If set to "required" and no hardware-accelerated encoder is
available, then the render will fail.

### `--offthreadvideo-cache-size-in-bytes` [v4.0.23](https://github.com/remotion-dev/remotion/releases/v4.0.23) [​](\#--offthreadvideo-cache-size-in-bytes "Direct link to --offthreadvideo-cache-size-in-bytes")

From v4.0, Remotion has a cache for [`<OffthreadVideo>`](https://remotion.dev/docs/offthreadvideo) frames. The default is `null`, corresponding to half of the system memory available when the render starts.

This option allows to override the size of the cache. The higher it is, the faster the render will be, but the more memory will be used.

The used value will be printed when running in verbose mode.

Default: `null`

### `--binaries-directory` [v4.0.120](https://github.com/remotion-dev/remotion/releases/v4.0.120) [​](\#--binaries-directory "Direct link to --binaries-directory")

The directory where the platform-specific binaries and libraries that Remotion needs are located. Those include an `ffmpeg` and `ffprobe` binary, a Rust binary for various tasks, and various shared libraries. If the value is set to `null`, which is the default, then the path of a platform-specific package located at `node_modules/@remotion/compositor-*` is selected.

This option is useful in environments where Remotion is not officially supported to run like bundled serverless functions or Electron.

### `--ffmpeg-executable` [​](\#--ffmpeg-executable "Direct link to --ffmpeg-executable")

_optional, removed in v4.0_

Inherited from [`npx remotion render`](/docs/cli/render#--ffmpeg-executable)

### `--ffprobe-executable` [​](\#--ffprobe-executable "Direct link to --ffprobe-executable")

_optional, removed in v4.0_

Inherited from [`npx remotion render`](/docs/cli/render#--ffprobe-executable-)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/benchmark.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
regions](/docs/cloudrun/cli/regions) [Next\
\
browser](/docs/cli/browser/)

- [Flags](#flags)
  - [`--runs`](#--runs)
  - [`--concurrencies`](#--concurrencies)
  - [`--codec`](#--codec)
  - [`--audio-codec`](#--audio-codec)
  - [`--crf`](#--crf)
  - [`--frames`](#--frames)
  - [`--image-format`](#--image-format)
  - [`--pixel-format`](#--pixel-format)
  - [`--props`](#--props)
  - [`--prores-profile`](#--prores-profile)
  - [`--jpeg-quality`](#--jpeg-quality)
  - [`--quality`](#--quality)
  - [`--log`](#--log)
  - [`--ignore-certificate-errors`](#--ignore-certificate-errors)
  - [`--disable-web-security`](#--disable-web-security)
  - [`--disable-headless?`](#--disable-headless)
  - [`--enable-multiprocess-on-linux`](#--enable-multiprocess-on-linux)
  - [`--gl`](#--gl)
  - [`--timeout`](#--timeout)
  - [`--scale`](#--scale)
  - [`--port`](#--port)
  - [`--number-of-gif-loops`](#--number-of-gif-loops)
  - [`--every-nth-frame`](#--every-nth-frame)
  - [`--log`](#--log-1)
  - [`--muted`](#--muted)
  - [`--enforce-audio-track`](#--enforce-audio-track)
  - [`--browser-executable`](#--browser-executable)
  - [`--public-dir`](#--public-dir)
  - [`--config`](#--config)
  - [`--bundle-cache`](#--bundle-cache)
  - [`--video-bitrate`](#--video-bitrate)
  - [`--audio-bitrate`](#--audio-bitrate)
  - [`--color-space`](#--color-space)
  - [`--hardware-acceleration`](#--hardware-acceleration)
  - [`--offthreadvideo-cache-size-in-bytes`](#--offthreadvideo-cache-size-in-bytes)
  - [`--binaries-directory`](#--binaries-directory)
  - [`--ffmpeg-executable`](#--ffmpeg-executable)
  - [`--ffprobe-executable`](#--ffprobe-executable)
