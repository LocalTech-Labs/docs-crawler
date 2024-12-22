FAQBrowser renderingOn this pageCan I render videos in the browser?Rendering videos in the browser is not supported. In order to render videos, you need to hook up server-side rendering, Remotion Lambda, or render videos locally.
Will it be supported in the future?​
Currently there is no browser API that allows to capture the viewport. There are some proposed APIs for it.
If such an API gets introduced, we can consider supporting browser rendering in the future.
Why not...​
a Chrome extension?​
Chrome extensions do get the privilege of capturing the viewport. We may explore this in the future, but the combination of asking the user to install an extension and slow render times means we are not prioritizing this feature.
canvas.getImageData()?​
It allows to capture the pixels from a canvas, however Remotion videos can be written with any web technology including HTML and SVG. Canvas elements are just a subset of what's supported in Remotion.
html2canvas?​
It implements it's own rendering engine which only supports a subset of CSS properties.
You would only have access to a very limited feature set.
SVG <foreignElement>?​
You can render HTML markup inside an SVG <foreignElement> and then render that SVG to a canvas and then call getImageData() to turn that into an image.
This is the closest thing to browser rendering, however there are problems with <img> tags and potentially other technologies as well. It's not yet fully out of the questions, but seems hacky so far.
How can I avoid server-rendering?​
We cannot fully avoid server-rendering, but we are taking initiatives to make it easier and efficient:

Remotion Lambda is a batteries-included renderer for Remotion that you only have to pay for when you use it.

See also​

<Player>: Displaying a Remotion video in the browser without encoding it
Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousAI System PromptNextAutomatic durationWill it be supported in the future?Why not...a Chrome extension?canvas.getImageData()?html2canvas?SVG <foreignElement>?How can I avoid server-rendering?See also