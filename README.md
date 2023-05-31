# htmlassetref

`htmlassetref` is a Go package that enables the modification of URL references within HTML content. It offers a straightforward and convenient approach to identify HTML tags with attributes like src, href, or srcset that reference URLs, allowing for easy customization based on a callback function. With `htmlassetref`, developers can efficiently update and manipulate URL references in HTML, facilitating tasks such as rewriting asset URLs, handling path adjustments, or implementing custom transformations.

## Usage

```go
package main

import (
	"fmt"
	"github.com/ryanburnette/htmlassetref"
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

The reliability of this package is maintained through the test. The
`_content.html` markup in this project should contain all possible examples of
HTML tags that reference a downloadable asset. The tests ensure that all
applicable tags are being found and their correct references returned and then
modified by the callback.

```shell
go test
```
