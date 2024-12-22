Migration guidesReact 19On this pageReact 19Updating packages​
To use React 19's newest features, you need at least version 4.0.0 of Remotion.
You need to upgrade both react and react-dom:
diff- "react": "18.3.1"- "react-dom": "18.3.1"+ "react": "19.0.0"+ "react-dom": "19.0.0"
diff- "react": "18.3.1"- "react-dom": "18.3.1"+ "react": "19.0.0"+ "react-dom": "19.0.0"
If you use TypeScript, update to the newest types as well:
diff- "@types/react": "18.3.1"- "@types/react-dom": "18.3.1"+ "@types/react": "19.0.0"+ "@types/react-dom": "19.0.0"
diff- "@types/react": "18.3.1"- "@types/react-dom": "18.3.1"+ "@types/react": "19.0.0"+ "@types/react-dom": "19.0.0"
Run npm i, bun i, yarn or pnpm i afterwards, matching your package manager.
Updated templates​
We have updated all templates to use React 19 (exception: React Native Skia).
See the source code of the templates for examples on how to upgrade it to React 19.
HTMLRefElement Type Change​
If you have type errors related to React Refs, upgrade to v4.0.236 of Remotion, where we aligned the types with React 19.
Ecosystem support​
Some libraries that are used with Remotion need upgrading.
React Three Fiber​
Update to 9.0.0-rc.1 and three to 0.171.0.
styled-components​
Update to v6.
Next.js​
Update to Next.js 15.
React Native Skia​
No React 19 support yet. Stay on React 18.Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousList of resourcesNextv5.0 MigrationUpdating packagesUpdated templatesHTMLRefElement Type ChangeEcosystem supportReact Three Fiberstyled-componentsNext.jsReact Native Skia