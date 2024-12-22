---
title: npx remotion studio
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion studio

- [Home page](/)
- [Command line](/docs/cli/)
- studio

On this page

# npx remotion studio

_Alias: npx remotion preview_

Start the [Remotion Studio](/docs/studio).

```

bash

npx remotion studio <entry-point>?
```

```

bash

npx remotion studio <entry-point>?
```

You may pass an [entry point](/docs/terminology/entry-point) as an argument, otherwise it will be [determined](/docs/terminology/entry-point#which-entry-point-is-being-used).

## Flags [​](\#flags "Direct link to Flags")

### `--props` [​](\#--props "Direct link to --props")

[Input Props to pass to the selected composition of your video.](/docs/passing-props#passing-input-props-in-the-cli).

We don't recommend passing this flag when using the Studio - use [`defaultProps`](/docs/composition#defaultprops) instead.

Must be a serialized JSON string ( `--props='{"hello": "world"}'`) or a path to a JSON file ( `./path/to/props.json`).

From the root component the props can be read using [`getInputProps()`](/docs/get-input-props).

You may transform input props using [`calculateMetadata()`](/docs/calculate-metadata).

note

Inline JSON string isn't supported on Windows shells because it removes the `"` character, use a file name instead.

### `--config` [v1.2.0](https://github.com/remotion-dev/remotion/releases/v1.2.0) [​](\#--config "Direct link to --config")

Specify a location for the Remotion config file.

### `--env-file` [v2.2.0](https://github.com/remotion-dev/remotion/releases/v2.2.0) [​](\#--env-file "Direct link to --env-file")

Specify a location for a dotenv file - Default `.env`. [Read about how environment variables work in Remotion.](/docs/env-variables)

### `--log` [​](\#--log "Direct link to --log")

[Set the log level](/docs/config#setlevel). Increase or decrease the amount of output. Acceptable values: `error`, `warn`, `info` ( _default_), `verbose`

### `--port` [​](\#--port "Direct link to --port")

[Set a custom HTTP server port to start the server on](/docs/config#setstudioport). If not defined, Remotion will try to find a free port.

### `--public-dir` [v3.2.13](https://github.com/remotion-dev/remotion/releases/v3.2.13) [​](\#--public-dir "Direct link to --public-dir")

The path of the URL where the bundle is going to be hosted. By default it is `/`, meaning that the bundle is going to be hosted at the root of the domain (e.g. `https://localhost:3000/`). If you are deploying to a subdirectory (e.g. `/sites/my-site/`), you should set this to the subdirectory.

### `--disable-keyboard-shortcuts` [v3.2.11](https://github.com/remotion-dev/remotion/releases/v3.2.11) [​](\#--disable-keyboard-shortcuts "Direct link to --disable-keyboard-shortcuts")

[Disables all keyboard shortcuts in the Studio](/docs/config#setkeyboardshortcutsenabled).

### `--webpack-poll` [v3.3.11](https://github.com/remotion-dev/remotion/releases/v3.3.11) [​](\#--webpack-poll "Direct link to --webpack-poll")

[Enables Webpack polling](/docs/config#setwebpackpollinginmilliseconds) instead of the file system event listeners for hot reloading. This is useful if you are inside a virtual machine or have a remote file system.
Pass a value in milliseconds, for example `--webpack-poll=1000`.

### `--no-open` [v3.3.19](https://github.com/remotion-dev/remotion/releases/v3.3.19) [​](\#--no-open "Direct link to --no-open")

[Prevents Remotion from trying to open a browser](/docs/config#setshouldopenbrowser). This is useful if you use a different browser for Remotion than the operating system default.

### `--browser` [v3.3.79](https://github.com/remotion-dev/remotion/releases/v3.3.79) [​](\#--browser "Direct link to --browser")

Specify the browser which should be used for opening tab - using the default browser by default.

Pass an absolute string or `"chrome"` to use Chrome.
If Chrome is selected as the browser and you are on macOS, Remotion will try to reuse an existing tab

For backwards compatibility, the `BROWSER` environment variable is also supported.

### `--browser-args` [v3.3.79](https://github.com/remotion-dev/remotion/releases/v3.3.79) [​](\#--browser-args "Direct link to --browser-args")

A set of command line flags that should be passed to the browser. Pass them like this:

```

sh

npx remotion studio --browser-args="--disable-web-security"
```

```

sh

npx remotion studio --browser-args="--disable-web-security"
```

### `--beep-on-finish` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#--beep-on-finish "Direct link to --beep-on-finish")

[Plays a beep sound when the video is finished rendering](/docs/config#setbeeponfinish). This is useful if you are rendering a video in the background and want to be notified when it is finished.

```

sh

npx remotion studio --beep-on-finish
```

```

sh

npx remotion studio --beep-on-finish
```

### `--ipv4` [v4.0.125](https://github.com/remotion-dev/remotion/releases/v4.0.125) [​](\#--ipv4 "Direct link to --ipv4")

Forces the Studio to be bound to an IPv4 interface, even if a IPv6 interface is available.

```

sh

npx remotion studio --ipv4
```

```

sh

npx remotion studio --ipv4
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/studio.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
CLI reference](/docs/cli/) [Next\
\
render](/docs/cli/render)

- [Flags](#flags)
  - [`--props`](#--props)
  - [`--config`](#--config)
  - [`--env-file`](#--env-file)
  - [`--log`](#--log)
  - [`--port`](#--port)
  - [`--public-dir`](#--public-dir)
  - [`--disable-keyboard-shortcuts`](#--disable-keyboard-shortcuts)
  - [`--webpack-poll`](#--webpack-poll)
  - [`--no-open`](#--no-open)
  - [`--browser`](#--browser)
  - [`--browser-args`](#--browser-args)
  - [`--beep-on-finish`](#--beep-on-finish)
  - [`--ipv4`](#--ipv4)
