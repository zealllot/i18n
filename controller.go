package i18n

import (
	"fmt"
	"strings"

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
	key := form.Get("Key")
	locale := form.Get("Locale")
	locale = strings.TrimLeft(locale, "admin_")
	locale = strings.TrimLeft(locale, "supplier_")

	adminValue := utils.HTMLSanitizer.Sanitize(form.Get("AdminValue"))
	supplierValue := utils.HTMLSanitizer.Sanitize(form.Get("SupplierValue"))
	displayId := utils.HTMLSanitizer.Sanitize(form.Get("Id"))
	descriprion := utils.HTMLSanitizer.Sanitize(form.Get("Description"))

	adminTranslation := Translation{Key: key, Locale: "admin_" + locale, Value: adminValue, DisplayId: displayId, Description: descriprion}
	supplierTranslation := Translation{Key: key, Locale: "supplier_" + locale, Value: supplierValue, DisplayId: displayId, Description: descriprion}

	fmt.Println("出现了")
	fmt.Println("adminTranslation", adminTranslation)
	fmt.Println("supplierTranslation", supplierTranslation)

	var hasError bool
	var errMessage string

	err := controller.I18n.SaveTranslation(&adminTranslation)
	if err != nil {
		hasError = true
		errMessage = err.Error()
	}
	err = controller.I18n.SaveTranslation(&supplierTranslation)
	if err != nil {
		hasError = true
		errMessage = errMessage + "\n" + err.Error()
	}

	if !hasError {
		context.Writer.Write([]byte("OK"))
	} else {
		context.Writer.WriteHeader(422)
		context.Writer.Write([]byte(err.Error()))
	}
}
