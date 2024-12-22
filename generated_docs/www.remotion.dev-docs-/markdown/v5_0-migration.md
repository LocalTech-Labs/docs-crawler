---
title: v5.0 Migration
source: Remotion Documentation
last_updated: 2024-12-22
---

# v5.0 Migration

- [Home page](/)
- Migration guides
- v5.0 Migration

On this page

# v5.0 Migration

note

Remotion 5.0 is not yet released. This is an incomplete list of breaking changes that are planned for the release.

## How to upgrade [​](\#how-to-upgrade "Direct link to How to upgrade")

See the [changelog](https://remotion.dev/changelog) to find the latest version.
Upgrade `remotion` and all packages starting with `@remotion` to the latest version, e.g. `5.0.0`:

```

diff

- "remotion": "4.0.141"
- "@remotion/bundler": "4.0.141"
- "@remotion/eslint-config": "4.0.141"
- "@remotion/eslint-plugin": "4.0.141"
- "@remotion/cli": "4.0.141"
- "@remotion/renderer": "4.0.141"
+ "remotion": "5.0.0"
+ "@remotion/bundler": "5.0.0"
+ "@remotion/eslint-config": "5.0.0"
+ "@remotion/eslint-plugin": "5.0.0"
+ "@remotion/cli": "5.0.0"
+ "@remotion/renderer": "5.0.0"
```

```

diff

- "remotion": "4.0.141"
- "@remotion/bundler": "4.0.141"
- "@remotion/eslint-config": "4.0.141"
- "@remotion/eslint-plugin": "4.0.141"
- "@remotion/cli": "4.0.141"
- "@remotion/renderer": "4.0.141"
+ "remotion": "5.0.0"
+ "@remotion/bundler": "5.0.0"
+ "@remotion/eslint-config": "5.0.0"
+ "@remotion/eslint-plugin": "5.0.0"
+ "@remotion/cli": "5.0.0"
+ "@remotion/renderer": "5.0.0"
```

Run `npm i `, `yarn`, `pnpm i` or `bun i` respectively afterwards.

## Runtime requirements [​](\#runtime-requirements "Direct link to Runtime requirements")

The minimum Node version is now 18.0.0. The minimum Bun version is 1.1.3.

## `selectComposition()` and `getCompositions()` now require `inputProps` [​](\#selectcomposition-and-getcompositions-now-require-inputprops "Direct link to selectcomposition-and-getcompositions-now-require-inputprops")

`inputProps` is now required in [`selectComposition()`](/docs/renderer/select-composition) and [`getCompositions()`](/docs/renderer/get-compositions).

A common footgun was the render was not working as intended because the input props were not passed.

**Required action**: Pass an empty object `{}` if you don't have any input props.

## `visualizeAudio()` yields different result [​](\#visualizeaudio-yields-different-result "Direct link to visualizeaudio-yields-different-result")

[`optimizeFor: "speed"`](/docs/visualize-audio#optimizefor) is now the default. This will yield slightly different results.

**Required action**: Review the visualization of your audio. If unsatisfactory, revert to the old behavior by setting `optimizeFor: "accuracy"`.

## TransitionSeries does not support `layout="none"` anymore [​](\#transitionseries-does-not-support-layoutnone-anymore "Direct link to transitionseries-does-not-support-layoutnone-anymore")

Having a [TransitionSeries](/docs/transitions/transitionseries) with `layout="none"` is not supported anymore.

It never made sense to have this prop as transitioned elements need to be positioned absolutely.

**Required action**: Remove the `layout` prop.

## Zod should be upgraded to 3.23.8 [​](\#zod-should-be-upgraded-to-3238 "Direct link to Zod should be upgraded to 3.23.8")

Remotion previously used the types of Zod 3.22.3. With Remotion 5.0, the types of 3.23.8 are used.

**Required action**: If you use Zod, Upgrade Zod to 3.23.8.

## `measureSpring()` does not accept `from` and `to` options anymore [​](\#measurespring-does-not-accept-from-and-to-options-anymore "Direct link to measurespring-does-not-accept-from-and-to-options-anymore")

The values passed in there did not influence the calculation at all. Therefore we removed those options.

**Required action**: Remove the `from` and `to` options from your code.

## `overwrite` is now `true` by default in `renderMediaOnLambda()` [​](\#overwrite-is-now-true-by-default-in-rendermediaonlambda "Direct link to overwrite-is-now-true-by-default-in-rendermediaonlambda")

The default value of `overwrite` has been changed to `true` in `renderMediaOnLambda()`. This skips a check that the file already exists in the S3 bucket, which makes the render start faster.

**Required action**: If you want to keep the old behavior, set `overwrite: false`, explicitly.

## `openBrowser()` now takes a `logLevel` instead of `shouldDumpIo` [​](\#openbrowser-now-takes-a-loglevel-instead-of-shoulddumpio "Direct link to openbrowser-now-takes-a-loglevel-instead-of-shoulddumpio")

The `shouldDumpIo` option has been be removed in 5.0.

Use `logLevel` instead.

## `diskSizeInMb` is now 10240 by default [​](\#disksizeinmb-is-now-10240-by-default "Direct link to disksizeinmb-is-now-10240-by-default")

For Remotion Lambda, the default disk size is now 10240 MB.

This will add a miniscule cost to your renders technically, but will lead to more reliable and faster renders, since Chrome is less likely to run out of disk cache.

**Required actions**:

- If you want to keep the old behavior, set `diskSizeInMb: 2048`, explicitly.
- If your Lambda function name is hardcoded to include `disk2048mb`, unhardcode it and use `speculateFunctionName()` to get the correct name.

## License changes [​](\#license-changes "Direct link to License changes")

Remotion 5.0 has an updated license. View the [license](https://github.com/remotion-dev/remotion/blob/5-0-license/LICENSE.md) here or compare the [changes](https://github.com/remotion-dev/remotion/pull/3750).

Besides wording changes, there are two effective changes in this license:

- Contractors also count towards team size. Previously, a company could only work with contractors and never have to get a company license.
- The company license is bound to our [terms and conditions](https://www.remotion.pro/terms) that will be introduced with Remotion 5.0.

Previously, our terms were generated by a Terms and conditions generator and did not make sense. We wrote the terms and conditions to properly define how we currently handle our licensing business and policies.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/5-0-migration.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
React 19](/docs/react-19) [Next\
\
v4.0 Migration](/docs/4-0-migration)

- [How to upgrade](#how-to-upgrade)
- [Runtime requirements](#runtime-requirements)
- [`selectComposition()` and `getCompositions()` now require `inputProps`](#selectcomposition-and-getcompositions-now-require-inputprops)
- [`visualizeAudio()` yields different result](#visualizeaudio-yields-different-result)
- [TransitionSeries does not support `layout="none"` anymore](#transitionseries-does-not-support-layoutnone-anymore)
- [Zod should be upgraded to 3.23.8](#zod-should-be-upgraded-to-3238)
- [`measureSpring()` does not accept `from` and `to` options anymore](#measurespring-does-not-accept-from-and-to-options-anymore)
- [`overwrite` is now `true` by default in `renderMediaOnLambda()`](#overwrite-is-now-true-by-default-in-rendermediaonlambda)
- [`openBrowser()` now takes a `logLevel` instead of `shouldDumpIo`](#openbrowser-now-takes-a-loglevel-instead-of-shoulddumpio)
- [`diskSizeInMb` is now 10240 by default](#disksizeinmb-is-now-10240-by-default)
- [License changes](#license-changes)
