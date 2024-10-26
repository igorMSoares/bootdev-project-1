package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocationAreas(pageUrl *string) (*LocationAreaRes, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	body := []byte{}

	if content, ok := c.cache.Get(url); ok {
		fmt.Println(">> Cache Hit!")
		body = content
	} else {
		var err error

		body, err = c.fetch(url)
		if err != nil {
			return nil, err
		}

		c.cache.Add(url, body)
	}

	results := LocationAreaRes{}

	err := json.Unmarshal(body, &results)
	if err != nil {
		return nil, fmt.Errorf("Error decoding response body: %w", err)
	}

	return &results, nil
}

func (c *Client) GetLocationByName(name string) (*NamedLocationAreaRes, error) {
	if name == "" {
		return nil, fmt.Errorf("Error fetching location area details: no location name provided")
	}

	url := baseUrl + "/location-area/" + name

	body := []byte{}

	if content, ok := c.cache.Get(url); ok {
		fmt.Println(">> Cache Hit!")
		body = content
	} else {
		var err error

		body, err = c.fetch(url)
		if err != nil {
			return nil, err
		}

		c.cache.Add(url, body)
	}

	results := NamedLocationAreaRes{}

	err := json.Unmarshal(body, &results)
	if err != nil {
		return nil, fmt.Errorf("Error decoding response body: %w", err)
	}

	return &results, nil
}

func (c *Client) GetPokemonByName(name string) (*Pokemon, error) {
	if name == "" {
		return nil, fmt.Errorf("Error fetching location area details: no location name provided")
	}

	url := baseUrl + "/pokemon/" + name

	body := []byte{}

	if content, ok := c.cache.Get(url); ok {
		fmt.Println(">> Cache Hit!")
		body = content
	} else {
		var err error

		body, err = c.fetch(url)
		if err != nil {
			return nil, err
		}

		c.cache.Add(url, body)
	}

	results := Pokemon{}

	err := json.Unmarshal(body, &results)
	if err != nil {
		return nil, fmt.Errorf("Error decoding response body: %w", err)
	}

	return &results, nil
}

func (c *Client) fetch(url string) ([]byte, error) {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error getting location areas: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Request error: %v", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	return body, nil
}
