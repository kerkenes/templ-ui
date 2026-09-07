# templ-ui

Beautifully designed components built with [templ](https://templ.guide) and
Tailwind CSS, distributed as a Go module: `go get` them, import them, upgrade
them like any other dependency.

templ-ui is a fork of [templUI](https://github.com/templui/templui) v1. templUI
v2 moved to copying component sources into your repository with a CLI; this
fork keeps the library. See [NOTICE](./NOTICE) for the attribution.

![hero](./assets/img/readme.png)

## Install

```bash
go get github.com/kerkenes/templ-ui
```

## Use

```go
import "github.com/kerkenes/templ-ui/components/button"
```

```templ
@button.Button() {
  Click me
}
```

Interactive components load their JavaScript in your layout, and your app
mounts the route those scripts are served from:

```templ
<head>
  @datepicker.Script()
</head>
```

```go
utils.SetupScriptRoutes(mux, isDevelopment)
```

Tailwind scans your templates and the module's components together. The full
setup, including the CSS entry file and the Taskfile that generates the scan
sources, is in the documentation.

## Documentation

https://templ-ui.muratkirazkaya.com/docs/introduction

## Contributing

Please read the [contributing guide](CONTRIBUTING.md).

## License

MIT, see [LICENSE](./LICENSE).
