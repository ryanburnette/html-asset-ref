# html-asset-ref

`htmlassetref` is a Go package designed to update URL references within HTML
content. It provides a simple and convenient way to search for HTML tags
containing attributes such as src, href, or srcset that reference URLs and
modify them based on a callback function.

## Usage

```go
package main

import (
	"fmt"
	htmlassetref "github.com/ryanburnette/html-asset-ref"
)

func main() {
	// Sample HTML content
	html := `<link rel="stylesheet" href="styles.css">
	         <script src="script.js"></script>
	         <img src="image.jpg" alt="Image">`

	// Call UpdateAssetRefs with a callback function to update the references
	updatedHTML := htmlassetref.UpdateAssetRefs(html, func(ref string) string {
		// Modify the reference as needed (e.g., add a version suffix)
		return ref + "?v=1.0"
	})

	// Print the updated HTML
	fmt.Println(updatedHTML)
}

```

## Test

The `_content.html` markup in this project should contain all possible examples
of HTML tags that reference a downloadable asset.

The test ensures that all applicable tags are being found and their correct
references returned and then modified by the callback.

```shell
go test
```
