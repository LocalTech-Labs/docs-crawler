---
title: Command line reference
source: Remotion Documentation
last_updated: 2024-12-22
---

# Command line reference

- [Home page](/)
- Command line

On this page

# Command line reference

## How to use [​](\#how-to-use "Direct link to How to use")

You can run the CLI by installing `@remotion/cli` and running:

- `npx remotion` inside a npm project
- `yarn remotion` inside a Yarn project
- `pnpm exec remotion` inside a pnpm project.
- `bunx remotion` inside a Bun project

For brevity, in the documentation we always say `npx remotion`.

Inside an npm script, you don't need the `npx` prefix:

```

package.json
json

{
  "scripts": {
    "render": "remotion render"
  }
}
```

```

package.json
json

{
  "scripts": {
    "render": "remotion render"
  }
}
```

### Using Bun [v4.0.118](https://github.com/remotion-dev/remotion/releases/v4.0.118) [​](\#using-bun "Direct link to using-bun")

By default, the `npx remotion` command is being executed using Node.

Even `bunx remotion` is using Node, unless you add the `--bun` flag.

To use Bun, replace `remotion` with `remotionb`.

```

package.json
json

{
  "scripts": {
    "render": "remotionb render"
  }
}
```

```

package.json
json

{
  "scripts": {
    "render": "remotionb render"
  }
}
```

### Using Deno [v4.0.227](https://github.com/remotion-dev/remotion/releases/v4.0.227) [​](\#using-deno "Direct link to using-deno")

Deno is not supported by Remotion.

If you like to experiment nonetheless, use `npx remotiond` to run the Deno version of the CLI.

```

package.json
json

{
  "scripts": {
    "render": "remotiond render"
  }
}
```

```

package.json
json

{
  "scripts": {
    "render": "remotiond render"
  }
}
```

## Commands [​](\#commands "Direct link to Commands")

The following commands are available - you can always run them using `npx remotion` or even without the `npx` prefix if you put the command inside an npm script.

[**studio** \
\
Start the Remotion Studio](/docs/cli/studio) [**render** \
\
Render video or audio](/docs/cli/render) [**still** \
\
Render a still image](/docs/cli/still) [**compositions** \
\
List available compositions](/docs/cli/compositions) [**lambda** \
\
Control Remotion Lambda](/docs/lambda/cli) [**bundle** \
\
Create a Remotion Bundle](/docs/cli/bundle) [**browser** \
\
Ensure Remotion has a browser to use](/docs/cli/browser) [**cloudrun** \
\
Control Remotion Cloud Run](/docs/cloudrun/cli) [**benchmark** \
\
Measure and optimize render time](/docs/cli/benchmark) [**versions** \
\
List and validate Remotion package versions](/docs/cli/versions) [**upgrade** \
\
Upgrade to a newer version](/docs/cli/upgrade) [**gpu** \
\
Print information about Chrome's usage of the GPU](/docs/cli/gpu) [**ffmpeg** \
\
Execute an `ffmpeg` command](/docs/cli/ffmpeg) [**ffprobe** \
\
Execute an `ffprobe` command](/docs/cli/ffprobe) [**help** \
\
Show CLI commands](/docs/cli/help)

## Example command [​](\#example-command "Direct link to Example command")

```

npx remotion render --codec=vp8 HelloWorld out/video.webm
```

```

npx remotion render --codec=vp8 HelloWorld out/video.webm
```

## Fig.io autocompletion [​](\#figio-autocompletion "Direct link to Fig.io autocompletion")

Install [Fig](https://fig.io) (macOS only) to add Remotion autocomplete to your terminal.

Type `npx remotion` to start getting suggestions.

## See also [​](\#see-also "Direct link to See also")

- [Render your video](/docs/render)
- [Configuration file](/docs/config)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/cli.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
API overview](/docs/api) [Next\
\
studio](/docs/cli/studio)

- [How to use](#how-to-use)
  - [Using Bun](#using-bun)
  - [Using Deno](#using-deno)
- [Commands](#commands)
- [Example command](#example-command)
- [Fig.io autocompletion](#figio-autocompletion)
- [See also](#see-also)
