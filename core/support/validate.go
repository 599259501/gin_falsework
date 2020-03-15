package support

import (
	"shop_mall/core"
)

func Validate(app *core.App, rules interface{}) {
	err := app.GetContext().ShouldBind(rules)
	if err != nil {
		ThrowException(app.GetContext()).SetCode(400).SetMsg(err.Error()).SetData(nil).Panic()
	}
}
