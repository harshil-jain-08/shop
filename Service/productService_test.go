package Service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"shop/Models"
	repomock "shop/Repo/mock"
	"shop/dto"
	"testing"
)

func TestCreateProduct_Success(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repomock.NewMockProdRepository(ctrl)
	NewService(repo)

	data := dto.Product{
		Name:     "Arvind",
		Quantity: 1,
		Price:    40,
	}

	expectedModel := Models.Product{
		Name:     data.Name,
		Quantity: data.Quantity,
		Price:    data.Price,
	}

	repo.EXPECT().SaveProduct(&expectedModel).Do(
		func(model *Models.Product) {
			model.ID = 1
		}).Return(nil).Times(1)

	out, err := ProdSer.CreateProduct(&data)
	assertion.Nil(err)
	assertion.Equal(out.Name, expectedModel.Name)
	assertion.Equal(out.Price, expectedModel.Price)
	assertion.Equal(out.Quantity, expectedModel.Quantity)
	assertion.NotEmpty(out.ID)
}
