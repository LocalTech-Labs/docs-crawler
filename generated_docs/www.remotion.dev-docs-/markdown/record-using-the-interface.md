---
title: Record using the interface
source: Remotion Documentation
last_updated: 2024-12-22
---

# Record using the interface

- [Home page](/)
- Recorder
- Record
- Using the interface

On this page

# Record using the interface

The Recorder ships with its own recording interface.

From it, you can record a clip and add it as a scene to your project.

If you haven't done so, start the Remotion Recorder:

```

tsx

bun start
```

```

tsx

bun start
```

A Vite app will listen on [`http://localhost:4000`](http://localhost:4000).

Open this URL in the browser.

Give the page access to the webcam and microphone.

## Select the composition [​](\#select-the-composition "Direct link to Select the composition")

Recordings need to be stored in a subfolder of `public` with the name of the composition ID.

For example, if you want to record for the composition `my-video`, recordings will be stored under `public/my-video`.

In the top right corner, use the dropdown to "Create a new folder".

Give it the same name as the ID of the composition you want to record for.

The folder will be created and your subsequent recordings will be saved in this folder.

## Select sources [​](\#select-sources "Direct link to Select sources")

In the top left panel, select your webcam and microphone.

In the top right panel, optionally select the screen you want to record.

You can record more sources, but out of the box, the Remotion Recorder does not do anything with them.

## Before starting [​](\#before-starting "Direct link to Before starting")

Say "Test One Two Three" and ensure the volume meter is not going into the red zone.

Make sure you are centered in the video. You can enable the crop target overlay to see how the video will crop when miniaturized.

Prefer making short recordings. You're less likely to make a mistake and it's easy to stitch them together.

There is no countdown – The recording will start immediately.

However, the Recorder will later automatically trim away the silence in the beginning.

Take a deep breath and think of what you are about to say and show.

You're ready to record!

## Start the recording [​](\#start-the-recording "Direct link to Start the recording")

To start recording, press `R` or click

Start recording

in the top left corner.

Once you are done with the scene, click

Stop recording

.

## Selecting a take [​](\#selecting-a-take "Direct link to Selecting a take")

After you have recorded a scene and you are happy with your take, click "Use this take".

The server will then do the following:

- Convert the videos from WebM to MP4
- Put them into your project under `public/[composition-id]`.
- Caption the files with Whisper.cpp

## Discard a take [​](\#discard-a-take "Direct link to Discard a take")

If you have taken a recording that you don't want to use, click "Discard and retake" to clear the existing recording from memory.

## record.remotion.dev [​](\#recordremotiondev "Direct link to record.remotion.dev")

You can also use the hosted version of the recoding interface at [record.remotion.dev](https://record.remotion.dev).

It will not allow you to select a save destination, but instead download the recordings to your `Downloads` folder.

They will have names like `webcam171234567890.mp4` and `display171234567890.mp4`.

To copy the recordings to your Remotion project, first open `copy.ts` in your code editor. Set the `prefix` at the top of your file to the ID of the composition that you gave when you first created a new video.

```

copy.ts
ts

const prefix = '<id-of-your-composition>';
```

```

copy.ts
ts

const prefix = '<id-of-your-composition>';
```

Then run the `copy.ts` script whenever you want to save a take and immediately generate captions:

```

bash

bun copy.ts
bun sub.ts
```

```

bash

bun copy.ts
bun sub.ts
```

If you want to not use a take, just make another recording. The `copy.ts` script will only take the recordings with the latest timestamp.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/record/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Create a new video](/docs/recorder/create) [Next\
\
Manually add recordings](/docs/recorder/record/manually)

- [Select the composition](#select-the-composition)
- [Select sources](#select-sources)
- [Before starting](#before-starting)
- [Start the recording](#start-the-recording)
- [Selecting a take](#selecting-a-take)
- [Discard a take](#discard-a-take)
- [record.remotion.dev](#recordremotiondev)