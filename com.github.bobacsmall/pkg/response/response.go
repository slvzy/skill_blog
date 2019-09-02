package response

import "github.com/kataras/iris"

type Response struct {
	Ctx iris.Context
}

/*func (r *Response) ResponseSuccess(status int, msg string, args ...interface{}) {
	result := make(map[string]interface{})
	result["status"] = status
	result["msg"] = msg

	key := ""

	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)

		default:
			result[key] = arg
		}
	}
	r.Ctx.JSON(result)
	r.Ctx.StopExecution()
	return
}

func (r *Response) ResponseFail(msg string, args ...interface{}) {
	result := make(map[string]interface{})
	result["status"] = 0
	result["msg"] = msg

	key := ""

	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)

		default:
			result[key] = arg
		}
	}
	r.Ctx.JSON(result)
	r.Ctx.StopExecution()
}*/
