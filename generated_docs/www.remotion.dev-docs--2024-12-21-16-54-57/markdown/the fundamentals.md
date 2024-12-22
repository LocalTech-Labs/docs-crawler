Getting startedThe fundamentalsOn this pageThe fundamentalsReact components​

The idea of Remotion is to give you a frame number and a blank canvas, to which you can render anything you want using React. This is an example React component that renders the current frame as text:
MyComposition.tsxtsximport { AbsoluteFill, useCurrentFrame } from "remotion"; export const MyComposition = () => {  const frame = useCurrentFrame();   return (    <AbsoluteFill      style={{        justifyContent: "center",        alignItems: "center",        fontSize: 100,        backgroundColor: "white",      }}    >      The current frame is {frame}.    </AbsoluteFill>  );};
MyComposition.tsxtsximport { AbsoluteFill, useCurrentFrame } from "remotion"; export const MyComposition = () => {  const frame = useCurrentFrame();   return (    <AbsoluteFill      style={{        justifyContent: "center",        alignItems: "center",        fontSize: 100,        backgroundColor: "white",      }}    >      The current frame is {frame}.    </AbsoluteFill>  );};
A video is a function of images over time. If you change content every frame, you'll end up with an animation.
Video properties​
A video has 4 properties:

width in pixels.
height in pixels.
durationInFrames: the number of frames which the video is long.
fps: framerate of the video.

You can obtain these values from the useVideoConfig() hook:
tsximport {AbsoluteFill, useVideoConfig} from 'remotion'; export const MyComposition = () => {  const {fps, durationInFrames, width, height} = useVideoConfig();   return (    <AbsoluteFill      style={{        justifyContent: 'center',        alignItems: 'center',        fontSize: 60,        backgroundColor: 'white',      }}    >      This {width}x{height}px video is {durationInFrames / fps} seconds long.    </AbsoluteFill>  );};
tsximport {AbsoluteFill, useVideoConfig} from 'remotion'; export const MyComposition = () => {  const {fps, durationInFrames, width, height} = useVideoConfig();   return (    <AbsoluteFill      style={{        justifyContent: 'center',        alignItems: 'center',        fontSize: 60,        backgroundColor: 'white',      }}    >      This {width}x{height}px video is {durationInFrames / fps} seconds long.    </AbsoluteFill>  );};
noteA video's first frame is 0 and its last frame is durationInFrames - 1.
Compositions​
A composition is the combination of a React component and video metadata.
By rendering a <Composition> component in src/Root.tsx, you can register a renderable video and make it visible in the Remotion sidebar.
src/Root.tsxtsxexport const RemotionRoot: React.FC = () => {  return (    <Composition      id="MyComposition"      durationInFrames={150}      fps={30}      width={1920}      height={1080}      component={MyComposition}    />  );};
src/Root.tsxtsxexport const RemotionRoot: React.FC = () => {  return (    <Composition      id="MyComposition"      durationInFrames={150}      fps={30}      width={1920}      height={1080}      component={MyComposition}    />  );};
noteYou can register multiple compositions in src/Root.tsx by wrapping them in a React Fragment:
<><Composition/><Composition/></>. See also: How to combine compositions?Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousInstallationNextAnimating propertiesReact componentsVideo propertiesCompositions