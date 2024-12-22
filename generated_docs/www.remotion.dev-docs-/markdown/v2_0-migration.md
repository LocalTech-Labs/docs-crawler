---
title: v2.0 Migration
source: Remotion Documentation
last_updated: 2024-12-22
---

# v2.0 Migration

- [Home page](/)
- Migration guides
- v2.0 Migration

On this page

# v2.0 Migration

The following is a list of breaking changes in Remotion 2.0, as a reference for projects wanting to upgrade.

## Sequences are 1 frame shorter [​](\#sequences-are-1-frame-shorter "Direct link to Sequences are 1 frame shorter")

Because of a mistake in v1, sequences were 1 frame too long. The new behavior sees each composition be 1 frame shorter, but be consistent with the duration of compositions.

The behavior of sequences is now the following, as explained by an example: If `durationInFrames` is 60 and `from` is 0, the sequence goes from frame 0 to 59 (60 frames in total), same as a composition with the same duration. In versions 1.x of Remotion, a sequence with the same attributes would go from frame 0 to 60 (61 frames in total).

**Upgrade path**: Check your sequence lengths and if necessary, increase the duration by 1 frame.

## Node.JS API changes [​](\#nodejs-api-changes "Direct link to Node.JS API changes")

- The `userProps` option of `renderFrames` has been renamed to `inputProps`.
- You now need to pass `assetsInfo` to `stitchFramesToVideo` if you want your video to have sound. The `assetsInfo` object will be returned by `renderFrames()`.

**Upgrade path**: See the updated examples on the [SSR page](/docs/ssr) and update your program accordingly.

## `--overwrite` is now default [​](\#--overwrite-is-now-default "Direct link to --overwrite-is-now-default")

If an output already exists, Remotion will overwrite it without asking now, [unless you disable this behavior](/docs/config#setoverwriteoutput).

**Upgrade path**: Do nothing or adjust the setting in the config file if you like.

## Webpack now uses ESBuild instead of Babel [​](\#webpack-now-uses-esbuild-instead-of-babel "Direct link to Webpack now uses ESBuild instead of Babel")

Hopefully you will not notice a difference besides it being much faster. There is a way [to go back to Babel, you can read about it here](/docs/legacy-babel)

**Upgrade path**: Do nothing - should something break, use the legacy Babel plugin and file an issue.

## `react-dom` is a peer dependency [​](\#react-dom-is-a-peer-dependency "Direct link to react-dom-is-a-peer-dependency")

`react-dom` is not anymore pre-installed, so you need to install manually if you upgrade.

## Upgrade to version 2.0 [​](\#upgrade-to-version-20 "Direct link to Upgrade to version 2.0")

Upgrade **all** dependencies containing "remotion" in your package.json to version `^2.0.0`.

```

diff

-"@remotion/bundler": "^1.5.4",
-"@remotion/cli": "^1.5.4",
-"@remotion/eslint-config": "^1.5.4",
-"@remotion/renderer": "^1.5.4",
+"@remotion/bundler": "^2.0.0",
+"@remotion/cli": "^2.0.0",
+"@remotion/eslint-config": "2.0.0",
+"@remotion/renderer": "^2.0.0",
"@types/express": "^4.17.9",
"@types/react": "^17.0.0",
"eslint": "^7.15.0",
"express": "^4.17.1",
"prettier": "^2.2.1",
"react": "^17.0.2",
+"react-dom": "^17.0.2",
-"remotion": "^1.5.4",
+"remotion": "^2.0.0",
```

```

diff

-"@remotion/bundler": "^1.5.4",
-"@remotion/cli": "^1.5.4",
-"@remotion/eslint-config": "^1.5.4",
-"@remotion/renderer": "^1.5.4",
+"@remotion/bundler": "^2.0.0",
+"@remotion/cli": "^2.0.0",
+"@remotion/eslint-config": "2.0.0",
+"@remotion/renderer": "^2.0.0",
"@types/express": "^4.17.9",
"@types/react": "^17.0.0",
"eslint": "^7.15.0",
"express": "^4.17.1",
"prettier": "^2.2.1",
"react": "^17.0.2",
+"react-dom": "^17.0.2",
-"remotion": "^1.5.4",
+"remotion": "^2.0.0",
```

Run `npm install` afterwards.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/2-0-migration.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
v3.0 Migration](/docs/3-0-migration) [Next\
\
Get help](/docs/get-help)

- [Sequences are 1 frame shorter](#sequences-are-1-frame-shorter)
- [Node.JS API changes](#nodejs-api-changes)
- [`--overwrite` is now default](#--overwrite-is-now-default)
- [Webpack now uses ESBuild instead of Babel](#webpack-now-uses-esbuild-instead-of-babel)
- [`react-dom` is a peer dependency](#react-dom-is-a-peer-dependency)
- [Upgrade to version 2.0](#upgrade-to-version-20)
