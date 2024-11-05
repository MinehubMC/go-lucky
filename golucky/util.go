package golucky

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getRequest[V any](url string, auth string) (*V, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Request error: %s", err)
		return nil, err
	}

	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Do error: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	byteValue, _ := io.ReadAll(resp.Body)

	fmt.Println(string(byteValue))

	var marshalled V
	err = json.Unmarshal(byteValue, &marshalled)
	if err != nil {
		return nil, err
	}

	return &marshalled, nil
}

func postRequestBody[V any](url string, v any, auth string) (*V, error) {
	marshalBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalBytes))
	if err != nil {
		fmt.Printf("Request error: %s", err)
		return nil, err
	}

	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Do error: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	byteValue, _ := io.ReadAll(resp.Body)

	fmt.Println(string(byteValue))

	var marshalled V
	err = json.Unmarshal(byteValue, &marshalled)
	if err != nil {
		return nil, err
	}

	return &marshalled, nil
}

func patchRequestBody[V any](url string, v any, auth string) (*V, error) {
	marshalBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(marshalBytes))
	if err != nil {
		fmt.Printf("Request error: %s", err)
		return nil, err
	}

	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Do error: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	byteValue, _ := io.ReadAll(resp.Body)

	var marshalled V
	err = json.Unmarshal(byteValue, &marshalled)
	if err != nil {
		return nil, err
	}

	return &marshalled, nil
}

func putRequestBody[V any](url string, v any, auth string) (*V, error) {
	marshalBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(marshalBytes))
	if err != nil {
		fmt.Printf("Request error: %s", err)
		return nil, err
	}

	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Do error: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	byteValue, _ := io.ReadAll(resp.Body)

	var marshalled V
	err = json.Unmarshal(byteValue, &marshalled)
	if err != nil {
		return nil, err
	}

	return &marshalled, nil
}

func deleteRequest[V any](url string, auth string) (*V, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Printf("Request error: %s", err)
		return nil, err
	}

	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Do error: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	byteValue, _ := io.ReadAll(resp.Body)

	var marshalled V
	err = json.Unmarshal(byteValue, &marshalled)
	if err != nil {
		return nil, err
	}

	return &marshalled, nil
}

func deleteRequestNoResponse(url string, auth string) error {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Printf("Request error: %s", err)
		return err
	}

	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Do error: %s", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

func deleteRequestBody[V any](url string, v any, auth string) (*V, error) {
	marshalBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(marshalBytes))
	if err != nil {
		fmt.Printf("Request error: %s", err)
		return nil, err
	}

	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Do error: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	byteValue, _ := io.ReadAll(resp.Body)

	var marshalled V
	err = json.Unmarshal(byteValue, &marshalled)
	if err != nil {
		return nil, err
	}

	return &marshalled, nil
}
