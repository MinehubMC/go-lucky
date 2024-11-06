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

func getRequest[V any](ctx context.Context, url string, auth string) (*V, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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

func requestWithBody[V any](ctx context.Context, method string, url string, v any, auth string) (*V, error) {
	marshalBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(marshalBytes))
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

func postRequestBody[V any](ctx context.Context, url string, v any, auth string) (*V, error) {
	return requestWithBody[V](ctx, http.MethodPost, url, v, auth)
}

func patchRequestBody[V any](ctx context.Context, url string, v any, auth string) (*V, error) {
	return requestWithBody[V](ctx, http.MethodPatch, url, v, auth)
}

func putRequestBody[V any](ctx context.Context, url string, v any, auth string) (*V, error) {
	return requestWithBody[V](ctx, http.MethodPut, url, v, auth)
}

func deleteRequestBody[V any](ctx context.Context, url string, v any, auth string) (*V, error) {
	return requestWithBody[V](ctx, http.MethodDelete, url, v, auth)
}

func deleteRequest[V any](ctx context.Context, url string, auth string) (*V, error) {
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
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

func deleteRequestNoResponse(ctx context.Context, url string, auth string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
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

	byteValue, _ := io.ReadAll(resp.Body)
	bodyStr := string(byteValue)

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return errors.New(fmt.Sprintf("Request returned %s: %s", resp.Status, bodyStr))
	}

	return nil
}
