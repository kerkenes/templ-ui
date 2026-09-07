# templ-ui

shadcn/ui components for Go and [templ](https://templ.guide), distributed as a
Go module: `go get` them, import them, upgrade them like any other dependency.

templ-ui is a fork of [shadcn-templ](https://github.com/axadrn/shadcn-templ),
which ports [shadcn/ui](https://ui.shadcn.com) to templ and hands you the
sources to copy into your own repository. This fork keeps the components and
changes the distribution: they stay in the module, and your project imports
them. See [NOTICE](./NOTICE) for the attribution chain.

## Install

```bash
go get github.com/kerkenes/templ-ui
```

## Use

Import a component package and render it:

```templ
import "github.com/kerkenes/templ-ui/components/button"

templ Page() {
	@button.Button(button.Props{Variant: button.VariantOutline}) {
		Click me
	}
}
```

Components need their stylesheet and their script bundle. Serve both from the
handlers the module ships:

```go
mux.Handle("GET /assets/", assets.Handler())
mux.Handle("GET /components/{bundle}", components.ScriptsHandler())
```

And reference them once, in your layout's `<head>`:

```templ
<link rel="stylesheet" href={ assets.StylesheetURL() }/>
@components.Scripts()
```

The `style-<name>` class on `<body>` picks one of the eight styles
(`style-vega` is the default one the components are drawn against).

`assets.StylesheetURL()` serves a stylesheet compiled from the components
alone. If your project already runs Tailwind, point it at
`assets/css/globals.css` inside the module instead and let your own build
scan both your templates and this module's components.

## Documentation

The component reference lives on the docs site, which is built from this
repository (`task dev`).

## Contributing

Please read the [contributing guide](./CONTRIBUTING.md).

## License

MIT, see [LICENSE](./LICENSE).
