---
title: Using randomness
source: Remotion Documentation
last_updated: 2024-12-22
---

# Using randomness

- [Home page](/)
- Designing visuals
- Randomness

On this page

# Using randomness

The following thing is an anti-pattern in Remotion:

```

tsx

const MyComp: React.FC = () => {
  const [randomValues] = useState(() =>
    new Array(10).fill(true).map((a, i) => {
      return {
        x: Math.random(),
        y: Math.random(),
      };
    }),
  );
  // Do something with coordinates
  return <></>;
};
```

```

tsx

const MyComp: React.FC = () => {
  const [randomValues] = useState(() =>
    new Array(10).fill(true).map((a, i) => {
      return {
        x: Math.random(),
        y: Math.random(),
      };
    }),
  );
  // Do something with coordinates
  return <></>;
};
```

While this will work during preview, it will break while rendering. The reason is that Remotion is spinning up multiple instances of the webpage to render frames in parallel, and the random values will be different on every instance.

## Fixing the problem [​](\#fixing-the-problem "Direct link to Fixing the problem")

Use the [`random()`](/docs/random) API from Remotion to get deterministic pseudorandom values. Pass in a seed (number or string) and as long as the seed is the same, the return value will be the same.

```

tsx

import { random } from "remotion";
const MyComp: React.FC = () => {
  // No need to use useState
  const randomValues = new Array(10).fill(true).map((a, i) => {
    return {
      x: random(`x-${i}`),
      y: random(`y-${i}`),
    };
  });

  return <></>;
};
```

```

tsx

import { random } from "remotion";
const MyComp: React.FC = () => {
  // No need to use useState
  const randomValues = new Array(10).fill(true).map((a, i) => {
    return {
      x: random(`x-${i}`),
      y: random(`y-${i}`),
    };
  });

  return <></>;
};
```

Now the random values will be the same on all threads.

## False positives [​](\#false-positives "Direct link to False positives")

Did you get an ESLint warning when using `Math.random()`, but you are fully aware of the circumstances described above? Use `random(null)` to get a true random value without getting a warning.

## Exception: Inside `calculateMetadata()` [​](\#exception-inside-calculatemetadata "Direct link to exception-inside-calculatemetadata")

It is safe to use true random values inside [`calculateMetadata()`](/docs/calculate-metadata), as it is only called once and not in parallel.

## See also [​](\#see-also "Direct link to See also")

- [`random()`](/docs/random)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/using-randomness.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Measuring items](/docs/measuring) [Next\
\
Audio visualization](/docs/audio-visualization)

- [Fixing the problem](#fixing-the-problem)
- [False positives](#false-positives)
- [Exception: Inside `calculateMetadata()`](#exception-inside-calculatemetadata)
- [See also](#see-also)
