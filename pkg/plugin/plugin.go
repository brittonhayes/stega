package plugin

import (
	"github.com/urfave/cli/v2"
)

type Encoder interface {
	Encode() cli.ActionFunc
}
