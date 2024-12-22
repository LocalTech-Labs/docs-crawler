RenderingOn this pageRender your videoThere are various ways to render your video.
Remotion Studio​
To render a video, click the " Render" button.
Choose your preferred settings, then confirm using the Render video button.
Remotion Studio deployment​
It’s possible to deploy the Remotion Studio onto a long-running server in the cloud, which can then be accessed by your non-technical team members using just a URL. Check out the Deploy the Remotion Studio guide to learn how to do this.
CLI​
Render a video using render CLI command:
bashnpx remotion render HelloWorld
bashnpx remotion render HelloWorld
Modify the composition ID to select a different video to render, or add an output path at the end if you want to override the default.
You can leave out the composition name and a picker will be shown:
bashnpx remotion render
bashnpx remotion render
SSR​
Remotion has a full-featured server-side rendering API. Read more about it on the server-side rendering API.
AWS Lambda​
Check out Remotion Lambda.
GitHub Actions​
You can also render a video using a GitHub Action.
Google Cloud Run​
An alpha version of Remotion for Cloud Run is available.
We plan to change it in the future so it shares a runtime with Remotion Lambda.
Variants​
Audio-only​
Instead of rendering a video, you can also just export the audio.
Image Sequence​
Instead of encoding as a video, you can use the --sequence command to output a series of image.
Still images​
If you want a single image, you can do so using the CLI or Node.JS API.
GIF​
See: Render as GIF
Transparent videos​
See: Creating overlays
See also​

CLI options
Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousHow props get resolvedNextEncoding GuideRemotion StudioRemotion Studio deploymentCLISSRAWS LambdaGitHub ActionsGoogle Cloud RunVariantsAudio-onlyImage SequenceStill imagesGIFTransparent videosSee also