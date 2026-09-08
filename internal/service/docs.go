package service

import (
	"embed"
	"fmt"
	"path"

	"github.com/kerkenes/templ-ui/internal/markdown"
	"github.com/kerkenes/templ-ui/internal/ui/modules"
)

//go:embed content/docs
var contentFS embed.FS

type DocsService struct {
	parser *markdown.Parser
}

type DocPage struct {
	Slug        string
	Title       string
	Description string
	Order       int
	Content     string
	TOC         []modules.TableOfContentsItem
}

func NewDocsService() *DocsService {
	return &DocsService{
		parser: markdown.NewParser(),
	}
}

// GetPage loads and parses a markdown document by slug
func (s *DocsService) GetPage(slug string) (*DocPage, error) {
	// embed.FS (io/fs) always wants forward slashes -- filepath.Join would
	// use \ on Windows and every lookup would fail with "file does not
	// exist", so path.Join it is.
	mdPath := path.Join("content/docs", slug+".md")

	// Read markdown file from embedded FS
	content, err := contentFS.ReadFile(mdPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read markdown file: %w", err)
	}

	// Parse with frontmatter
	html, meta, err := s.parser.ParseWithFrontmatter(content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse markdown: %w", err)
	}

	// Extract table of contents
	toc := s.parser.ExtractTableOfContents(content)

	// Extract metadata
	page := &DocPage{
		Slug:    slug,
		Content: string(html),
		TOC:     toc,
	}

	if title, ok := meta["title"].(string); ok {
		page.Title = title
	}
	if desc, ok := meta["description"].(string); ok {
		page.Description = desc
	}
	if order, ok := meta["order"].(int); ok {
		page.Order = order
	}

	return page, nil
}
