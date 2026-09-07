---
title: "Installation"
description: "How to add templ-ui to a templ project."
order: 2
---

templ-ui is a Go module. Components are imported from it, not copied into your
repository, and `go get -u` is what updates them.

## Prerequisites

A Go project with [templ](https://templ.guide) set up. Tailwind CSS is
optional: the module ships a compiled stylesheet, and you only need your own
Tailwind build if your app has utilities of its own.

<Steps>

<Step>Add the module</Step>

```shell
go get github.com/kerkenes/templ-ui@latest
```

<Step>Serve the stylesheet, the fonts and the component scripts</Step>

```go
import (
  "github.com/kerkenes/templ-ui/assets"
  "github.com/kerkenes/templ-ui/components"
)

mux.Handle("GET /assets/", assets.Handler())
mux.Handle("GET /components/{bundle}", components.ScriptsHandler())
```

The `/assets/` prefix is not a preference: the stylesheet's `@font-face` rules
address the fonts there.

<Step>Reference them once, in your layout</Step>

```templ
<head>
  <link rel="stylesheet" href={ assets.StylesheetURL() }/>
  @components.Scripts()
</head>
```

<Step>Pick a style</Step>

Components carry `cn-*` classes; the style class on `<body>` decides how they
render. The eight styles are `style-vega`, `style-nova`, `style-maia`,
`style-lyra`, `style-mira`, `style-luma`, `style-sera` and `style-rhea`.

```templ
<body class="style-vega">
```

<Step>Import a component</Step>

```go
import "github.com/kerkenes/templ-ui/components/button"
```

```templ
@button.Button() {
  Click me
}
```

</Steps>

## Running Tailwind yourself

`assets.StylesheetURL()` serves a stylesheet compiled from the components
alone, which is everything the components need and nothing your own markup
uses. If your app has its own utility classes, run Tailwind over both instead
of shipping two stylesheets: point it at the module's `globals.css` and add
the module's components as a scan source.

The module directory is wherever the Go module cache put it, so ask the
toolchain rather than hardcoding a version path:

```yaml
version: "3"

tasks:
  tailwind:
    desc: Watch Tailwind CSS changes
    cmds:
      - |
        TEMPL_UI_PATH="$(go list -mod=mod -m -f {{`'{{.Dir}}'`}} github.com/kerkenes/templ-ui)" && \
        printf '%s\n' \
          "@import \"$TEMPL_UI_PATH/assets/css/globals.css\";" \
          '@source "./**/*.templ";' \
          "@source \"$TEMPL_UI_PATH/components/**/*.templ\";" \
          > ./assets/css/sources.generated.css && \
        tailwindcss -i ./assets/css/globals.css -o ./assets/css/output.css --watch

  templ:
    desc: Run templ with integrated server and hot reload
    cmds:
      - templ generate --watch --proxy="http://localhost:8090" --cmd="go run ./main.go" --open-browser=false

  dev:
    desc: Start development server with hot reload
    cmds:
      - task --parallel tailwind templ
```

Your own `assets/css/globals.css` then imports the generated file, and your
`<link>` points at your `output.css` rather than `assets.StylesheetURL()`. The
fonts still come from `assets.Handler()`.

For custom themes and color palettes, see the [theming docs](/docs/theming).

## JavaScript

Every component's behavior ships as one script bundle, mounted at
`/components/{bundle}` above and loaded by `@components.Scripts()`. There are
no per-component script tags.

The bundle is built from the JavaScript embedded in the module and served with
immutable caching under a content-hashed name. Running inside a checkout of
this repository, where a `components` directory sits next to the process, it is
instead rebuilt from those files on every request so edits hot-reload.

## Component Props

Every component accepts three universal props that are left out of the
per-component API tables:

| Prop         | Type               | Description                                         |
| ------------ | ------------------ | --------------------------------------------------- |
| `ID`         | `string`           | HTML id for the rendered element.                   |
| `Class`      | `string`           | Additional CSS classes, merged with the defaults.    |
| `Attributes` | `templ.Attributes` | Additional HTML attributes spread onto the element. |

Standard HTML behavior (`Disabled`, `Type`, `Href`, ...) works the way the
platform defines it; the API tables only document what a component adds on top.
