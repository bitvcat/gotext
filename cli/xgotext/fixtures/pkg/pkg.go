package pkg

import "github.com/bitvcat/gotext"

type SubTranslate struct {
	L gotext.Locale
}

type Translate struct {
	L gotext.Locale
	S SubTranslate
}

func test() {
	gotext.Get("inside sub package")
}
