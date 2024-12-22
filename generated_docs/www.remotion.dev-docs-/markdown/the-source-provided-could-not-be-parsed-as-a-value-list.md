---
title: The source provided could not be parsed as a value list
source: Remotion Documentation
last_updated: 2024-12-22
---

# The source provided could not be parsed as a value list

- [Home page](/)
- Troubleshooting
- Source could not be parsed

# The source provided could not be parsed as a value list

The following error message:

```

The source provided ([...]) could not be parsed as a value list.
```

```

The source provided ([...]) could not be parsed as a value list.
```

can happen when the wrong syntax is passed to the `FontFace` API.

Older versions on Chrome, including Chrome 104 which runs in Remotion Lambda does not support the FontFace syntax without quotes.

```

❌ Without quotes
txt

src: url(font.woff2) format(woff2);
```

```

❌ Without quotes
txt

src: url(font.woff2) format(woff2);
```

```

✅ With quotes
txt

src: url('font.woff2') format('woff2');
```

```

✅ With quotes
txt

src: url('font.woff2') format('woff2');
```

To fix the issue, always include quotes.

Use [`@remotion/fonts`](/docs/fonts-api) to use an abstraction that helps you avoid this mistake.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/troubleshooting/could-not-be-parsed-as-a-value-list.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
defaultProps too big](/docs/troubleshooting/defaultprops-too-big) [Next\
\
Fast Refresh not working](/docs/troubleshooting/broken-fast-refresh)
