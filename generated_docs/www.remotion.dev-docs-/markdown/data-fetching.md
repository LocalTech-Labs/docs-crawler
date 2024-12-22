---
title: Data fetching
source: Remotion Documentation
last_updated: 2024-12-22
---

# Data fetching

- [Home page](/)
- [Parameterized videos](/docs/parameterized-rendering)
- Data fetching

On this page

# Data fetching

In Remotion, you may fetch data from an API to use it in your video. On this page, we document recipes and best practices.

## Fetching data before the render [v4.0.0](https://github.com/remotion-dev/remotion/releases/v4.0.0) [​](\#fetching-data-before-the-render "Direct link to fetching-data-before-the-render")

You may use the [`calculateMetadata`](/docs/composition#calculatemetadata) prop of the [`<Composition />`](/docs/composition) component to alter the props that get passed to your React component.

### When to use [​](\#when-to-use "Direct link to When to use")

The data being fetched using `calculateMetadata()` must be JSON-serializable. That means it is useful for API responses, but not for assets in binary format.

### Usage [​](\#usage "Direct link to Usage")

Pass a callback function which takes the untransformed `props`, and return an object with the new props.

```

src/Root.tsx
tsx

import { Composition } from "remotion";

type ApiResponse = {
  title: string;
  description: string;
};
type MyCompProps = {
  id: string;
  data: ApiResponse | null;
};

const MyComp: React.FC<MyCompProps> = () => null;

export const Root: React.FC = () => {
  return (
    <Composition
      id="MyComp"
      component={MyComp}
      durationInFrames={300}
      fps={30}
      width={1920}
      height={1080}
      defaultProps={{
        id: "1",
        data: null,
      }}
      calculateMetadata={async ({ props }) => {
        const data = await fetch(`https://example.com/api/${props.id}`);
        const json = await data.json();

        return {
          props: {
            ...props,
            data: json,
          },
        };
      }}
    />
  );
};
```

```

src/Root.tsx
tsx

import { Composition } from "remotion";

type ApiResponse = {
  title: string;
  description: string;
};
type MyCompProps = {
  id: string;
  data: ApiResponse | null;
};

const MyComp: React.FC<MyCompProps> = () => null;

export const Root: React.FC = () => {
  return (
    <Composition
      id="MyComp"
      component={MyComp}
      durationInFrames={300}
      fps={30}
      width={1920}
      height={1080}
      defaultProps={{
        id: "1",
        data: null,
      }}
      calculateMetadata={async ({ props }) => {
        const data = await fetch(`https://example.com/api/${props.id}`);
        const json = await data.json();

        return {
          props: {
            ...props,
            data: json,
          },
        };
      }}
    />
  );
};
```

The `props` being passed to `calculateMetadata()` are the [input props merged together with the default props](/docs/props-resolution).

In addition to `props`, `defaultProps` can also be read from the same object.

When transforming, the input and the output must be the same TypeScript type.
Consider using a nullable type for your data and throw an error inside your component to deal with the `null` type:

```

MyComp.tsx
tsx

type MyCompProps = {
  id: string;
  data: ApiResponse | null;
};

const MyComp: React.FC<MyCompProps> = ({ data }) => {
  if (data === null) {
    throw new Error("Data was not fetched");
  }

  return <div>{data.title}</div>;
};
```

```

MyComp.tsx
tsx

type MyCompProps = {
  id: string;
  data: ApiResponse | null;
};

const MyComp: React.FC<MyCompProps> = ({ data }) => {
  if (data === null) {
    throw new Error("Data was not fetched");
  }

  return <div>{data.title}</div>;
};
```

### TypeScript typing [​](\#typescript-typing "Direct link to TypeScript typing")

You may use the `CalculateMetadataFunction` type from `remotion` to type your callback function. Pass as the generic value ( `<>`) the type of your props.

```

src/Root.tsx
tsx

import { CalculateMetadataFunction } from "remotion";

export const calculateMyCompMetadata: CalculateMetadataFunction<
  MyCompProps
> = ({ props }) => {
  return {
    props: {
      ...props,
      data: {
        title: "Hello world",
        description: "This is a description",
      },
    },
  };
};

export const MyComp: React.FC<MyCompProps> = () => null;
```

```

