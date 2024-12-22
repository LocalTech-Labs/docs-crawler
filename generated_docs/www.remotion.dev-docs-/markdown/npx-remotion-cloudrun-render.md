---
title: npx remotion cloudrun render
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion cloudrun render

- [Home page](/)
- [Command line](/docs/cli/)
- [cloudrun](/docs/cloudrun/cli)
- render

On this page

# npx remotion cloudrun render

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs
may change in any version and documentation is not yet finished. See the

[changelog to stay up to date with breaking changes](https://remotion.dev/changelog)

.

Using the `npx remotion cloudrun render` command, you can render a video on GCP.

The structure of a command is as follows:

```

npx remotion cloudrun render <serve-url> [<composition-id>] [<output-location>]
```

```

npx remotion cloudrun render <serve-url> [<composition-id>] [<output-location>]
```

- The serve URL is obtained by deploying a Remotion project to a GCP Storage Bucket using the [`sites create`](/docs/cloudrun/cli/sites#create) command or calling [`deployService()`](/docs/cloudrun/deployservice).
- The [composition ID](/docs/terminology/composition#composition-id). If not specified, the list of compositions will be fetched and you can choose a composition.
- The `output-location` parameter is optional. If you don't specify it, the video is stored in your Cloud Storage bucket. If you specify a location, it gets downloaded to your device in an additional step.

## Example commands [​](\#example-commands "Direct link to Example commands")

Rendering a video, passing the service name:

```

npx remotion cloudrun render https://storage.googleapis.com/remotioncloudrun-123asd321/sites/abcdefgh/index.html tiles --service-name=remotion--3-3-82--mem512mi--cpu1-0--t-800
```

```

npx remotion cloudrun render https://storage.googleapis.com/remotioncloudrun-123asd321/sites/abcdefgh/index.html tiles --service-name=remotion--3-3-82--mem512mi--cpu1-0--t-800
```

Using the site name as opposed to the full serve-url:

```

npx remotion cloudrun render test-site tiles --service-name=remotion--3-3-82--mem512mi--cpu1-0--t-800
```

```

npx remotion cloudrun render test-site tiles --service-name=remotion--3-3-82--mem512mi--cpu1-0--t-800
```

Passing in input props:

```

npx remotion cloudrun render test-site tiles --service-name=remotion--3-3-82--mem512mi--cpu1-0--t-800 --props='{"hi": "there"}'
```

```

npx remotion cloudrun render test-site tiles --service-name=remotion--3-3-82--mem512mi--cpu1-0--t-800 --props='{"hi": "there"}'
```

## Flags [​](\#flags "Direct link to Flags")

### `--region` [​](\#--region "Direct link to --region")

The [GCP region](/docs/cloudrun/region-selection) to select. For lowest latency, the service, site and output bucket should be in the same region.

### `--props` [​](\#--props "Direct link to --props")

[Input Props to pass to the selected composition of your video.](/docs/passing-props#passing-input-props-in-the-cli).

Must be a serialized JSON string ( `--props='{"hello": "world"}'`) or a path to a JSON file ( `./path/to/props.json`).

From the root component the props can be read using [`getInputProps()`](/docs/get-input-props).

You may transform input props using [`calculateMetadata()`](/docs/calculate-metadata).

note

Inline JSON string isn't supported on Windows shells because it removes the `"` character, use a file name instead.

### `--privacy` [​](\#--privacy "Direct link to --privacy")

One of:

- `"public"` ( _default_): The rendered media is publicly accessible under the Cloud Storage URL.
- `"private"`: The rendered media is not publicly available, but is available within the GCP project to those with the correct permissions.

### `--force-bucket-name` [​](\#--force-bucket-name "Direct link to --force-bucket-name")

Specify a specific bucket name to be used for the output. The resulting Google Cloud Storage URL will be in the format `gs://{bucket-name}/renders/{render-id}/{file-name}`. If not set, Remotion will choose the right bucket to use based on the region.

### `--concurrency` [​](\#--concurrency "Direct link to --concurrency")

A number or a string describing how many browser tabs should be opened. Default "50%".

note

Before v4.0.76, this was "100%" by default. It is now aligned to the other server-side rendering APIs.

### `--jpeg-quality` [​](\#--jpeg-quality "Direct link to --jpeg-quality")

[Value between 0 and 100 for JPEG rendering quality](/docs/config#setjpegquality). Doesn't work when PNG frames are rendered.

### `--image-format` [​](\#--image-format "Direct link to --image-format")

[`jpeg` or `png` \- JPEG is faster, but doesn't support transparency.](/docs/config#setvideoimageformat) The default image format is `jpeg`.

### `--scale` [​](\#--scale "Direct link to --scale")

[Scales the output frames by the factor you pass in.](/docs/scaling) For example, a 1280x720px frame will become a 1920x1080px frame with a scale factor of `1.5`. Vector elements like fonts and HTML markups will be rendered with extra details.

### `--env-file` [​](\#--env-file "Direct link to --env-file")

Specify a location for a dotenv file - Default `.env`. [Read about how environment variables work in Remotion.](/docs/env-variables)

### `--out-name` [​](\#--out-name "Direct link to --out-name")

The file name of the media output as stored in the Cloud Storage bucket. By default, it is `out` plus the appropriate file extension, for example: `out.mp4`. Must match `/([0-9a-zA-Z-!_.*'()/]+)/g`.

### `--cloud-run-url` [​](\#--cloud-run-url "Direct link to --cloud-run-url")

Specify the url of the service which should be used to perform the render. You must set either `cloud-run-url` _or_ `service-name`, but not both

### `--service-name` [​](\#--service-name "Direct link to --service-name")

Specify the name of the service which should be used to perform the render. This is used in conjunction with the region to determine the service endpoint, as the same service name can exist across multiple regions. You must set either `cloud-run-url` _or_ `service-name`, but not both

### `--codec` [​](\#--codec "Direct link to --codec")

[`h264` or `h265` or `png` or `vp8` or `mp3` or `aac` or `wav` or `prores`](/docs/config#setcodec). If you don't supply `--codec`, it will use `h264`.

### `--audio-codec` [​](\#--audio-codec "Direct link to --audio-codec")

Set the format of the audio that is embedded in the video. Not all codec and audio codec combinations are supported and certain combinations require a certain file extension and container format. See the table in the docs to see possible combinations.

### `--audio-bitrate` [​](\#--audio-bitrate "Direct link to --audio-bitrate")

Specify the target bitrate for the generated video. The syntax for FFmpeg's `-b:a` parameter should be used. FFmpeg may encode the video in a way that will not result in the exact audio bitrate specified. Example values: `512K` for 512 kbps, `1M` for 1 Mbps. Default: `320k`

### `--video-bitrate` [​](\#--video-bitrate "Direct link to --video-bitrate")

Specify the target bitrate for the generated video. The syntax for FFmpeg's `-b:v` parameter should be used. FFmpeg may encode the video in a way that will not result in the exact video bitrate specified. Example values: `512K` for 512 kbps, `1M` for 1 Mbps.

### `--prores-profile` [​](\#--prores-profile "Direct link to --prores-profile")

[Set the ProRes profile](/docs/config#setproresprofile). This option is only valid if the [`codec`](#--codec) has been set to `prores`. Possible values: `4444-xq`, `4444`, `hq`, `standard`, `light`, `proxy`. See [here](https://video.stackexchange.com/a/14715) for explanation of possible values. Default: `hq`.

### `--x264-preset` [​](\#--x264-preset "Direct link to --x264-preset")

Sets a x264 preset profile. Only applies to videos rendered with `h264` codec.

Possible values: `superfast`, `veryfast`, `faster`, `fast`, `medium`, `slow`, `slower`, `veryslow`, `placebo`.

Default: `medium`

### `--crf` [​](\#--crf "Direct link to --crf")

[To set Constant Rate Factor (CRF) of the output](/docs/config#setcrf). Minimum 0. Use this rate control mode if you want to keep the best quality and care less about the file size.

### `--pixel-format` [​](\#--pixel-format "Direct link to --pixel-format")

[Set a custom pixel format. See here for available values.](/docs/config#setpixelformat)

### `--every-nth-frame` [​](\#--every-nth-frame "Direct link to --every-nth-frame")

[Render only every nth frame.](/docs/config#seteverynthframe) This option may only be set when rendering GIFs. This allows you to lower the FPS of the GIF.

For example only every second frame, every third frame and so on. Only works for rendering GIFs. [See here for more details.](/docs/render-as-gif)

### `--number-of-gif-loops` [​](\#--number-of-gif-loops "Direct link to --number-of-gif-loops")

Allows you to set the number of loops as follows:

- `null` (or omitting in the CLI) plays the GIF indefinitely.
- `0` disables looping
- `1` loops the GIF once (plays twice in total)
- `2` loops the GIF twice (plays three times in total) and so on.

### `--frames` [​](\#--frames "Direct link to --frames")

[Render a subset of a video](/docs/config#setframerange). Example: `--frames=0-9` to select the first 10 frames.

### `--offthreadvideo-cache-size-in-bytes` [v4.0.23](https://github.com/remotion-dev/remotion/releases/v4.0.23) [​](\#--offthreadvideo-cache-size-in-bytes "Direct link to --offthreadvideo-cache-size-in-bytes")

From v4.0, Remotion has a cache for [`<OffthreadVideo>`](https://remotion.dev/docs/offthreadvideo) frames. The default is `null`, corresponding to half of the system memory available when the render starts.

This option allows to override the size of the cache. The higher it is, the faster the render will be, but the more memory will be used.

The used value will be printed when running in verbose mode.

Default: `null`

### `--color-space` [v4.0.28](https://github.com/remotion-dev/remotion/releases/v4.0.28) [​](\#--color-space "Direct link to --color-space")

Color space to use for the video. Acceptable values: `"default"`(default since 5.0), `"bt709"`(since v4.0.28), `"bt2020-ncl"`(since v4.0.88), `"bt2020-cl"`(since v4.0.88), .

For best color accuracy, it is recommended to also use `"png"` as the image format to have accurate color transformations throughout.

Only since v4.0.83, colorspace conversion is actually performed, previously it would only tag the metadata of the video.

### `--metadata` [v4.0.216](https://github.com/remotion-dev/remotion/releases/v4.0.216) [​](\#--metadata "Direct link to --metadata")

Metadata to be embedded in the video. See [here](/docs/metadata) for which metadata is accepted.

The parameter must be in the format of `--metadata key=value` and can be passed multiple times.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/cli/render.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
services](/docs/cloudrun/cli/services) [Next\
\
still](/docs/cloudrun/cli/still)

- [Example commands](#example-commands)
- [Flags](#flags)
  - [`--region`](#--region)
  - [`--props`](#--props)
  - [`--privacy`](#--privacy)
  - [`--force-bucket-name`](#--force-bucket-name)
  - [`--concurrency`](#--concurrency)
  - [`--jpeg-quality`](#--jpeg-quality)
  - [`--image-format`](#--image-format)
  - [`--scale`](#--scale)
  - [`--env-file`](#--env-file)
  - [`--out-name`](#--out-name)
  - [`--cloud-run-url`](#--cloud-run-url)
  - [`--service-name`](#--service-name)
  - [`--codec`](#--codec)
  - [`--audio-codec`](#--audio-codec)
  - [`--audio-bitrate`](#--audio-bitrate)
  - [`--video-bitrate`](#--video-bitrate)
  - [`--prores-profile`](#--prores-profile)
  - [`--x264-preset`](#--x264-preset)
  - [`--crf`](#--crf)
  - [`--pixel-format`](#--pixel-format)
  - [`--every-nth-frame`](#--every-nth-frame)
  - [`--number-of-gif-loops`](#--number-of-gif-loops)
  - [`--frames`](#--frames)
  - [`--offthreadvideo-cache-size-in-bytes`](#--offthreadvideo-cache-size-in-bytes)
  - [`--color-space`](#--color-space)
  - [`--metadata`](#--metadata)