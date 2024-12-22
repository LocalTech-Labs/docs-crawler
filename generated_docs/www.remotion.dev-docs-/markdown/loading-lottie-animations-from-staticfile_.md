---
title: Loading Lottie animations from staticFile()
source: Remotion Documentation
last_updated: 2024-12-22
---

# Loading Lottie animations from staticFile()

- [Home page](/)
- [@remotion/lottie](/docs/lottie/)
- Loading from staticFile()

On this page

# Loading Lottie animations from staticFile()

In order to load a Lottie animation from a file that has been put into the `public/` folder, use [`staticFile()`](/docs/staticfile) in combination with [`fetch`](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API) and Remotion's [`delayRender()`](/docs/delay-render) function.

Use the `LottieAnimationData` type to keep a state using React's `useState()` and only render the [`<Lottie>`](/docs/lottie/lottie) component once the data has been fetched.

```

Animation.tsx
tsx

import { Lottie, LottieAnimationData } from "@remotion/lottie";
import { useEffect, useState } from "react";
import {
  cancelRender,
  continueRender,
  delayRender,
  staticFile,
} from "remotion";

const Square = () => {
  const [handle] = useState(() => delayRender("Loading Lottie animation"));

  const [animationData, setAnimationData] =
    useState<LottieAnimationData | null>(null);

  useEffect(() => {
    fetch(staticFile("data.json"))
      .then((data) => data.json())
      .then((json) => {
        setAnimationData(json);
        continueRender(handle);
      })
      .catch((err) => {
        cancelRender(err);
      });
  }, [handle]);

  if (!animationData) {
    return null;
  }

  return <Lottie animationData={animationData} />;
};
```

```

Animation.tsx
tsx

import { Lottie, LottieAnimationData } from "@remotion/lottie";
import { useEffect, useState } from "react";
import {
  cancelRender,
  continueRender,
  delayRender,
  staticFile,
} from "remotion";

const Square = () => {
  const [handle] = useState(() => delayRender("Loading Lottie animation"));

  const [animationData, setAnimationData] =
    useState<LottieAnimationData | null>(null);

  useEffect(() => {
    fetch(staticFile("data.json"))
      .then((data) => data.json())
      .then((json) => {
        setAnimationData(json);
        continueRender(handle);
      })
      .catch((err) => {
        cancelRender(err);
      });
  }, [handle]);

  if (!animationData) {
    return null;
  }

  return <Lottie animationData={animationData} />;
};
```

## See also [​](\#see-also "Direct link to See also")

- [`<Lottie>`](/docs/lottie/lottie)
- [Loading from a URL](/docs/lottie/remote)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lottie/lottie-staticfile.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getLottieMetadata()](/docs/lottie/getlottiemetadata) [Next\
\
Loading from a URL](/docs/lottie/remote)

- [See also](#see-also)
