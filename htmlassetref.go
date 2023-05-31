package htmlassetref

import (
	"regexp"
)

// UpdateAssetRefs updates asset references in the HTML using the provided callback.
func UpdateAssetRefs(html string, callback func(ref string) string) string {
	re := regexp.MustCompile(`(src|href|srcset)=(['"])?([^>\s'"\s]+)(['"\s>])?`)
	return re.ReplaceAllStringFunc(html, func(matched string) string {

		attr := re.ReplaceAllString(matched, "$1")
		beforeRef := re.ReplaceAllString(matched, "$2")
		ref := re.ReplaceAllString(matched, "$3")
		afterRef := re.ReplaceAllString(matched, "$4")

		// fmt.Printf("%v %v %v %v %v\n", matched, attr, beforeRef, ref, afterRef)

		ref = callback(ref)

		// Reconstruct the attribute with the updated reference
		return attr + "=" + beforeRef + ref + afterRef
	})
}
