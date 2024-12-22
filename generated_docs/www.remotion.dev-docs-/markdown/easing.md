---
title: Easing
source: Remotion Documentation
last_updated: 2024-12-22
---

# Easing

- [Home page](/)
- [remotion](/docs/remotion)
- Easing

On this page

# Easing

The `Easing` module implements common easing functions. You can use it with the [`interpolate()`](/docs/interpolate) API.

You can find a visualization of some common easing functions at [http://easings.net/](http://easings.net/)

### Predefined animations [​](\#predefined-animations "Direct link to Predefined animations")

The `Easing` module provides several predefined animations through the following methods:

- [`back`](/docs/easing#back) provides a basic animation where the object goes slightly back before moving forward
- [`bounce`](/docs/easing#bounce) provides a bouncing animation
- [`ease`](/docs/easing#ease) provides a basic inertial animation
- [`elastic`](/docs/easing#elastic) provides a basic spring interaction

### Standard functions [​](\#standard-functions "Direct link to Standard functions")

Three standard easing functions are provided:

- [`linear`](/docs/easing#linear)
- [`quad`](/docs/easing#quad)
- [`cubic`](/docs/easing#cubic)

The [`poly`](/docs/easing#poly) function can be used to implement quartic, quintic, and other higher power functions.

### Additional functions [​](\#additional-functions "Direct link to Additional functions")

Additional mathematical functions are provided by the following methods:

- [`bezier`](/docs/easing#bezier) provides a cubic bezier curve
- [`circle`](/docs/easing#circle) provides a circular function
- [`sin`](/docs/easing#sin) provides a sinusoidal function
- [`exp`](/docs/easing#exp) provides an exponential function

The following helpers are used to modify other easing functions.

- [`in`](/docs/easing#ineasing) runs an easing function forwards
- [`inOut`](/docs/easing#inout) makes any easing function symmetrical
- [`out`](/docs/easing#out) runs an easing function backwards

## Example [​](\#example "Direct link to Example")

```

tsx

import { Easing, interpolate } from "remotion";

const MyVideo: React.FC = () => {
  const frame = useCurrentFrame();
  const interpolated = interpolate(frame, [0, 100], [0, 1], {
    easing: Easing.bezier(0.8, 0.22, 0.96, 0.65),
    extrapolateLeft: "clamp",
    extrapolateRight: "clamp",
  });
  return (
    <AbsoluteFill
      style={{
        transform: `scale(${interpolated})`,
        backgroundColor: "red",
      }}
    />
  );
};
```

```

tsx

import { Easing, interpolate } from "remotion";

const MyVideo: React.FC = () => {
  const frame = useCurrentFrame();
  const interpolated = interpolate(frame, [0, 100], [0, 1], {
    easing: Easing.bezier(0.8, 0.22, 0.96, 0.65),
    extrapolateLeft: "clamp",
    extrapolateRight: "clamp",
  });
  return (
    <AbsoluteFill
      style={{
        transform: `scale(${interpolated})`,
        backgroundColor: "red",
      }}
    />
  );
};
```

* * *

# Reference

## Methods [​](\#methods "Direct link to Methods")

### `step0` [​](\#step0 "Direct link to step0")

```

jsx

static step0(n): number
```

```

jsx

static step0(n): number
```

A stepping function, returns 1 for any positive value of `n`.

* * *

### `step1` [​](\#step1 "Direct link to step1")

```

jsx

static step1(n): number
```

```

jsx

static step1(n): number
```

A stepping function, returns 1 if `n` is greater than or equal to 1.

* * *

### `linear` [​](\#linear "Direct link to linear")

```

jsx

static linear(t): number
```

```

jsx

static linear(t): number
```

A linear function, `f(t) = t`. Position correlates to elapsed time one to one.

[http://cubic-bezier.com/#0,0,1,1](http://cubic-bezier.com/#0,0,1,1)

* * *

### `ease` [​](\#ease "Direct link to ease")

```

jsx

static ease(t): number
```

```

jsx

static ease(t): number
```

A basic inertial interaction, similar to an object slowly accelerating to speed.

[http://cubic-bezier.com/#.42,0,1,1](http://cubic-bezier.com/#.42,0,1,1)

* * *

### `quad` [​](\#quad "Direct link to quad")

```

jsx

static quad(t): number
```

```

jsx

static quad(t): number
```

A quadratic function, `f(t) = t * t`. Position equals the square of elapsed time.

[http://easings.net/#easeInQuad](http://easings.net/#easeInQuad)

* * *

### `cubic` [​](\#cubic "Direct link to cubic")

```

jsx

static cubic(t): number
```

```

jsx

static cubic(t): number
```

A cubic function, `f(t) = t * t * t`. Position equals the cube of elapsed time.

[http://easings.net/#easeInCubic](http://easings.net/#easeInCubic)

* * *

### `poly()` [​](\#poly "Direct link to poly")

```

jsx

static poly(n): (t) => number
```

```

jsx

static poly(n): (t) => number
```

A power function. Position is equal to the Nth power of elapsed time.

n = 4: [http://easings.net/#easeInQuart](http://easings.net/#easeInQuart) n = 5: [http://easings.net/#easeInQuint](http://easings.net/#easeInQuint)

* * *

### `sin` [​](\#sin "Direct link to sin")

```

jsx

static sin(t): number
```

```

jsx

static sin(t): number
```

A sinusoidal function.

[http://easings.net/#easeInSine](http://easings.net/#easeInSine)

* * *

### `circle` [​](\#circle "Direct link to circle")

```

jsx

static circle(t): number
```

```

jsx

static circle(t): number
```

A circular function.

[http://easings.net/#easeInCirc](http://easings.net/#easeInCirc)

* * *

### `exp` [​](\#exp "Direct link to exp")

```

jsx

static exp(t): number
```

```

jsx

static exp(t): number
```

An exponential function.

[http://easings.net/#easeInExpo](http://easings.net/#easeInExpo)

* * *

### `elastic()` [​](\#elastic "Direct link to elastic")

```

jsx

static elastic(bounciness): (t) =>  number
```

```

jsx

static elastic(bounciness): (t) =>  number
```

A basic elastic interaction, similar to a spring oscillating back and forth.

Default bounciness is 1, which overshoots a little bit once. 0 bounciness doesn't overshoot at all, and bounciness of N > 1 will overshoot about N times.

[http://easings.net/#easeInElastic](http://easings.net/#easeInElastic)

* * *

### `back()` [​](\#back "Direct link to back")

```

jsx

static back(s): (t) => number
```

```

jsx

static back(s): (t) => number
```

Use with `Animated.parallel()` to create a basic effect where the object animates back slightly as the animation starts.

* * *

### `bounce` [​](\#bounce "Direct link to bounce")

```

jsx

static bounce(t): number
```

```

jsx

static bounce(t): number
```

Provides a basic bouncing effect.

[http://easings.net/#easeInBounce](http://easings.net/#easeInBounce)

See an example of how you might use it below

```

jsx

interpolate(0.5, [0, 1], [0, 1], {
  easing: Easing.bounce,
});
```

```

jsx

interpolate(0.5, [0, 1], [0, 1], {
  easing: Easing.bounce,
});
```

* * *

### `bezier()` [​](\#bezier "Direct link to bezier")

```

jsx

static bezier(x1, y1, x2, y2) => (t): number
```

```

jsx

static bezier(x1, y1, x2, y2) => (t): number
```

Provides a cubic bezier curve, equivalent to CSS Transitions' `transition-timing-function`.

A useful tool to visualize cubic bezier curves can be found at [http://cubic-bezier.com/](http://cubic-bezier.com/)

```

jsx

interpolate(0.5, [0, 1], [0, 1], {
  easing: Easing.bezier(0.5, 0.01, 0.5, 1),
});
```

```

jsx

interpolate(0.5, [0, 1], [0, 1], {
  easing: Easing.bezier(0.5, 0.01, 0.5, 1),
});
```

* * *

### `in(easing)` [​](\#ineasing "Direct link to ineasing")

```

jsx

static in(easing: (t: number) => number): (t: number) => number;
```

```

jsx

static in(easing: (t: number) => number): (t: number) => number;
```

Runs an easing function forwards.

```

jsx

{
  easing: Easing.in(Easing.ease);
}
```

```

jsx

{
  easing: Easing.in(Easing.ease);
}
```

* * *

### `out()` [​](\#out "Direct link to out")

```

jsx

static out(easing: (t: number) => number): (t: number) => number;
```

```

jsx

static out(easing: (t: number) => number): (t: number) => number;
```

Runs an easing function backwards.

```

jsx

{
  easing: Easing.out(Easing.ease);
}
```

```

jsx

{
  easing: Easing.out(Easing.ease);
}
```

* * *

### `inOut()` [​](\#inout "Direct link to inout")

```

jsx

static inOut(easing: (t: number) => number): (t: number) => number;
```

```

jsx

static inOut(easing: (t: number) => number): (t: number) => number;
```

```

jsx

{
  easing: Easing.inOut(Easing.ease);
}
```

```

jsx

{
  easing: Easing.inOut(Easing.ease);
}
```

Makes any easing function symmetrical. The easing function will run forwards for half of the duration, then backwards for the rest of the duration.

## Credits [​](\#credits "Direct link to Credits")

The Easing API is the exact same as the one from [React Native](https://reactnative.dev/docs/easing) and the documentation has been copied over. Credit goes to them for this excellent API.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this helper](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/easing.ts)

CONTRIBUTORS

[![kaf-lamed-beyt](https://github.com/kaf-lamed-beyt.png)\
\
**kaf-lamed-beyt** \
\
Improved function signatures](https://github.com/kaf-lamed-beyt)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/easing.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
delayRender()](/docs/delay-render) [Next\
\
<Folder>](/docs/folder)

- [Predefined animations](#predefined-animations)
- [Standard functions](#standard-functions)
- [Additional functions](#additional-functions)
- [Example](#example)
- [Methods](#methods)
  - [`step0`](#step0)
  - [`step1`](#step1)
  - [`linear`](#linear)
  - [`ease`](#ease)
  - [`quad`](#quad)
  - [`cubic`](#cubic)
  - [`poly()`](#poly)
  - [`sin`](#sin)
  - [`circle`](#circle)
  - [`exp`](#exp)
  - [`elastic()`](#elastic)
  - [`back()`](#back)
  - [`bounce`](#bounce)
  - [`bezier()`](#bezier)
  - [`in(easing)`](#ineasing)
  - [`out()`](#out)
  - [`inOut()`](#inout)
- [Credits](#credits)
- [See also](#see-also)
