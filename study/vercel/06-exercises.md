# Module 6: Hands-on Exercises

Complete these in order. Each builds on the previous.

---

## Exercise 1: Deploy a Go "Hello World" Function
**Goal:** Experience a first Go deployment end-to-end.

Steps:
1. Create a new project locally:
   ```bash
   mkdir hello-vercel-go && cd hello-vercel-go
   go mod init github.com/yourname/hello-vercel-go
   mkdir api
   ```
2. Create `api/hello.go`:
   ```go
   package handler

   import (
       "fmt"
       "net/http"
   )

   func Handler(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintf(w, "Hello from Go on Vercel!")
   }
   ```
3. Push it to a new GitHub repo
4. Import the repo into Vercel (New Project → Import, no framework preset)
5. Visit `https://<your-project>.vercel.app/api/hello`

**Checkpoint:** You see "Hello from Go on Vercel!" at the live URL.

---

## Exercise 2: Preview Deployment via Branch
**Goal:** Understand preview deployments and branch workflows.

Steps:
1. Create a branch: `git checkout -b feature/update-greeting`
2. Change the response text in `api/hello.go` to `"Hello from the feature branch!"`
3. `git push origin feature/update-greeting`
4. Go to Vercel dashboard — find the preview deployment
5. Open a GitHub PR — verify Vercel posts a status check with the preview URL
6. Visit `https://<preview-url>/api/hello` — confirm the new greeting appears there only

**Checkpoint:** Production still shows old text; preview shows the new greeting.

---

## Exercise 3: Add an Environment Variable
**Goal:** Use env vars in a Go function.

Steps:
1. In Vercel dashboard → Project → Settings → Environment Variables
2. Add `GREETING=Hello from Vercel Go` for all environments
3. Update `api/hello.go` to read it:
   ```go
   package handler

   import (
       "fmt"
       "net/http"
       "os"
   )

   func Handler(w http.ResponseWriter, r *http.Request) {
       greeting := os.Getenv("GREETING")
       if greeting == "" {
           greeting = "Hello (no env var set)"
       }
       fmt.Fprintf(w, greeting)
   }
   ```
4. Push to your branch → verify the greeting appears in preview
5. Merge to main → verify it appears in production

**Checkpoint:** The env var value is returned by the live Go function.

---

## Exercise 4: Ship a Feature End-to-End
**Goal:** Complete the full GitHub → PR → review → merge → production cycle using Go.

Steps:
1. Create a branch `feature/status-endpoint`
2. Add a new Go function `api/status.go`:
   ```go
   package handler

   import (
       "encoding/json"
       "net/http"
       "os"
   )

   func Handler(w http.ResponseWriter, r *http.Request) {
       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(map[string]string{
           "status": "ok",
           "region": os.Getenv("VERCEL_REGION"),
       })
   }
   ```
3. Push and open a PR
4. Test on the preview URL: `https://<preview-url>/api/status`
5. Merge the PR
6. Verify it works on production

**Checkpoint:** `/api/status` returns JSON and is live on your production domain.

---

## Exercise 5: Rollback a Deployment
**Goal:** Practice instant rollback.

Steps:
1. Deploy a breaking change to `main` (e.g., throw an error in a component)
2. Go to Vercel dashboard → Deployments
3. Find the last known-good deployment
4. Click "..." → "Promote to Production"
5. Verify production is restored without a rebuild

**Checkpoint:** Production is back to the good state in under 30 seconds.

---

## Exercise 6: Token-Protected Go Endpoint (Stretch)
**Goal:** Protect a Go function with a bearer token check.

Steps:
1. Add `API_TOKEN=secret123` as an env var in Vercel (Preview + Production)
2. Create `api/protected.go`:
   ```go
   package handler

   import (
       "net/http"
       "os"
       "strings"
   )

   func Handler(w http.ResponseWriter, r *http.Request) {
       token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
       if token != os.Getenv("API_TOKEN") {
           http.Error(w, "Unauthorized", http.StatusUnauthorized)
           return
       }
       w.Write([]byte("Access granted"))
   }
   ```
3. Push and deploy
4. Test without a token:
   ```bash
   curl https://<your-project>.vercel.app/api/protected
   # → 401 Unauthorized
   ```
5. Test with the token:
   ```bash
   curl -H "Authorization: Bearer secret123" https://<your-project>.vercel.app/api/protected
   # → Access granted
   ```

**Checkpoint:** The endpoint returns 401 without the token and 200 with it.

---

## Study Complete
After finishing these exercises you will have:
- Deployed a real Go serverless function on Vercel
- Used preview deployments and PR checks with a Go project
- Read environment variables inside a Go handler
- Shipped a feature through the full GitHub → Vercel pipeline using Go
- Practiced instant rollback
- (Stretch) Protected a Go endpoint with bearer token auth
