---
title: Runtime support
source: Remotion Documentation
last_updated: 2024-12-22
---

# Runtime support

- [Home page](/)
- [Media Parser](/docs/media-parser/)
- Runtime support

On this page

# Runtime support

`@remotion/media-parser` works in the Browser, Node.js and Bun.

However, all of these require some quite modern versions (considered modern as of August 2024).

## For parsing [​](\#for-parsing "Direct link to For parsing")

For parsing a video, the following minimum versions are required:

- Node.js: 20.0.0
- Bun 1.0.0
- Chrome 111
- Edge 111
- Safari 16.4
- Firefox 128

### Feature detection [​](\#feature-detection "Direct link to Feature detection")

Use this to check if the runtime supports `parseMedia()`:

```

tsx

const canUseMediaParser = typeof fetch === 'function' && typeof new ArrayBuffer().resize === 'function';
```

```

tsx

const canUseMediaParser = typeof fetch === 'function' && typeof new ArrayBuffer().resize === 'function';
```

## For using WebCodecs [​](\#for-using-webcodecs "Direct link to For using WebCodecs")

WebCodecs support is not tied to `@remotion/media-parser` itself, but if you use it to extract samples, you need to use the following minimum versions:

- Chrome 94
- Edge 94
- Safari 16.4 - No support for `AudioDecoder`
- Firefox - No support, but in development

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/support.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Extract ID3 tags](/docs/media-parser/tags) [Next\
\
WebCodecs](/docs/media-parser/webcodecs)

- [For parsing](#for-parsing)
  - [Feature detection](#feature-detection)
- [For using WebCodecs](#for-using-webcodecs)
