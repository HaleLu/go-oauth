package dao

import (
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/HaleLu/go-oauth/model"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_RawUserByNameAndPwd(t *testing.T) {
	Convey("test RawUserByNameAndPwd", t, func() {
		once.Do(startService)
		user, err := d.RawUserByNameAndPwd(ctx, "admin", fmt.Sprintf("%X", md5.Sum([]byte("123456"))))
		So(err, ShouldBeNil)
		So(user, ShouldBeNil)
	})
}

func Test_AddUser(t *testing.T) {
	Convey("test AddUser", t, func() {
		once.Do(startService)
		user := &model.User{
			Username:  "admin",
			Password:  fmt.Sprintf("%X", md5.Sum([]byte("123456"))),
			Nickname:  "小明",
			IsTwoStep: false,
		}
		_, err := d.AddUser(ctx, user)
		So(err, ShouldBeNil)
	})
}

func Test_UpdateUser(t *testing.T) {
	Convey("test UpdateUser", t, func() {
		once.Do(startService)
		user := &model.User{
			ID:        1,
			Password:  fmt.Sprintf("%X", md5.Sum([]byte("1234561"))),
			Nickname:  "小明1",
			IsTwoStep: true,
			Secret:    "",
		}
		_, err := d.UpdateUser(ctx, user)
		So(err, ShouldBeNil)
	})
}
