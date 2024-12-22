---
title: Import from After Effects
source: Remotion Documentation
last_updated: 2024-12-22
---

# Import from After Effects

- [Home page](/)
- Tooling
- Import from After Effects

On this page

# Import from After Effects

If you are a After Effects user, you might find it useful to convert your After Effects compositions to Remotion compositions. You can use the [`@remotion/lottie`](/docs/lottie) package for this.

note

Remotion compositions got their name because After Effects coined this term!

### Install the Bodymovin plugin [​](\#install-the-bodymovin-plugin "Direct link to Install the Bodymovin plugin")

- Make sure After Effects is closed.
- Go to [this site](https://aescripts.com/learn/zxp-installer/) and download the ZXP installer for your platform.
- Click [here](https://github.com/airbnb/lottie-web/blob/master/build/extension/bodymovin.zxp?raw=true) to download the latest Bodymovin plugin.
- Open the ZXP installer and drag the bodymovin file into it.

### Create a composition [​](\#create-a-composition "Direct link to Create a composition")

Open After Effects and create a new project and then click `New composition`.

![](https://remotion-assets.s3.eu-central-1.amazonaws.com/new-composition.png)

### Create your animation [​](\#create-your-animation "Direct link to Create your animation")

Design your animation in After Effects. In this basic example, we used the rounded rectangle tool to draw a blue rounded square and then opened the transform menu and clicked the stopwatch icon to set keyframes for position and rotation to create a simple entrance effect.

![](https://remotion-assets.s3.eu-central-1.amazonaws.com/animation.png)

### Allow export as JSON [​](\#allow-export-as-json "Direct link to Allow export as JSON")

In the After Effects menu, go to `Preferences -> Scripting & Expressions...`. Enable the first option: `Allow Scripts to Write Files and Access Network`. You only need to do this once.

![](https://remotion-assets.s3.eu-central-1.amazonaws.com/scripting.png)

### Open the Bodymovin plugin [​](\#open-the-bodymovin-plugin "Direct link to Open the Bodymovin plugin")

In the After Effects menu, go to `Window -> Extensions -> Bodymovin`.

![](https://remotion-assets.s3.eu-central-1.amazonaws.com/bodymovin.png)

### Export the animation as JSON [​](\#export-the-animation-as-json "Direct link to Export the animation as JSON")

First, select the composition

1

. Then press the export icon

2

. You will be prompted for a location to save the JSON file.
Click Render

3

to write the file.

![](https://remotion-assets.s3.eu-central-1.amazonaws.com/bodymovin-tutorial.png)

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
        console.log("Animation failed to load", err);
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
        console.log("Animation failed to load", err);
      });
  }, [handle]);

  if (!animationData) {
    return null;
  }

  return <Lottie animationData={animationData} />;
};
```

## Finetuning [​](\#finetuning "Direct link to Finetuning")

It is advised to make your composition [the same size and duration](/docs/lottie/getlottiemetadata) as the original composition in After Effects. Congrats, you're playing an After Effects animation in Remotion! 🎉

![](https://remotion-assets.s3.eu-central-1.amazonaws.com/result.gif)

## See also [​](\#see-also "Direct link to See also")

- [Using LottieFiles](/docs/lottie/lottiefiles)
- [`getLottieMetadata()`](/docs/lottie/getlottiemetadata)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/after-effects.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Import from Spline](/docs/spline) [Next\
\
Debugging render failures](/docs/troubleshooting/debug-failed-render)

- [Install the Bodymovin plugin](#install-the-bodymovin-plugin)
- [Create a composition](#create-a-composition)
- [Create your animation](#create-your-animation)
- [Allow export as JSON](#allow-export-as-json)
- [Open the Bodymovin plugin](#open-the-bodymovin-plugin)
- [Export the animation as JSON](#export-the-animation-as-json)
- [Import the file into Remotion](#import-the-file-into-remotion)
- [Finetuning](#finetuning)
- [See also](#see-also)
