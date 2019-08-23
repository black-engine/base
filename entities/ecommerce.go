package entities

type Category struct {
	TimelessModel
	Nameable

	ParentCategory *Category
	ParentCategoryId *string `gorm:"type:UUID"`
}
