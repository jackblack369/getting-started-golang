package stage1

type Animal struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。
}

type AnimalCategory struct {
	kingdom string // 界。
}

func (AnimalCategory) string() string {
	return "AnimalCategory"
}

func main() {
	animal := Animal{
		"dog",
		AnimalCategory{
			"a",
		},
	}

	println("animal scientificName is ", animal.scientificName, " and kingdom is ", animal.kingdom)
	println("animal's AnimalCategory is :", animal.AnimalCategory.string())
}
