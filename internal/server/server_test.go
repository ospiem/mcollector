package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockStorage struct{}

func (s *MockStorage) InsertGauge(name string, value float64) error {
	return nil
}

func (s *MockStorage) InsertCounter(name string, value int64) error {
	return nil
}
func (s *MockStorage) SelectGauge(k string) (float64, error) {
	return 0, nil
}
func (s *MockStorage) SelectCounter(k string) (int64, error) {
	return 0, nil
}

//func TestUpdateGaugeHandler(t *testing.T) {
//	// Create a mock storage
//	storage := &MockStorage{}
//
//	// Create a request
//	req := httptest.NewRequest("GET", "/update/gauge/metricName/123.45", nil)
//	w := httptest.NewRecorder()
//
//	// Call the handler
//	handler := updateGauge(storage)
//	handler(w, req)
//
//	// Check the response status code
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
//	}
//
//	// You can add more specific tests here if needed
//}
//
//func TestUpdateCounterHandler(t *testing.T) {
//	// Create a mock storage
//	storage := &MockStorage{}
//
//	// Create a request
//	req := httptest.NewRequest("GET", "/update/counter/metricName/123", nil)
//	w := httptest.NewRecorder()
//
//	// Call the handler
//	handler := updateCounter(storage)
//	handler(w, req)
//
//	// Check the response status code
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
//	}
//
//	// You can add more specific tests here if needed
//}

func TestUpdateGaugeHandlerInvalidURL(t *testing.T) {
	// Create a mock storage
	storage := &MockStorage{}

	// Create a request with an invalid URL
	req := httptest.NewRequest("GET", "/update/metric", nil)
	w := httptest.NewRecorder()

	// Call the handler
	handler := updateGauge(storage)
	handler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
}

func TestUpdateCounterHandlerInvalidValue(t *testing.T) {
	// Create a mock storage
	storage := &MockStorage{}

	// Create a request with an invalid counter value
	req := httptest.NewRequest("GET", "/update/counter/metricName/invalidValue", nil)
	w := httptest.NewRecorder()

	// Call the handler
	handler := updateCounter(storage)
	handler(w, req)

	// Check the response status code, it should return a 400 (Bad Request) error
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
}
