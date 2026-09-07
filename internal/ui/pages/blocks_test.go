package pages

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/kerkenes/templ-ui/internal/config"
	"github.com/kerkenes/templ-ui/internal/ctxkeys"
)

func TestBlocksPageExplainsItsBaseUIRegistryReference(t *testing.T) {
	t.Setenv("GO_ENV", "production")
	previousConfig := config.AppConfig
	config.AppConfig = &config.Config{GoEnv: "test"}
	t.Cleanup(func() { config.AppConfig = previousConfig })

	ctx := context.WithValue(context.Background(), ctxkeys.URLPathValue, "/blocks")
	var output bytes.Buffer
	if err := Blocks("").Render(ctx, &output); err != nil {
		t.Fatal(err)
	}
	html := output.String()
	for _, want := range []string{
		"apps/v4/registry/bases/base/blocks",
		"Registry reference",
		"shadcn/ui Base UI blocks",
		"legacy new-york-v4/Radix variants",
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("blocks page is missing registry reference %q", want)
		}
	}
}
