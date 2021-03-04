package repository

import (
	"bytes"
	"go-products-microservice/products-api/models"
	"reflect"
	"testing"
)

func TestProductRepo_AddProduct(t *testing.T) {
	type args struct {
		p *models.Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Successful Add Product",
			 args{&models.Product{ID: 3, Name: "test", Description: "test cat", PriceGBP: 2.00, SKU: "jk", ImageURL: "test.com"}},
			 false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &ProductRepo{}
			if err := pr.AddProduct(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("AddProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductRepo_DeleteProduct(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Successfully Delete Product",
			args{1 },
			false,
		},
		{
			"Failed Delete Product",
			args{-1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &ProductRepo{}
			if err := pr.DeleteProduct(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductRepo_GetProducts(t *testing.T) {
	tests := []struct {
		name string
		want Products
	}{
		{
			"Successfully retrieve list of products",
			productList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &ProductRepo{}
			if got := pr.GetProducts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepo_UpdateProduct(t *testing.T) {
	type args struct {
		id int
		p  *models.Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Successfully update product",
			args{1, &models.Product{1, "update", "update product 1", 2, "LD70", "levee.com"}},
			false,
		},
		{
			"Successfully update product",
			args{10, &models.Product{1, "update", "update product 1", 2, "LD70", "levee.com"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &ProductRepo{}
			if err := pr.UpdateProduct(tt.args.id, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProducts_ToJSON(t *testing.T) {
	tests := []struct {
		name    string
		p       Products
		wantW   string
		wantErr bool
	}{
		{
			"Successful JSON marshall",
			Products{{1,"test", "test cat", 2.00, "tech69", "test.com"}},
			"[{\"id\":1,\"name\":\"test\",\"description\":\"test cat\",\"price\":2,\"sku\":\"jk69\",\"image_url\":\"test.com\"}]\n",
			false,
		},
		{
			"Successful Empty JSON marshall",
			Products{{0,"", "", 0, "", ""}},
			"[{\"id\":0,\"name\":\"\",\"description\":\"\",\"price\":0,\"sku\":\"\",\"image_url\":\"\"}]\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := tt.p.ToJSON(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ToJSON() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