src/Root.tsx
tsx

import { CalculateMetadataFunction } from "remotion";

export const calculateMyCompMetadata: CalculateMetadataFunction<
  MyCompProps
> = ({ props }) => {
  return {
    props: {
      ...props,
      data: {
        title: "Hello world",
        description: "This is a description",
      },
    },
  };
};

export const MyComp: React.FC<MyCompProps> = () => null;
```

### Colocation [​](\#colocation "Direct link to Colocation")

Here is an example of how you could define a schema, a component and a fetcher function in the same file:

```

MyComp.tsx
tsx

import { CalculateMetadataFunction } from "remotion";
import { z } from "zod";

const apiResponse = z.object({ title: z.string(), description: z.string() });

export const myCompSchema = z.object({
  id: z.string(),
  data: z.nullable(apiResponse),
});

type Props = z.infer<typeof myCompSchema>;

export const calcMyCompMetadata: CalculateMetadataFunction<Props> = async ({
  props,
}) => {
  const data = await fetch(`https://example.com/api/${props.id}`);
  const json = await data.json();

  return {
    props: {
      ...props,
      data: json,
    },
  };
};

export const MyComp: React.FC<Props> = ({ data }) => {
  if (data === null) {
    throw new Error("Data was not fetched");
  }

  return <div>{data.title}</div>;
};
```

```

MyComp.tsx
tsx

import { CalculateMetadataFunction } from "remotion";
import { z } from "zod";

const apiResponse = z.object({ title: z.string(), description: z.string() });

export const myCompSchema = z.object({
  id: z.string(),
  data: z.nullable(apiResponse),
});

type Props = z.infer<typeof myCompSchema>;

export const calcMyCompMetadata: CalculateMetadataFunction<Props> = async ({
  props,
}) => {
  const data = await fetch(`https://example.com/api/${props.id}`);
  const json = await data.json();

  return {
    props: {
      ...props,
      data: json,
    },
  };
};

export const MyComp: React.FC<Props> = ({ data }) => {
  if (data === null) {
    throw new Error("Data was not fetched");
  }

  return <div>{data.title}</div>;
};
```

```

src/Root.tsx
tsx

import React from "react";
import { Composition } from "remotion";
import { MyComp, calcMyCompMetadata, myCompSchema } from "./MyComp";

export const Root = () => {
  return (
    <Composition
      id="MyComp"
      component={MyComp}
      durationInFrames={300}
      fps={30}
      width={1920}
      height={1080}
      defaultProps={{
        id: "1",
        data: null,
      }}
      schema={myCompSchema}
      calculateMetadata={calcMyCompMetadata}
    />
  );
};
```

```

src/Root.tsx
tsx

import React from "react";
import { Composition } from "remotion";
import { MyComp, calcMyCompMetadata, myCompSchema } from "./MyComp";

export const Root = () => {
  return (
    <Composition
      id="MyComp"
      component={MyComp}
      durationInFrames={300}
      fps={30}
      width={1920}
      height={1080}
      defaultProps={{
        id: "1",
        data: null,
      }}
      schema={myCompSchema}
      calculateMetadata={calcMyCompMetadata}
    />
  );
};
```

By implementing this pattern, The `id` in the [props editor](/docs/visual-editing) can now be tweaked and Remotion will refetch the data whenever it changes.

### Setting the duration based on data [​](\#setting-the-duration-based-on-data "Direct link to Setting the duration based on data")

You may set the `durationInFrames`, `fps`, `width` and `height` by returning those keys in the callback function:

```

tsx

import { CalculateMetadataFunction } from "remotion";

type MyCompProps = {
  durationInSeconds: number;
};

export const calculateMyCompMetadata: CalculateMetadataFunction<
  MyCompProps
> = ({ props }) => {
  const fps = 30;
  const durationInSeconds = props.durationInSeconds;

  return {
    durationInFrames: durationInSeconds * fps,
    fps,
  };
};
```

```

tsx

import { CalculateMetadataFunction } from "remotion";

type MyCompProps = {
  durationInSeconds: number;
};

export const calculateMyCompMetadata: CalculateMetadataFunction<
  MyCompProps
