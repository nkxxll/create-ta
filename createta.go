package createta

import (
	"strings"

	"github.com/google/uuid"
)

const MODE = 644

type Ctactx struct {
	name      string
	upperName string
	lowerName string
	uuid      uuid.UUID
}

func NewCtactx(name string) *Ctactx {
	return &Ctactx{
		name:      cleanup(name),
		upperName: strings.ToUpper(name),
		lowerName: strings.ToUpper(name),
		uuid:      uuid.New(),
	}
}

func (ctx *Ctactx) Create() {
    // call all the template files
    ctx.WriteAndroid()
    ctx.WriteCmakeList()
}

func cleanup(name string) string {
	// replace all '-',' ' to _ because of C naming convention
	name = strings.Replace(name, "-", "_", -1)
	name = strings.Replace(name, " ", "_", -1)
	return name
}
