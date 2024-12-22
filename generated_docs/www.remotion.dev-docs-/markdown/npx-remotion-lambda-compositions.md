---
title: npx remotion lambda compositions
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion lambda compositions

- [Home page](/)
- [Command line](/docs/cli/)
- [lambda](/docs/lambda/cli)
- compositions

On this page

# npx remotion lambda compositions

_Available from v3.3.2._

Print list of composition IDs from a serve URL, fetched from inside a Lambda function.

## API [​](\#api "Direct link to API")

```

bash

npx remotion lambda compositions <serve-url>
```

```

bash

npx remotion lambda compositions <serve-url>
```

- Obtain a [Serve URL](/docs/terminology/serve-url) using the [`sites create`](/docs/lambda/cli/sites#create) command or by calling [`deploySite()`](/docs/lambda/deploysite).

Show output

```
looped 60 1080x1080 200 (3.33 sec)
cancel-render 30 920x720 100 (3.33 sec)
iframe 30 1080x1080 10 (0.33 sec)
stagger-test 30 1280x720 100 (3.33 sec)
freeze-example 30 1280x720 300 (10.00 sec)
base-spring 30 1080x1080 100 (3.33 sec)
spring-with-duration 30 1080x1080 100 (3.33 sec)
missing-img 30 1080x1080 10 (0.33 sec)
ten-frame-tester 30 1080x1080 10 (0.33 sec)
framer 30 1080x1080 100 (3.33 sec)
skip-zero-frame 30 1280x720 100 (3.33 sec)
scripts 30 1280x720 100 (3.33 sec)
many-audio 30 1280x720 30 (1.00 sec)
error-on-frame-10 30 1280x720 1000000 (33333.33 sec)
wrapped-in-context 1280x720 Still
drop-dots 30 1080x1080 5400 (180.00 sec)

```

## `remotion lambda compositions` vs. `remotion compositions` [​](\#remotion-lambda-compositions-vs-remotion-compositions "Direct link to remotion-lambda-compositions-vs-remotion-compositions")

You can also get the compositions of a site that is hosted on S3 locally using [`npx remotion compositions`](/docs/cli/compositions).

Vice versa, you can also get the compositions from a [Serve URL](/docs/terminology/serve-url) that is not hosted on AWS Lambda using `npx remotion lambda compositions`.

You should use `npx remotion lambda compositions` if you cannot use [`npx remotion compositions`](/docs/cli/compositions) because the machine cannot run Chrome.

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

### `--env-file` [​](\#--env-file "Direct link to --env-file")

Specify a location for a dotenv file - Default `.env`. [Read about how environment variables work in Remotion.](/docs/env-variables)

### `--log` [​](\#--log "Direct link to --log")

[Set the log level](/docs/config#setlevel). Increase or decrease the amount of output. Acceptable values: `error`, `warn`, `info` ( _default_), `verbose`

info

If you don't feel like passing command line flags every time, consider creating a `remotion.config.ts` [config file](/docs/config).

### `--timeout` [​](\#--timeout "Direct link to --timeout")

Define how long it may take to resolve all [`delayRender()`](/docs/delay-render) calls before the composition fetching times out in milliseconds. Default: `30000`.

info

Not to be confused with the [`--timeout` flag when deploying a Lambda function](/docs/lambda/cli/functions#--timeout).

### `--ignore-certificate-errors` [​](\#--ignore-certificate-errors "Direct link to --ignore-certificate-errors")

Results in invalid SSL certificates in Chrome, such as self-signed ones, being ignored.

### `--disable-web-security` [​](\#--disable-web-security "Direct link to --disable-web-security")

This will most notably disable CORS in Chrome among other security features.

### `--disable-headless` [​](\#--disable-headless "Direct link to --disable-headless")

Deprecated - will be removed in 5.0.0. With the migration to [Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell), this option is not functional anymore.

If disabled, the render will open an actual Chrome window where you can see the render happen. The default is headless mode.

### `--quiet`, `--q` [​](\#--quiet---q "Direct link to --quiet---q")

Only prints the composition IDs, separated by a space.

### `--force-bucket-name` [v3.3.42](https://github.com/remotion-dev/remotion/releases/v3.3.42) [​](\#--force-bucket-name "Direct link to --force-bucket-name")

Specify a specific bucket name to be used. [This is not recommended](/docs/lambda/multiple-buckets), instead let Remotion discover the right bucket automatically.

### `--user-agent` [v3.3.83](https://github.com/remotion-dev/remotion/releases/v3.3.83) [​](\#--user-agent "Direct link to --user-agent")

Lets you set a custom user agent that the headless Chrome browser assumes.

### `--offthreadvideo-cache-size-in-bytes` [v4.0.23](https://github.com/remotion-dev/remotion/releases/v4.0.23) [​](\#--offthreadvideo-cache-size-in-bytes "Direct link to --offthreadvideo-cache-size-in-bytes")

From v4.0, Remotion has a cache for [`<OffthreadVideo>`](https://remotion.dev/docs/offthreadvideo) frames. The default is `null`, corresponding to half of the system memory available when the render starts.

This option allows to override the size of the cache. The higher it is, the faster the render will be, but the more memory will be used.

The used value will be printed when running in verbose mode.

Default: `null`

### `--force-path-style` [v4.0.202](https://github.com/remotion-dev/remotion/releases/v4.0.202) [​](\#--force-path-style "Direct link to --force-path-style")

Passes `forcePathStyle` to the AWS S3 client. If you don't know what this is, you probably don't need it.

## See also [​](\#see-also "Direct link to See also")

- [`getCompositionsOnLambda()`](/docs/lambda/getcompositionsonlambda)
- [`npx remotion compositions`](/docs/cli/compositions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/cli/compositions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
policies](/docs/lambda/cli/policies) [Next\
\
regions](/docs/lambda/cli/regions)

- [API](#api)
- [`remotion lambda compositions` vs. `remotion compositions`](#remotion-lambda-compositions-vs-remotion-compositions)
- [Flags](#flags)
  - [`--props`](#--props)
  - [`--config`](#--config)
  - [`--env-file`](#--env-file)
  - [`--log`](#--log)
  - [`--timeout`](#--timeout)
  - [`--ignore-certificate-errors`](#--ignore-certificate-errors)
  - [`--disable-web-security`](#--disable-web-security)
  - [`--disable-headless`](#--disable-headless)
  - [`--quiet`, `--q`](#--quiet---q)
  - [`--force-bucket-name`](#--force-bucket-name)
  - [`--user-agent`](#--user-agent)
  - [`--offthreadvideo-cache-size-in-bytes`](#--offthreadvideo-cache-size-in-bytes)
  - [`--force-path-style`](#--force-path-style)
- [See also](#see-also)
