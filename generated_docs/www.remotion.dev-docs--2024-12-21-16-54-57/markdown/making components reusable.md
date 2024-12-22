Getting startedReuse componentsOn this pageMaking components reusable
React components allow us to encapsulate video logic and reuse the same visuals multiple times.
Consider a title - to make it reusable, factor it out into its own component:
MyComposition.tsxtsximport {AbsoluteFill, interpolate, useCurrentFrame} from 'remotion'; const Title: React.FC<{title: string}> = ({title}) => {  const frame = useCurrentFrame();  const opacity = interpolate(frame, [0, 20], [0, 1], {    extrapolateRight: 'clamp',  });   return (    <div style={{opacity, textAlign: 'center', fontSize: '7em'}}>{title}</div>  );}; export const MyVideo = () => {  return (    <AbsoluteFill>      <Title title="Hello World" />    </AbsoluteFill>  );};
MyComposition.tsxtsximport {AbsoluteFill, interpolate, useCurrentFrame} from 'remotion'; const Title: React.FC<{title: string}> = ({title}) => {  const frame = useCurrentFrame();  const opacity = interpolate(frame, [0, 20], [0, 1], {    extrapolateRight: 'clamp',  });   return (    <div style={{opacity, textAlign: 'center', fontSize: '7em'}}>{title}</div>  );}; export const MyVideo = () => {  return (    <AbsoluteFill>      <Title title="Hello World" />    </AbsoluteFill>  );};
To render multiple instances of the title, duplicate the <Title> component.
You can also use the <Sequence> component to limit the duration of the first title and time-shift the appearance of the second title.
tsximport {Sequence} from 'remotion'; export const MyVideo = () => {  return (    <AbsoluteFill>      <Sequence durationInFrames={40}>        <Title title="Hello" />      </Sequence>      <Sequence from={40}>        <Title title="World" />      </Sequence>    </AbsoluteFill>  );};
tsximport {Sequence} from 'remotion'; export const MyVideo = () => {  return (    <AbsoluteFill>      <Sequence durationInFrames={40}>        <Title title="Hello" />      </Sequence>      <Sequence from={40}>        <Title title="World" />      </Sequence>    </AbsoluteFill>  );};
You should see two titles appearing after each other.
Note that the value of useCurrentFrame() has been shifted in the second instance, so that it returns 0 only when the absolute time is 40. Before that, the sequence was not mounted at all.
Sequences by default are absolutely positioned - you can use layout="none" to make them render like a regular <div>.
See also​

<Sequence> component
How do I combine compositions?
Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousAnimating propertiesNextPreview your videoSee also