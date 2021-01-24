package ui

import (
	"github.com/Z-Bolt/OctoScreen/interfaces"
	"github.com/Z-Bolt/OctoScreen/octoprintApis"
	// "github.com/Z-Bolt/OctoScreen/uiWidgets"
)

type customItemsPanel struct {
	CommonPanel
	items			[]octoprintApis.MenuItem
}

func CustomItemsPanel(
	ui				*UI,
	parentPanel		interfaces.IPanel,
	items			[]octoprintApis.MenuItem,
) *customItemsPanel {
	instance := &customItemsPanel {
		CommonPanel: NewCommonPanel(ui, parentPanel),
		items:       items,
	}
	instance.initialize()

	return instance
}

func (this *customItemsPanel) initialize() {
	defer this.Initialize()
	this.arrangeMenuItems(this.grid, this.items, 4)
}
