package components

import (
	"slices"

	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

type Grid struct {
	*ContainerComponent
	columns    *document.Property[[]*gridSize]
	rows       *document.Property[[]*gridSize]
	background *document.Property[ansi.FormattedRune]
}

func NewGrid(element *document.Element) *Grid {
	return &Grid{
		ContainerComponent: NewContainerComponent(element),
		columns:            document.NewSliceProperty([]*gridSize{}),
		rows:               document.NewSliceProperty([]*gridSize{}),
		background:         document.AppendState(element, document.NewPropertyFunc(ansi.NewSimpleRune(' '), ansi.FormattedRuneEquals)),
	}
}

func (c *Creator) NewGrid() *Grid {
	return CreateWithTheme(c, NewGrid)
}

func (g *Grid) ApplyTheme(t *style.Theme) {
	g.SetBackground(t.Background)
}

func (g *Grid) SetBackground(r ansi.FormattedRune) *Grid {
	g.background.Set(r)
	return g
}

func (g *Grid) Focusable() bool {
	return false
}

func (g *Grid) AddColumn(columns int) *Grid {
	if columns < 1 {
		panic("columns should be positive")
	}
	g.columns.Map(func(oldVal []*gridSize) []*gridSize {
		return append(oldVal, &gridSize{
			fixed: columns,
		})
	})
	return g
}

func (g *Grid) AddColumnOfRatio(ratio int) *Grid {
	if ratio < 1 {
		panic("ratio should be positive")
	}
	g.columns.Map(func(oldVal []*gridSize) []*gridSize {
		return append(oldVal, &gridSize{
			ratio: ratio,
		})
	})
	return g
}

func (g *Grid) AddRow(rows int) *Grid {
	if rows < 1 {
		panic("rows should be positive")
	}
	g.rows.Map(func(oldVal []*gridSize) []*gridSize {
		return append(oldVal, &gridSize{
			fixed: rows,
		})
	})
	return g
}

func (g *Grid) AddRowOfRatio(ratio int) *Grid {
	if ratio < 1 {
		panic("ratio should be positive")
	}
	g.rows.Map(func(oldVal []*gridSize) []*gridSize {
		return append(oldVal, &gridSize{
			ratio: ratio,
		})
	})
	return g
}

func (g *Grid) Render(c *document.RenderWriter) {
	rows := g.rows.Value()
	columns := g.columns.Value()
	totalRows := 0
	totalRowRatios := 0
	for _, r := range rows {
		totalRows += r.fixed
		totalRowRatios += r.ratio
	}
	totalColumns := 0
	totalColumnRatios := 0
	for _, col := range columns {
		totalColumns += col.fixed
		totalColumnRatios += col.ratio
	}
	_, h := c.Size()
	restRows := h - totalRows
	ratioRow := 0
	if restRows > 0 {
		ratioRow = restRows / totalRowRatios
	}
	i := 0
	children := slices.Collect(g.Children())
	for _, r := range rows {
		height := max(r.ratio*ratioRow, r.fixed)
		if height > 0 {
			childLines := []document.RenderOutput{}
			for range columns {
				if i >= len(children) {
					continue
				}
				childLines = append(childLines, children[i].RenderFill(g.background.Value()))
				i++
			}
			for i := range height {
				line := ansi.FormattedText{}
				for _, cl := range childLines {
					line = line.Concat(cl[i])
				}
				c.WriteLineFormattedText(line)
			}
		}
		if i >= len(children) {
			break
		}
	}
}

func (g *Grid) Layout(c document.LayoutContext) {
	rows := g.rows.Value()
	columns := g.columns.Value()
	totalRows := 0
	totalRowRatios := 0
	for _, r := range rows {
		totalRows += r.fixed
		totalRowRatios += r.ratio
	}
	totalColumns := 0
	totalColumnRatios := 0
	for _, col := range columns {
		totalColumns += col.fixed
		totalColumnRatios += col.ratio
	}
	w, h := c.Size()
	restRows := h - totalRows
	ratioRow := 0
	if restRows > 0 {
		ratioRow = restRows / totalRowRatios
	}
	restColumns := w - totalColumns
	ratioColumn := 0
	if restColumns > 0 {
		ratioColumn = restColumns / totalColumnRatios
	}
	children := slices.Collect(g.Children())
	i := 0
	for _, r := range rows {
		height := max(r.ratio*ratioRow, r.fixed)
		if height > 0 {
			for _, col := range columns {
				if i >= len(children) {
					break
				}
				width := max(col.ratio*ratioColumn, col.fixed)
				c.Add(children[i], document.NewRenderContext(width, height))
				i++
			}
		}
		if i >= len(children) {
			break
		}
	}
}

func (g *Grid) OnEvent(e *document.Event) {

}

type gridSize struct {
	fixed          int
	ratio          int
	calculatedSize int
}
