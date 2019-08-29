package entities

type Category struct {
	TimelessModel
	Nameable

	ParentCategory *Category
	ParentCategoryID *string `gorm:"type:UUID"`
}
