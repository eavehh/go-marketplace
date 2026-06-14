package entity

type Catalog_item struct {
	Base_entity
	Short_description *string
	Full_description  *string
	Image_url         *string
	Brand             *Brand
	Category          *Category
	Price             float64
}
