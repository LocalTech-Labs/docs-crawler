---
title: getInputProps()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getInputProps()

- [Home page](/)
- [remotion](/docs/remotion)
- getInputProps()

On this page

# getInputProps()

_Available from v2.0_.

Using this method, you can retrieve inputs that you pass in from the command line using [`--props`](/docs/cli), or the [`inputProps`](/docs/ssr) parameter if you are using the Node.JS API.

You should whenever possible prefer to retrieve props directly in your composition, like you would read props from a component if you were to code a React application, but this method is useful if you want to retrieve the input props outside of a composition.

info

This method is not available when inside a Remotion Player. Instead, get the props as React props from the component you passed as the `component` prop to the player.

## API [​](\#api "Direct link to API")

Pass in a [parseable](/docs/cli) JSON representation using the `--props` flag to either `remotion studio` or `remotion render`:

```

bash

npx remotion render --props='{"hello": "world"}'
```

```

bash

npx remotion render --props='{"hello": "world"}'
```

To simulate how it behaves, you can also pass props when using the Remotion Studio:

```

bash

npx remotion studio --props='{"hello": "world"}'
```

```

bash

npx remotion studio --props='{"hello": "world"}'
```

You may also specify a file containing JSON and Remotion will parse the file for you:

```

bash

npx remotion render --props=./path/to/props.json
```

```

bash

npx remotion render --props=./path/to/props.json
```

You can then access the props in JavaScript:

```

tsx

const { hello } = getInputProps(); // "world"
```

```

tsx

const { hello } = getInputProps(); // "world"
```

In this example, the props also get passed to the component of the composition with the id `my-composition`.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/config/input-props.ts)
- [Dynamic duration, FPS & dimensions](/docs/dynamic-metadata)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/get-input-props.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Freeze>](/docs/freeze) [Next\
\
getRemotionEnvironment()](/docs/get-remotion-environment)

- [API](#api)
- [See also](#see-also)
