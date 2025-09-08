package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/micro-name/internal/constant"
	"github.com/cynx-io/micro-name/internal/dependencies/config"
	"io"
	"net/http"
	"strconv"
)

func HttpRequest[Rp interface{}](ctx context.Context, url string, body interface{}, headers map[string]string, method constant.HttpMethod) (Rp, error) {

	var response Rp

	client := &http.Client{}

	// Convert body to JSON
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return response, err
	}

	var req *http.Request
	if method == constant.GET {
		req, err = http.NewRequest(string(method), url, nil)
		if err != nil {
			return response, err
		}
	} else {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return response, err
		}
		req, err = http.NewRequest(string(method), url, bytes.NewReader(bodyBytes))
		if err != nil {
			return response, err
		}
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	logger.Debug(ctx, "HTTP", method, "request to", url, "with body:", TruncateString(string(bodyBytes), config.Config.Log.MaxLength))
	responseByte, err := client.Do(req)
	if err != nil {
		return response, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error(ctx, "error closing response body on ", url, " :", err)
		}
	}(responseByte.Body)

	if responseByte.StatusCode != http.StatusOK {
		err = errors.New("HTTP " + string(method) + " request failed with status code:" + strconv.Itoa(responseByte.StatusCode))
		logger.Error(ctx, err.Error())
		return response, err
	}

	responseBody, err := io.ReadAll(responseByte.Body)
	if err != nil {
		logger.Error(ctx, "error reading response body on ", url, " :", err)
		return response, err
	}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		logger.Error(ctx, "error unmarshalling response body on ", url, " :", err, " response body:", TruncateString(string(responseBody), config.Config.Log.MaxLength))
		return response, errors.New("error unmarshalling response body: " + err.Error())
	}

	responseStr, _ := json.Marshal(response)
	logger.Debug(ctx, "HTTP ", method, " request to", url, "succeeded with response:", TruncateString(string(responseStr), config.Config.Log.MaxLength))
	return response, nil
}
