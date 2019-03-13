package i18n

import (
	"fmt"

	"github.com/qor/admin"
	"github.com/qor/qor/utils"
)

type i18nController struct {
	*I18n
}

func (controller *i18nController) Index(context *admin.Context) {
	context.Execute("index", controller.I18n)
}

func (controller *i18nController) Update(context *admin.Context) {
	form := context.Request.Form
	translation := Translation{Key: form.Get("Key"), Locale: form.Get("Locale"), Value: utils.HTMLSanitizer.Sanitize(form.Get("Value")), DisplayId: utils.HTMLSanitizer.Sanitize(form.Get("Id")), Description: utils.HTMLSanitizer.Sanitize(form.Get("Description"))}

	fmt.Println("出现了")
	fmt.Println(translation)

	if err := controller.I18n.SaveTranslation(&translation); err == nil {
		context.Writer.Write([]byte("OK"))
	} else {
		context.Writer.WriteHeader(422)
		context.Writer.Write([]byte(err.Error()))
	}
}
