package vo

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"` // 你还可以给字段指定特定规则的修饰符，如果一个字段用binding:"required"修饰，并且在绑定时该字段的值为空，那么将返回一个错误。
}
