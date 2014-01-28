package pogo

import (
	"database/sql"
	//"fmt"
	//"reflect"
)

type Rows struct {
	*sql.Rows
}

// func (rs *Rows) Scan(dest ...interface{}) error {
// 	if len(dest) == 1 {
// 		fmt.Println("REFLECT", reflect.ValueOf(dest[0]))
// 	}
// 	return rs.Rows.Scan(dest...)
// }

func (rs *Rows) ScanStruct(dest interface{}) error {
	colMapper := &mapper{}
	cols, err := rs.Columns()
	if err != nil {
		return err
	}
	// Scan values into containers
	cont := make([]interface{}, 0, len(cols))
	for i := 0; i < cap(cont); i++ {
		cont = append(cont, new(interface{}))
	}
	err = rs.Rows.Scan(cont...)
	if err != nil {
		return err
	}
	// Map values
	err = colMapper.unpack(cols, cont, dest)
	if err != nil {
		return err
	}
	return nil
}
