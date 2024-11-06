package golucky

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func requestNoResponse(ctx context.Context, method string, url string, body any, auth string) error {
	var req *http.Request
	var err error
	if body != nil {
		marshalBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(marshalBytes))
	} else {
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
	}

	if req == nil || err != nil {
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

	byteValue, _ := io.ReadAll(resp.Body)
	bodyStr := string(byteValue)

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return errors.New(fmt.Sprintf("Request returned %s: %s", resp.Status, bodyStr))
	}

	return nil
}

func request[V any](ctx context.Context, method string, url string, body any, auth string) (*V, error) {
	var req *http.Request
	var err error
	if body != nil {
		marshalBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(marshalBytes))
	} else {
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
	}

	if req == nil || err != nil {
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
	bodyStr := string(byteValue)

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return nil, errors.New(fmt.Sprintf("Request returned %s: %s", resp.Status, bodyStr))
	}

	var marshalled V
	err = json.Unmarshal(byteValue, &marshalled)
	if err != nil {
		return nil, err
	}

	return &marshalled, nil
}

func getRequest[V any](ctx context.Context, url string, auth string) (*V, error) {
	return request[V](ctx, http.MethodGet, url, nil, auth)
}

func postRequestBody[V any](ctx context.Context, url string, body any, auth string) (*V, error) {
	return request[V](ctx, http.MethodPost, url, body, auth)
}

func patchRequestBody[V any](ctx context.Context, url string, body any, auth string) (*V, error) {
	return request[V](ctx, http.MethodPatch, url, body, auth)
}
func patchRequestNoResponse(ctx context.Context, url string, body any, auth string) error {
	return requestNoResponse(ctx, http.MethodPatch, url, body, auth)
}

func putRequestNoResponse(ctx context.Context, url string, body any, auth string) error {
	return requestNoResponse(ctx, http.MethodPut, url, body, auth)
}

func deleteRequestNoResponse(ctx context.Context, url string, body any, auth string) error {
	return requestNoResponse(ctx, http.MethodDelete, url, body, auth)
}
