package gitwrapper

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestConfig_GetConfig(t *testing.T) {

	Convey("Test SetConfig", t, func() {

		Convey("Set config should not return an error", func() {

			err := NewConfigurator().SetConfig("user", "john", func(string, ...string) (string, error) {
				return "john", nil
			})

			So(err, ShouldBeNil)
		})

	})
}

func TestConfig_SetConfig(t *testing.T) {

	Convey("Test GetConfig", t, func() {
		Convey("Get Config should return a value", func() {

			value := NewConfigurator().GetConfig("user", func(string, ...string) (string, error) {
				return "name", nil
			})

			So(value, ShouldEqual, "name")
		})
	})

}
