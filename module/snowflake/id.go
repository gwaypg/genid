package snowflake

import (
	"database/sql"
	"database/sql/driver"
	"strconv"

	"github.com/gwaylib/errors"
)

type ID int64

func ParseID(in string) (ID, error) {
	id, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		return 0, errors.As(err, in)
	}
	return ID(id), nil
}

func (u ID) String() string {
	return strconv.FormatInt(int64(u), 10)
}

// for database
func (u *ID) Scan(v interface{}) error {
	val := sql.NullInt64{}
	if err := val.Scan(v); err != nil {
		return errors.As(err)
	}
	*u = ID(val.Int64)
	return nil
}
func (c *ID) Value() (driver.Value, error) {
	return int64(*c), nil
}

// 读取一个新的ID
func New() (ID, error) {
	return gen()
}

// 读取多个新的ID
// 若返回成功，返回的数组长度与请求的长度一致
func NewBuffer(l int) ([]ID, error) {
	if l <= 0 {
		panic("need buffer len > 0")
	}
	result := []ID{}
	for i := l; i > 0; i-- {
		id, err := gen()
		if err != nil {
			return nil, errors.As(err, l)
		}
		result = append(result, id)
	}
	return result, nil
}
