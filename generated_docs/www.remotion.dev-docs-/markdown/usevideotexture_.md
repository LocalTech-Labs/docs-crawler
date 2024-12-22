---
title: useVideoTexture()
source: Remotion Documentation
last_updated: 2024-12-22
---

# useVideoTexture()

- [Home page](/)
- [@remotion/three](/docs/three)
- useVideoTexture()

On this page

# useVideoTexture()

Allows you to use a video in React Three Fiber that is synchronized with Remotion's `useCurrentFrame()`.

To use a video in a Three.JS context, you first have to render it and assign it a ref. If you only want to use it in a React Three Fiber Scene, you can make it invisible by adding a `{position: "absolute", opacity: 0}` style.

```

tsx

import { useRef } from "react";
import { Video } from "remotion";
import src from "./vid.mp4";

const MyVideo = () => {
  const videoRef = useRef<HTMLVideoElement | null>(null);

  return (
    <Video
      ref={videoRef}
      src={src}
      style={{ position: "absolute", opacity: 0 }}
    />
  );
};
```

```

tsx

import { useRef } from "react";
import { Video } from "remotion";
import src from "./vid.mp4";

const MyVideo = () => {
  const videoRef = useRef<HTMLVideoElement | null>(null);

  return (
    <Video
      ref={videoRef}
      src={src}
      style={{ position: "absolute", opacity: 0 }}
    />
  );
};
```

To convert the video to a video texture, place the `useVideoTexture()` hook in the same component.

```

tsx

import { useVideoTexture } from "@remotion/three";

// ...

const texture = useVideoTexture(videoRef);
```

```

tsx

import { useVideoTexture } from "@remotion/three";

// ...

const texture = useVideoTexture(videoRef);
```

The return type of it is a `THREE.VideoTexture | null` which you can assign as a `map` to for example `meshBasicMaterial`. We recommend to only render the material when the texture is not `null` to prevent bugs.

```

tsx

{
  videoTexture ? <meshBasicMaterial map={videoTexture} /> : null;
}
```

```

tsx

{
  videoTexture ? <meshBasicMaterial map={videoTexture} /> : null;
}
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/three/src/use-video-texture.ts)
- [Example usage: React Three Fiber template](/templates/three)
- [`<ThreeCanvas />`](/docs/three-canvas)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/use-video-texture.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<ThreeCanvas>](/docs/three-canvas) [Next\
\
useOffthreadVideoTexture()](/docs/use-offthread-video-texture)

- [See also](#see-also)
