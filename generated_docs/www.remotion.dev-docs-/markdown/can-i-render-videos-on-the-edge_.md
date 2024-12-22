---
title: Can I render videos on the edge?
source: Remotion Documentation
last_updated: 2024-12-22
---

# Can I render videos on the edge?

- [Home page](/)
- FAQ
- Rendering on the edge

# Can I render videos on the edge?

Edge runtimes such as Vercel Edge functions and Cloudflare Workers promise to change how we compute in the future.

Unfortunately, it is not possible to render Remotion videos on the “edge” and will most likely not be possible in the foreseeable future due to very heavy constraints being placed in edge functions.

The following constraints currently make it impossible to render videos on the edge:

- _Lack of filesystem APIs_: The `fs` module is not available on the edge.
- _Limited execution duration_: On Vercel, an edge function must respond within 30 seconds, on Cloudflare Workers, a function may only run for 10ms.
- _Code size limit_: On Vercel, the maximum function size in 4MB, on Cloudflare Workers it is 5MB. Remotion currently needs ~150MB of space due to Chrome and FFmpeg being dependencies.
- _Memory limit_: Vercel Edge Functions only have 128MB of RAM available. Remotion currently works best with 2GB+ of RAM.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/render-on-edge.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Vercel Functions / Next.js](/docs/miscellaneous/vercel-functions) [Next\
\
Embedding Remotion Studio](/docs/miscellaneous/embed-studio)
