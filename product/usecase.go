package product

type ProductUseCase struct {
	repo ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository}

}

func (uc *ProductUseCase) GetProduct(id int) (Product, error) {
	return uc.repo.GetProduct(id)
}
