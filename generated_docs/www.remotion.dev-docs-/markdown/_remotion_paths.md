---
title: @remotion/paths
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/paths

- [Home page](/)
- @remotion/paths

On this page

# @remotion/paths

A package providing utility functions for dealing with SVG paths. This package includes code from [`svg-path-properties`](https://www.npmjs.com/package/svg-path-properties), [`svg-path-reverse`](https://github.com/Pomax/svg-path-reverse#readme), [`svgpath`](https://github.com/fontello/svgpath), [`svg-path-bbox`](https://github.com/mondeja/svg-path-bbox), [`translate-svg-path`](https://github.com/michaelrhodes/translate-svg-path) and [`d3-interpolate-path`](https://github.com/pbeshai/d3-interpolate-path) with the following improvements:

- Functional style APIs
- First class Typescript types
- Documentation with examples
- ESM import style

This package has no dependencies, meaning this package can be used without Remotion.

- npm
- pnpm
- yarn

```

bash

npm i @remotion/paths
```

```

bash

npm i @remotion/paths
```

```

bash

pnpm i @remotion/paths
```

```

bash

pnpm i @remotion/paths
```

```

bash

yarn add @remotion/paths
```

```

bash

yarn add @remotion/paths
```

## Functions [​](\#functions "Direct link to Functions")

[**getLength()** \
\
Obtain length of an SVG path](/docs/paths/get-length) [**getPointAtLength()** \
\
Get coordinates at a certain point of an SVG path](/docs/paths/get-point-at-length) [**getTangentAtLength()** \
\
Gets tangents `x` and `y` of a point which is on an SVG path](/docs/paths/get-tangent-at-length) [**reversePath()** \
\
Switch direction of an SVG path](/docs/paths/reverse-path) [**normalizePath()** \
\
Replace relative with absolute coordinates](/docs/paths/normalize-path) [**interpolatePath()** \
\
Interpolates between two SVG paths](/docs/paths/interpolate-path) [**evolvePath()** \
\
Animate an SVG path](/docs/paths/evolve-path) [**translatePath()** \
\
Translates the position of an path against X/Y coordinates](/docs/paths/translate-path) [**warpPath()** \
\
Remap the coordinates of a path](/docs/paths/warp-path) [**scalePath()** \
\
Grow or shrink the size of the path](/docs/paths/scale-path) [**getBoundingBox()** \
\
Get the bounding box of a SVG path](/docs/paths/get-bounding-box) [**resetPath()** \
\
Translates an SVG path to `(0, 0)`](/docs/paths/reset-path) [**extendViewBox()** \
\
Widen an SVG viewBox in all directions](/docs/paths/extend-viewbox) [**getSubpaths()** \
\
Split SVG path into its parts](/docs/paths/get-subpaths) [**parsePath()** \
\
Parse a string into an array of instructions](/docs/paths/parse-path) [**serializeInstructions()** \
\
Turn an array of instructions into a SVG path](/docs/paths/serialize-instructions) [**reduceInstructions()** \
\
Reduce the amount of instruction types](/docs/paths/reduce-instructions)

## License [​](\#license "Direct link to License")

MIT

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
resolveRedirect()](/docs/preload/resolve-redirect) [Next\
\
getLength()](/docs/paths/get-length)

- [Functions](#functions)
- [License](#license)
