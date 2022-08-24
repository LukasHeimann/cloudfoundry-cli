package i18n

import "github.com/LukasHeimann/cloudfoundrycli/v8/util/ui"

var T ui.TranslateFunc

type LocaleReader interface {
	Locale() string
}

func Init(config LocaleReader) ui.TranslateFunc {
	t, _ := ui.GetTranslationFunc(config)
	return t
}
