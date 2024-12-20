package uiWidgets

import (
	"github.com/Z-Bolt/OctoScreen/logger"
)

type AmountToMoveStepButton struct {
	*StepButton
}

func CreateAmountToMoveStepButton() *AmountToMoveStepButton {
	base, err := CreateStepButton(
		1,
		Step{"100mm", "move-step.svg", nil, 100.00},
		Step{"10mm", "move-step.svg", nil, 10.00},
		Step{"1mm", "move-step.svg", nil, 1.00},
		Step{"0.1mm", "move-step.svg", nil, 0.1},
		Step{"0.02mm", "move-step.svg", nil, 0.02},
		Step{"100mm", "move-step.svg", nil, 100.00},
	)
	if err != nil {
		logger.LogError("PANIC!!! - CreateAmountToMoveStepButton()", "CreateStepButton()", err)
		panic(err)
	}

	instance := &AmountToMoveStepButton{
		StepButton: base,
	}

	return instance
}

func (this *AmountToMoveStepButton) Value() float64  {
	return this.StepButton.Value().(float64)
}
