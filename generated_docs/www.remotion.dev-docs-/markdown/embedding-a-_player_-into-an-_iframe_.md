---
title: Embedding a <Player> into an <iframe>
source: Remotion Documentation
last_updated: 2024-12-22
---

# Embedding a <Player> into an <iframe>

- [Home page](/)
- Snippets
- Embedding a <Player> into an <iframe>

On this page

# Embedding a <Player> into an <iframe>

info

Credit to [@marcusstenbeck](https://twitter.com/marcusstenbeck) for creating this snippet.

This snippet is useful if you want to isolate the global styles of your homepage from the global styles of the [`<Player>`](/docs/player).

```

Usage
diff

- import { Player } from '@remotion/player';
+ import { IframePlayer } from 'path/to/IframePlayer';

- <Player {/* ... */} />
+ <IframePlayer {/* ... */} />
```

```

Usage
diff

- import { Player } from '@remotion/player';
+ import { IframePlayer } from 'path/to/IframePlayer';

- <Player {/* ... */} />
+ <IframePlayer {/* ... */} />
```

```

IframePlayer.tsx
tsx

import { Player, PlayerProps, PlayerRef } from "@remotion/player";
import React, { forwardRef, useEffect, useRef, useState } from "react";
import ReactDOM from "react-dom";
import { AnyZodObject } from "zod";

const className = "__player";
const borderNone: React.CSSProperties = {
  border: "none",
};

const IframePlayerWithoutRef = <T extends Record<string, unknown>>(
  props: PlayerProps<AnyZodObject, T>,
  ref: React.Ref<PlayerRef>
) => {
  const [contentRef, setContentRef] = useState<HTMLIFrameElement | null>(null);
  const resizeObserverRef = useRef<ResizeObserver | null>(null);
  const mountNode = contentRef?.contentDocument?.body;
  useEffect(() => {
    if (!contentRef || !contentRef.contentDocument) return;
    // Remove margin and padding so player fits snugly
    contentRef.contentDocument.body.style.margin = "0";
    contentRef.contentDocument.body.style.padding = "0";
    // When player div is resized also resize iframe
    resizeObserverRef.current = new ResizeObserver(([playerEntry]) => {
      const playerRect = playerEntry.contentRect;
      contentRef.width = String(playerRect.width);
      contentRef.height = String(playerRect.height);
    });
    // The remotion player element
    const playerElement = contentRef.contentDocument.querySelector(
      "." + className
    );
    if (!playerElement) {
      throw new Error(
        'Player element not found. Add a "' +
          className +
          '" class to the <Player>.'
      );
    }
    // Watch the player element for size changes
    resizeObserverRef.current.observe(playerElement as Element);
    return () => {
      // ContentRef changed: unobserve!
      (resizeObserverRef.current as ResizeObserver).unobserve(
        playerElement as Element
      );
    };
  }, [contentRef]);
  const combinedClassName = `${className} ${props.className ?? ""}`.trim();
  return (
    // eslint-disable-next-line @remotion/warn-native-media-tag
    <iframe ref={setContentRef} style={borderNone}>
      {mountNode &&
        ReactDOM.createPortal(
          // @ts-expect-error PlayerProps are incorrectly typed
          <Player<AnyZodObject, T>
            {...props}
            ref={ref}
            className={combinedClassName}
          />,
          mountNode
        )}
    </iframe>
  );
};
export const IframePlayer = forwardRef(IframePlayerWithoutRef);
```

```

IframePlayer.tsx
tsx

import { Player, PlayerProps, PlayerRef } from "@remotion/player";
import React, { forwardRef, useEffect, useRef, useState } from "react";
import ReactDOM from "react-dom";
import { AnyZodObject } from "zod";

const className = "__player";
const borderNone: React.CSSProperties = {
  border: "none",
};

const IframePlayerWithoutRef = <T extends Record<string, unknown>>(
  props: PlayerProps<AnyZodObject, T>,
  ref: React.Ref<PlayerRef>
) => {
  const [contentRef, setContentRef] = useState<HTMLIFrameElement | null>(null);
  const resizeObserverRef = useRef<ResizeObserver | null>(null);
  const mountNode = contentRef?.contentDocument?.body;
  useEffect(() => {
    if (!contentRef || !contentRef.contentDocument) return;
    // Remove margin and padding so player fits snugly
    contentRef.contentDocument.body.style.margin = "0";
    contentRef.contentDocument.body.style.padding = "0";
    // When player div is resized also resize iframe
    resizeObserverRef.current = new ResizeObserver(([playerEntry]) => {
      const playerRect = playerEntry.contentRect;
      contentRef.width = String(playerRect.width);
      contentRef.height = String(playerRect.height);
    });
    // The remotion player element
    const playerElement = contentRef.contentDocument.querySelector(
      "." + className
    );
    if (!playerElement) {
      throw new Error(
        'Player element not found. Add a "' +
          className +
          '" class to the <Player>.'
      );
    }
    // Watch the player element for size changes
    resizeObserverRef.current.observe(playerElement as Element);
    return () => {
      // ContentRef changed: unobserve!
      (resizeObserverRef.current as ResizeObserver).unobserve(
        playerElement as Element
      );
    };
  }, [contentRef]);
  const combinedClassName = `${className} ${props.className ?? ""}`.trim();
  return (
    // eslint-disable-next-line @remotion/warn-native-media-tag
    <iframe ref={setContentRef} style={borderNone}>
      {mountNode &&
        ReactDOM.createPortal(
          // @ts-expect-error PlayerProps are incorrectly typed
          <Player<AnyZodObject, T>
            {...props}
            ref={ref}
            className={combinedClassName}
          />,
          mountNode
        )}
    </iframe>
  );
};
export const IframePlayer = forwardRef(IframePlayerWithoutRef);
```

## See also [​](\#see-also "Direct link to See also")

- [`<Player>` examples](/docs/player/examples)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/snippets/player-in-iframe.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Different segments at different speeds](/docs/miscellaneous/snippets/different-segments-at-different-speeds) [Next\
\
useDelayRender()](/docs/miscellaneous/snippets/use-delay-render)

- [See also](#see-also)
