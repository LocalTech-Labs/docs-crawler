---
title: Finding Lottie files to use
source: Remotion Documentation
last_updated: 2024-12-22
---

# Finding Lottie files to use

- [Home page](/)
- [@remotion/lottie](/docs/lottie/)
- Finding Lottie files

On this page

# Finding Lottie files to use

[LottieFiles](https://lottiefiles.com) is a website where people can share and download Lottie files.

![](https://remotion-assets.s3.eu-central-1.amazonaws.com/lottiefiles.png)

If you find a file that you like, click on it, then click `Download`

1

and choose `Lottie JSON`

2

as the format.

![](https://remotion-assets.s3.eu-central-1.amazonaws.com/lottiefiles-instructions.png)

### Import the file into Remotion [​](\#import-the-file-into-remotion "Direct link to Import the file into Remotion")

Copy the file into the Remotion project. The recommended way is to put the JSON inside the `public/` folder of Remotion (create it if necessary) and then load it using [`staticFile()`](/docs/staticfile):

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

const Balloons = () => {
  const [handle] = useState(() => delayRender("Loading Lottie animation"));

  const [animationData, setAnimationData] =
    useState<LottieAnimationData | null>(null);

  useEffect(() => {
    fetch(staticFile("animation.json"))
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

const Balloons = () => {
  const [handle] = useState(() => delayRender("Loading Lottie animation"));

  const [animationData, setAnimationData] =
    useState<LottieAnimationData | null>(null);

  useEffect(() => {
    fetch(staticFile("animation.json"))
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

- [Import from After Effects](/docs/after-effects)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lottie/lottie-lottiefiles.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Loading from a URL](/docs/lottie/remote) [Next\
\
Overview](/docs/preload/)

- [Import the file into Remotion](#import-the-file-into-remotion)
- [See also](#see-also)
