//+build wireinject

package mysqldb

import (
	"github.com/google/wire"
)

func InitMySQL() error {
	panic(wire.Build(ReadMySQLConfig, InitMySQLClient))
}

func InitViper() error {
	panic(wire.Build(NewViper))
}
