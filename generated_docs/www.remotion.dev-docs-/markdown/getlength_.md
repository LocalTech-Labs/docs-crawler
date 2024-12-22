---
title: getLength()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getLength()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- getLength()

On this page

# getLength()

_Part of the [`@remotion/paths`](/docs/paths) package._

Gets the length of an SVG path. The argument must be a valid SVG path property.

A number is returned if the path is valid:

```

tsx

import { getLength } from "@remotion/paths";

const length = getLength("M 0 0 L 100 0");
console.log(length); // 100
```

```

tsx

import { getLength } from "@remotion/paths";

const length = getLength("M 0 0 L 100 0");
console.log(length); // 100
```

The function will throw if the path is invalid:

```

tsx

getLength("remotion"); // Error: Malformed path data: ...
```

```

tsx

getLength("remotion"); // Error: Malformed path data: ...
```

## Credits [​](\#credits "Direct link to Credits")

Source code stems mostly from [svg-path-properties](https://www.npmjs.com/package/svg-path-properties).

## See also [​](\#see-also "Direct link to See also")

- [getPointAtLength()](/docs/paths/get-point-at-length)
- [getTangentAtLength()](/docs/paths/get-point-at-length)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/get-length.ts)
- [`@remotion/paths`](/docs/paths)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/get-length.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
@remotion/paths](/docs/paths/) [Next\
\
getPointAtLength()](/docs/paths/get-point-at-length)

- [Credits](#credits)
- [See also](#see-also)
