---
title: preloadGif()
source: Remotion Documentation
last_updated: 2024-12-22
---

# preloadGif()

- [Home page](/)
- [@remotion/gif](/docs/gif/)
- preloadGif()

On this page

# preloadGif()

_available from v3.3.38_

Call `preloadGif(src)` with the URL of the GIF that you would like to load and the GIF will be prepared for display in the [`<Player>`](/docs/player).

The function returns an object with two entries: `waitUntilDone()` that returns a Promise which can be awaited and `free()` which will cancel preloading or free up the memory if the GIF is not being used anymore.

```

ts

import { preloadGif } from "@remotion/gif";

const { waitUntilDone, free } = preloadGif(
  "https://media.giphy.com/media/xT0GqH01ZyKwd3aT3G/giphy.gif"
);

waitUntilDone().then(() => {
  console.log("The GIF is now ready to play!");

  // Later, free the memory of the GIF
  free();
});
```

```

ts

import { preloadGif } from "@remotion/gif";

const { waitUntilDone, free } = preloadGif(
  "https://media.giphy.com/media/xT0GqH01ZyKwd3aT3G/giphy.gif"
);

waitUntilDone().then(() => {
  console.log("The GIF is now ready to play!");

  // Later, free the memory of the GIF
  free();
});
```

## See also [​](\#see-also "Direct link to See also")

- [`<Gif>`](/docs/gif)
- [Preloading](/docs/preload)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/gif/preload-gif.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getGifDurationInSeconds()](/docs/gif/get-gif-duration-in-seconds) [Next\
\
@remotion/media-utils](/docs/media-utils/)

- [See also](#see-also)
