package scalar

import (
	"strings"
	"testing"
)

func TestApiReferenceHTMLUsesEmbeddedDefaultScript(t *testing.T) {
	html, err := ApiReferenceHTML(&Options{SpecContent: `{}`})
	if err != nil {
		t.Fatalf("ApiReferenceHTML() error = %v", err)
	}

	if !strings.Contains(html, "<script>"+embeddedAPIReference+"</script>") {
		t.Fatal("ApiReferenceHTML() did not inline the embedded Scalar script")
	}
	if strings.Contains(html, `<script src="`+DefaultCDN+`"></script>`) {
		t.Fatal("ApiReferenceHTML() still referenced the default remote CDN")
	}
}

func TestApiReferenceHTMLUsesCustomScriptURL(t *testing.T) {
	const customCDN = "https://example.com/scalar.js"

	html, err := ApiReferenceHTML(&Options{
		CDN:         customCDN,
		SpecContent: `{}`,
	})
	if err != nil {
		t.Fatalf("ApiReferenceHTML() error = %v", err)
	}

	if !strings.Contains(html, `<script src="`+customCDN+`"></script>`) {
		t.Fatal("ApiReferenceHTML() did not use the custom script URL")
	}
	if strings.Contains(html, "<script>"+embeddedAPIReference+"</script>") {
		t.Fatal("ApiReferenceHTML() unexpectedly inlined the embedded script for a custom URL")
	}
}
