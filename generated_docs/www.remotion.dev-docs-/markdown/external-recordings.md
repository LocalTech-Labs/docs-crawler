---
title: External recordings
source: Remotion Documentation
last_updated: 2024-12-22
---

# External recordings

- [Home page](/)
- Recorder
- External recordings

On this page

# External recordings

You can also include recordings that you recorded externally in your Recorder video.

For example, you may edit a vlog with it.

## Add files [​](\#add-files "Direct link to Add files")

Transfer the recordings from your camera to your computer and put them in the `public/<composition-id>` folder.

The file structure should look like this:

```

Assuming the composition id is 'vlog'
bash

public
├── vlog
│ ├── camera10.mov
│ ├── camera20.mov
│ └── camera30.mov

```

```

Assuming the composition id is 'vlog'
bash

public
├── vlog
│ ├── camera10.mov
│ ├── camera20.mov
│ └── camera30.mov

```

The prefix must be `"webcam"` for every video.

The suffixes must be ascending in the order you want to use the videos.

## Generate captions [​](\#generate-captions "Direct link to Generate captions")

Run

```

bash

bun sub.ts
```

```

bash

bun sub.ts
```

to generate captions for your recordings.

## Use the videos [​](\#use-the-videos "Direct link to Use the videos")

Now you can add more scenes in the right sidebar to start including the recordings.

## iPhone recordings [​](\#iphone-recordings "Direct link to iPhone recordings")

iPhone recordings are HEVC videos, which during rendering Chrome cannot get the dimensions of the video.

You will get a render failure.

Here is a script to convert the recordings to MP4:

```

convert.ts
ts

import { $ } from "bun";
import { renameSync, unlinkSync } from "fs";

const compositionId = "vlog";

const ls = (await $`ls public/${compositionId}`).stdout.toString("utf-8");
const files = ls.split("\n");

for (const file of files) {
  if (
    file.toLowerCase().endsWith(".mov") ||
    file.toLowerCase().endsWith(".mp4")
  ) {
    await $`bunx remotion ffmpeg -i public/${compositionId}/${file} public/${compositionId}/_${file}.mp4`;
    unlinkSync(`public/${compositionId}/${file}`);
    renameSync(
      `public/${compositionId}/_${file}.mp4`,
      `public/${compositionId}/${file.replace(/ /g, "_")}`,
    );
  }
}
```

```

convert.ts
ts

import { $ } from "bun";
import { renameSync, unlinkSync } from "fs";

const compositionId = "vlog";

const ls = (await $`ls public/${compositionId}`).stdout.toString("utf-8");
const files = ls.split("\n");

for (const file of files) {
  if (
    file.toLowerCase().endsWith(".mov") ||
    file.toLowerCase().endsWith(".mp4")
  ) {
    await $`bunx remotion ffmpeg -i public/${compositionId}/${file} public/${compositionId}/_${file}.mp4`;
    unlinkSync(`public/${compositionId}/${file}`);
    renameSync(
      `public/${compositionId}/_${file}.mp4`,
      `public/${compositionId}/${file.replace(/ /g, "_")}`,
    );
  }
}
```

Run it using

```

sh

bun convert.ts
```

```

sh

bun convert.ts
```

to fix the issue.

## Examples [​](\#examples "Direct link to Examples")

In the [`our-recorder`](https://github.com/remotion-dev/our-recorder) repository, the compositions `marathon` and `airportrun` are using real-life footage.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/external-recordings.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Source Control](/docs/recorder/source-control) [Next\
\
Upgrading](/docs/recorder/upgrading)

- [Add files](#add-files)
- [Generate captions](#generate-captions)
- [Use the videos](#use-the-videos)
- [iPhone recordings](#iphone-recordings)
- [Examples](#examples)
