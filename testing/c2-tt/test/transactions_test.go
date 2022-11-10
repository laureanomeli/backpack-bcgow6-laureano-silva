package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/cmd/server/handler"
	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/internal/transactions"
	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/test/mocks"
	"github.com/stretchr/testify/assert"
)

func createServer(mockStore mocks.MockStorage) *gin.Engine {
	_ = os.Setenv("TOKEN", "1234")

	repo := transactions.NewRepository(&mockStore)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)
	r := gin.Default()

	pr := r.Group("/transactions/")
	pr.PUT(":id", t.Update())
	pr.DELETE(":id", t.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "1234")

	return req, httptest.NewRecorder()
}

func Test_UpdateProduct_OK(t *testing.T) {
	mockStorage := mocks.MockStorage{
		DataMock: []transactions.Transaction{{
			Id:       1,
			Code:     "algo",
			Currency: "ars",
			Amount:   100.,
			Emitter:  "uno",
			Receiver: "otro",
			Date:     "hoy",
		},
		},
	}
	r := createServer(mockStorage)

	req, rr := createRequestTest(http.MethodPut, "/transactions/1", `{
        "code": "otra cosa","currency": "usd","amount": 10,"emitter": "alguno", "receiver": "vuh", "date": "mañana"
    }`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_DeleteProduct_OK(t *testing.T) {
	mockStorage := mocks.MockStorage{
		DataMock: []transactions.Transaction{{
			Id:       1,
			Code:     "algo",
			Currency: "ars",
			Amount:   100.,
			Emitter:  "uno",
			Receiver: "otro",
			Date:     "hoy",
		},
			{
				Id:       2,
				Code:     "algo",
				Currency: "ars",
				Amount:   100.,
				Emitter:  "uno",
				Receiver: "otro",
				Date:     "hoy",
			},
		},
	}
	r := createServer(mockStorage)

	req, rr := createRequestTest(http.MethodDelete, "/transactions/1", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
