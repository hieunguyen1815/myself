# Module 5: Functions & Edge

## Two Compute Models

| | Serverless Functions | Edge Functions |
|--|---------------------|----------------|
| Runtime | Node.js, Go, Python, Ruby | V8 isolates (no Node.js APIs) |
| Cold start | ~100–500ms | ~0ms (always warm) |
| Location | Single region (your config) | Every edge node globally |
| Use case | DB queries, auth, heavy logic | Auth checks, redirects, A/B, geo |
| Max duration | 10s (Hobby), 300s (Pro) | 30s |
| File system | Read-only `/tmp` | None |

## Serverless Functions

### In Next.js (App Router)
```ts
// app/api/hello/route.ts
export async function GET(request: Request) {
  return Response.json({ message: "hello" });
}
```

### In Next.js (Pages Router)
```ts
// pages/api/hello.ts
export default function handler(req, res) {
  res.json({ message: "hello" });
}
```

### Standalone (non-Next.js / Node.js)
Place files in `/api` directory:
```ts
// api/hello.ts
export default function handler(req, res) {
  res.send("hello");
}
```

### Configuration (Node.js)
```ts
export const config = {
  runtime: "nodejs",       // or "edge"
  maxDuration: 30,
  regions: ["iad1"],       // deploy to specific region
};
```

## Go Serverless Functions

### How it works
- Place `.go` files inside the `/api` directory
- Each file is compiled and deployed as its own serverless function
- The file must be in `package handler` and export an `Handler` func with the standard `http.HandlerFunc` signature

### Project structure
```
my-project/
├── api/
│   ├── hello.go        → /api/hello
│   └── status.go       → /api/status
├── vercel.json
└── go.mod
```

### Basic handler
```go
// api/hello.go
package handler

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Go on Vercel!")
}
```

### Return JSON
```go
// api/status.go
package handler

import (
    "encoding/json"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "ok",
        "lang":   "go",
    })
}
```

### Read environment variables
```go
import "os"

dbURL := os.Getenv("DATABASE_URL")
```

### Read query params & body
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")   // ?name=foo
    // For POST body:
    // defer r.Body.Close()
    // body, _ := io.ReadAll(r.Body)
    fmt.Fprintf(w, "Hello, %s", name)
}
```

### `vercel.json` for Go
```json
{
  "functions": {
    "api/**/*.go": {
      "runtime": "go1.x"
    }
  }
}
```

### `go.mod` (required)
```
module github.com/yourname/my-project

go 1.21
```
Dependencies in `go.mod` / `go.sum` are installed automatically during Vercel's build.

### Go vs Node.js on Vercel — when to pick Go

| Prefer Go when | Prefer Node.js when |
|---------------|---------------------|
| CPU-bound processing | Using npm ecosystem libraries |
| Low-overhead JSON APIs | Next.js API routes |
| Teams already writing Go | Shared types with a TS frontend |
| Binary size / cold start matters | Rapid prototyping |

## Edge Functions / Middleware

### Next.js Middleware
Runs before every request, at the edge:
```ts
// middleware.ts (project root)
import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export function middleware(request: NextRequest) {
  // Example: redirect unauthenticated users
  const token = request.cookies.get("token");
  if (!token) {
    return NextResponse.redirect(new URL("/login", request.url));
  }
  return NextResponse.next();
}

export const config = {
  matcher: ["/dashboard/:path*"],  // only run on these paths
};
```

### Common Edge Use Cases
- Authentication checks (verify JWT at edge, no DB call)
- Geo-based routing (`request.geo.country`)
- A/B testing (set a cookie, rewrite to variant)
- Bot protection
- Rate limiting (with Vercel KV)

## Logs & Monitoring
- Dashboard → Project → Functions: real-time invocation counts, errors, durations
- Dashboard → Deployment → Logs: streaming build + runtime logs
- Use `console.log` in functions — output appears in Vercel logs

## Storage (Brief Overview)
Vercel offers first-party storage products for use with functions:

| Product | Type | Use case |
|---------|------|---------|
| Vercel KV | Redis (Upstash) | Sessions, rate limiting, flags |
| Vercel Postgres | Postgres (Neon) | Relational data |
| Vercel Blob | Object storage | Images, files, uploads |
| Edge Config | Ultra-fast key-value | Feature flags, config (read at edge) |

## Next
→ [Module 6: Hands-on Exercises](./06-exercises.md)
