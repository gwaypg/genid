package snowflake

import "github.com/gwaylib/errors"

const (
	BUFFER_LEN = 100
)

var (
	idBuffer = make(chan ID, BUFFER_LEN)
)

func init() {
	idBuffer <- 0
}

func gen() (ID, error) {
	id := <-idBuffer
	if id > 0 {
		// 响应给本次的请求的值
		return id, nil
	}

	// 若取不到数据，从生成中心中取

	defer func() {
		// 生成本次请求的事件
		idBuffer <- 0
	}()

	ret, err := Gen(&GenArg{BUFFER_LEN})
	if err != nil {
		// 响应给本次的请求的值
		return 0, errors.As(err)
	}

	// 填充缓存
	for _, id := range ret.ID[1:] {
		idBuffer <- ID(id)
	}

	// 响应给本次的请求的值
	return ID(ret.ID[0]), nil
}