> = ({ props }) => {
  const fps = 30;
  const durationInSeconds = props.durationInSeconds;

  return {
    durationInFrames: durationInSeconds * fps,
    fps,
  };
};
```

Learn more about this feature in the [Variable metadata](/docs/dynamic-metadata) page.

### Aborting stale requests [​](\#aborting-stale-requests "Direct link to Aborting stale requests")

The props in the props editor may rapidly change for example by typing fast.

It is a good practice to cancel requests which are stale using the `abortSignal` that gets passed to the [`calculateMetadata()`](/docs/composition#calculatemetadata) function:

```

src/MyComp.tsx
tsx

export const calculateMyCompMetadata: CalculateMetadataFunction<
  MyCompProps
> = async ({ props, abortSignal }) => {
  const data = await fetch(`https://example.com/api/${props.id}`, {
    signal: abortSignal,
  });
  const json = await data.json();

  return {
    props: {
      ...props,
      data: json,
    },
  };
};

export const MyComp: React.FC<MyCompProps> = () => null;
```

```

src/MyComp.tsx
tsx

export const calculateMyCompMetadata: CalculateMetadataFunction<
  MyCompProps
> = async ({ props, abortSignal }) => {
  const data = await fetch(`https://example.com/api/${props.id}`, {
    signal: abortSignal,
  });
  const json = await data.json();

  return {
    props: {
      ...props,
      data: json,
    },
  };
};

export const MyComp: React.FC<MyCompProps> = () => null;
```

This `abortSignal` is created by Remotion using the [`AbortController`](https://developer.mozilla.org/en-US/docs/Web/API/AbortController) API.

### Debouncing requests [​](\#debouncing-requests "Direct link to Debouncing requests")

If you are making requests to an expensive API, you might want to only fire a request after the user has stopped typing for a while. You may use the following function for doing so:

```

src/wait-for-no-input.ts
tsx

import { getRemotionEnvironment } from "remotion";
export const waitForNoInput = (signal: AbortSignal, ms: number) => {
  // Don't wait during rendering
  if (getRemotionEnvironment().isRendering) {
    return Promise.resolve();
  }

  if (signal.aborted) {
    return Promise.reject(new Error("stale"));
  }

  return Promise.race<void>([
    new Promise<void>((_, reject) => {
      signal.addEventListener("abort", () => {
        reject(new Error("stale"));
      });
    }),
    new Promise<void>((resolve) => {
      setTimeout(() => {
        resolve();
      }, ms);
    }),
  ]);
};
```

```

src/wait-for-no-input.ts
tsx

import { getRemotionEnvironment } from "remotion";
export const waitForNoInput = (signal: AbortSignal, ms: number) => {
  // Don't wait during rendering
  if (getRemotionEnvironment().isRendering) {
    return Promise.resolve();
  }

  if (signal.aborted) {
    return Promise.reject(new Error("stale"));
  }

  return Promise.race<void>([
    new Promise<void>((_, reject) => {
      signal.addEventListener("abort", () => {
        reject(new Error("stale"));
      });
    }),
    new Promise<void>((resolve) => {
      setTimeout(() => {
        resolve();
      }, ms);
    }),
  ]);
};
```

```

src/MyComp.tsx
tsx

export const calculateMyCompMetadata: CalculateMetadataFunction<
  MyCompProps
> = async ({ props, abortSignal }) => {
  await waitForNoInput(abortSignal, 750);
  const data = await fetch(`https://example.com/api/${props.id}`, {
    signal: abortSignal,
  });
  const json = await data.json();

  return {
    props: {
      ...props,
      data: json,
    },
  };
};

export const MyComp: React.FC<MyCompProps> = () => null;
```

```

src/MyComp.tsx
tsx

export const calculateMyCompMetadata: CalculateMetadataFunction<
  MyCompProps
> = async ({ props, abortSignal }) => {
  await waitForNoInput(abortSignal, 750);
  const data = await fetch(`https://example.com/api/${props.id}`, {
    signal: abortSignal,
  });
  const json = await data.json();

  return {
    props: {
      ...props,
      data: json,
    },
  };
};

