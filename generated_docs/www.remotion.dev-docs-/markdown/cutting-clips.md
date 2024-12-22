---
title: Cutting clips
source: Remotion Documentation
last_updated: 2024-12-22
---

# Cutting clips

- [Home page](/)
- Recorder
- Editing
- Cutting clips

# Cutting clips

The recommended way is to record many small bite-sized clips with the Recording interface and create many scenes.

You will have to do less re-takes due to mistakes and having multiple scenes allows adding transitions inbetween.

If this is unfeasible and you already have a long recording that you would like to split after the fact, you can use the following script:

```

cut.ts
tsx

import { $ } from "bun";
import { existsSync, readdirSync, renameSync, rmSync } from "fs";
import { join } from "path";
import {
  ALTERNATIVE1_PREFIX,
  ALTERNATIVE2_PREFIX,
  CAPTIONS_PREFIX,
  DISPLAY_PREFIX,
  WEBCAM_PREFIX,
} from "./config/cameras";

const compositionId = process.argv[2];
if (!compositionId) {
  throw new Error("Expected composition ID as first argument");
}

console.log("Composition ID", compositionId);
const timestamp = process.argv[3];
if (!timestamp) {
  throw new Error("Expected timestamp as second argument");
}

const trimPoint = process.argv[4];
if (!trimPoint) {
  throw new Error("Expected trim point as third argument");
}

console.log("Timestamp", timestamp);
console.log("Trim point", trimPoint);

const folder = join("public", compositionId);
const files = readdirSync(folder);

const webcamFile = files.find((f) => f.startsWith(`webcam${timestamp}`));

if (!webcamFile) {
  throw new Error(
    `Expected file ${compositionId}/webcam${timestamp}.* to exist`,
  );
}

const webcamTimestamp = webcamFile.match(/\d{1,14}/g)?.[0] ?? null;

if (!webcamTimestamp) {
  throw new Error(
    `Expected file ${compositionId}/webcam${timestamp}.* to have timestamp ${timestamp}`,
  );
}

const allFilesWithTimestamp = files.filter((f) => {
  return (
    f.startsWith(`${WEBCAM_PREFIX}${timestamp}`) ||
    f.startsWith(`${DISPLAY_PREFIX}${timestamp}`) ||
    f.startsWith(`${ALTERNATIVE1_PREFIX}${timestamp}.mp4`) ||
    f.startsWith(`${ALTERNATIVE2_PREFIX}${timestamp}.mp4`)
  );
});

const allTimestamps = [
  ...new Set(
    files.map((f) => {
      return f.match(/\d{1,14}/g)?.[0] ?? null;
    }),
  ),
];

const timestampBefore = Number(
  allTimestamps.filter((t) => Number(t) < Number(webcamTimestamp) ?? null) ??
    "0",
);

const newTimestamp =
  (Number(webcamTimestamp) - timestampBefore) / 2 + timestampBefore;

for (const file of allFilesWithTimestamp) {
  const newFile = file.replace(
    new RegExp(`${webcamTimestamp}`),
    String(newTimestamp),
  );

  // Create video before trim point
  await $`bunx ffmpeg -i ${file} -t ${trimPoint} -c:v copy -c:a copy ${newFile}`.cwd(
    folder,
  );

  // Create video after trim point
  const temporaryFile = `temp-${newFile}`;
  await $`bunx ffmpeg -ss ${trimPoint} -accurate_seek -i ${file} ${temporaryFile}`.cwd(
    folder,
  );
  rmSync(join(folder, file));
  renameSync(join(folder, temporaryFile), join(folder, file));
}

// Remove the captions file
const captionsFile = join(folder, `${CAPTIONS_PREFIX}${timestamp}.json`);
if (existsSync(captionsFile)) {
  rmSync(captionsFile);
}
```

