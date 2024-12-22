Embedding videosAdding a videoOn this pageEmbedding videos into RemotionYou can embed existing videos into Remotion by using the <OffthreadVideo> component.
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return (    <OffthreadVideo src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />  );};
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return (    <OffthreadVideo src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />  );};
Using a local file​
Put a file into the public folder and reference it using staticFile().
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return <OffthreadVideo src={staticFile('video.mp4')} />;};
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return <OffthreadVideo src={staticFile('video.mp4')} />;};
Trimming​
By using the startFrom prop, you can remove the first few seconds of the video.
In the example below, the first two seconds of the video are skipped (assuming a composition FPS of 30).
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return <OffthreadVideo src={staticFile('video.mp4')} startFrom={60} />;};
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return <OffthreadVideo src={staticFile('video.mp4')} startFrom={60} />;};
Similarly, you can use endAt to trim the end of the video.
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return (    <OffthreadVideo src={staticFile('video.mp4')} startFrom={60} endAt={120} />  );};
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return (    <OffthreadVideo src={staticFile('video.mp4')} startFrom={60} endAt={120} />  );};
Delaying​
Use the <Sequence> component to delay the appearance of a video.
In the example below, the video will start playing at frame 60.
tsximport React from 'react';import {OffthreadVideo, staticFile, Sequence} from 'remotion'; export const MyComp: React.FC = () => {  return (    <Sequence from={60}>      <OffthreadVideo src={staticFile('video.mp4')} />    </Sequence>  );};
tsximport React from 'react';import {OffthreadVideo, staticFile, Sequence} from 'remotion'; export const MyComp: React.FC = () => {  return (    <Sequence from={60}>      <OffthreadVideo src={staticFile('video.mp4')} />    </Sequence>  );};
Size and Position​
You can size and position the video using CSS.
You'll find the properties width, height, position, left, top, right, bottom, margin, aspectRatio and objectFit useful.
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return (    <OffthreadVideo      src={staticFile('video.mp4')}      style={{        width: 640,        height: 360,        position: 'absolute',        top: 100,        left: 100,      }}    />  );};
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return (    <OffthreadVideo      src={staticFile('video.mp4')}      style={{        width: 640,        height: 360,        position: 'absolute',        top: 100,        left: 100,      }}    />  );};
Volume​
You can set the volume of the video using the volume prop.
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return <OffthreadVideo src={staticFile('video.mp4')} volume={0.5} />;};
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return <OffthreadVideo src={staticFile('video.mp4')} volume={0.5} />;};
You can also mute a video using the muted prop.
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return <OffthreadVideo src={staticFile('video.mp4')} muted />;};
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return <OffthreadVideo src={staticFile('video.mp4')} muted />;};
See Using Audio for more ways you can control volume.
Speed​
You can use the playbackRate prop to change the speed of the video.
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return <OffthreadVideo src={staticFile('video.mp4')} playbackRate={2} />;};
tsximport React from 'react';import {OffthreadVideo, staticFile} from 'remotion'; export const MyComp: React.FC = () => {  return <OffthreadVideo src={staticFile('video.mp4')} playbackRate={2} />;};
This only works if the speed is constant. See also: Changing the speed of a video over time.
Looping​
See: Looping an <OffthreadVideo>
GIFs​
See: Using GIFs
See also​

<OffthreadVideo>
Using Audio
Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousAnimation mathNextMake timeline duration the sameUsing a local fileTrimmingDelayingSize and PositionVolumeSpeedLoopingGIFsSee also