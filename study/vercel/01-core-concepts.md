# Module 1: Core Concepts

## What is Vercel?
Vercel is a cloud platform for deploying and hosting frontend and fullstack web applications. It is the company behind Next.js and optimized heavily for it, but supports many frameworks (Remix, SvelteKit, Astro, Nuxt, Vite, etc.).

## Core Primitives

### 1. Project
The top-level unit. Linked to a Git repository. Each project has:
- A production deployment (from your main branch)
- Preview deployments (from every PR/branch)
- Environment variables (per environment)
- Domain configuration

### 2. Deployment
A snapshot of your app at a point in time. Immutable — once built, it never changes.

Types:
- **Production** — triggered by push to the production branch (default: `main`)
- **Preview** — triggered by push to any other branch or PR
- **Instant Rollback** — re-promote any previous deployment to production with one click

### 3. Builds
Vercel runs `npm run build` (or your configured build command) in a sandboxed environment.

Build outputs are split into:
- **Static assets** — served from Vercel's global CDN
- **Serverless functions** — deployed as isolated function units
- **Edge functions** — deployed to Vercel's edge network (runs close to user)

### 4. Domains
- Each deployment gets a unique URL (`my-app-abc123.vercel.app`)
- Each project gets a stable preview alias and a production domain
- You can attach custom domains (e.g., `myapp.com`)

### 5. Environment Variables
Scoped to three environments:
- `Development` — used with `vercel dev`
- `Preview` — used in all preview deployments
- `Production` — used in production deployments

Variables can be plain text or encrypted secrets.

### 6. Teams & Organizations
- Personal accounts and team accounts
- Role-based access: Owner, Member, Viewer
- Team-level billing, shared projects

## How a Request is Served

```
User Request
    │
    ▼
Vercel Edge Network (CDN)
    │
    ├── Static asset? → Serve from CDN (fastest)
    │
    ├── Edge Function? → Run at the edge node closest to the user
    │
    └── Serverless Function? → Route to a function instance
```

## Key Terms Cheat Sheet

| Term | Meaning |
|------|---------|
| Deployment | An immutable build artifact |
| Preview URL | Unique URL for a branch/PR deployment |
| Production URL | The live URL for your project |
| Serverless Function | Code that runs on-demand in a managed runtime |
| Edge Function | Code that runs at the CDN edge, near the user |
| Build Cache | Cached node_modules/build output to speed up builds |
| ISR | Incremental Static Regeneration (Next.js specific) |

## Next
→ [Module 2: Project Configuration](./02-project-configuration.md)
