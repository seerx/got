package got

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

//UUID 生成 uuid
func UUID() string {
	uid, _ := uuid.NewV4()
	u := uuid.Must(uid, nil)
	return strings.Replace(u.String(), "-", "", -1)
}
