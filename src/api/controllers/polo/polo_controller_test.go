package polo

import (
	"github.com/golanshy/golang-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPolo(t *testing.T) {
	resonse := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/marco", nil)

	c := test_utils.GetMockContext(request, resonse)
	Marco(c)

	assert.EqualValues(t, http.StatusOK, resonse.Code)
	assert.EqualValues(t, "polo", resonse.Body.String())

}
