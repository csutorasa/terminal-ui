package document

import (
	"math"
	"slices"
)

// Rendering context
type RenderContext struct {
	w int
	h int
}

func NewRenderContext(w int, h int) RenderContext {
	return RenderContext{
		w: w,
		h: h,
	}
}

// Creates an empty [RenderContext].
func NewEmptyRenderContext() RenderContext {
	return NewRenderContext(0, 0)
}

// Creates an infinitely large [RenderContext].
func NewInfiniteRenderContext() RenderContext {
	return NewRenderContext(math.MaxInt, math.MaxInt)
}

// Gets the width and the height of the [RenderContext].
func (rc *RenderContext) Size() (int, int) {
	return rc.w, rc.h
}

// Gets if [RenderContext] is empty.
func (rc *RenderContext) Empty() bool {
	return rc.w < 1 || rc.h < 1
}

// Layout calculation context.
type LayoutContext struct {
	RenderContext
	layout  map[*Element]RenderContext
	allowed []*Element
}

// Creates a new [LayoutContext].
func NewLayoutContext(rc RenderContext) LayoutContext {
	return LayoutContext{
		RenderContext: rc,
		layout:        map[*Element]RenderContext{},
	}
}

// Adds the layout for an [Element].
func (lc LayoutContext) Add(e *Element, rc RenderContext) {
	if !slices.Contains(lc.allowed, e) {
		panic("element is not children")
	}
	lc.add(e, rc)
}

func (lc LayoutContext) add(e *Element, rc RenderContext) {
	lc.layout[e] = rc
}

func (lc *LayoutContext) setCurrentParent(e *Element) {
	lc.allowed = e.properties.children.Value()
}

func (lc LayoutContext) apply() {
	for e, ctx := range lc.layout {
		e.properties.renderContext.Set(ctx)
	}
}
