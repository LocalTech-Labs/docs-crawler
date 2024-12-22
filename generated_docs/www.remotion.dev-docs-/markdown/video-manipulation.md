---
title: Video manipulation
source: Remotion Documentation
last_updated: 2024-12-22
---

# Video manipulation

- [Home page](/)
- Embedding videos
- Manipulating pixels

On this page

# Video manipulation

You can draw frames of an [`<OffthreadVideo>`](/docs/offthreadvideo) or a [`<Video>`](/docs/video) onto a `<canvas>` element using the [`drawImage()`](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/drawImage) API.

note

During preview, makes use of the [`requestVideoFrameCallback()`](https://developer.mozilla.org/en-US/docs/Web/API/HTMLVideoElement/requestVideoFrameCallback) API.

Browser support: Firefox 130 (August 2024), Chrome 83, Safari 15.4.

## Basic example [​](\#basic-example "Direct link to Basic example")

In this example, an [`<OffthreadVideo>`](/docs/offthreadvideo) is rendered and made invisible.

Every frame that is emitted is drawn to a Canvas and a grayscale [`filter`](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/filter) is applied.

0:00 / 0:05

```

tsx

export const VideoOnCanvas: React.FC = () => {
  const video = useRef<HTMLVideoElement>(null);
  const canvas = useRef<HTMLCanvasElement>(null);
  const {width, height} = useVideoConfig();

  // Process a frame
  const onVideoFrame = useCallback(
    (frame: CanvasImageSource) => {
      if (!canvas.current) {
        return;
      }
      const context = canvas.current.getContext('2d');

      if (!context) {
        return;
      }

      context.filter = 'grayscale(100%)';
      context.drawImage(frame, 0, 0, width, height);
    },
    [height, width],
  );

  return (
    <AbsoluteFill>
      <AbsoluteFill>
        <OffthreadVideo
          // Hide the original video tag
          style={{opacity: 0}}
          onVideoFrame={onVideoFrame}
          src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"
        />
      </AbsoluteFill>
      <AbsoluteFill>
        <canvas ref={canvas} width={width} height={height} />
      </AbsoluteFill>
    </AbsoluteFill>
  );
};
```

```

tsx

export const VideoOnCanvas: React.FC = () => {
  const video = useRef<HTMLVideoElement>(null);
  const canvas = useRef<HTMLCanvasElement>(null);
  const {width, height} = useVideoConfig();

  // Process a frame
  const onVideoFrame = useCallback(
    (frame: CanvasImageSource) => {
      if (!canvas.current) {
        return;
      }
      const context = canvas.current.getContext('2d');

      if (!context) {
        return;
      }

      context.filter = 'grayscale(100%)';
      context.drawImage(frame, 0, 0, width, height);
    },
    [height, width],
  );

  return (
    <AbsoluteFill>
      <AbsoluteFill>
        <OffthreadVideo
          // Hide the original video tag
          style={{opacity: 0}}
          onVideoFrame={onVideoFrame}
          src="http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"
        />
      </AbsoluteFill>
      <AbsoluteFill>
        <canvas ref={canvas} width={width} height={height} />
      </AbsoluteFill>
    </AbsoluteFill>
  );
};
```

## Greenscreen example [​](\#greenscreen-example "Direct link to Greenscreen example")

In this example, we loop over each pixel in the image buffer and if it's green, we transparentize it. Drag the slider below to turn the video transparent.

0:00 / 0:05

Slide to adjust transparency:

```

tsx

export const Greenscreen: React.FC<{
  opacity: number;
}> = ({opacity}) => {
  const canvas = useRef<HTMLCanvasElement>(null);
  const {width, height} = useVideoConfig();

  // Process a frame
  const onVideoFrame = useCallback(
    (frame: CanvasImageSource) => {
      if (!canvas.current) {
        return;
      }
      const context = canvas.current.getContext('2d');

      if (!context) {
        return;
      }

      context.drawImage(frame, 0, 0, width, height);
      const imageFrame = context.getImageData(0, 0, width, height);
      const {length} = imageFrame.data;

      // If the pixel is very green, reduce the alpha channel
      for (let i = 0; i < length; i += 4) {
        const red = imageFrame.data[i + 0];
        const green = imageFrame.data[i + 1];
        const blue = imageFrame.data[i + 2];
        if (green > 100 && red < 100 && blue < 100) {
          imageFrame.data[i + 3] = opacity * 255;
        }
      }
      context.putImageData(imageFrame, 0, 0);
    },
    [height, width],
  );

  return (
    <AbsoluteFill>
      <AbsoluteFill>
        <OffthreadVideo
          style={{opacity: 0}}
          onVideoFrame={onVideoFrame}
          src="https://remotion-assets.s3.eu-central-1.amazonaws.com/just-do-it-short.mp4"
        />
      </AbsoluteFill>
      <AbsoluteFill>
        <canvas ref={canvas} width={width} height={height} />
      </AbsoluteFill>
    </AbsoluteFill>
  );
};
```

```

tsx

export const Greenscreen: React.FC<{
  opacity: number;
}> = ({opacity}) => {
  const canvas = useRef<HTMLCanvasElement>(null);
  const {width, height} = useVideoConfig();

  // Process a frame
  const onVideoFrame = useCallback(
    (frame: CanvasImageSource) => {
      if (!canvas.current) {
        return;
      }
      const context = canvas.current.getContext('2d');

      if (!context) {
        return;
      }

      context.drawImage(frame, 0, 0, width, height);
      const imageFrame = context.getImageData(0, 0, width, height);
      const {length} = imageFrame.data;

      // If the pixel is very green, reduce the alpha channel
      for (let i = 0; i < length; i += 4) {
        const red = imageFrame.data[i + 0];
        const green = imageFrame.data[i + 1];
        const blue = imageFrame.data[i + 2];
        if (green > 100 && red < 100 && blue < 100) {
          imageFrame.data[i + 3] = opacity * 255;
        }
      }
      context.putImageData(imageFrame, 0, 0);
    },
    [height, width],
  );

  return (
    <AbsoluteFill>
      <AbsoluteFill>
        <OffthreadVideo
          style={{opacity: 0}}
          onVideoFrame={onVideoFrame}
          src="https://remotion-assets.s3.eu-central-1.amazonaws.com/just-do-it-short.mp4"
        />
      </AbsoluteFill>
      <AbsoluteFill>
        <canvas ref={canvas} width={width} height={height} />
      </AbsoluteFill>
    </AbsoluteFill>
  );
};
```

## Before v4.0.190 [​](\#before-v40190 "Direct link to Before v4.0.190")

Before v4.0.190, the `onVideoFrame` prop of [`<OffthreadVideo>`](/docs/offthreadvideo) and [`<Video>`](/docs/video) was not supported.

You could only manipulate a `<Video>` using the `requestVideoFrameCallback` API.

Click [here](https://github.com/remotion-dev/remotion/blob/b966d0e99cfb91478ca697675f822284da1f9055/packages/docs/docs/video-manipulation.mdx) to see the old version of this page.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/video-manipulation.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Adding a transparent video](/docs/videos/transparency) [Next\
\
Changing speed over time](/docs/miscellaneous/snippets/accelerated-video)

- [Basic example](#basic-example)
- [Greenscreen example](#greenscreen-example)
- [Before v4.0.190](#before-v40190)
