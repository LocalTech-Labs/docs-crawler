Server-side renderingOn this pageServer-Side RenderingRemotion's rendering engine is built with Node.JS, which makes it easy to render a video in the cloud.
Render a video on AWS Lambda​
The easiest and fastest way to render videos in the cloud is to use @remotion/lambda.
Render a video using Node.js APIs​
We provide a set of APIs to render videos using Node.js and Bun.
See an example or the API reference for more information.
Render using GitHub Actions​
The Hello World starter template includes a GitHub Actions workflow file .github/workflows/render-video.yml.
1  Commit the template to a GitHub repository.
2  On GitHub, click the Actions tab.
3  Select the Render video workflow on the left.
4  A Run workflow button should appear. Click it.
5  Fill in the props of the root component and click Run workflow. 
6  After the rendering is finished, you can download the video under Artifacts.
Note that running the workflow may incur costs. However, the workflow will only run if you actively trigger it.
See also: Passing input props in GitHub Actions
Render a video using Docker​
See: Dockerizing a Remotion project
Render a video using GCP Cloud Run (Alpha)​
Check out the experimental Cloud Run package.
Our plan is to port the Lambda runtime to Cloud Run instead of maintaining a separate implementation.
API reference​
getCompositions()List available compositionsselectComposition()Get a compositionrenderMedia()Render a video or audiorenderFrames()Render a series of imagesrenderStill()Render a single imagestitchFramesToVideo()Turn images into a videoopenBrowser()Open a Chrome browser to reuse across rendersensureBrowser()Open a Chrome browser to reuse across rendersmakeCancelSignal()Create token to later cancel a rendergetVideoMetadata()Get metadata from a video file in Node.jsgetSilentParts()Obtain silent portions of a video or audioensureFfmpeg()Check for ffmpeg binary and install if not existingensureFfprobe()Check for ffprobe binary and install if not existinggetCanExtractFramesFast()Probes for fast extraction for <OffthreadVideo>Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousDeploy to a VPSNextOverviewRender a video on AWS LambdaRender a video using Node.js APIsRender using GitHub ActionsRender a video using DockerRender a video using GCP Cloud Run (Alpha)API reference