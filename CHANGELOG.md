# Changelog

## v1.14.0

The first release of the fork. templ-ui is distributed as a Go module, where
[shadcn-templ](https://github.com/axadrn/shadcn-templ) is distributed by
copying sources into your repository.

Read this as a fork point, not as an upgrade: the version numbers on either
side of it are not comparable, and everything about installing, updating and
serving the components changed between 1.13 and 1.14. Expect to redo your
setup rather than to bump a dependency.

### Breaking

- **The module path is `github.com/kerkenes/templ-ui`.** Every import changes,
  and the `/v2` suffix is gone with the old path: the fork starts its own
  version line.
- **The CLI is gone.** There is no `shadcn-templ init` / `add`, no
  `components.json`, no registry to install from. Components are imported from
  the module, and `go get -u` is what updates them. The one thing the CLI could
  do that importing cannot is let you edit a component's source in place; fork
  the module if you need that.
- **The per-component install instructions are gone** from the docs for the
  same reason. The import line under each component's Usage is the whole story.

### Added

- `assets/css/templ-ui.css`, the compiled stylesheet, shipped in the module and
  served by `assets.Handler()` with `assets.StylesheetURL()` for the `<link>`.
  Projects running their own Tailwind can scan `assets/css/globals.css` in the
  module instead.
- `task build-css` compiles that stylesheet against the components alone. The
  docs-only class safelist moved to the new `assets/css/docs.css`, which is
  what the docs site's Tailwind build now reads.

### Fixed

- The component script bundle came out empty in any project that imported the
  components without setting `GO_ENV=production`: it read the component
  sources from a `components` directory next to the running process, which
  only ever existed because the CLI had copied one there. It now falls back to
  the sources embedded in the module, and no component script goes missing.
