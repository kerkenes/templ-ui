---
title: "How To Use"
description: "Learn how to install templ-ui and use its components."
order: 2
---

## Tools

The documented setup uses these tools in every workflow.

### Go

```shell
go version  # Check if installed
```

> **📝 Note:** Download from [golang.org/dl](https://golang.org/dl) if not installed.

### templ

```shell
go install github.com/a-h/templ/cmd/templ@latest
```

> **📝 Note:** Learn more at [templ.guide](https://templ.guide)

### Tailwind CSS v4.1+

The Tailwind CSS standalone CLI is required:
- Download from [GitHub Releases](https://github.com/tailwindlabs/tailwindcss/releases/latest)
- Or use your package manager

### Task

```shell
go install github.com/go-task/task/v3/cmd/task@latest
```

> **📝 Note:** Learn more at [taskfile.dev](https://taskfile.dev)

## Setup

Components are imported from the module. Nothing is copied into your repository, and `go get -u` is what updates them.

### 1. Add templ-ui

```shell
go get github.com/kerkenes/templ-ui@latest
```

You can also just import a component package and run `go mod tidy`.

### 2. Base Styles

Create `assets/css/input.css`:

This is your Tailwind entry file. Tailwind reads it and writes the compiled result to `assets/css/output.css`.

```css
@import "tailwindcss";
@import "./sources.generated.css";

@custom-variant dark (&:where(.dark, .dark *));

@theme inline {
  --breakpoint-3xl: 1600px;
  --breakpoint-4xl: 2000px;
  --radius-sm: calc(var(--radius) - 4px);
  --radius-md: calc(var(--radius) - 2px);
  --radius-lg: var(--radius);
  --radius-xl: calc(var(--radius) + 4px);
  --color-background: var(--background);
  --color-foreground: var(--foreground);
  --color-card: var(--card);
  --color-card-foreground: var(--card-foreground);
  --color-popover: var(--popover);
  --color-popover-foreground: var(--popover-foreground);
  --color-primary: var(--primary);
  --color-primary-foreground: var(--primary-foreground);
  --color-secondary: var(--secondary);
  --color-secondary-foreground: var(--secondary-foreground);
  --color-muted: var(--muted);
  --color-muted-foreground: var(--muted-foreground);
  --color-accent: var(--accent);
  --color-accent-foreground: var(--accent-foreground);
  --color-destructive: var(--destructive);
  --color-border: var(--border);
  --color-input: var(--input);
  --color-ring: var(--ring);
  --color-sidebar: var(--sidebar);
  --color-sidebar-foreground: var(--sidebar-foreground);
  --color-sidebar-primary: var(--sidebar-primary);
  --color-sidebar-primary-foreground: var(--sidebar-primary-foreground);
  --color-sidebar-accent: var(--sidebar-accent);
  --color-sidebar-accent-foreground: var(--sidebar-accent-foreground);
  --color-sidebar-border: var(--sidebar-border);
  --color-sidebar-ring: var(--sidebar-ring);
}

:root {
  --radius: 0.65rem;
  --background: oklch(1 0 0);
  --foreground: oklch(0.145 0 0);
  --card: oklch(1 0 0);
  --card-foreground: oklch(0.145 0 0);
  --popover: oklch(1 0 0);
  --popover-foreground: oklch(0.145 0 0);
  --primary: oklch(0.205 0 0);
  --primary-foreground: oklch(0.985 0 0);
  --secondary: oklch(0.97 0 0);
  --secondary-foreground: oklch(0.205 0 0);
  --muted: oklch(0.97 0 0);
  --muted-foreground: oklch(0.556 0 0);
  --accent: oklch(0.97 0 0);
  --accent-foreground: oklch(0.205 0 0);
  --destructive: oklch(0.577 0.245 27.325);
  --border: oklch(0.922 0 0);
  --input: oklch(0.922 0 0);
  --ring: oklch(0.708 0 0);
  --sidebar: oklch(0.985 0 0);
  --sidebar-foreground: oklch(0.145 0 0);
  --sidebar-primary: oklch(0.205 0 0);
  --sidebar-primary-foreground: oklch(0.985 0 0);
  --sidebar-accent: oklch(0.97 0 0);
  --sidebar-accent-foreground: oklch(0.205 0 0);
  --sidebar-border: oklch(0.922 0 0);
  --sidebar-ring: oklch(0.708 0 0);
}

.dark {
  --background: oklch(0.145 0 0);
  --foreground: oklch(0.985 0 0);
  --card: oklch(0.205 0 0);
  --card-foreground: oklch(0.985 0 0);
  --popover: oklch(0.205 0 0);
  --popover-foreground: oklch(0.985 0 0);
  --primary: oklch(0.922 0 0);
  --primary-foreground: oklch(0.205 0 0);
  --secondary: oklch(0.269 0 0);
  --secondary-foreground: oklch(0.985 0 0);
  --muted: oklch(0.269 0 0);
  --muted-foreground: oklch(0.708 0 0);
  --accent: oklch(0.269 0 0);
  --accent-foreground: oklch(0.985 0 0);
  --destructive: oklch(0.704 0.191 22.216);
  --border: oklch(1 0 0 / 10%);
  --input: oklch(1 0 0 / 15%);
  --ring: oklch(0.556 0 0);
  --sidebar: oklch(0.205 0 0);
  --sidebar-foreground: oklch(0.985 0 0);
  --sidebar-primary: oklch(0.488 0.243 264.376);
  --sidebar-primary-foreground: oklch(0.985 0 0);
  --sidebar-accent: oklch(0.269 0 0);
  --sidebar-accent-foreground: oklch(0.985 0 0);
  --sidebar-border: oklch(1 0 0 / 10%);
  --sidebar-ring: oklch(0.556 0 0);
}

@layer base {
  * {
    @apply border-border;
    scrollbar-width: thin;
    scrollbar-color: var(--color-muted-foreground) transparent;
  }
  *::-webkit-scrollbar {
    width: 8px;
    height: 8px;
  }
  *::-webkit-scrollbar-thumb {
    background: var(--color-muted-foreground);
    border-radius: 4px;
  }
  *::-webkit-scrollbar-thumb:hover {
    background: var(--color-foreground);
  }
  body {
    @apply bg-background text-foreground;
  }
}
```

> **💡 Tip:** For custom themes and color palettes, visit [/docs/themes](/docs/themes).

### 3. Create Taskfile

```yaml
version: "3"

tasks:
  templ:
    desc: Run templ with integrated server and hot reload
    cmds:
      - templ generate --watch --proxy="http://localhost:8090" --cmd="go run ./main.go" --open-browser=false

  tailwind:
    desc: Watch Tailwind CSS changes
    cmds:
      - |
        TEMPL_UI_PATH="$(go list -mod=mod -m -f {{`'{{.Dir}}'`}} github.com/kerkenes/templ-ui)" && \
        printf '%s\n' \
          '@source "./**/*.templ";' \
          "@source \"$TEMPL_UI_PATH/components/**/*.templ\";" \
          > ./assets/css/sources.generated.css && \
        tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

  dev:
    desc: Start development server with hot reload
    cmds:
      - task --parallel tailwind templ
```

Run everything with:

```shell
task dev
```

### 4. Import and use a component

```go
import "github.com/kerkenes/templ-ui/components/button"
```

```templ
@button.Button() {
  Click me
}
```

### 5. Load JavaScript

Interactive components load JavaScript explicitly in your layout.

```go
import (
  "github.com/kerkenes/templ-ui/components/datepicker"
)
```

```templ
<head>
  @datepicker.Script()
</head>
```

`@datepicker.Script()` loads the `datepicker` script and its direct dependencies like `calendar` and `popover`.

For debugging, you can switch to the unminified scripts during app startup:

```go
func init() {
  utils.UseUnminifiedScripts = true
}
```

```go
func main() {
  utils.UseUnminifiedScripts = true
  // setup routes, then start server
}
```

### 6. Serve Assets

Use `setupAssetsRoutes(...)` to serve your app assets like Tailwind CSS output, fonts, images, and local files. In the import workflow, this is also where you mount templ-ui's embedded component scripts.

```go
func setupAssetsRoutes(mux *http.ServeMux) {
  isDevelopment := os.Getenv("GO_ENV") != "production"

  // Your app assets (CSS, fonts, images, ...)
  assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if isDevelopment {
      w.Header().Set("Cache-Control", "no-store")
    } else {
      w.Header().Set("Cache-Control", "public, max-age=31536000")
    }

    var fs http.Handler
    if isDevelopment {
      fs = http.FileServer(http.Dir("./assets"))
    } else {
      fs = http.FileServer(http.FS(assets.Assets))
    }

    fs.ServeHTTP(w, r)
  })

  mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))

  // templ-ui embedded component scripts
  utils.SetupScriptRoutes(mux, isDevelopment)
}
```

Your Go app must serve `/assets/...` so the browser can load `assets/css/output.css`, fonts, images, and local files. For import-based apps, `utils.SetupScriptRoutes(...)` adds templ-ui's embedded component scripts.
