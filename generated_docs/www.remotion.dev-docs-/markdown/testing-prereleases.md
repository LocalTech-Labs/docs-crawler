---
title: Testing prereleases
source: Remotion Documentation
last_updated: 2024-12-22
---

# Testing prereleases

- [Home page](/)
- Contributing
- Testing prereleases

# Testing prereleases

If you have received a prerelease version of Remotion, such as `4.0.1`, this is how you install it:

Replace all packages that are part of Remotion, such as `remotion`, `@remotion/renderer`, `@remotion/lambda`, etc with the version that you have received:

```

- "@remotion/bundler": "4.0.0"
+ "@remotion/bundler": "4.0.1"

- "@remotion/renderer": "4.0.0"
+ "@remotion/renderer": "4.0.1"

- "@remotion/lambda": "4.0.0"
+ "@remotion/lambda": "4.0.1"

- "remotion": "4.0.0"
+ "remotion": "4.0.1"

```

warning

Make sure that you remove the `^` character from the version. If you don't, you get the version with the alphabetically highest hash, which is a random version of Remotion rather than the one you want.

Afterwards, run `yarn`, `npm i` or `pnpm i` respectively.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/prereleases.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Authoring a library](/docs/authoring-packages) [Next\
\
Ineligible for bounties](/docs/contributing/ineligible)
