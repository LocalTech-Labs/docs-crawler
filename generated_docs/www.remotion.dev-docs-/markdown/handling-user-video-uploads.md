---
title: Handling user video uploads
source: Remotion Documentation
last_updated: 2024-12-22
---

# Handling user video uploads

- [Home page](/)
- Building apps
- Handling user video uploads

On this page

# Handling user video uploads

In an app where users can upload videos and edit them, we can create a better user experience by loading the video into a player even before the upload is finished. Good news: This can be done pretty easily!

## Allowing user uploads [​](\#allowing-user-uploads "Direct link to Allowing user uploads")

We have a component which returns a [`<Video>`](/docs/video) tag that includes an URL as source.

```

MyComposition.tsx
tsx

import {AbsoluteFill, Video} from 'remotion';

type VideoProps = {
  videoURL: string;
};

export const MyComponent: React.FC<VideoProps> = ({videoURL}) => {
  return (
    <AbsoluteFill>
      <Video src={videoURL} />
    </AbsoluteFill>
  );
};
```

```

MyComposition.tsx
tsx

import {AbsoluteFill, Video} from 'remotion';

type VideoProps = {
  videoURL: string;
};

export const MyComponent: React.FC<VideoProps> = ({videoURL}) => {
  return (
    <AbsoluteFill>
      <Video src={videoURL} />
    </AbsoluteFill>
  );
};
```

The video URL will be passed from the Remotion Player to our component.

Using a `<input type="file">` element, we allow a user upload.

As soon as a file is fully uploaded to the cloud, the URL will be set and can be used by the component to display the video.

```

App.tsx
tsx

import {Player} from '@remotion/player';
import {useState} from 'react';

export const RemotionPlayer: React.FC = () => {
  const [videoUrl, setVideoUrl] = useState<string | null>(null);

  const handleChange = useCallback(
    async (event: React.ChangeEvent<HTMLInputElement>) => {
      if (event.target.files === null) {
        return;
      }

      const file = event.target.files[0];
      //upload is an example function  & returns a URL when a file is uploaded on the cloud.
      const cloudURL = await upload(file);
      // E.g., cloudURL = https://exampleBucketName.s3.ExampleAwsRegion.amazonaws.com
      setVideoUrl(cloudURL);
    },
    [],
  );

  return (
    <div>
      {videoUrl === null ? null : (
        <Player
          component={MyComposition}
          durationInFrames={120}
          compositionWidth={1920}
          compositionHeight={1080}
          fps={30}
          inputProps={{videoUrl}}
        />
      )}

      <input type="file" onChange={handleChange} />
    </div>
  );
};
```

```

App.tsx
tsx

import {Player} from '@remotion/player';
import {useState} from 'react';

export const RemotionPlayer: React.FC = () => {
  const [videoUrl, setVideoUrl] = useState<string | null>(null);

  const handleChange = useCallback(
    async (event: React.ChangeEvent<HTMLInputElement>) => {
      if (event.target.files === null) {
        return;
      }

      const file = event.target.files[0];
      //upload is an example function  & returns a URL when a file is uploaded on the cloud.
      const cloudURL = await upload(file);
      // E.g., cloudURL = https://exampleBucketName.s3.ExampleAwsRegion.amazonaws.com
      setVideoUrl(cloudURL);
    },
    [],
  );

  return (
    <div>
      {videoUrl === null ? null : (
        <Player
          component={MyComposition}
          durationInFrames={120}
          compositionWidth={1920}
          compositionHeight={1080}
          fps={30}
          inputProps={{videoUrl}}
        />
      )}

      <input type="file" onChange={handleChange} />
    </div>
  );
};
```

The implementation of the `upload()` function is provider-specific and we do not show an implementation in this article. We assume it is a function that takes a file, uploads it and returns an URL.

## Optimistic updates [​](\#optimistic-updates "Direct link to Optimistic updates")

When we start the upload of the file, we can create a blob URL using `URL.createObjectURL()` which can be used to display the local file in a [`<Video>`](/docs/video) tag. When the file is done uploading and we get the remote URL, the component shall use remote URL as source.

```

App.tsx
tsx

import {Player} from '@remotion/player';
import {useCallback, useState} from 'react';

type VideoState =
  | {
      type: 'empty';
    }
  | {
      type: 'blob' | 'cloud';
      url: string;
    };

export const RemotionPlayer: React.FC = () => {
  const [videoState, setVideoState] = useState<VideoState>({
    type: 'empty',
  });

  const handleChange = useCallback(
    async (event: React.ChangeEvent<HTMLInputElement>) => {
      if (event.target.files === null) {
        return;
      }

      const file = event.target.files[0];
      const blobUrl = URL.createObjectURL(file);
      setVideoState({type: 'blob', url: blobUrl});
      const cloudUrl = await upload(file);
      setVideoState({type: 'cloud', url: cloudUrl});
      URL.revokeObjectURL(blobUrl);
    },
    [],
  );

  return (
    <div>
      {videoState.type !== 'empty' ? (
        <Player
          component={MyComposition}
          durationInFrames={120}
          compositionWidth={1920}
          compositionHeight={1080}
          fps={30}
          inputProps={{videoUrl: videoState.url}}
        />
      ) : null}
      <input type="file" onChange={handleChange} />
    </div>
  );
};
```

```

App.tsx
tsx

import {Player} from '@remotion/player';
import {useCallback, useState} from 'react';

type VideoState =
  | {
      type: 'empty';
    }
  | {
      type: 'blob' | 'cloud';
      url: string;
    };

export const RemotionPlayer: React.FC = () => {
  const [videoState, setVideoState] = useState<VideoState>({
    type: 'empty',
  });

  const handleChange = useCallback(
    async (event: React.ChangeEvent<HTMLInputElement>) => {
      if (event.target.files === null) {
        return;
      }

      const file = event.target.files[0];
      const blobUrl = URL.createObjectURL(file);
      setVideoState({type: 'blob', url: blobUrl});
      const cloudUrl = await upload(file);
      setVideoState({type: 'cloud', url: cloudUrl});
      URL.revokeObjectURL(blobUrl);
    },
    [],
  );

  return (
    <div>
      {videoState.type !== 'empty' ? (
        <Player
          component={MyComposition}
          durationInFrames={120}
          compositionWidth={1920}
          compositionHeight={1080}
          fps={30}
          inputProps={{videoUrl: videoState.url}}
        />
      ) : null}
      <input type="file" onChange={handleChange} />
    </div>
  );
};
```

This will result in the user immediately seeing the video as they drag it into the input field. It is a good practice to call `URL.revokeObjectURL()` after the local video is not used anymore to free up the used memory.

note

As soon as possible, replace the `blob:` URL with the real URL.

Since `blob:` URLs don't support the HTTP `Range` header, video seeking performance is degraded when using `blob:` URLs.

## See also [​](\#see-also "Direct link to See also")

- [Uploading with presigned URLs](/docs/presigned-urls)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/video-uploads.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Turn a <Player> into a Remotion project](/docs/player-into-remotion-project) [Next\
\
Upload with a presigned URL](/docs/presigned-urls)

- [Allowing user uploads](#allowing-user-uploads)
- [Optimistic updates](#optimistic-updates)
- [See also](#see-also)
