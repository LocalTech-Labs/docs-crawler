---
title: How much does Remotion Lambda cost?
source: Remotion Documentation
last_updated: 2024-12-22
---

# How much does Remotion Lambda cost?

- [Home page](/)
- [Lambda](/docs/lambda)
- Cost example

On this page

# How much does Remotion Lambda cost?

This page shows estimations to help you better expect how much Remotion Lambda will cost.

All renders are done with the following Lambda configuration:

- 2048 MB RAM (default)
- 10GB Disk size (default from Remotion 5.0)
- Default [concurrency](/docs/lambda/concurrency)
- Warm Lambda functions (expect the first render to be more expensive)
- `us-east-1` region

Prices may also vary based on the region, heavyness of your bundle, fluctuations in latency.

We always recommend to measure the cost of your composition yourself.

## Rendering the Hello World project [​](\#rendering-the-hello-world-project "Direct link to Rendering the Hello World project")

Rendering the [`HelloWorld`](https://github.com/remotion-dev/template-helloworld) composition of our default template.

**Cost**: $0.001

## Rendering a 1 minute Video in the same S3 bucket [​](\#rendering-a-1-minute-video-in-the-same-s3-bucket "Direct link to Rendering a 1 minute Video in the same S3 bucket")

```

OffthreadRemoteVideo.tsx
tsx

import {CalculateMetadataFunction, OffthreadVideo, staticFile} from 'remotion';

export const calculateMetadataFn: CalculateMetadataFunction<{}> = async () => {
  return {
    durationInFrames: 60 * 30,
    fps: 30,
    width: 1280,
    height: 720,
  };
};

export const OffthreadRemoteVideo: React.FC = () => {
  return <OffthreadVideo src={staticFile('bigbuckbunny.mp4')} />;
};
```

```

OffthreadRemoteVideo.tsx
tsx

import {CalculateMetadataFunction, OffthreadVideo, staticFile} from 'remotion';

export const calculateMetadataFn: CalculateMetadataFunction<{}> = async () => {
  return {
    durationInFrames: 60 * 30,
    fps: 30,
    width: 1280,
    height: 720,
  };
};

export const OffthreadRemoteVideo: React.FC = () => {
  return <OffthreadVideo src={staticFile('bigbuckbunny.mp4')} />;
};
```

[This](https://github.com/remotion-dev/remotion/blob/main/packages/example/public/bigbuckbunny.mp4) is the video file used.

**Cost**: $0.021

Rendering a video inside your video increases the cost significantly.

## Rendering a 10 minute Remote HD Video [​](\#rendering-a-10-minute-remote-hd-video "Direct link to Rendering a 10 minute Remote HD Video")

```

OffthreadRemoteVideo.tsx
tsx

import {CalculateMetadataFunction, OffthreadVideo} from 'remotion';

export const calculateMetadataFn: CalculateMetadataFunction<{}> = async () => {
  return {
    durationInFrames: 10 * 60 * 30,
    fps: 30,
    width: 1280,
    height: 720,
  };
};

export const OffthreadRemoteVideo: React.FC = () => {
  return (
    <OffthreadVideo src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />
  );
};
```

```

OffthreadRemoteVideo.tsx
tsx

import {CalculateMetadataFunction, OffthreadVideo} from 'remotion';

export const calculateMetadataFn: CalculateMetadataFunction<{}> = async () => {
  return {
    durationInFrames: 10 * 60 * 30,
    fps: 30,
    width: 1280,
    height: 720,
  };
};

export const OffthreadRemoteVideo: React.FC = () => {
  return (
    <OffthreadVideo src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />
  );
};
```

**Cost**: $0.162

## Rendering a 10 second Remote 4K Video [​](\#rendering-a-10-second-remote-4k-video "Direct link to Rendering a 10 second Remote 4K Video")

```

OffthreadRemoteVideo.tsx
tsx

import {CalculateMetadataFunction, OffthreadVideo} from 'remotion';

export const calculateMetadataFn: CalculateMetadataFunction<{}> = async () => {
  return {
    durationInFrames: 10 * 30,
    fps: 30,
    width: 3840,
    height: 2160,
  };
};

export const OffthreadRemoteVideo: React.FC = () => {
  return (
    <OffthreadVideo src="https://videos.pexels.com/video-files/5530402/5530402-uhd_3840_2160_25fps.mp4" />
  );
};
```

```

OffthreadRemoteVideo.tsx
tsx

import {CalculateMetadataFunction, OffthreadVideo} from 'remotion';

export const calculateMetadataFn: CalculateMetadataFunction<{}> = async () => {
  return {
    durationInFrames: 10 * 30,
    fps: 30,
    width: 3840,
    height: 2160,
  };
};

export const OffthreadRemoteVideo: React.FC = () => {
  return (
    <OffthreadVideo src="https://videos.pexels.com/video-files/5530402/5530402-uhd_3840_2160_25fps.mp4" />
  );
};
```

**Cost**: $0.017

Increasing the resolution to 4K does significantly increase the time to render and therfore the cost.

## Additional cost [​](\#additional-cost "Direct link to Additional cost")

The calculation above only factors in Lambda computation cost.

This is usually the majority of the cost.

However, you also incur additional cost for:

- Bandwidth: Pulling in assets happens always over HTTP and incurs S3 egress cost.
- Storage: Storing the sites and renders in S3 incurs S3 storage.
- CloudWatch Logs: Renders are logged to CloudWatch by default.
- [Remotion Licensing](https://remotion.pro/license): Teams of 4+ people also need to acquire a Remotion Company License in addition to the AWS cost.

## See also [​](\#see-also "Direct link to See also")

- [Optimizing for cost](/docs/lambda/optimizing-cost)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/cost-example.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Webhooks](/docs/lambda/webhooks) [Next\
\
Auto-delete renders](/docs/lambda/autodelete)

- [Rendering the Hello World project](#rendering-the-hello-world-project)
- [Rendering a 1 minute Video in the same S3 bucket](#rendering-a-1-minute-video-in-the-same-s3-bucket)
- [Rendering a 10 minute Remote HD Video](#rendering-a-10-minute-remote-hd-video)
- [Rendering a 10 second Remote 4K Video](#rendering-a-10-second-remote-4k-video)
- [Additional cost](#additional-cost)
- [See also](#see-also)