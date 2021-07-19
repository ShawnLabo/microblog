package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	metadataServerHost = "metadata.google.internal"
	metadataRegionPath = "/computeMetadata/v1/instance/region"
)

type metadata struct {
	region string
}

func getMetadata(ctx context.Context) (*metadata, error) {
	c := &http.Client{}
	md := &metadata{}

	var err error

	md.region, err = getMetadataValue(ctx, c, metadataRegionPath)
	if err != nil {
		return nil, err
	}

	return md, nil
}

func getMetadataValue(ctx context.Context, c *http.Client, path string) (string, error) {
	req, err := http.NewRequest("GET", "http://"+metadataServerHost+path, nil)
	if err != nil {
		return "", fmt.Errorf("failed http.NewRequest: %w", err)
	}

	req.Header.Add("Metadata-Flavor", "Google")

	resp, err := c.Do(req.WithContext(ctx))
	if err != nil {
		return "", fmt.Errorf("failed to get metadata %s: %w", path, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body of %s: %w", path, err)
	}

	return string(body), nil
}
