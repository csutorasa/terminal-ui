package document

import (
	"math"
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

// Creates an empty context.
func NewEmptyRenderContext() RenderContext {
	return NewRenderContext(0, 0)
}

// Creates an infinitely large context.
func NewInfiniteRenderContext() RenderContext {
	return NewRenderContext(math.MaxInt, math.MaxInt)
}

// Gets the width and the height of the context.
func (c *RenderContext) Size() (int, int) {
	return c.w, c.h
}

// Layout calculation context.
type LayoutContext struct {
	RenderContext
	layout map[*Element]RenderContext
}

func NewLayoutContext(rc RenderContext) LayoutContext {
	return LayoutContext{
		RenderContext: rc,
		layout:        map[*Element]RenderContext{},
	}
}

// Adds the layout for an element.
func (lc LayoutContext) Add(e *Element, rc RenderContext) {
	_, ok := lc.layout[e]
	if ok {
		panic("element layout already exists")
	}
	lc.layout[e] = rc
}

func (lc LayoutContext) apply() {
	for e, ctx := range lc.layout {
		e.properties.renderContext.Set(ctx)
	}
}
