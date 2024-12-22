ContributingGeneral informationOn this pageContributing to RemotionIssues and pull requests of all sorts are welcome!
For bigger projects, please coordinate with Jonny Burger (jonny@remotion.dev, Discord: @jonnyburger) to make sure your changes get merged.
Please note that since we charge for Remotion when companies are using it, this is a commercial project.
By sending pull requests, you agree that we can use your code changes in a commercial context.
Furthermore, also note that you cannot redistribute this project. Please see LICENSE.md for what's allowed and what's not.
This project is released with a Contributor Code of Conduct. By participating in this project you agree to abide by its terms.
Setup​
1  Remotion uses 
pnpm v8.10.2
 as the package manager for development in this repository. We recommend
using Corepack so you get the same version of pnpm as we have. 
shcorepack enable
shcorepack enable
If you don't have corepack, install pnpm@8.10.2 manually:
shnpm i -g pnpm@8.10.2
shnpm i -g pnpm@8.10.2
Prefix with sudo if necessary.
The version must be 8.10.2.
We use Bun to speed up some parts of the build.
Install at least Bun 1.1.7 on your system - see https://bun.sh/ for instructions.
2  Clone the Remotion repository:
shgit clone --depth=1 https://github.com/remotion-dev/remotion.git && cd remotion
shgit clone --depth=1 https://github.com/remotion-dev/remotion.git && cd remotion
noteThe full Git history of Remotion is large. To save time and disk space, we recommend adding --depth=1 to only clone the most recent main branch.
3  Install all dependencies:
shpnpm i
shpnpm i
4  Build the project initially:
shpnpm build
shpnpm build
5  Rebuild whenever a file changes:
shpnpm watch
shpnpm watch
6  You can start making changes!
Testing your changes​
You can start the Testbed using
shcd packages/examplepnpm run dev
shcd packages/examplepnpm run dev
You can render a test video using
shcd packages/examplepnpm exec remotion render
shcd packages/examplepnpm exec remotion render
You can run tests using
shpnpm test
shpnpm test
in either a subpackage to run tests for that package or in the root to run all tests.
Running the @remotion/player testbed​
You can test changes to @remotion/player by starting the Player testbed:
shcd packages/player-examplepnpm run dev
shcd packages/player-examplepnpm run dev
For information about testing, please consult TESTING.md. Issues and pull requests of all sorts are welcome!
Running documentation​
You can run the Docusaurus server that powers our docs using:
shcd packages/docspnpm start
shcd packages/docspnpm start
Running the CLI​
You can test changes to the CLI by moving to packages/examples directory and using pnpm exec to execute the CLI:
shcd packages/examples# Example - Get available compositionspnpm exec remotion compositions# Example - Render commandpnpm exec remotion render ten-frame-tester --output ../../out/video.mp4
shcd packages/examples# Example - Get available compositionspnpm exec remotion compositions# Example - Render commandpnpm exec remotion render ten-frame-tester --output ../../out/video.mp4
Testing Remotion Lambda​
In packages/example, there is a runlambda.sh script that will rebuild the code for the Lambda function, remove any deployed Lambda functions, deploy a new one and render a video.
You need to put you AWS credentials in a .env file of the packages/example directory.
shcd packages/examplesh ./runlambda.sh
shcd packages/examplesh ./runlambda.sh
noteThis will delete any Lambda functions in your account!
Testing Remotion Cloud Run​
In packages/example, there is a runcloudrun.sh script that will rebuild the code for the Cloud Run function, remove any deployed Cloud Run services, deploy a new one and render a video.
You need to put you GCP credentials in a .env file of the packages/example directory.
shcd packages/examplesh ./runcloudrun.sh
shcd packages/examplesh ./runcloudrun.sh
noteThis will delete any Cloud Run services in your account!
Troubleshooting​
If your pnpm build throws errors, oftentimes it is because of caching issues. You can resolve many of these errors by running
tspnpm run clean
tspnpm run clean
in the root directory. Make sure to beforehand kill any pnpm watch commands, as those might regenerate files as you clean them!
Developing Rust parts​
To develop the Rust parts of Remotion, see here.
See also​

Implementing a new feature
Implementing a new option
Writing documentation
How to take a bounty issue
Formatting guidelines
Authoring Remotion libraries
Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousDifference to Motion CanvasNextImplementing a new featureSetupTesting your changesRunning the @remotion/player testbedRunning documentationRunning the CLITesting Remotion LambdaTesting Remotion Cloud RunTroubleshootingDeveloping Rust partsSee also