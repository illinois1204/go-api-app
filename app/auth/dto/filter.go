package authDto

type Filter struct {
	Param string `form:"param" validate:"required"`
}
