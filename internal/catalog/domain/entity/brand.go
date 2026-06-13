package entity

type Brand struct {
	Base_entity
	Image_url *string // указатель потому что при присваивании значении
	// NULL из sql строкаломается а вот указатель может хранить
}
