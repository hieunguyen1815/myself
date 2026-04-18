# Module 2: Project Configuration

## `vercel.json`
The optional configuration file at your project root. Controls routing, build settings, headers, redirects, and more.

### Minimal example (Next.js / Node)
```json
{
  "buildCommand": "npm run build",
  "outputDirectory": "dist",
  "installCommand": "npm ci",
  "devCommand": "npm run dev",
  "framework": "nextjs"
}
```

### Go project example
Go projects don't use a framework preset. Vercel treats `.go` files in `/api` as serverless functions automatically.
```json
{
  "functions": {
    "api/**/*.go": {
      "runtime": "go1.x"
    }
  }
}
```
No `buildCommand` or `outputDirectory` needed — Vercel compiles each Go file in `/api` independently.

### Redirects
```json
{
  "redirects": [
    { "source": "/old-path", "destination": "/new-path", "permanent": true }
  ]
}
```

### Rewrites (proxy without changing URL)
```json
{
  "rewrites": [
    { "source": "/api/:path*", "destination": "https://api.example.com/:path*" }
  ]
}
```

### Headers
```json
{
  "headers": [
    {
      "source": "/(.*)",
      "headers": [
        { "key": "X-Content-Type-Options", "value": "nosniff" }
      ]
    }
  ]
}
```

### Environment-specific overrides
Not done in `vercel.json` — use the dashboard or CLI for env vars.

## Environment Variables

### Via Dashboard
Project → Settings → Environment Variables

Set a variable for one or more environments:
- Development
- Preview
- Production

### Via CLI
```bash
vercel env add MY_VAR          # interactive
vercel env add MY_VAR production
vercel env ls
vercel env rm MY_VAR
```

### In code
```js
process.env.MY_VAR             // Node.js / serverless functions
import.meta.env.MY_VAR         // Vite-based frameworks
```

For Next.js, prefix with `NEXT_PUBLIC_` to expose to the browser.

### In Go functions
```go
import "os"

val := os.Getenv("MY_VAR")
```
Environment variables are injected at runtime — no build-time embedding needed.

## Build & Output Settings (Dashboard)
Found under Project → Settings → General:

| Setting | Purpose |
|---------|---------|
| Framework Preset | Tells Vercel which framework to detect/optimize for |
| Build Command | Override default (e.g., `npm run build`) |
| Output Directory | Where built files are (e.g., `dist`, `.next`) |
| Install Command | Override default (`npm install`) |
| Root Directory | For monorepos — subdirectory containing your app |

## Ignored Build Step
Prevent unnecessary builds with a script:
```json
{
  "ignoreCommand": "git diff HEAD^ HEAD --quiet ./src"
}
```
Exit code 1 = run the build. Exit code 0 = skip.

## Vercel CLI Basics
```bash
npm i -g vercel

vercel login
vercel link              # link local dir to a Vercel project
vercel dev               # run locally with Vercel's runtime
vercel                   # deploy (preview by default)
vercel --prod            # deploy to production
vercel ls                # list deployments
vercel inspect [url]     # inspect a deployment
```

## Next
→ [Module 3: GitHub Integration](./03-github-integration.md)
