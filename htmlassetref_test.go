package htmlassetref_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ryanburnette/htmlassetref"
)

func TestRefString(t *testing.T) {
	// Read the content from _content.html
	content, err := ioutil.ReadFile("_content.html")
	if err != nil {
		t.Fatalf("failed to read _content.html: %v", err)
	}

	i := 0
	_ = htmlassetref.UpdateAssetRefs(string(content), func(ref string) string {
		// Check that we are getting the correct ref by checking the first ref string
		if i == 0 {
			if !strings.EqualFold(ref, "styles.css") {
				t.Logf("First iteration reference: %s", ref)
			}
		}

		i = i + 1

		// Check if any of the reference strings contain any quotes
		if strings.ContainsAny(ref, `"'`) {
			t.Errorf("Reference contains quotes: %s", ref)
		}

		return ref
	})
}

func TestUpdateAssetRefs(t *testing.T) {
	// Read the original HTML from _content.html
	originalHTML, err := ioutil.ReadFile("_content.html")
	if err != nil {
		t.Fatalf("failed to read _content.html: %v", err)
	}

	// Define the expected modified HTML from _modified.html
	expectedHTML, err := ioutil.ReadFile("_modified.html")
	if err != nil {
		t.Fatalf("failed to read _modified.html: %v", err)
	}

	// Define the callback function to add _x before the extension
	callback := func(ref string) string {
		// Extract the filename and extension using the filepath package
		filename := filepath.Base(ref)
		extension := filepath.Ext(ref)

		// Add _x before the extension
		newFilename := strings.TrimSuffix(filename, extension) + "_x" + extension

		// Reconstruct the reference with the updated filename
		return filepath.Join(filepath.Dir(ref), newFilename)
	}

	// Update the asset references in the original HTML
	modifiedHTML := htmlassetref.UpdateAssetRefs(string(originalHTML), callback)

	// Compare the modified HTML with the expected HTML
	if modifiedHTML != string(expectedHTML) {
		t.Errorf("modified HTML does not match the expected HTML")

		// Print the expected and modified HTML to the console
		fmt.Println("--- Expected HTML ---")
		fmt.Println(string(expectedHTML))
		fmt.Println("--- Modified HTML ---")
		fmt.Println(modifiedHTML)
	}
}
