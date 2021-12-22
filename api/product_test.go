package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	mockdb "order-demo/db/mock"
	db "order-demo/db/sqlc"
	"order-demo/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestListProductsAPI(t *testing.T) {
	n := 5
	products := make([]db.Product, n)
	for i := 0; i < n; i++ {
		products[i] = randomProduct()
	}
	type Query struct {
		pageID   int
		pageSize int
	}
	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListProductsParams{
					Limit:  int32(n),
					Offset: 0,
				}
				store.EXPECT().ListProducts(gomock.Any(), gomock.Eq(arg)).Times(1).Return(products, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProducts(t, recorder.Body, products)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListProducts(gomock.Any(), gomock.Any()).Times(1).Return([]db.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   -1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListProducts(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: Query{
				pageID:   1,
				pageSize: 100000,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListProducts(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			server := NewServer(store)
			recorder := httptest.NewRecorder()
			url := "/products"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomProduct() db.Product {
	return db.Product{
		ID:       util.RandomInt(1, 1000),
		Name:     util.RandomName(),
		Cost:     util.RandomCost(),
		Quantity: util.RandomQuantity(),
	}
}

func requireBodyMatchProduct(t *testing.T, body *bytes.Buffer, product db.Product) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotProduct db.Product
	err = json.Unmarshal(data, &gotProduct)
	require.NoError(t, err)
	require.Equal(t, gotProduct, product)
}

func requireBodyMatchProducts(t *testing.T, body *bytes.Buffer, products []db.Product) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotProducts []db.Product
	err = json.Unmarshal(data, &gotProducts)
	require.NoError(t, err)
	require.Equal(t, products, gotProducts)
}

func TestCreateProductAPI(t *testing.T) {
	product := randomProduct()
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":     product.Name,
				"cost":     product.Cost,
				"quantity": product.Quantity,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateProductParams{
					Name:     product.Name,
					Cost:     product.Cost,
					Quantity: product.Quantity,
				}
				store.EXPECT().CreateProduct(gomock.Any(), gomock.Eq(arg)).Times(1).Return(product, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProduct(t, recorder.Body, product)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"name":     product.Name,
				"cost":     product.Cost,
				"quantity": product.Quantity,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Times(1).Return(db.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidQuantity",
			body: gin.H{
				"name":     product.Name,
				"cost":     product.Cost,
				"quantity": -1,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidCost",
			body: gin.H{
				"name":     product.Name,
				"cost":     -1,
				"quantity": product.Quantity,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			server := NewServer(store)
			recorder := httptest.NewRecorder()
			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			url := "/products"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
