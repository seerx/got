package got

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

//UUID 生成 uuid
func UUID() string {
	u := uuid.Must(uuid.NewV4(), nil)
	return strings.Replace(u.String(), "-", "", -1)
}
