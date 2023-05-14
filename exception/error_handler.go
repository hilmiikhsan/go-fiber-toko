package exception

// func ErrorHandler(ctx *fiber.Ctx, err error) error {
// 	_, validationError := err.(ValidationError)
// 	if validationError {
// 		data := err.Error()
// 		var messages []map[string]interface{}

// 		errJson := json.Unmarshal([]byte(data), &messages)
// 		PanicLogging(errJson)
// 		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
// 			Status: false,
// 			// Message: m,
// 		})
// 	}
// }
