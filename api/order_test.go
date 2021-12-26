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

func TestListOrdersAPI(t *testing.T) {
	n := 5
	orders := make([]db.Order, n)
	for i := 0; i < n; i++ {
		orders[i] = randomOrder(t)
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
				arg := db.ListOrdersParams{
					Limit:  int32(n),
					Offset: 0,
				}
				store.EXPECT().ListOrders(gomock.Any(), gomock.Eq(arg)).Times(1).Return(orders, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchOrders(t, recorder.Body, orders)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().ListOrders(gomock.Any(), gomock.Any()).Times(1).Return([]db.Order{}, sql.ErrConnDone)
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
				store.EXPECT().ListOrders(gomock.Any(), gomock.Any()).Times(0)
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
				store.EXPECT().ListOrders(gomock.Any(), gomock.Any()).Times(0)
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
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()
			url := "/orders"
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

func randomOrder(t *testing.T) db.Order {
	user, _ := randomUser(t)
	product := randomProduct()
	quantity := util.RandomQuantity()
	price := product.Cost * quantity
	return db.Order{
		ID:        util.RandomInt(1, 1000),
		Owner:     user.Username,
		ProductID: product.ID,
		Quantity:  quantity,
		Price:     price,
	}
}

// func requireBodyMatchOrder(t *testing.T, body *bytes.Buffer, order db.Order) {
// 	data, err := ioutil.ReadAll(body)
// 	require.NoError(t, err)
// 	var gotOrder db.Order
// 	err = json.Unmarshal(data, &gotOrder)
// 	require.NoError(t, err)
// 	require.Equal(t, gotOrder, order)
// }

func requireBodyMatchOrders(t *testing.T, body *bytes.Buffer, orders []db.Order) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotOrders []db.Order
	err = json.Unmarshal(data, &gotOrders)
	require.NoError(t, err)
	require.Equal(t, orders, gotOrders)
}

func TestCreateOrderAPI(t *testing.T) {
	user, _ := randomUser(t)
	product := randomProduct()
	quantity := util.RandomQuantity()
	// correctPrice := product.Cost * quantity
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":   user.Username,
				"product_id": product.ID,
				"quantity":   quantity,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(user, nil)
				store.EXPECT().GetProduct(gomock.Any(), gomock.Eq(product.ID)).Times(1).Return(product, nil)
				arg := db.OrderTxParams{
					User:     user,
					Product:  product,
					Quantity: quantity,
				}
				store.EXPECT().OrderTx(gomock.Any(), gomock.Eq(arg)).Times(1)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NegativeQuantity",
			body: gin.H{
				"username":   user.Username,
				"product_id": product.ID,
				"quantity":   -1,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Any()).Times(0)
				store.EXPECT().GetProduct(gomock.Any(), gomock.Any()).Times(0)
				store.EXPECT().OrderTx(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "OrderTxError",
			body: gin.H{
				"username":   user.Username,
				"product_id": product.ID,
				"quantity":   quantity,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(user, nil)
				store.EXPECT().GetProduct(gomock.Any(), gomock.Eq(product.ID)).Times(1).Return(product, nil)
				store.EXPECT().OrderTx(gomock.Any(), gomock.Any()).Times(1).Return(db.OrderTxResult{}, sql.ErrTxDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()
			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			url := "/orders"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
