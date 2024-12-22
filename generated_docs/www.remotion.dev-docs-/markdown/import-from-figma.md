---
title: Import from Figma
source: Remotion Documentation
last_updated: 2024-12-22
---

# Import from Figma

- [Home page](/)
- Tooling
- Import from Figma

On this page

# Import from Figma

You can export a design from Figma as an SVG file and import it as a React component in Remotion and then animate it.

## Open the Figma design [​](\#open-the-figma-design "Direct link to Open the Figma design")

We are going to use a copy of [Streamline's Vector illustrations](https://www.figma.com/community/file/1118919399684035468).

![Thumbnail](/assets/images/banner-ff73ec91ab2352435ba2488a6df3446c.png)

## Grouping vector elements in Figma [​](\#grouping-vector-elements-in-figma "Direct link to Grouping vector elements in Figma")

If not already done, group the items you want to export together or frame them in Figma.

![figma-grouping](/assets/images/figma-grouping-5844456151c6a2d9cea885af173c5e92.gif)

Groups will be represented as `<g>` elements in SVG and if you want to animate multiple items together, it can be useful to group them.

## Export as SVG [​](\#export-as-svg "Direct link to Export as SVG")

You can export any design by copying it as SVG markup — you can do that by right-clicking on the design itself and selecting the **Copy/Paste as** option.

![exporting as SVG](/assets/images/copy-as-svg-d88d15061af200010bdf8cacab284636.png)

Next, you need to convert the SVG into a React component. Often, you can just paste the SVG into React markup and it will work.

Alternatively, use the [SVGR playground](https://react-svgr.com/playground/) to reliably convert SVG into React components.

## Importing SVG in Remotion [​](\#importing-svg-in-remotion "Direct link to Importing SVG in Remotion")

Paste the component into a Remotion project and [register a `<Composition>`](/docs/composition).

An example can be found in this [repository](https://github.com/kaf-lamed-beyt/remo-sample).

![](/assets/images/editor-pink-e7c4673d031d133d5a0558a6eac5ce24.png)

## Animate SVG markup [​](\#animate-svg-markup "Direct link to Animate SVG markup")

Locate the element that you want to animate and add a style property to it.
In this case, let's animate the `<g>` element that contains the rocket.

```

tsx

import { AbsoluteFill, useCurrentFrame, useVideoConfig } from "remotion";

export const Rocket: React.FC = () => {
  const frame = useCurrentFrame();
  const { fps } = useVideoConfig();

  return (
    <AbsoluteFill
      style={{
        backgroundColor: "pink",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <svg
        width="800"
        height="800"
        viewBox="0 0 394 394"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <g
          id="vehicle"
          style={{
            transformOrigin: "center center",
            transformBox: "fill-box",
          }}
        >
          // vehicle's paths
        </g>
      </svg>
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, useCurrentFrame, useVideoConfig } from "remotion";

export const Rocket: React.FC = () => {
  const frame = useCurrentFrame();
  const { fps } = useVideoConfig();

  return (
    <AbsoluteFill
      style={{
        backgroundColor: "pink",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <svg
        width="800"
        height="800"
        viewBox="0 0 394 394"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <g
          id="vehicle"
          style={{
            transformOrigin: "center center",
            transformBox: "fill-box",
          }}
        >
          // vehicle's paths
        </g>
      </svg>
    </AbsoluteFill>
  );
};
```

Adding `{transformOrigin: "center center", transformBox: "fill-box"}` will ensure that the transformations center is it's own center.

Let's create two spring animations, one for scale and one for transformation:

```

tsx

const up = spring({
  fps,
  frame: frame - 20,
  config: {
    damping: 20,
    mass: 15,
  },
});

const scale = spring({
  fps,
  frame,
  config: {
    stiffness: 200,
  },
});

const launch = `translateY(${interpolate(up, [0, 1], [0, -3000])}px)`;
```

```

tsx

const up = spring({
  fps,
  frame: frame - 20,
  config: {
    damping: 20,
    mass: 15,
  },
});

const scale = spring({
  fps,
  frame,
  config: {
    stiffness: 200,
  },
});

const launch = `translateY(${interpolate(up, [0, 1], [0, -3000])}px)`;
```

The `scale` will go from 0 to 1 and `launch` animates from `0` to `-3000px`. Apply the styles to the element:

```

tsx

<g
  id="vehicle"
  style={{
    transform: `scale(${scale}) ${launch}`,
    transformOrigin: "center center",
    transformBox: "fill-box",
  }}
>
  {/* ... */}
</g>
```

```

tsx

<g
  id="vehicle"
  style={{
    transform: `scale(${scale}) ${launch}`,
    transformOrigin: "center center",
    transformBox: "fill-box",
  }}
>
  {/* ... */}
</g>
```

![rocket svg](https://res.cloudinary.com/meje/image/upload/v1665432945/article%20assets/rocket_clhn8w.gif)

You have animated a rocket! 🚀

## See also [​](\#see-also "Direct link to See also")

- [Blog - SVG Animation with Remotion](https://meje.dev/blog/svg-animtion-with-remotion)
- [Video - Animating Figma Mockups with Remotion](https://twitter.com/jnybgr/status/1496748768821133312)
- [Locofy - Export Figma mockups as React components](https://www.locofy.ai)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/figma.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
TypeScript aliases](/docs/typescript-aliases) [Next\
\
Import from Spline](/docs/spline)

- [Open the Figma design](#open-the-figma-design)
- [Grouping vector elements in Figma](#grouping-vector-elements-in-figma)
- [Export as SVG](#export-as-svg)
- [Importing SVG in Remotion](#importing-svg-in-remotion)
- [Animate SVG markup](#animate-svg-markup)
- [See also](#see-also)
