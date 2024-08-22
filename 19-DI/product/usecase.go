package product

type ProductUseCase struct {
	repo ProductRepositoryInterface
}

func NewProductUseCase(repo ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) GetProduct(id int) (Product, error) {
	return uc.repo.GetProduct(id)
}
