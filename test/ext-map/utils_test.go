package ext_map

import (
	. "github.com/lonchura/go-ext-lib/ext-map"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStringMapContains(t *testing.T) {
	config := map[string]string {
		"[": "]",
		"(": ")",
		"{": "}",
	}

	Convey("StringMapContains should return true when val contain", t, func() {
		var isContain bool

		_, isContain = StringMapContains("]", config)
		So(isContain, ShouldBeTrue)
		_, isContain = StringMapContains(")", config)
		So(isContain, ShouldBeTrue)
		_, isContain = StringMapContains("}", config)
		So(isContain, ShouldBeTrue)
	})
}
