package router

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

// makeRange generates a sequence of numbers
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// TestLimit tests the limit option
func TestLimit(t *testing.T) {
	r := GetRouter()

	// Random seed
	rand.Seed(time.Now().UnixNano())

	// Generate shuffled limits from 1-100
	limits := makeRange(1, 100)
	rand.Shuffle(len(limits), func(i, j int) { limits[i], limits[j] = limits[j], limits[i] })

	for _, limit := range limits {
		w := httptest.NewRecorder()
		params := url.Values{"limit": []string{strconv.Itoa(limit)}}
		req, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("%s?%s", "/api/courses", params.Encode()),
			nil,
		)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		resp := map[string]interface{}{}
		err := json.Unmarshal([]byte(w.Body.String()), &resp)
		assert.Nil(t, err)

		assert.Equal(t, "success", resp["status_message"])
		assert.Equal(t, limit, len(resp["response"].([]interface{})))
	}
}