```

cut.ts
tsx

import { $ } from "bun";
import { existsSync, readdirSync, renameSync, rmSync } from "fs";
import { join } from "path";
import {
  ALTERNATIVE1_PREFIX,
  ALTERNATIVE2_PREFIX,
  CAPTIONS_PREFIX,
  DISPLAY_PREFIX,
  WEBCAM_PREFIX,
} from "./config/cameras";

const compositionId = process.argv[2];
if (!compositionId) {
  throw new Error("Expected composition ID as first argument");
}

console.log("Composition ID", compositionId);
const timestamp = process.argv[3];
if (!timestamp) {
  throw new Error("Expected timestamp as second argument");
}

const trimPoint = process.argv[4];
if (!trimPoint) {
  throw new Error("Expected trim point as third argument");
}

console.log("Timestamp", timestamp);
console.log("Trim point", trimPoint);

const folder = join("public", compositionId);
const files = readdirSync(folder);

const webcamFile = files.find((f) => f.startsWith(`webcam${timestamp}`));

if (!webcamFile) {
  throw new Error(
    `Expected file ${compositionId}/webcam${timestamp}.* to exist`,
  );
}

const webcamTimestamp = webcamFile.match(/\d{1,14}/g)?.[0] ?? null;

if (!webcamTimestamp) {
  throw new Error(
    `Expected file ${compositionId}/webcam${timestamp}.* to have timestamp ${timestamp}`,
  );
}

const allFilesWithTimestamp = files.filter((f) => {
  return (
    f.startsWith(`${WEBCAM_PREFIX}${timestamp}`) ||
    f.startsWith(`${DISPLAY_PREFIX}${timestamp}`) ||
    f.startsWith(`${ALTERNATIVE1_PREFIX}${timestamp}.mp4`) ||
    f.startsWith(`${ALTERNATIVE2_PREFIX}${timestamp}.mp4`)
  );
});

const allTimestamps = [
  ...new Set(
    files.map((f) => {
      return f.match(/\d{1,14}/g)?.[0] ?? null;
    }),
  ),
];

const timestampBefore = Number(
  allTimestamps.filter((t) => Number(t) < Number(webcamTimestamp) ?? null) ??
    "0",
);

const newTimestamp =
  (Number(webcamTimestamp) - timestampBefore) / 2 + timestampBefore;

for (const file of allFilesWithTimestamp) {
  const newFile = file.replace(
    new RegExp(`${webcamTimestamp}`),
    String(newTimestamp),
  );

  // Create video before trim point
  await $`bunx ffmpeg -i ${file} -t ${trimPoint} -c:v copy -c:a copy ${newFile}`.cwd(
    folder,
  );

  // Create video after trim point
  const temporaryFile = `temp-${newFile}`;
  await $`bunx ffmpeg -ss ${trimPoint} -accurate_seek -i ${file} ${temporaryFile}`.cwd(
    folder,
  );
  rmSync(join(folder, file));
  renameSync(join(folder, temporaryFile), join(folder, file));
}

// Remove the captions file
const captionsFile = join(folder, `${CAPTIONS_PREFIX}${timestamp}.json`);
if (existsSync(captionsFile)) {
  rmSync(captionsFile);
}
```

The following is a destructive action - commit your changes before you run this.

To use this script, find the file that you would like to split, e.g. `my-video/webcam123456789012345.mp4`.

Execute the following on the command line:

```

shell

bun cut.ts my-video 123456789012345 2.0
#          ^ comp   ^timestamp      ^ cut point in seconds
```

```

shell

bun cut.ts my-video 123456789012345 2.0
#          ^ comp   ^timestamp      ^ cut point in seconds
```

This will split all recordings into two.

It will also delete the captions file – to regenerate it, run:

```

shell

bun sub.ts
```

```

shell

bun sub.ts
```

Now, in the Studio, [add a new scene](/docs/recorder/editing/#add-new-scenes-to-the-composition) after the one you just edited.

We plan on adding a GUI shortcut for this workflow in the future.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/editing/cutting-clips.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
B-Roll](/docs/recorder/editing/b-roll) [Next\
\
Normalizing audio levels](/docs/recorder/editing/normalizing-audio)
