MiscellaneousBun supportOn this pageBun supportRemotion is excited about Bun, and we mostly support it (from v1.0.3).
As a package manager​
You can use bun i to initialize all of our Remotion templates.
To scaffold a new project with bun, use:
bun create video
bun create video
This command sets all scripts to use bunx remotionb which will use Bun as a runtime.
Change remotionb to remotion if you want to use Node.js as a runtime.
Remotion CLI​
If you want to run the Remotion CLI using Bun, use remotionb instead of the remotion command.
It doesn't matter if you prefix remotionb with npx, bunx or another runner command.
npx remotionb render
npx remotionb render
As a runtime​
As of Bun 1.0.24 and Remotion 4.0.88, the following issues are known:

⚠️ The lazyComponent prop on <Composition> and <Player> does not work, and this feature is automatically disabled.
⚠️ A server-side rendering script may not quit automatically after it is done running.

Feel free to file more issues with Remotion if you find them.
Previous issues listed here have been resolved as of Bun 1.0.24.
For contributors​
Start the example testbed using bun run start-bun.
See also​

Deno support
Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024Previous--gl flagNextDeno supportAs a package managerRemotion CLIAs a runtimeFor contributorsSee also