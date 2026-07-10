package document

import (
	"fmt"
)

// Error that happens when the document has no root set.
var ErrDocumentRootMissing error = fmt.Errorf("document root missing")

// Error that happens when the document interacts with an element that was not created by the same document.
var ErrDocumentMismatch error = fmt.Errorf("document mismatch")

// Error that happens when the document has a corrupted DOM.
var ErrCorruptDom error = fmt.Errorf("dom is corrupt")

// Error that happens when invalid operation is attempted on an infinite sized [RenderContext].
var ErrInfiniteRenderContext error = fmt.Errorf("infinite render context")
