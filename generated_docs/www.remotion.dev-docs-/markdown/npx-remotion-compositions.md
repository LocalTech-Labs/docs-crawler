---
title: npx remotion compositions
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion compositions

- [Home page](/)
- [Command line](/docs/cli/)
- compositions

On this page

# npx remotion compositions

_Available from v2.6.12._

Print list of composition IDs based on a path of an entry point.

```

bash

npx remotion compositions <serve-url|entry-file>?
```

```

bash

npx remotion compositions <serve-url|entry-file>?
```

You may pass a [Serve URL](/docs/terminology/serve-url) or an [entry point](/docs/terminology/entry-point) as the first argument, otherwise the entry point will be [determined](/docs/terminology/entry-point#which-entry-point-is-being-used).

## Flags [​](\#flags "Direct link to Flags")

### `--props` [​](\#--props "Direct link to --props")

[Input Props to pass to the selected composition of your video.](/docs/passing-props#passing-input-props-in-the-cli).

Must be a serialized JSON string ( `--props='{"hello": "world"}'`) or a path to a JSON file ( `./path/to/props.json`).

From the root component the props can be read using [`getInputProps()`](/docs/get-input-props).

You may transform input props using [`calculateMetadata()`](/docs/calculate-metadata).

note

Inline JSON string isn't supported on Windows shells because it removes the `"` character, use a file name instead.

### `--config` [​](\#--config "Direct link to --config")

Specify a location for the Remotion config file.

### `--env-file` [v2.2.0](https://github.com/remotion-dev/remotion/releases/v2.2.0) [​](\#--env-file "Direct link to --env-file")

Specify a location for a dotenv file - Default `.env`. [Read about how environment variables work in Remotion.](/docs/env-variables)

### `--bundle-cache` [​](\#--bundle-cache "Direct link to --bundle-cache")

[Enable or disable Webpack caching](/docs/config#setcachingenabled). This flag is enabled by default, use `--bundle-cache=false` to disable caching.

### `--log` [​](\#--log "Direct link to --log")

[Set the log level](/docs/config#setlevel). Increase or decrease the amount of output. Acceptable values: `error`, `warn`, `info` ( _default_), `verbose`

info

If you don't feel like passing command line flags every time, consider creating a `remotion.config.ts` [config file](/docs/config).

### `--port` [​](\#--port "Direct link to --port")

[Set a custom HTTP server port to host the Webpack bundle](/docs/config#setport). If not defined, Remotion will try to find a free port.

### `--public-dir` [v3.2.13](https://github.com/remotion-dev/remotion/releases/v3.2.13) [​](\#--public-dir "Direct link to --public-dir")

The path of the URL where the bundle is going to be hosted. By default it is `/`, meaning that the bundle is going to be hosted at the root of the domain (e.g. `https://localhost:3000/`). If you are deploying to a subdirectory (e.g. `/sites/my-site/`), you should set this to the subdirectory.

### `--timeout` [​](\#--timeout "Direct link to --timeout")

Define how long it may take to resolve all [`delayRender()`](/docs/delay-render) calls before the composition fetching times out in milliseconds. Default: `30000`.

info

Not to be confused with the [`--timeout` flag when deploying a Lambda function](/docs/lambda/cli/functions#--timeout).

### `--ignore-certificate-errors` [​](\#--ignore-certificate-errors "Direct link to --ignore-certificate-errors")

Results in invalid SSL certificates in Chrome, such as self-signed ones, being ignored.

### `--disable-web-security` [​](\#--disable-web-security "Direct link to --disable-web-security")

_available since v2.6.5_

This will most notably disable CORS in Chrome among other security features.

### `--disable-headless?` [​](\#--disable-headless "Direct link to --disable-headless")

Deprecated - will be removed in 5.0.0. With the migration to [Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell), this option is not functional anymore.

If disabled, the render will open an actual Chrome window where you can see the render happen. The default is headless mode.

### `--enable-multiprocess-on-linux` [v4.0.42](https://github.com/remotion-dev/remotion/releases/v4.0.42) [​](\#--enable-multiprocess-on-linux "Direct link to --enable-multiprocess-on-linux")

Removes the `--single-process` flag that gets passed to Chromium on Linux by default. This will make the render faster because multiple processes can be used, but may cause issues with some Linux distributions or if window server libraries are missing.

Default: `false` until v4.0.136, then `true` from v4.0.137 on because newer Chrome versions don't allow rendering with the `--single-process` flag.

This flag will be removed in Remotion v5.0.

### `--user-agent` [v3.3.83](https://github.com/remotion-dev/remotion/releases/v3.3.83) [​](\#--user-agent "Direct link to --user-agent")

Lets you set a custom user agent that the headless Chrome browser assumes.

### `--offthreadvideo-cache-size-in-bytes` [v4.0.23](https://github.com/remotion-dev/remotion/releases/v4.0.23) [​](\#--offthreadvideo-cache-size-in-bytes "Direct link to --offthreadvideo-cache-size-in-bytes")

From v4.0, Remotion has a cache for [`<OffthreadVideo>`](https://remotion.dev/docs/offthreadvideo) frames. The default is `null`, corresponding to half of the system memory available when the render starts.

This option allows to override the size of the cache. The higher it is, the faster the render will be, but the more memory will be used.

The used value will be printed when running in verbose mode.

Default: `null`

### `--binaries-directory` [v4.0.120](https://github.com/remotion-dev/remotion/releases/v4.0.120) [​](\#--binaries-directory "Direct link to --binaries-directory")

The directory where the platform-specific binaries and libraries that Remotion needs are located. Those include an `ffmpeg` and `ffprobe` binary, a Rust binary for various tasks, and various shared libraries. If the value is set to `null`, which is the default, then the path of a platform-specific package located at `node_modules/@remotion/compositor-*` is selected.

This option is useful in environments where Remotion is not officially supported to run like bundled serverless functions or Electron.

### `--quiet`, `--q` [​](\#--quiet---q "Direct link to --quiet---q")

Only prints the composition IDs, separated by a space.

### `--ffmpeg-executable` [​](\#--ffmpeg-executable "Direct link to --ffmpeg-executable")

_removed in v4.0_

[Set a custom `ffmpeg` executable](/docs/config#setffmpegexecutable). If not defined, a `ffmpeg` executable will be searched in `PATH`.

### `--ffprobe-executable` [​](\#--ffprobe-executable "Direct link to --ffprobe-executable")

_removed in v4.0_

[Set a custom `ffprobe` executable](/docs/config#setffprobeexecutable). If not defined, a `ffprobe` executable will be searched in `PATH`.

## See also [​](\#see-also "Direct link to See also")

- [`getCompositions()`](/docs/renderer/get-compositions)
- [`getCompositionsOnLambda()`](/docs/lambda/getcompositionsonlambda)
- [`npx remotion lambda compositions`](/docs/lambda/cli/compositions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/compositions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
still](/docs/cli/still) [Next\
\
bundle](/docs/cli/bundle)

- [Flags](#flags)
  - [`--props`](#--props)
  - [`--config`](#--config)
  - [`--env-file`](#--env-file)
  - [`--bundle-cache`](#--bundle-cache)
  - [`--log`](#--log)
  - [`--port`](#--port)
  - [`--public-dir`](#--public-dir)
  - [`--timeout`](#--timeout)
  - [`--ignore-certificate-errors`](#--ignore-certificate-errors)
  - [`--disable-web-security`](#--disable-web-security)
  - [`--disable-headless?`](#--disable-headless)
  - [`--enable-multiprocess-on-linux`](#--enable-multiprocess-on-linux)
  - [`--user-agent`](#--user-agent)
  - [`--offthreadvideo-cache-size-in-bytes`](#--offthreadvideo-cache-size-in-bytes)
  - [`--binaries-directory`](#--binaries-directory)
  - [`--quiet`, `--q`](#--quiet---q)
  - [`--ffmpeg-executable`](#--ffmpeg-executable)
  - [`--ffprobe-executable`](#--ffprobe-executable)
- [See also](#see-also)
