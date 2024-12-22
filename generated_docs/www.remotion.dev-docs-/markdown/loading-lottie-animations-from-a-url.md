---
title: Loading Lottie animations from a URL
source: Remotion Documentation
last_updated: 2024-12-22
---

# Loading Lottie animations from a URL

- [Home page](/)
- [@remotion/lottie](/docs/lottie/)
- Loading from a URL

On this page

# Loading Lottie animations from a URL

In order to load a Lottie animation from a URL that has been put into the `public/` folder, use [`fetch`](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API) and Remotion's [`delayRender()`](/docs/delay-render) function.

The resource must support [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS).

Use the `LottieAnimationData` type to keep a state using React's `useState()` and only render the [`<Lottie>`](/docs/lottie/lottie) component once the data has been fetched.

```

Animation.tsx
tsx

import { Lottie, LottieAnimationData } from "@remotion/lottie";
import { useEffect, useState } from "react";
import { cancelRender, continueRender, delayRender } from "remotion";

const Balloons = () => {
  const [handle] = useState(() => delayRender("Loading Lottie animation"));

  const [animationData, setAnimationData] =
    useState<LottieAnimationData | null>(null);

  useEffect(() => {
    fetch("https://assets4.lottiefiles.com/packages/lf20_zyquagfl.json")
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
import { cancelRender, continueRender, delayRender } from "remotion";

const Balloons = () => {
  const [handle] = useState(() => delayRender("Loading Lottie animation"));

  const [animationData, setAnimationData] =
    useState<LottieAnimationData | null>(null);

  useEffect(() => {
    fetch("https://assets4.lottiefiles.com/packages/lf20_zyquagfl.json")
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
- [Loading from `staticFile()`](/docs/staticfile)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lottie/lottie-remote.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Loading from staticFile()](/docs/lottie/staticfile) [Next\
\
Finding Lottie files](/docs/lottie/lottiefiles)

- [See also](#see-also)
