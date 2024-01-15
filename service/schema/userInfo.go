package schema

type UserInfo struct {
	// json 校验结构体
	// 在使用 bind 功能时候, 需要给结构体加上标签 ps: json form uri xml yaml
	// 验证信息 binding
	Name     string `json:"name"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserInfo struct {
	// json 校验结构体
	// 在使用 bind 功能时候, 需要给结构体加上标签 ps: json form uri xml yaml
	// 验证信息 binding
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}
