package repository

import (
	"bytes"
	"go-products-microservice/products-api/models"
	"reflect"
	"testing"
)

func TestCategories_ToJSON(t *testing.T) {
	var tests = []struct {
		name    string
		c       Categories
		wantW   string
		wantErr bool
	}{
		{
			"Successful JSON marshall",
			Categories{{1,"test", "test cat", "testUrl"}},
			"[{\"id\":1,\"name\":\"test\",\"description\":\"test cat\",\"image_url\":\"testUrl\"}]\n",
			false,
		},
		{
			"Successful Empty JSON marshall",
			Categories{{0,"", "", ""}},
			"[{\"id\":0,\"name\":\"\",\"description\":\"\",\"image_url\":\"\"}]\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := tt.c.ToJSON(w)
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

func TestCategoryRepo_AddCategory(t *testing.T) {
	type args struct {
		c *models.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &CategoryRepo{}
			if err := cr.AddCategory(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AddCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCategoryRepo_DeleteCategory(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &CategoryRepo{}
			if err := cr.DeleteCategory(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCategoryRepo_GetCategories(t *testing.T) {
	tests := []struct {
		name string
		want Categories
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &CategoryRepo{}
			if got := cr.GetCategories(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCategories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCategoryRepo_UpdateCategory(t *testing.T) {
	type args struct {
		id int
		c  *models.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &CategoryRepo{}
			if err := cr.UpdateCategory(tt.args.id, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCategoryRepo(t *testing.T) {
	tests := []struct {
		name string
		want *CategoryRepo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCategoryRepo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCategoryRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}