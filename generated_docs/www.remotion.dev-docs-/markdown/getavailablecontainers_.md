---
title: getAvailableContainers()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getAvailableContainers()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- getAvailableContainers()

On this page

# getAvailableContainers()

Returns an array of available containers that can be used with the `convertMedia` function.

```

Getting available containers
tsx

import {getAvailableContainers} from '@remotion/webcodecs';

const containers = getAvailableContainers();
console.log(containers);
```

```

Getting available containers
tsx

import {getAvailableContainers} from '@remotion/webcodecs';

const containers = getAvailableContainers();
console.log(containers);
```

note

New containers may be added to this function and it will not be considered a breaking change.

## As a type [​](\#as-a-type "Direct link to As a type")

If you need a TypeScript type that covers the available output containers, you can import the type definition:

```

Type definition
tsx

import type {ConvertMediaContainer} from '@remotion/webcodecs';

(alias) type ConvertMediaContainer = "webm" | "mp4" | "wav"
import ConvertMediaContainer
```

```

Type definition
tsx

import type {ConvertMediaContainer} from '@remotion/webcodecs';

(alias) type ConvertMediaContainer = "webm" | "mp4" | "wav"
import ConvertMediaContainer
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/get-available-containers.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/get-available-containers.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
convertMedia()](/docs/webcodecs/convert-media) [Next\
\
canReencodeAudioTrack()](/docs/webcodecs/can-reencode-audio-track)

- [As a type](#as-a-type)
- [See also](#see-also)
