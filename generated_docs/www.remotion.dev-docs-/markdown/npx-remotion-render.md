---
title: npx remotion render
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion render

- [Home page](/)
- [Command line](/docs/cli/)
- render

On this page

# npx remotion render

Render a video or audio based on the entry point, the composition ID and save it to the output location.

```

bash

npx remotion render <entry-point|serve-url>? <composition-id> <output-location>
```

```

bash

npx remotion render <entry-point|serve-url>? <composition-id> <output-location>
```

You may pass a [Serve URL](/docs/terminology/serve-url) or an [entry point](/docs/terminology/entry-point) as the first argument, otherwise the entry point will be [determined](/docs/terminology/entry-point#which-entry-point-is-being-used).

If `composition-id` is not passed, Remotion will ask you to select a composition.

If `output-location` is not passed, the media will be rendered into the `out` folder.

## Flags [​](\#flags "Direct link to Flags")

Besides choosing a video and output location with the command line arguments, the following flags are supported:

### `--props` [​](\#--props "Direct link to --props")

[Input Props to pass to the selected composition of your video.](/docs/passing-props#passing-input-props-in-the-cli).

Must be a serialized JSON string ( `--props='{"hello": "world"}'`) or a path to a JSON file ( `./path/to/props.json`).

From the root component the props can be read using [`getInputProps()`](/docs/get-input-props).

You may transform input props using [`calculateMetadata()`](/docs/calculate-metadata).

note

Inline JSON string isn't supported on Windows shells because it removes the `"` character, use a file name instead.

### `--height` [v3.2.40](https://github.com/remotion-dev/remotion/releases/v3.2.40) [​](\#--height "Direct link to --height")

[Overrides composition height.](/docs/config#overrideheight)

### `--width` [v3.2.40](https://github.com/remotion-dev/remotion/releases/v3.2.40) [​](\#--width "Direct link to --width")

[Overrides composition width.](/docs/config#overridewidth)

### `--concurrency` [​](\#--concurrency "Direct link to --concurrency")

[How many CPU threads to use.](/docs/config#setconcurrency) Minimum 1. The maximum is the amount of threads you have (In Node.JS `os.cpus().length`). You can also provide a percentage value (e.g. 50%).

### `--pixel-format` [​](\#--pixel-format "Direct link to --pixel-format")

[Set a custom pixel format. See here for available values.](/docs/config#setpixelformat)

### `--image-format` [v1.4.0](https://github.com/remotion-dev/remotion/releases/v1.4.0) [​](\#--image-format "Direct link to --image-format")

[`jpeg` or `png` \- JPEG is faster, but doesn't support transparency.](/docs/config#setvideoimageformat) The default image format is `jpeg` since v1.1.

### `--config` [v1.2.0](https://github.com/remotion-dev/remotion/releases/v1.2.0) [​](\#--config "Direct link to --config")

Specify a location for the Remotion config file.

### `--env-file` [v2.2.0](https://github.com/remotion-dev/remotion/releases/v2.2.0) [​](\#--env-file "Direct link to --env-file")

Specify a location for a dotenv file. Default `.env`.

### `--jpeg-quality` [v4.0.0](https://github.com/remotion-dev/remotion/releases/v4.0.0) [​](\#--jpeg-quality "Direct link to --jpeg-quality")

[Value between 0 and 100 for JPEG rendering quality](/docs/config#setjpegquality). Doesn't work when PNG frames are rendered.

### `--quality` [v1.4.0](https://github.com/remotion-dev/remotion/releases/v1.4.0) [​](\#--quality "Direct link to --quality")

Renamed to `--jpeg-quality` in v4.0.0

### `--output` [v4.0.0](https://github.com/remotion-dev/remotion/releases/v4.0.0) [​](\#--output- "Direct link to --output-")

Sets the output file path, as an alternative to the `output-location` positional argument.

### `--overwrite` [​](\#--overwrite "Direct link to --overwrite")

[Write to output even if file already exists.](/docs/config#setoverwriteoutput). This flag is enabled by default, use `--overwrite=false` to disable it.

### `--sequence` [v1.4.0](https://github.com/remotion-dev/remotion/releases/v1.4.0) [​](\#--sequence "Direct link to --sequence")

[Pass this flag if you want an image sequence as the output instead of a video.](/docs/config#setimagesequence)

### `--codec` [v1.4.0](https://github.com/remotion-dev/remotion/releases/v1.4.0) [​](\#--codec "Direct link to --codec")

[`h264` or `h265` or `png` or `vp8` or `vp9` or `mp3` or `aac` or `wav` or `prores` or `h264-mkv`](/docs/config#setcodec). If you don't supply `--codec`, it will use the H.264 encoder.

### `--audio-codec` [v3.3.42](https://github.com/remotion-dev/remotion/releases/v3.3.42) [​](\#--audio-codec "Direct link to --audio-codec")

Set the format of the audio that is embedded in the video. Not all codec and audio codec combinations are supported and certain combinations require a certain file extension and container format. See the table in the docs to see possible combinations.

### `--audio-bitrate` [v3.2.32](https://github.com/remotion-dev/remotion/releases/v3.2.32) [​](\#--audio-bitrate "Direct link to --audio-bitrate")

Specify the target bitrate for the generated video. The syntax for FFmpeg's `-b:a` parameter should be used. FFmpeg may encode the video in a way that will not result in the exact audio bitrate specified. Example values: `512K` for 512 kbps, `1M` for 1 Mbps. Default: `320k`

### `--video-bitrate` [v3.2.32](https://github.com/remotion-dev/remotion/releases/v3.2.32) [​](\#--video-bitrate "Direct link to --video-bitrate")

Specify the target bitrate for the generated video. The syntax for FFmpeg's `-b:v` parameter should be used. FFmpeg may encode the video in a way that will not result in the exact video bitrate specified. Example values: `512K` for 512 kbps, `1M` for 1 Mbps.

### `--buffer-size` [v4.0.78](https://github.com/remotion-dev/remotion/releases/v4.0.78) [​](\#--buffer-size "Direct link to --buffer-size")

The value for the `-bufsize` flag of FFmpeg. Should be used in conjunction with the encoding max rate flag.

### `--max-rate` [v4.0.78](https://github.com/remotion-dev/remotion/releases/v4.0.78) [​](\#--max-rate "Direct link to --max-rate")

The value for the `-maxrate` flag of FFmpeg. Should be used in conjunction with the encoding buffer size flag.

### `--prores-profile` [v2.1.6](https://github.com/remotion-dev/remotion/releases/v2.1.6) [​](\#--prores-profile "Direct link to --prores-profile")

[Set the ProRes profile](/docs/config#setproresprofile). This option is only valid if the [`codec`](#--codec) has been set to `prores`. Possible values: `4444-xq`, `4444`, `hq`, `standard`, `light`, `proxy`. See [here](https://video.stackexchange.com/a/14715) for explanation of possible values. Default: `hq`.

### `--x264-preset` [v4.2.2](https://github.com/remotion-dev/remotion/releases/v4.2.2) [​](\#--x264-preset "Direct link to --x264-preset")

Sets a x264 preset profile. Only applies to videos rendered with `h264` codec.

Possible values: `superfast`, `veryfast`, `faster`, `fast`, `medium`, `slow`, `slower`, `veryslow`, `placebo`.

Default: `medium`

### `--crf` [v1.4.0](https://github.com/remotion-dev/remotion/releases/v1.4.0) [​](\#--crf "Direct link to --crf")

[To set Constant Rate Factor (CRF) of the output](/docs/config#setcrf). Minimum 0. Use this rate control mode if you want to keep the best quality and care less about the file size. This option cannot be set if `--video-bitrate` is set.

### `--browser-executable` [v1.5.0](https://github.com/remotion-dev/remotion/releases/v1.5.0) [​](\#--browser-executable "Direct link to --browser-executable")

[Path to a Chrome executable](/docs/config#setbrowserexecutable). If not specified and Remotion cannot find one, it will download one during rendering.

### `--scale` [​](\#--scale "Direct link to --scale")

[Scales the output frames by the factor you pass in.](/docs/scaling) For example, a 1280x720px frame will become a 1920x1080px frame with a scale factor of `1.5`. Vector elements like fonts and HTML markups will be rendered with extra details. `scale` must be greater than 0 and less than equal to 16. Default: `1`.

### `--frames` [v2.0.0](https://github.com/remotion-dev/remotion/releases/v2.0.0) [​](\#--frames "Direct link to --frames")

[Render a subset of a video](/docs/config#setframerange). Example: `--frames=0-9` to select the first 10 frames. To render a still, use the `still` command.

### `--every-nth-frame` [v3.1.0](https://github.com/remotion-dev/remotion/releases/v3.1.0) [​](\#--every-nth-frame "Direct link to --every-nth-frame")

[Render only every nth frame.](/docs/config#seteverynthframe) This option may only be set when rendering GIFs. This allows you to lower the FPS of the GIF.

For example only every second frame, every third frame and so on. Only works for rendering GIFs. [See here for more details.](/docs/render-as-gif#reducing-frame-rate)

### `--muted` [v3.2.1](https://github.com/remotion-dev/remotion/releases/v3.2.1) [​](\#--muted "Direct link to --muted")

[Disables audio output.](/docs/cli/render#--muted) This option may only be used when rendering a video.

### `--enforce-audio-track` [v3.2.1](https://github.com/remotion-dev/remotion/releases/v3.2.1) [​](\#--enforce-audio-track "Direct link to --enforce-audio-track")

[Render a silent audio track if there wouldn't be one otherwise.](/docs/cli/render#--enforce-audio-track).

### `--number-of-gif-loops` [v3.1.0](https://github.com/remotion-dev/remotion/releases/v3.1.0) [​](\#--number-of-gif-loops "Direct link to --number-of-gif-loops")

Allows you to set the number of loops as follows:

- `null` (or omitting in the CLI) plays the GIF indefinitely.
- `0` disables looping
- `1` loops the GIF once (plays twice in total)
- `2` loops the GIF twice (plays three times in total) and so on.

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

### `--bundle-cache` [v2.0.0](https://github.com/remotion-dev/remotion/releases/v2.0.0) [​](\#--bundle-cache "Direct link to --bundle-cache")

[Enable or disable Webpack caching](/docs/config#setcachingenabled). This flag is enabled by default, use `--bundle-cache=false` to disable caching.

### `--log` [​](\#--log "Direct link to --log")

[Set the log level](/docs/config#setlevel). Increase or decrease the amount of output. Acceptable values: `error`, `warn`, `info` ( _default_), `verbose`

### `--port` [​](\#--port "Direct link to --port")

[Set a custom HTTP server port that will be used to host the Webpack bundle](/docs/config#setrendererport). If not defined, Remotion will try to find a free port.

### `--public-dir` [v3.2.13](https://github.com/remotion-dev/remotion/releases/v3.2.13) [​](\#--public-dir "Direct link to --public-dir")

The path of the URL where the bundle is going to be hosted. By default it is `/`, meaning that the bundle is going to be hosted at the root of the domain (e.g. `https://localhost:3000/`). If you are deploying to a subdirectory (e.g. `/sites/my-site/`), you should set this to the subdirectory.

### `--timeout` [​](\#--timeout "Direct link to --timeout")

Define how long a single frame may take to resolve all [`delayRender()`](/docs/delay-render) calls [before it times out](/docs/timeout) in milliseconds. Default: `30000`.

info

Not to be confused with the [`--timeout` flag when deploying a Lambda function](/docs/lambda/cli/functions#--timeout).

### `--ignore-certificate-errors` [v2.6.5](https://github.com/remotion-dev/remotion/releases/v2.6.5) [​](\#--ignore-certificate-errors "Direct link to --ignore-certificate-errors")

Results in invalid SSL certificates in Chrome, such as self-signed ones, being ignored.

### `--disable-web-security` [v2.6.5](https://github.com/remotion-dev/remotion/releases/v2.6.5) [​](\#--disable-web-security "Direct link to --disable-web-security")

This will most notably disable CORS in Chrome among other security features.

### `--disable-headless` [v2.6.5](https://github.com/remotion-dev/remotion/releases/v2.6.5) [​](\#--disable-headless "Direct link to --disable-headless")

Deprecated - will be removed in 5.0.0. With the migration to [Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell), this option is not functional anymore.

If disabled, the render will open an actual Chrome window where you can see the render happen. The default is headless mode.

### `--gl` [​](\#--gl "Direct link to --gl")

Changelog

- From Remotion v2.6.7 until v3.0.7, the default for Remotion Lambda was `swiftshader`, but from v3.0.8 the default is `swangle` (Swiftshader on Angle) since Chrome 101 added support for it.
- From Remotion v2.4.3 until v2.6.6, the default was `angle`, however it turns out to have a small memory leak that could crash long Remotion renders.

Select the OpenGL renderer backend for Chromium.

Accepted values:

- `"angle"`
- `"egl"`
- `"swiftshader"`
- `"swangle"`
- `"vulkan"` ( _from Remotion v4.0.41_)
- `"angle-egl"` ( _from Remotion v4.0.51_)

The default is `null`, letting Chrome decide, except on Lambda where the default is `"swangle"`

### `--user-agent` [v3.3.83](https://github.com/remotion-dev/remotion/releases/v3.3.83) [​](\#--user-agent "Direct link to --user-agent")

Lets you set a custom user agent that the headless Chrome browser assumes.

### `--offthreadvideo-cache-size-in-bytes` [v4.0.23](https://github.com/remotion-dev/remotion/releases/v4.0.23) [​](\#--offthreadvideo-cache-size-in-bytes "Direct link to --offthreadvideo-cache-size-in-bytes")

From v4.0, Remotion has a cache for [`<OffthreadVideo>`](https://remotion.dev/docs/offthreadvideo) frames. The default is `null`, corresponding to half of the system memory available when the render starts.

This option allows to override the size of the cache. The higher it is, the faster the render will be, but the more memory will be used.

The used value will be printed when running in verbose mode.

Default: `null`

### `--enable-multiprocess-on-linux` [v4.0.42](https://github.com/remotion-dev/remotion/releases/v4.0.42) [​](\#--enable-multiprocess-on-linux "Direct link to --enable-multiprocess-on-linux")

Removes the `--single-process` flag that gets passed to Chromium on Linux by default. This will make the render faster because multiple processes can be used, but may cause issues with some Linux distributions or if window server libraries are missing.

Default: `false` until v4.0.136, then `true` from v4.0.137 on because newer Chrome versions don't allow rendering with the `--single-process` flag.

This flag will be removed in Remotion v5.0.

### `--repro` [v4.0.88](https://github.com/remotion-dev/remotion/releases/v4.0.88) [​](\#--repro "Direct link to --repro")

Create a ZIP that you can submit to Remotion if asked for a reproduction.

### `--binaries-directory` [v4.0.120](https://github.com/remotion-dev/remotion/releases/v4.0.120) [​](\#--binaries-directory "Direct link to --binaries-directory")

The directory where the platform-specific binaries and libraries that Remotion needs are located. Those include an `ffmpeg` and `ffprobe` binary, a Rust binary for various tasks, and various shared libraries. If the value is set to `null`, which is the default, then the path of a platform-specific package located at `node_modules/@remotion/compositor-*` is selected.

This option is useful in environments where Remotion is not officially supported to run like bundled serverless functions or Electron.

### `--for-seamless-aac-concatenation` [v4.0.123](https://github.com/remotion-dev/remotion/releases/v4.0.123) [​](\#--for-seamless-aac-concatenation "Direct link to --for-seamless-aac-concatenation")

If enabled, the audio is trimmed to the nearest AAC frame, which is required for seamless concatenation of AAC files. This is a requirement if you later want to combine multiple video snippets seamlessly.

This option is used internally. There is currently no documentation yet for to concatenate the audio chunks.

### `--separate-audio-to` [v4.0.123](https://github.com/remotion-dev/remotion/releases/v4.0.123) [​](\#--separate-audio-to "Direct link to --separate-audio-to")

If set, the audio will not be included in the main output but rendered as a separate file at the location you pass. It is recommended to use an absolute path. If a relative path is passed, it is relative to the Remotion Root.

### `--metadata` [v4.0.216](https://github.com/remotion-dev/remotion/releases/v4.0.216) [​](\#--metadata "Direct link to --metadata")

Metadata to be embedded in the video. See [here](/docs/metadata) for which metadata is accepted.

The parameter must be in the format of `--metadata key=value` and can be passed multiple times.

### `--ffmpeg-executable` [​](\#--ffmpeg-executable "Direct link to --ffmpeg-executable")

_removed in v4.0_

[Set a custom `ffmpeg` executable](/docs/config#setffmpegexecutable). If not defined, a `ffmpeg` executable will be searched in `PATH`.

### `--ffprobe-executable` [v3.0.17](https://github.com/remotion-dev/remotion/releases/v3.0.17) [​](\#--ffprobe-executable- "Direct link to --ffprobe-executable-")

_removed in v4.0_

[Set a custom `ffprobe` executable](/docs/config#setffprobeexecutable). If not defined, a `ffprobe` executable will be searched in `PATH`.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/render.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
studio](/docs/cli/studio) [Next\
\
still](/docs/cli/still)

- [Flags](#flags)
  - [`--props`](#--props)
  - [`--height`](#--height)
  - [`--width`](#--width)
  - [`--concurrency`](#--concurrency)
  - [`--pixel-format`](#--pixel-format)
  - [`--image-format`](#--image-format)
  - [`--config`](#--config)
  - [`--env-file`](#--env-file)
  - [`--jpeg-quality`](#--jpeg-quality)
  - [`--quality`](#--quality)
  - [`--output`](#--output-)
  - [`--overwrite`](#--overwrite)
  - [`--sequence`](#--sequence)
  - [`--codec`](#--codec)
  - [`--audio-codec`](#--audio-codec)
  - [`--audio-bitrate`](#--audio-bitrate)
  - [`--video-bitrate`](#--video-bitrate)
  - [`--buffer-size`](#--buffer-size)
  - [`--max-rate`](#--max-rate)
  - [`--prores-profile`](#--prores-profile)
  - [`--x264-preset`](#--x264-preset)
  - [`--crf`](#--crf)
  - [`--browser-executable`](#--browser-executable)
  - [`--scale`](#--scale)
  - [`--frames`](#--frames)
  - [`--every-nth-frame`](#--every-nth-frame)
  - [`--muted`](#--muted)
  - [`--enforce-audio-track`](#--enforce-audio-track)
  - [`--number-of-gif-loops`](#--number-of-gif-loops)
  - [`--color-space`](#--color-space)
  - [`--hardware-acceleration`](#--hardware-acceleration)
  - [`--bundle-cache`](#--bundle-cache)
  - [`--log`](#--log)
  - [`--port`](#--port)
  - [`--public-dir`](#--public-dir)
  - [`--timeout`](#--timeout)
  - [`--ignore-certificate-errors`](#--ignore-certificate-errors)
  - [`--disable-web-security`](#--disable-web-security)
  - [`--disable-headless`](#--disable-headless)
  - [`--gl`](#--gl)
  - [`--user-agent`](#--user-agent)
  - [`--offthreadvideo-cache-size-in-bytes`](#--offthreadvideo-cache-size-in-bytes)
  - [`--enable-multiprocess-on-linux`](#--enable-multiprocess-on-linux)
  - [`--repro`](#--repro)
  - [`--binaries-directory`](#--binaries-directory)
  - [`--for-seamless-aac-concatenation`](#--for-seamless-aac-concatenation)
  - [`--separate-audio-to`](#--separate-audio-to)
  - [`--metadata`](#--metadata)
  - [`--ffmpeg-executable`](#--ffmpeg-executable)
  - [`--ffprobe-executable`](#--ffprobe-executable-)
