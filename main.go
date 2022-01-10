package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/montanaflynn/stats"
)

// Structure for reading data from the request
type Data struct {
	List []int
	Quantifier int
}

// Structure to write data to response, int or float depending on the output format
type Result struct {
	Ints []int
	Floats []float64
}

// Informational page
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to Math API!</h1>")
}

func avg(w http.ResponseWriter, r *http.Request) {
	var data Data

	// Read data from request and unmarshal
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Calculate average
	avg := Avg(data.List)
	result := Result{nil, []float64{avg}}

	// Marshal result to JSON format
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func median(w http.ResponseWriter, r *http.Request) {
	var data Data

	// Read data from the request and unmarshal
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Calculate median
	inputData := stats.LoadRawData(data.List)
	median, err := stats.Median(inputData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	result := Result{nil, []float64{median}}

	// Marshal result to JSON format
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func percentile(w http.ResponseWriter, r *http.Request) {
	var data Data
	// Read data from the request and unmarshal
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Calculate Percentile using standard library
	inputData := stats.LoadRawData(data.List)
	quant := float64(data.Quantifier)
	percentile, err := stats.PercentileNearestRank(inputData, quant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	result := Result{nil, []float64{percentile}}

	// Marshal result to JSON format
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func min(w http.ResponseWriter, r *http.Request) {
	var data Data
	// Read data from the request and unmarshal
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Calculate min
	minList, err := Min(data.List, data.Quantifier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	result := Result{minList, nil}

	// Marshal result to JSON format
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func max(w http.ResponseWriter, r *http.Request) {
	var data Data
	// Read data from the request and unmarshal
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Calculate max
	minList, err := Max(data.List, data.Quantifier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	result := Result{minList, nil}

	// Marshal result to JSON format
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func custom404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "custom 404 not found")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/avg", avg)
	r.HandleFunc("/median", median)
	r.HandleFunc("/percentile", percentile)
	r.HandleFunc("/min", min)
	r.HandleFunc("/max", max)
	r.NotFoundHandler = http.HandlerFunc(custom404)
	http.ListenAndServe(":3000", r)
}
