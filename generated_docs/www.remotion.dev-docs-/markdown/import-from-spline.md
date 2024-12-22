---
title: Import from Spline
source: Remotion Documentation
last_updated: 2024-12-22
---

# Import from Spline

- [Home page](/)
- Tooling
- Import from Spline

On this page

# Import from Spline

_Authors: [Dhroov Makwana](https://github.com/pabloescoder) and [Jonny Burger](https://twitter.com/JNYBGR)_

[Spline](https://app.spline.design/) allows to design 3D models on the web. It allows exporting these models as React Three Fiber code, meaning that can be directly used and animated in Remotion.

This tutorial shows how to design, export and animate a scene with Spline and Remotion.

## Create the 3D model in Spline [​](\#create-the-3d-model-in-spline "Direct link to Create the 3D model in Spline")

In this part I will go over all the steps for creating a 3D model in Spline, if you wish to, you can also skip this part and directly access the built model [here](https://app.spline.design/file/e954db42-8eb2-4130-a5df-d1a1f9fbdc2a).

1. Go to [https://app.spline.design/](https://app.spline.design/), create a new account if necessary.

2. Click on the `New File` button on the top right to create a new project.

![Blank Workspace](/assets/images/1_blank_workspace-326f774a138bc725ea700811b412025b.png)

3. Click the box in the scene, and press `Backspace`, because we don't need it. We will also delete the directional light later, but keep it for now.

4. New shapes can be added using the **+** icon on the toolbar at the top. Click on it, scroll down and select **Torus**.

![Top Menu](data:image/png;base64,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)![Left Menu](/assets/images/3_left_menu_torus-5ea764c4f87ca13556b35e62fb06dc04.png)

5. Once your cursor turns to a **+**, click anywhere on the screen to place the Torus.

![Blank Torus](/assets/images/4_blank_torus-51217df626f705d78df07fcdcfe58d8e.png)

note

The camera rotation can be changed by clicking and dragging while holding the `Alt` (Windows) or `Option` (Mac) key.

7. Center the Torus by right clicking on it, scrolling down and selecting `Reset Position`. Alternatively, you can enter 0 in all three fields of the Position row on the right.

8. Scale the Torus by a factor of two by entering `2` in all three fields of the Scale row.

![Position Options](data:image/png;base64,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)

9. You can play around with the different fields in the **Shape** section to modify the Torus, however we will continue with the default properties.

![Shape Options](/assets/images/6_shape_options-a161b10ce43142eee4be53fd4a06a0fc.png)

10. Scroll down to the **Material** section in the right column, click on the textbox beside the colour picker and enter `3489DC` or any color you like. Effects on the 3D object in Spline are applied as layers, let's add another effect.

11. Click on the **+** icon in the **Material** section, and from the dropdown select **Noise**

![Noise Layer](/assets/images/7_noise_layer-666d03266d35b2c66fa4f7b3eac9b296.png)

12. You'll notice that the colour you previously chose has been replaced by black and white noise, this is because the noise layer is covering all the layers below it. In order to fix this, click on the icon to the right of 100 in the noise row and select **overlay**.

![Overlay](/assets/images/8_overlay-5ad3958552bb2de34af4013afccd7385.png)

13. Finally, let's change the noise type, just for better visuals. Click on the squiggly lines beside noise, select **Simplex Fractal** and set **Movement** to 2.

![Change Noise Type](/assets/images/9_change_noise_type-0d3bb8f5a9aa26f25df9e79c6628df93.png)

Your 3D model should look like this:

![Completed Torus](/assets/images/10_completed_torus-2cdca9f3c25f0f0a2ce1c7a18c1508a1.png)

14. Click on **Directional Light** in the left column, and press `Backspace` key because we don't need it anymore.
15. Click **Export** in the top menu. Click on the first dropdown and select **Code (Experimental)**. Select **react-three-fiber** in the second dropdown. Finally, click on the Copy icon in the code section to copy the generated code.

![Export Model Code](/assets/images/11_export_model_code-ca272998643b6c912a95e78252921501.png)

## Animating 3D Model using Remotion [​](\#animating-3d-model-using-remotion "Direct link to Animating 3D Model using Remotion")

1. Create a new Remotion project from the React Three Fiber template by running `npx create-video@latest`.

![Project Starter](/assets/images/12_project_starter-8ccfa5f4414adaa11a70e5f949e9ee7c.png)

2. Install the [`@splinetools/r3f-spline`](https://github.com/splinetool/r3f-spline) package:



```

shell

npm install @splinetool/r3f-spline
```


```

shell

npm install @splinetool/r3f-spline
```

3. Create a new file at `src/Torus.tsx` with the code that you have copied from Spline.

4. Remove the `OrthographicCamera` import, and the `<OrthographicCamera>` component as we will be using the `useThree` hook for the camera.

5. In `src/Scene.tsx`, replace the default scene with:

```

src/Root.tsx
tsx

import {ThreeCanvas} from '@remotion/three';
import {useVideoConfig} from 'remotion';
import Torus from './Torus';

export const Scene: React.FC = () => {
  const {width, height} = useVideoConfig();

  return (
    <ThreeCanvas width={width} height={height}>
      <Torus />
    </ThreeCanvas>
  );
};
```

```

src/Root.tsx
tsx

import {ThreeCanvas} from '@remotion/three';
import {useVideoConfig} from 'remotion';
import Torus from './Torus';

export const Scene: React.FC = () => {
  const {width, height} = useVideoConfig();

  return (
    <ThreeCanvas width={width} height={height}>
      <Torus />
    </ThreeCanvas>
  );
};
```

7. Add the following imports to `src/Torus.tsx`:

```

src/Torus.tsx
ts

import React, {useEffect} from 'react';
import {useThree} from '@react-three/fiber';
import {interpolate, spring, useCurrentFrame, useVideoConfig} from 'remotion';
```

```

src/Torus.tsx
ts

import React, {useEffect} from 'react';
import {useThree} from '@react-three/fiber';
import {interpolate, spring, useCurrentFrame, useVideoConfig} from 'remotion';
```

8. Add the following code inside the Torus function below the `useSpline()` call:

```

src/Torus.tsx
tsx

const frame = useCurrentFrame();
const {fps, durationInFrames} = useVideoConfig();
```

```

src/Torus.tsx
tsx

const frame = useCurrentFrame();
const {fps, durationInFrames} = useVideoConfig();
```

Position the camera to look at the center point from `200` z offset.

```

src/Torus.tsx
tsx

const camera = useThree((state) => state.camera);

useEffect(() => {
  camera.position.set(0, 0, -400);
  camera.near = 0.2;
  camera.far = 1000;
  camera.lookAt(0, 0, 0);
}, [camera]);
```

```

src/Torus.tsx
tsx

const camera = useThree((state) => state.camera);

useEffect(() => {
  camera.position.set(0, 0, -400);
  camera.near = 0.2;
  camera.far = 1000;
  camera.lookAt(0, 0, 0);
}, [camera]);
```

9. Let's add some animation. Below the above code, insert the following:

```

src/Torus.tsx
tsx

const constantRotation = interpolate(
  frame,
  [0, durationInFrames],
  [0, Math.PI * 6],
);

const entranceAnimation = spring({
  frame,
  fps,
  config: {
    damping: 200,
  },
});
```

```

src/Torus.tsx
tsx

const constantRotation = interpolate(
  frame,
  [0, durationInFrames],
  [0, Math.PI * 6],
);

const entranceAnimation = spring({
  frame,
  fps,
  config: {
    damping: 200,
  },
});
```

- `constantRotation` will cause the Torus to rotate forever. Math.PI \* 2 equals a complete 360 degree rotation.
- `entranceAnimation` animates from 0 to 1 over time using a spring animation.

10. Find the `<mesh>` element and replace it's `scale` and `rotation` parameters with:

```

tsx

scale={entranceAnimation + 3}
rotation={[constantRotation / 2, constantRotation, 0]}
```

```

tsx

scale={entranceAnimation + 3}
rotation={[constantRotation / 2, constantRotation, 0]}
```

note

The X rotation is half of the Y rotation, so for every 2 turns in the Y axis, our Torus will turn once in the X axis.

11. That's it! See a rotating Torus in the Remotion Studio.

12. If you want to export your video to a MP4 video file, run:

```

shell

npx remotion render
```

```

shell

npx remotion render
```

## Final Result / Generated Video [​](\#final-result--generated-video "Direct link to Final Result / Generated Video")

## Credits [​](\#credits "Direct link to Credits")

CONTRIBUTORS

[![pabloescoder](https://github.com/pabloescoder.png)\
\
**pabloescoder** \
\
Author](https://github.com/pabloescoder)

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/three`](/docs/three)
- [Remotion Three Template](https://github.com/remotion-dev/template-three)
- [Spline](https://docs.spline.design/)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/spline.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Import from Figma](/docs/figma) [Next\
\
Import from After Effects](/docs/after-effects)

- [Create the 3D model in Spline](#create-the-3d-model-in-spline)
- [Animating 3D Model using Remotion](#animating-3d-model-using-remotion)
- [Final Result / Generated Video](#final-result--generated-video)
- [Credits](#credits)
- [See also](#see-also)
