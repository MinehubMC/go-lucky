package golucky

import (
	"bytes"
	"context"
	"encoding/json"
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

	var marshalled V
	err = json.Unmarshal(byteValue, &marshalled)
	if err != nil {
		return nil, err
	}

	return &marshalled, nil
}

func postRequestBody[V any](ctx context.Context, url string, v any, auth string) (*V, error) {
	marshalBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(marshalBytes))
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

func patchRequestBody[V any](ctx context.Context, url string, v any, auth string) (*V, error) {
	marshalBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bytes.NewBuffer(marshalBytes))
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

func putRequestBody[V any](ctx context.Context, url string, v any, auth string) (*V, error) {
	marshalBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(marshalBytes))
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

	return nil
}

func deleteRequestBody[V any](ctx context.Context, url string, v any, auth string) (*V, error) {
	marshalBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, bytes.NewBuffer(marshalBytes))
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
