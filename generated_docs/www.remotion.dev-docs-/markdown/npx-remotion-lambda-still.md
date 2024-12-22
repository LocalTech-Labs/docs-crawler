---
title: npx remotion lambda still
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion lambda still

- [Home page](/)
- [Command line](/docs/cli/)
- [lambda](/docs/lambda/cli)
- still

On this page

# npx remotion lambda still

Using the `npx remotion lambda still` command, you can render a still frame in the cloud.

The command has the following structure:

```

npx remotion lambda still <serve-url>? [<composition-id>] [<output-location>]
```

```

npx remotion lambda still <serve-url>? [<composition-id>] [<output-location>]
```

- Obtain a [Serve URL](/docs/terminology/serve-url) using the [`sites create`](/docs/lambda/cli/sites#create) command or by calling [`deploySite()`](/docs/lambda/deploysite).
- The [Composition ID](/docs/terminology/composition#composition-id). If not specified, the list of compositions will be fetched and you can choose a composition.
- The `output-location` parameter is optional. If you don't specify it, the still is stored in your S3 bucket. If you specify a location, it gets downloaded to your device in an additional step.

## Example commands [​](\#example-commands "Direct link to Example commands")

Rendering a still:

```

npx remotion lambda still https://remotionlambda-abcdef.s3.eu-central-1.amazonaws.com/sites/testbed/index.html my-comp
```

```

npx remotion lambda still https://remotionlambda-abcdef.s3.eu-central-1.amazonaws.com/sites/testbed/index.html my-comp
```

Rendering using the serve URL shorthand:

```

npx remotion lambda still testbed my-comp
```

```

npx remotion lambda still testbed my-comp
```

info

If you are using the shorthand serve URL, you have to pass a composition ID. Available compositions can only be fetched if a complete serve URL is passed.

Rendering the 10th frame of a composition:

```

npx remotion lambda still --frame=10 testbed my-comp
```

```

npx remotion lambda still --frame=10 testbed my-comp
```

Downloading the result to a `out.png` file:

```

npx remotion lambda still testbed my-comp out.png
```

```

npx remotion lambda still testbed my-comp out.png
```

## Flags [​](\#flags "Direct link to Flags")

### `--frame` [​](\#--frame "Direct link to --frame")

Render a specific frame of a composition. Default `0`

### `--region` [​](\#--region "Direct link to --region")

The [AWS region](/docs/lambda/region-selection) to select. Both project and function should be in this region.

### `--props` [​](\#--props "Direct link to --props")

[Input Props to pass to the selected composition of your video.](/docs/passing-props#passing-input-props-in-the-cli).

Must be a serialized JSON string ( `--props='{"hello": "world"}'`) or a path to a JSON file ( `./path/to/props.json`).

From the root component the props can be read using [`getInputProps()`](/docs/get-input-props).

You may transform input props using [`calculateMetadata()`](/docs/calculate-metadata).

note

Inline JSON string isn't supported on Windows shells because it removes the `"` character, use a file name instead.

### `--scale` [​](\#--scale "Direct link to --scale")

[Scales the output frames by the factor you pass in.](/docs/scaling) For example, a 1280x720px frame will become a 1920x1080px frame with a scale factor of `1.5`. Vector elements like fonts and HTML markups will be rendered with extra details.

### `--log` [​](\#--log "Direct link to --log")

Log level to be used inside the Lambda function. Also, if you set it to `verbose`, a link to CloudWatch will be printed where you can inspect logs.

### `--privacy` [​](\#--privacy "Direct link to --privacy")

Defines if the output media is accessible for everyone or not. Either `public` or `private`, default `public`.

### `--max-retries` [​](\#--max-retries "Direct link to --max-retries")

How many times a single chunk is being retried if it fails to render. Default `1`.

### `--out-name` [​](\#--out-name "Direct link to --out-name")

The file name of the media output as stored in the S3 bucket. By default, it is `out` plus the appropriate file extension, for example: `out.png`. Must match `/([0-9a-zA-Z-!_.*'()/]+)/g`.

### `--jpeg-quality` [​](\#--jpeg-quality "Direct link to --jpeg-quality")

[Value between 0 and 100 for JPEG rendering quality](/docs/config#setjpegquality). Doesn't work when rendering an image format other than JPEG.

### `--quality` [​](\#--quality "Direct link to --quality")

Renamed to `jpegQuality` in `v4.0.0`.

### `--ignore-certificate-errors` [​](\#--ignore-certificate-errors "Direct link to --ignore-certificate-errors")

Results in invalid SSL certificates in Chrome, such as self-signed ones, being ignored.

### `--disable-web-security` [​](\#--disable-web-security "Direct link to --disable-web-security")

This will most notably disable CORS in Chrome among other security features.

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

### `--force-path-style` [v4.0.202](https://github.com/remotion-dev/remotion/releases/v4.0.202) [​](\#--force-path-style "Direct link to --force-path-style")

Passes `forcePathStyle` to the AWS S3 client. If you don't know what this is, you probably don't need it.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/cli/still.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
render](/docs/lambda/cli/render) [Next\
\
policies](/docs/lambda/cli/policies)

- [Example commands](#example-commands)
- [Flags](#flags)
  - [`--frame`](#--frame)
  - [`--region`](#--region)
  - [`--props`](#--props)
  - [`--scale`](#--scale)
  - [`--log`](#--log)
  - [`--privacy`](#--privacy)
  - [`--max-retries`](#--max-retries)
  - [`--out-name`](#--out-name)
  - [`--jpeg-quality`](#--jpeg-quality)
  - [`--quality`](#--quality)
  - [`--ignore-certificate-errors`](#--ignore-certificate-errors)
  - [`--disable-web-security`](#--disable-web-security)
  - [`--user-agent`](#--user-agent)
  - [`--offthreadvideo-cache-size-in-bytes`](#--offthreadvideo-cache-size-in-bytes)
  - [`--delete-after`](#--delete-after)
  - [`--force-path-style`](#--force-path-style)