export const MyComp: React.FC<MyCompProps> = () => null;
```

### Time limit [​](\#time-limit "Direct link to Time limit")

When Remotion calls the `calculateMetadata()` function, it wraps it in a [`delayRender()`](/docs/delay-render), which by default times out after 30 seconds.

## Fetching data during the render [​](\#fetching-data-during-the-render "Direct link to Fetching data during the render")

Using [`delayRender()`](/docs/delay-render) and [`continueRender()`](/docs/continue-render) you can tell Remotion to wait for asynchronous operations to finish before rendering a frame.

### When to use [​](\#when-to-use-1 "Direct link to When to use")

Use this approach to load assets which are not JSON serializable or if you are using a Remotion version lower than 4.0.

### Usage [​](\#usage-1 "Direct link to Usage")

Call [`delayRender()`](/docs/delay-render) as soon as possible, for example when initializing the state inside your component.

```

tsx

import { useCallback, useEffect, useState } from "react";
import { cancelRender, continueRender, delayRender } from "remotion";

export const MyComp = () => {
  const [data, setData] = useState(null);
  const [handle] = useState(() => delayRender());

  const fetchData = useCallback(async () => {
    try {
      const response = await fetch("http://example.com/api");
      const json = await response.json();
      setData(json);

      continueRender(handle);
    } catch (err) {
      cancelRender(err);
    }
  }, [handle]);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  return (
    <div>
      {data ? (
        <div>This video has data from an API! {JSON.stringify(data)}</div>
      ) : null}
    </div>
  );
};
```

```

tsx

import { useCallback, useEffect, useState } from "react";
import { cancelRender, continueRender, delayRender } from "remotion";

export const MyComp = () => {
  const [data, setData] = useState(null);
  const [handle] = useState(() => delayRender());

  const fetchData = useCallback(async () => {
    try {
      const response = await fetch("http://example.com/api");
      const json = await response.json();
      setData(json);

      continueRender(handle);
    } catch (err) {
      cancelRender(err);
    }
  }, [handle]);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  return (
    <div>
      {data ? (
        <div>This video has data from an API! {JSON.stringify(data)}</div>
      ) : null}
    </div>
  );
};
```

Once the data is fetched, you can call [`continueRender()`](/docs/continue-render) to tell Remotion to continue rendering the video.

In case the data fetching fails, you can call [`cancelRender()`](/docs/cancel-render) to cancel the render without waiting for the timeout.

### Time limit [​](\#time-limit-1 "Direct link to Time limit")

You need to clear all handles created by [`delayRender()`](/docs/delay-render) within 30 seconds after the page is opened. You may [increase the timeout](/docs/timeout#increase-timeout).

### Prevent overfetching [​](\#prevent-overfetching "Direct link to Prevent overfetching")

During rendering, multiple headless browser tabs are opened to speed up rendering.

In Remotion Lambda, the [rendering concurrency](/docs/lambda/concurrency) may be up to 200x.

This means that if you are fetching data inside your component, the data fetching will be performed many times.

[1](#1)

Prefer fetching the data before rendering if possible.

[2](#2) Ensure you are entitled to a high request rate without running into a rate limit.

[3](#3) The data returned by the API must be the same on all threads, otherwise [flickering](/docs/flickering) may occur.

[4](#4) Make sure to not have `frame` as a dependency of the `useEffect()`, directly or indirectly, otherwise data will be fetched every frame leading to slowdown and potentially running into rate limits.

## See also [​](\#see-also "Direct link to See also")

- [`delayRender()`](/docs/delay-render)
- [How props get resolved](/docs/props-resolution)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/data-fetching.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Visual editing](/docs/visual-editing) [Next\
\
Variable duration](/docs/dynamic-metadata)

- [Fetching data before the render](#fetching-data-before-the-render)
  - [When to use](#when-to-use)
  - [Usage](#usage)
  - [TypeScript typing](#typescript-typing)
  - [Colocation](#colocation)
  - [Setting the duration based on data](#setting-the-duration-based-on-data)
  - [Aborting stale requests](#aborting-stale-requests)
  - [Debouncing requests](#debouncing-requests)
  - [Time limit](#time-limit)
- [Fetching data during the render](#fetching-data-during-the-render)
  - [When to use](#when-to-use-1)
  - [Usage](#usage-1)
  - [Time limit](#time-limit-1)
  - [Prevent overfetching](#prevent-overfetching)
- [See also](#see-also)
