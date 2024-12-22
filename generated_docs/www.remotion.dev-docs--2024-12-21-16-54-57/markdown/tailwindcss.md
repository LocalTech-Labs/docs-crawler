ToolingTailwindCSSOn this pageTailwindCSSUsing the CLI​
The easiest way to get started with Tailwind is by scaffolding a new video using the CLI and selecting a template that supports adding Tailwind automatically.
npmbunpnpmyarnbashnpx create-video@latest
bashnpx create-video@latestbashpnpm create video
bashpnpm create videobashbun create video
bashbun create videobashyarn create video
bashyarn create video
The following templates support scaffolding with TailwindCSS: Hello WorldBlankHello World (JavaScript)AudiogramOverlayStargazerTikTok
Install in existing project​
from v3.3.95, see instructions for older versions

Install @remotion/tailwind package and TailwindCSS dependencies.

npmyarnpnpmbashnpm i -D @remotion/tailwind
bashnpm i -D @remotion/tailwindbashyarn add -D @remotion/tailwind
bashyarn add -D @remotion/tailwindbashpnpm i -D @remotion/tailwind
bashpnpm i -D @remotion/tailwind

Add the Webpack override from @remotion/tailwind to your config file:

remotion.config.tstsimport {Config} from '@remotion/cli/config';import {enableTailwind} from '@remotion/tailwind'; Config.overrideWebpackConfig((currentConfiguration) => {  return enableTailwind(currentConfiguration);});
remotion.config.tstsimport {Config} from '@remotion/cli/config';import {enableTailwind} from '@remotion/tailwind'; Config.overrideWebpackConfig((currentConfiguration) => {  return enableTailwind(currentConfiguration);});
notePrior to v3.3.39, the option was called Config.Bundling.overrideWebpackConfig().


If you use the bundle() or deploySite() Node.JS API, add the Webpack override to it as well.


Create a file src/style.css with the following content:


src/style.csscss@tailwind base;@tailwind components;@tailwind utilities;
src/style.csscss@tailwind base;@tailwind components;@tailwind utilities;

Import the stylesheet in your src/Root.tsx file. Add to the top of the file:

src/Root.tsxjsimport './style.css';
src/Root.tsxjsimport './style.css';

Add a tailwind.config.js file to the root of your project:

tailwind.config.jsjs/* eslint-env node */module.exports = {  content: ['./src/**/*.{ts,tsx}'],  theme: {    extend: {},  },  plugins: [],};
tailwind.config.jsjs/* eslint-env node */module.exports = {  content: ['./src/**/*.{ts,tsx}'],  theme: {    extend: {},  },  plugins: [],};

Ensure your package.json does not have "sideEffects": false set. If it has, declare that CSS files have a side effect:

package.jsondiff{// Only if `"sideEffects": false` exists in your package.json.-  "sideEffects": false+  "sideEffects": ["*.css"]}
package.jsondiff{// Only if `"sideEffects": false` exists in your package.json.-  "sideEffects": false+  "sideEffects": ["*.css"]}

Start using TailwindCSS! You can verify that it's working by adding className="bg-red-900" to any element.

noteYour package manager might show a peer dependency warning. You can safely ignore it or add strict-peer-dependencies=false to your .npmrc to suppress it.
See also​

TailwindCSS v2 (legacy)
Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousIntegration into VueNextEnvironment variablesUsing the CLIInstall in existing projectSee also