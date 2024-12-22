---
title: npx remotion lambda render
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion lambda render

- [Home page](/)
- [Command line](/docs/cli/)
- [lambda](/docs/lambda/cli)
- render

On this page

# npx remotion lambda render

Using the `npx remotion lambda render` command, you can render a video in the cloud.

The structure of a command is as follows:

```

npx remotion lambda render <serve-url> [<composition-id>] [<output-location>]
```

```

npx remotion lambda render <serve-url> [<composition-id>] [<output-location>]
```

Arguments:

- Obtain a [Serve URL](/docs/terminology/serve-url) using the [`sites create`](/docs/lambda/cli/sites#create) command or by calling [`deploySite()`](/docs/lambda/deploysite).
- The [Composition ID](/docs/terminology/composition#composition-id). If not specified, the list of compositions will be fetched and you can choose a composition.
- The `output-location` parameter is optional. If you don't specify it, the video is stored in your S3 bucket. If you specify a location, it gets downloaded to your device in an additional step.

## Example commands [​](\#example-commands "Direct link to Example commands")

Rendering a video:

```

npx remotion lambda render https://remotionlambda-abcdef.s3.eu-central-1.amazonaws.com/sites/testbed/index.html my-comp
```

```

npx remotion lambda render https://remotionlambda-abcdef.s3.eu-central-1.amazonaws.com/sites/testbed/index.html my-comp
```

Rendering a video and saving it to `out/video.mp4`:

```

npx remotion lambda render https://remotionlambda-abcdef.s3.eu-central-1.amazonaws.com/sites/testbed/index.html my-comp out/video.mp4
```

```

npx remotion lambda render https://remotionlambda-abcdef.s3.eu-central-1.amazonaws.com/sites/testbed/index.html my-comp out/video.mp4
```

Using the shorthand serve URL:

```

npx remotion lambda render testbed my-comp
```

```

npx remotion lambda render testbed my-comp
```

info

If you are using the shorthand serve URL, you have to pass a composition ID. Available compositions can only be fetched if a complete serve URL is passed.

Passing in input props:

```

npx remotion lambda render --props='{"hi": "there"}' testbed my-comp
```

```

npx remotion lambda render --props='{"hi": "there"}' testbed my-comp
```

Printing debug information including a CloudWatch link:

```

npx remotion lambda render --log=verbose testbed my-comp
```

```

npx remotion lambda render --log=verbose testbed my-comp
```

Keeping the output video private:

```

npx remotion lambda render --privacy=private testbed my-comp
```

```

npx remotion lambda render --privacy=private testbed my-comp
```

Rendering only the audio:

```

npx remotion lambda render --codec=mp3 testbed my-comp
```

```

npx remotion lambda render --codec=mp3 testbed my-comp
```

## Flags [​](\#flags "Direct link to Flags")

### `--region` [​](\#--region "Direct link to --region")

The [AWS region](/docs/lambda/region-selection) to select. Both project and function should be in this region.

### `--props` [​](\#--props "Direct link to --props")

[Input Props to pass to the selected composition of your video.](/docs/passing-props#passing-input-props-in-the-cli).

Must be a serialized JSON string ( `--props='{"hello": "world"}'`) or a path to a JSON file ( `./path/to/props.json`).

From the root component the props can be read using [`getInputProps()`](/docs/get-input-props).

You may transform input props using [`calculateMetadata()`](/docs/calculate-metadata).

note

Inline JSON string isn't supported on Windows shells because it removes the `"` character, use a file name instead.

### `--log` [​](\#--log "Direct link to --log")

Log level to be used inside the Lambda function. Also, if you set it to `verbose`, a link to CloudWatch will be printed where you can inspect logs.

### `--privacy` [​](\#--privacy "Direct link to --privacy")

One of:

- `"public"` ( _default_): The rendered media is publicly accessible under the S3 URL.
- `"private"`: The rendered media is not publicly available, but signed links can be created using [presignUrl()](/docs/lambda/presignurl).
- `"no-acl"` ( _available from v.3.1.7_): The ACL option is not being set at all, this option is useful if you are writing to another bucket that does not support ACL using [`outName`](/docs/lambda/cli/render#--out-name).

### `--max-retries` [​](\#--max-retries "Direct link to --max-retries")

How many times a single chunk is being retried if it fails to render. Default `1`.

### `--frames-per-lambda` [​](\#--frames-per-lambda "Direct link to --frames-per-lambda")

How many frames should be rendered in a single Lambda function. Increase it to require less Lambda functions to render the video, decrease it to make the render faster.

Default value: [Dependant on video length](/docs/lambda/concurrency)

Minimum value: `4`

note

The `framesPerLambda` parameter cannot result in more than 200 functions being spawned. See: [Concurrency](/docs/lambda/concurrency)

### `--concurrency-per-lambda` [v3.0.30](https://github.com/remotion-dev/remotion/releases/v3.0.30) [​](\#--concurrency-per-lambda "Direct link to --concurrency-per-lambda")

By default, each Lambda function renders with concurrency 1 (one open browser tab). You may use the option to customize this value.

### `--jpeg-quality` [​](\#--jpeg-quality "Direct link to --jpeg-quality")

[Value between 0 and 100 for JPEG rendering quality](/docs/config#setjpegquality). Doesn't work when PNG frames are rendered.

### `--quality` [​](\#--quality "Direct link to --quality")

Renamed to `jpegQuality` in `v4.0.0`.

### `--muted` [v3.2.1](https://github.com/remotion-dev/remotion/releases/v3.2.1) [​](\#--muted "Direct link to --muted")

[Disables audio output.](/docs/cli/render#--muted) This option may only be used when rendering a video.

### `--codec` [​](\#--codec "Direct link to --codec")

[`h264` or `h265` (supported since v4.0.32) or `png` or `vp8` or `mp3` or `aac` or `wav` or `prores`](/docs/config#setcodec). If you don't supply `--codec`, it will use `h264`.

### `--audio-codec` [v3.3.42](https://github.com/remotion-dev/remotion/releases/v3.3.42) [​](\#--audio-codec "Direct link to --audio-codec")

Set the format of the audio that is embedded in the video. Not all codec and audio codec combinations are supported and certain combinations require a certain file extension and container format. See the table in the docs to see possible combinations.

### `--audio-bitrate` [v3.2.32](https://github.com/remotion-dev/remotion/releases/v3.2.32) [​](\#--audio-bitrate "Direct link to --audio-bitrate")

Specify the target bitrate for the generated video. The syntax for FFmpeg's `-b:a` parameter should be used. FFmpeg may encode the video in a way that will not result in the exact audio bitrate specified. Example values: `512K` for 512 kbps, `1M` for 1 Mbps. Default: `320k`

### `--video-bitrate` [v3.2.32](https://github.com/remotion-dev/remotion/releases/v3.2.32) [​](\#--video-bitrate "Direct link to --video-bitrate")

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

### `--image-format` [​](\#--image-format "Direct link to --image-format")

[`jpeg` or `png` \- JPEG is faster, but doesn't support transparency.](/docs/config#setvideoimageformat) The default image format is `jpeg`.

### `--scale` [​](\#--scale "Direct link to --scale")

[Scales the output frames by the factor you pass in.](/docs/scaling) For example, a 1280x720px frame will become a 1920x1080px frame with a scale factor of `1.5`. Vector elements like fonts and HTML markups will be rendered with extra details.

### `--env-file` [​](\#--env-file "Direct link to --env-file")

Specify a location for a dotenv file - Default `.env`. [Read about how environment variables work in Remotion.](/docs/env-variables)

### `--frames` [​](\#--frames "Direct link to --frames")

[Render a subset of a video](/docs/config#setframerange). Example: `--frames=0-9` to select the first 10 frames. To render a still, use the [`still`](/docs/lambda/cli/still) command.

### `--every-nth-frame` [v3.1.0](https://github.com/remotion-dev/remotion/releases/v3.1.0) [​](\#--every-nth-frame "Direct link to --every-nth-frame")

[Render only every nth frame.](/docs/config#seteverynthframe) This option may only be set when rendering GIFs. This allows you to lower the FPS of the GIF.

For example only every second frame, every third frame and so on. Only works for rendering GIFs. [See here for more details.](/docs/render-as-gif)

### `--number-of-gif-loops` [v3.1.0](https://github.com/remotion-dev/remotion/releases/v3.1.0) [​](\#--number-of-gif-loops "Direct link to --number-of-gif-loops")

Allows you to set the number of loops as follows:

- `null` (or omitting in the CLI) plays the GIF indefinitely.
- `0` disables looping
- `1` loops the GIF once (plays twice in total)
- `2` loops the GIF twice (plays three times in total) and so on.

### `--timeout` [​](\#--timeout "Direct link to --timeout")

A number describing how long the render may take to resolve all [`delayRender()`](/docs/delay-render) calls [before it times out](/docs/timeout). Default: `30000`

### `--out-name` [​](\#--out-name "Direct link to --out-name")

The file name of the media output as stored in the S3 bucket. By default, it is `out` plus the appropriate file extension, for example: `out.mp4`. Must match `/([0-9a-zA-Z-!_.*'()/]+)/g`.

### `--overwrite` [v3.2.25](https://github.com/remotion-dev/remotion/releases/v3.2.25) [​](\#--overwrite "Direct link to --overwrite")

If a custom out name is specified and a file already exists at this key in the S3 bucket, decide whether that file will be deleted before the render begins. Default `false`.

An existing file at the output S3 key will conflict with the render and must be deleted beforehand. If this setting is `false` and a conflict occurs, an error will be thrown.

### `--webhook` [v3.2.30](https://github.com/remotion-dev/remotion/releases/v3.2.30) [​](\#--webhook "Direct link to --webhook")

Sets a webhook to be called when the render finishes or fails. [`renderMediaOnLambda() -> webhook.url`](/docs/lambda/rendermediaonlambda#webhook). To be used together with `--webhook-secret`.

### `--webhook-secret` [v3.2.30](https://github.com/remotion-dev/remotion/releases/v3.2.30) [​](\#--webhook-secret "Direct link to --webhook-secret")

Sets a webhook secret for the webhook (see above). [`renderMediaOnLambda() -> webhook.secret`](/docs/lambda/rendermediaonlambda#webhook). To be used together with `--webhook`.

### `--height` [v3.2.40](https://github.com/remotion-dev/remotion/releases/v3.2.40) [​](\#--height "Direct link to --height")

[Overrides composition height.](/docs/config#overrideheight)

### `--width` [v3.2.40](https://github.com/remotion-dev/remotion/releases/v3.2.40) [​](\#--width "Direct link to --width")

[Overrides composition width.](/docs/config#overridewidth)

### `--function-name` [v3.3.38](https://github.com/remotion-dev/remotion/releases/v3.3.38) [​](\#--function-name "Direct link to --function-name")

Specify the name of the function which should be used to invoke and orchestrate the render. You only need to pass it if there are multiple functions with different configurations.

### `--renderer-function-name` [v3.3.38](https://github.com/remotion-dev/remotion/releases/v3.3.38) [​](\#--renderer-function-name "Direct link to --renderer-function-name")

If specified, this function will be used for rendering the individual chunks. This is useful if you want to use a function with higher or lower power for rendering the chunks than the main orchestration function.

If you want to use this option, the function must be in the same region, the same account and have the same version as the main function.

### `--force-bucket-name` [v3.3.42](https://github.com/remotion-dev/remotion/releases/v3.3.42) [​](\#--force-bucket-name "Direct link to --force-bucket-name")

Specify a specific bucket name to be used. [This is not recommended](/docs/lambda/multiple-buckets), instead let Remotion discover the right bucket automatically.

### `--ignore-certificate-errors` [​](\#--ignore-certificate-errors "Direct link to --ignore-certificate-errors")

Results in invalid SSL certificates in Chrome, such as self-signed ones, being ignored.

### `--disable-web-security` [​](\#--disable-web-security "Direct link to --disable-web-security")

This will most notably disable CORS in Chrome among other security features.

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

### `--delete-after` [v4.0.32](https://github.com/remotion-dev/remotion/releases/v4.0.32) [​](\#--delete-after "Direct link to --delete-after")

Automatically delete the render after a certain period. Accepted values are `1-day`, `3-days`, `7-days` and `30-days`.

For this to work, your bucket needs to have [lifecycles enabled](/docs/lambda/autodelete).

### `--webhook-custom-data` [v4.0.25](https://github.com/remotion-dev/remotion/releases/v4.0.25) [​](\#--webhook-custom-data "Direct link to --webhook-custom-data")

Pass up to 1,024 bytes of a JSON-serializable object to the webhook. This data will be included in the webhook payload.Alternatively, pass a file path pointing to a JSON file

### `--color-space` [v4.0.28](https://github.com/remotion-dev/remotion/releases/v4.0.28) [​](\#--color-space "Direct link to --color-space")

Color space to use for the video. Acceptable values: `"default"`(default since 5.0), `"bt709"`(since v4.0.28), `"bt2020-ncl"`(since v4.0.88), `"bt2020-cl"`(since v4.0.88), .

For best color accuracy, it is recommended to also use `"png"` as the image format to have accurate color transformations throughout.

Only since v4.0.83, colorspace conversion is actually performed, previously it would only tag the metadata of the video.

### `--prefer-lossless` [v4.0.110](https://github.com/remotion-dev/remotion/releases/v4.0.110) [​](\#--prefer-lossless "Direct link to --prefer-lossless")

Uses a lossless audio codec, if one is available for the codec. If you set `audioCodec`, it takes priority over `preferLossless`.

### `--metadata` [v4.0.216](https://github.com/remotion-dev/remotion/releases/v4.0.216) [​](\#--metadata "Direct link to --metadata")

Metadata to be embedded in the video. See [here](/docs/metadata) for which metadata is accepted.

The parameter must be in the format of `--metadata key=value` and can be passed multiple times.

### `--force-path-style` [v4.0.202](https://github.com/remotion-dev/remotion/releases/v4.0.202) [​](\#--force-path-style "Direct link to --force-path-style")

Passes `forcePathStyle` to the AWS S3 client. If you don't know what this is, you probably don't need it.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/cli/render.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
functions](/docs/lambda/cli/functions) [Next\
\
still](/docs/lambda/cli/still)

- [Example commands](#example-commands)
- [Flags](#flags)
  - [`--region`](#--region)
  - [`--props`](#--props)
  - [`--log`](#--log)
  - [`--privacy`](#--privacy)
  - [`--max-retries`](#--max-retries)
  - [`--frames-per-lambda`](#--frames-per-lambda)
  - [`--concurrency-per-lambda`](#--concurrency-per-lambda)
  - [`--jpeg-quality`](#--jpeg-quality)
  - [`--quality`](#--quality)
  - [`--muted`](#--muted)
  - [`--codec`](#--codec)
  - [`--audio-codec`](#--audio-codec)
  - [`--audio-bitrate`](#--audio-bitrate)
  - [`--video-bitrate`](#--video-bitrate)
  - [`--prores-profile`](#--prores-profile)
  - [`--x264-preset`](#--x264-preset)
  - [`--crf`](#--crf)
  - [`--pixel-format`](#--pixel-format)
  - [`--image-format`](#--image-format)
  - [`--scale`](#--scale)
  - [`--env-file`](#--env-file)
  - [`--frames`](#--frames)
  - [`--every-nth-frame`](#--every-nth-frame)
  - [`--number-of-gif-loops`](#--number-of-gif-loops)
  - [`--timeout`](#--timeout)
  - [`--out-name`](#--out-name)
  - [`--overwrite`](#--overwrite)
  - [`--webhook`](#--webhook)
  - [`--webhook-secret`](#--webhook-secret)
  - [`--height`](#--height)
  - [`--width`](#--width)
  - [`--function-name`](#--function-name)
  - [`--renderer-function-name`](#--renderer-function-name)
  - [`--force-bucket-name`](#--force-bucket-name)
  - [`--ignore-certificate-errors`](#--ignore-certificate-errors)
  - [`--disable-web-security`](#--disable-web-security)
  - [`--gl`](#--gl)
  - [`--user-agent`](#--user-agent)
  - [`--offthreadvideo-cache-size-in-bytes`](#--offthreadvideo-cache-size-in-bytes)
  - [`--delete-after`](#--delete-after)
  - [`--webhook-custom-data`](#--webhook-custom-data)
  - [`--color-space`](#--color-space)
  - [`--prefer-lossless`](#--prefer-lossless)
  - [`--metadata`](#--metadata)
  - [`--force-path-style`](#--force-path-style)
