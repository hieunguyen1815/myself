# Module 3: GitHub Integration

## How the Integration Works

Vercel installs a **GitHub App** on your repository. This app:
- Listens for push and pull_request events
- Triggers builds automatically
- Posts deployment status checks back to GitHub PRs
- Creates preview URLs per branch/PR

## Setting Up

### Step 1: Connect GitHub to Vercel
1. Go to vercel.com → New Project
2. Click "Import Git Repository"
3. Authorize the Vercel GitHub App on your account/org
4. Select the repository

### Step 2: Configure the Project
Vercel auto-detects the framework. Review:
- Framework preset
- Build command
- Output directory
- Root directory (for monorepos)

Then click **Deploy**. Your first production build runs.

## How Deployments Are Triggered

| Git Event | Result |
|-----------|--------|
| Push to production branch (`main`) | Production deployment |
| Push to any other branch | Preview deployment |
| Pull request opened/updated | Preview deployment + PR status check |
| PR merged to `main` | New production deployment |

## Preview Deployments

Every branch and PR gets a unique, shareable URL:
```
https://my-app-git-feature-xyz-myteam.vercel.app
```

This URL is:
- Stable per branch (updates on each push to that branch)
- Posted as a status check on the GitHub PR
- Can be shared with teammates or stakeholders for review

## GitHub PR Status Checks

When a PR is opened, Vercel posts:
- A "Vercel" check with a link to the preview URL
- Build logs accessible from the check detail view
- Pass/fail status that can be required before merging (branch protection)

To require the Vercel check before merge:
GitHub repo → Settings → Branches → Branch protection rules → "Require status checks to pass"

## Deploy Hooks
Trigger a deployment from outside GitHub (e.g., a CMS publish event):

1. Project → Settings → Git → Deploy Hooks
2. Create a hook URL for a specific branch
3. POST to that URL to trigger a build:
```bash
curl -X POST "https://api.vercel.com/v1/integrations/deploy/YOUR_HOOK_ID"
```

## Environment Variables Per Branch
You can have preview-specific variables:
- All preview deployments share the `Preview` env vars
- You can also override per branch via the dashboard (Project → Settings → Environment Variables → add a specific git branch)

## Ignored Builds for Monorepos
If only one package changed, skip the build for others:
```json
{
  "ignoreCommand": "npx turbo-ignore"
}
```

## Next
→ [Module 4: End-to-End Feature Shipping](./04-e2e-feature-shipping.md)
