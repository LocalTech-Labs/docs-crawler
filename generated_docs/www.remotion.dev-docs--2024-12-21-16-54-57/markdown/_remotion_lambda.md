LambdaOn this page@remotion/lambdaAlso available as a 11min video Integrate Remotion Lambda into your app
Render Remotion videos on AWS Lambda. This is the fastest and most scalable way to render Remotion videos.
You only pay while you are rendering, making it very cost-effective.

When should I use it?​

Your videos are less than 80 minutes long at Full HD. approximately until the 15min AWS Timeout limit is hit
You stay within the (AWS Lambda Concurrency Limit) or you are requesting an increase from AWS.
You are fine with using Amazon Web Services in one of the supported regions.

If one of those constraints is a dealbreaker for you, resort to normal server-side rendering or consider using Cloud Run.
See also: Comparison of server-side rendering options
How it works​

A Lambda function and a S3 bucket is created on AWS.
A Remotion project gets deployed to a S3 bucket as a website.
The Lambda function gets invoked and opens the Remotion project.
A lot of Lambda functions are created in parallel which each render a small part of the video
The initial Lambda function downloads the videos and stitches them together.
The final video gets uploaded to S3 and is available for download.

See in more detail: How Remotion Lambda works
Architecture​

Lambda function: Requires a layer with Chromium, currently hosted by Remotion. Only one Lambda function is required, but it can execute different actions.
S3 bucket: Stores the projects, the renders, and render metadata.
CLI: Allows to control the overall architecture from the command line. Is installed by adding @remotion/lambda to a project.
Node.JS API: Has the same features as the CLI but is easier to use programmatically

Setup / Installation​
See here
Region selection​
The following regions are available for Remotion Lambda:
eu-central-1 eu-west-1 eu-west-2 eu-west-3 eu-south-1 eu-north-1 us-east-1 us-east-2 us-west-1 us-west-2 af-south-1 ap-south-1 ap-east-1 ap-southeast-1 ap-southeast-2 ap-northeast-1 ap-northeast-2 ap-northeast-3 ca-central-1 me-south-1 sa-east-1 
See here for configurations and considerations.
Limitations​

You only have up to 10GB of storage available in a Lambda function. This must be sufficient for both the separated chunks and the concatenated output, therefore the output file can only be about 5GB maximum, limiting the maximum video length to around 2 hours of Full HD video.
Lambda has a global limit of 1000 concurrent lambdas per region by default, although it can be increased.

Cost​
Most of our users render multiple minutes of video for just a few pennies. The exact cost is dependent on the region, assigned memory, type of video, parallelization and other parameters. For each render, we estimate a cost and display it to you. You might also need a Remotion license (see below).
AWS permissions​
Remotion Lambda requires you to create an AWS account and create a user with some permissions attached to it. We require only the minimal amount of permissions required for operating Remotion Lambda.
Read more about permissions
CLI​
You can control Remotion Lambda using the npx remotion lambda command.
Read more about the CLI
Node.JS API​
Everything you can do using the CLI, you can also control using Node.JS APIs. See the reference here.
License​
The standard Remotion license applies. https://github.com/remotion-dev/remotion/blob/main/LICENSE.md
Companies need to buy 1 cloud rendering seat per 2000 renders per month - see https://remotion.pro
Uninstalling​
We make it easy to remove all Remotion resources from your AWS account without leaving any traces or causing further costs.
How to uninstall Remotion LambdaImprove this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousMedia KeysNextSetupWhen should I use it?How it worksArchitectureSetup / InstallationRegion selectionLimitationsCostAWS permissionsCLINode.JS APILicenseUninstalling