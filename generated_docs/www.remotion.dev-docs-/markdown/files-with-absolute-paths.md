---
title: Files with absolute paths
source: Remotion Documentation
last_updated: 2024-12-22
---

# Files with absolute paths

- [Home page](/)
- FAQ
- Absolute paths

On this page

# Files with absolute paths

Why is it not possible to use absolute paths to require assets?

```

This does not work
tsx

import { Img } from "remotion";

const MyComp: React.FC = () => {
  return <Img src="C://Users/Joe/Documents/image.png" />;
};
```

```

This does not work
tsx

import { Img } from "remotion";

const MyComp: React.FC = () => {
  return <Img src="C://Users/Joe/Documents/image.png" />;
};
```

## Browsers don't have access to your filesystem [​](\#browsers-dont-have-access-to-your-filesystem "Direct link to Browsers don't have access to your filesystem")

Remotion runs in a browser during preview and in a headless browser during rendering.

By default, browsers don't have access to your filesystem. To read any files, they need to be served via an HTTP server.

## Only the `public` folder is supported by web frameworks [​](\#only-the-public-folder-is-supported-by-web-frameworks "Direct link to only-the-public-folder-is-supported-by-web-frameworks")

To allow importing assets, frameworks for making web apps such as Next.js, Vite, Remix or Create React App expose the `public` folder.

We follow this convention in Remotion to allow for seamless code reuse across frameworks - assets that show up in the [Remotion Studio](/docs/studio) also show up in the [Remotion Player](/docs/player).

Allowing assets to be imported from any folder would break this convention, and they could not be displayed in the Remotion Player.

The `public` folder convention

Remotion implements a convention where the `public` folder in your project is exposed via HTTP server and can be accessed on the same domain as your development server.

In Next.js, Remix, and Vite, if you have a file `public/image.png`, you can access it via `"http://localhost:3000/image.png"`, or just `"/image.png"`

In Remotion you need to wrap the path in [`staticFile()`](/docs/staticfile), so it would be `staticFile("image.png")`.

This ensures that assets are loaded correctly if the webpage is not served on the root path `/`, which is the case on [Remotion Lambda](/docs/lambda) and [Remotion Cloud Run](/docs/cloudrun).

By using `staticFile()`, you can import assets from the public folder, and it will automatically also work in Remotion Studio as well as Remix, Vite, Next.js and Create React App in case you are using the [Remotion Player](/docs/player).

## Absolute paths are not portable [​](\#absolute-paths-are-not-portable "Direct link to Absolute paths are not portable")

Before rendering a video, code is bundled to reduce the size of the assets that need to be sent to the browser. Your project with it's `node_modules` folder would be hundreds of megabytes large, but by bundling, the size is reduced drastically.

A bundle can also be uploaded and hosted on the internet, which is what [Remotion Lambda](/docs/lambda) and [Remotion Cloud Run](/docs/cloudrun) do and what is also possible using [regular rendering](/docs/terminology/serve-url).

Because the goal is to reduce the size of the bundle, only the `public` gets copied into the bundle.

Files with an absolute path would not be copied into the bundle, and therefore not be available when rendering on another device.

## Serving the filesystem via HTTP is dangerous [​](\#serving-the-filesystem-via-http-is-dangerous "Direct link to Serving the filesystem via HTTP is dangerous")

While Remotion could theoretically serve the full filesystem via HTTP so it is accessible via browser, this would be a security risk.

Any application or website could then read any file on your computer by making a HTTP request.

## What to do instead [​](\#what-to-do-instead "Direct link to What to do instead")

### Use the `public` folder [​](\#use-the-public-folder "Direct link to use-the-public-folder")

Preferrably, copy files into the `public` folder and [`staticFile()`](/docs/staticfile) to [import assets](/docs/assets).

### Workaround: Use `npx serve` [​](\#workaround-use-npx-serve "Direct link to workaround-use-npx-serve")

Sometimes, it is preferrable to not copy all assets into your Remotion project because they are too large.

In this case, you can use [`npx serve`](https://www.npmjs.com/package/serve) to serve a folder on your computer via HTTP.

```

bash

npx serve --cors C://Users/Joe/Documents
```

```

bash

npx serve --cors C://Users/Joe/Documents
```

You will get a URL that you can use to access the folder via HTTP.

Pass URLs of assets to `<Img>`, `<Video>`, `<OffthreadVideo>`, `<Audio>`, `<Gif>` tags to use them in Remotion.

If you need to spawn a server programmatically, use [`serve-handler`](https://www.npmjs.com/package/serve-handler).

## See also [​](\#see-also "Direct link to See also")

- [Importing assets](/docs/assets)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/absolute-paths.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Embedding Remotion Studio](/docs/miscellaneous/embed-studio) [Next\
\
Live streaming](/docs/miscellaneous/live-streaming)

- [Browsers don't have access to your filesystem](#browsers-dont-have-access-to-your-filesystem)
- [Only the `public` folder is supported by web frameworks](#only-the-public-folder-is-supported-by-web-frameworks)
- [Absolute paths are not portable](#absolute-paths-are-not-portable)
- [Serving the filesystem via HTTP is dangerous](#serving-the-filesystem-via-http-is-dangerous)
- [What to do instead](#what-to-do-instead)
  - [Use the `public` folder](#use-the-public-folder)
  - [Workaround: Use `npx serve`](#workaround-use-npx-serve)
- [See also](#see-also)
