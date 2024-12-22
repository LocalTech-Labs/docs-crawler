---
title: reversePath()
source: Remotion Documentation
last_updated: 2024-12-22
---

# reversePath()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- reversePath()

On this page

# reversePath()

_Part of the [`@remotion/paths`](/docs/paths) package._

Reverses a path so the end and start are switched.

```

tsx

import { reversePath } from "@remotion/paths";

const reversedPath = reversePath("M 0 0 L 100 0");
console.log(reversedPath); // "L 100 0 M 0 0"
```

```

tsx

import { reversePath } from "@remotion/paths";

const reversedPath = reversePath("M 0 0 L 100 0");
console.log(reversedPath); // "L 100 0 M 0 0"
```

The function will throw if the path is invalid:

```

tsx

reversePath("remotion"); // Error: Malformed path data: ...
```

```

tsx

reversePath("remotion"); // Error: Malformed path data: ...
```

## Credits [​](\#credits "Direct link to Credits")

Source code stems mostly from [svg-path-reverse](https://www.npmjs.com/package/svg-path-reverse).

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/reverse-path.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/reverse-path.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getInstructionIndexAtLength()](/docs/paths/get-instruction-index-at-length) [Next\
\
normalizePath()](/docs/paths/normalize-path)

- [Credits](#credits)
- [See also](#see-also)
