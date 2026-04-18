# Module 4: End-to-End Feature Shipping

This module walks through the complete lifecycle of shipping a feature from a local branch to production on Vercel.

## The Full Workflow

```
Local branch
    │
    ▼
git push origin feature/my-feature
    │
    ▼
Vercel builds preview deployment
    │
    ▼
GitHub PR opened → Preview URL posted as status check
    │
    ▼
Reviewer tests on preview URL
    │
    ▼
PR approved + merged to main
    │
    ▼
Vercel builds production deployment
    │
    ▼
Production URL updated (zero-downtime)
```

## Step-by-Step

### 1. Local Development
```bash
vercel link          # link your local project to Vercel
vercel env pull      # pull env vars into .env.local
vercel dev           # start local dev server with Vercel runtime
```

### 2. Create a Feature Branch
```bash
git checkout -b feature/my-feature
# make changes...
git add .
git commit -m "feat: add my feature"
git push origin feature/my-feature
```

### 3. Preview Deployment (Automatic)
- Vercel detects the push and starts a build
- Go to your Vercel dashboard → project → the new deployment appears
- A unique URL is generated: `https://my-app-git-feature-my-feature-team.vercel.app`

### 4. Open a Pull Request
- Open a PR on GitHub: `feature/my-feature` → `main`
- Vercel posts a status check with the preview URL
- Team members can click the link and test the feature live

### 5. Review & Iterate
- Push more commits to the branch → preview URL auto-updates
- Use preview URL for QA, design review, stakeholder sign-off

### 6. Merge to Production
```bash
# On GitHub: click "Merge pull request"
```
- Vercel triggers a production build from `main`
- Once complete, your production domain serves the new build
- Previous deployment is retained for instant rollback

### 7. Rollback (if needed)
- Vercel dashboard → project → Deployments
- Find a previous deployment → click "..." → Promote to Production
- Takes effect in seconds — no rebuild needed

## Environment Variables in This Workflow

| Stage | Env used |
|-------|----------|
| `vercel dev` | Development |
| Preview deployment | Preview |
| Production deployment | Production |

Always set sensitive keys (API secrets, DB URLs) via the dashboard, not in code.

## Branch Protection + Vercel (Best Practice)

Set up branch protection on `main`:
1. Require the Vercel preview check to pass
2. Require at least 1 approval
3. This ensures no broken build ever reaches production

## Deployment URL Patterns

| Deployment | URL Pattern |
|------------|-------------|
| Production | `your-domain.com` or `project.vercel.app` |
| Branch (stable) | `project-git-branchname-team.vercel.app` |
| Commit (unique) | `project-abc123xyz.vercel.app` |

## Key Vercel Dashboard Pages to Know

| Page | Path | Purpose |
|------|------|---------|
| Deployments | Project → Deployments | See all builds, promote, rollback |
| Functions | Project → Functions | Monitor serverless function invocations |
| Logs | Deployment → Logs | Build and runtime logs |
| Analytics | Project → Analytics | Web vitals, traffic |
| Settings | Project → Settings | Env vars, domains, git config |

## Next
→ [Module 5: Functions & Edge](./05-functions-and-edge.md)
