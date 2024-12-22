Getting startedAnimating propertiesOn this pageAnimating propertiesAnimation works by changing properties over time.
Let's create a simple fade in animation.
If we want to fade the text in over 60 frames, we need to gradually change the opacity over time so that it goes from 0 to 1.
FadeIn.tsxtsxexport const FadeIn = () => {  const frame = useCurrentFrame();   const opacity = Math.min(1, frame / 60);   return (    <AbsoluteFill      style={{        justifyContent: "center",        alignItems: "center",        backgroundColor: "white",        fontSize: 80,      }}    >      <div style={{ opacity: opacity }}>Hello World!</div>    </AbsoluteFill>  );};
FadeIn.tsxtsxexport const FadeIn = () => {  const frame = useCurrentFrame();   const opacity = Math.min(1, frame / 60);   return (    <AbsoluteFill      style={{        justifyContent: "center",        alignItems: "center",        backgroundColor: "white",        fontSize: 80,      }}    >      <div style={{ opacity: opacity }}>Hello World!</div>    </AbsoluteFill>  );};

Using the interpolate helper function​
Using the interpolate() function can make animations more readable. The above animation can also be written as:
tsximport { interpolate } from "remotion"; const opacity = interpolate(frame, [0, 60], [0, 1], {  /*                        ^^^^^   ^^^^^    ^^^^  Variable to interpolate ----|       |       |  Input range ------------------------|       |  Output range -------------------------------|  */  extrapolateRight: "clamp",});
tsximport { interpolate } from "remotion"; const opacity = interpolate(frame, [0, 60], [0, 1], {  /*                        ^^^^^   ^^^^^    ^^^^  Variable to interpolate ----|       |       |  Input range ------------------------|       |  Output range -------------------------------|  */  extrapolateRight: "clamp",});
In this example, we map the frames 0 to 60 to their opacity values (0, 0.0166, 0.033, 0.05 ...) and use the extrapolateRight setting to clamp the output so that it never becomes bigger than 1.
Using spring animations​
Spring animations are a natural animation primitive. By default, they animate from 0 to 1 over time. This time, let's animate the scale of the text.
tsximport { spring, useCurrentFrame, useVideoConfig } from "remotion"; export const MyVideo = () => {  const frame = useCurrentFrame();  const { fps } = useVideoConfig();   const scale = spring({    fps,    frame,  });   return (    <div      style={{        flex: 1,        textAlign: "center",        fontSize: "7em",      }}    >      <div style={{ transform: `scale(${scale})` }}>Hello World!</div>    </div>  );};
tsximport { spring, useCurrentFrame, useVideoConfig } from "remotion"; export const MyVideo = () => {  const frame = useCurrentFrame();  const { fps } = useVideoConfig();   const scale = spring({    fps,    frame,  });   return (    <div      style={{        flex: 1,        textAlign: "center",        fontSize: "7em",      }}    >      <div style={{ transform: `scale(${scale})` }}>Hello World!</div>    </div>  );};
You should see the text jump in.


The default spring configuration leads to a little bit of overshoot, meaning the text will bounce a little bit. See the documentation page for spring() to learn how to customize it.
Always animate using useCurrentFrame()​
Watch out for flickering issues during rendering that arise if you write animations that are not driven by useCurrentFrame() – for example CSS transitions.
Read more about how Remotion's rendering works - understanding it will help you avoid issues down the road.Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousThe fundamentalsNextReuse componentsUsing the interpolate helper functionUsing spring animationsAlways animate using useCurrentFrame()