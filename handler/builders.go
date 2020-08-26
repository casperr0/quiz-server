package handler

import (
	"encoding/json"
)

// HATEOAS describe the restful response structure
type HATEOAS struct {
	LinksSection    map[string]LinkDetail `json:"_links"`
	StatusSection   Status                `json:"status"`
	DataSection     interface{}           `json:"data"`
	EmbeddedSection []HATEOAS             `json:"_embedded"`
}

// LinkDetail describe details of each link in _links
type LinkDetail struct {
	Href string `json:"href"`
}

// Status describe status of the response
type Status struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
}

// BuildHATEOAS build the response object
func BuildHATEOAS(links map[string]LinkDetail, status Status, data interface{}, embedded []HATEOAS) (string, error) {
	content := HATEOAS{
		LinksSection:    links,
		StatusSection:   status,
		DataSection:     data,
		EmbeddedSection: embedded,
	}
	result, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// func init() {

// 	link0 := LinkDetail{
// 		Href: "google.com",
// 		Rel:  "home",
// 	}
// 	link1 := LinkDetail{
// 		Href: "facebook.com",
// 		Rel:  "fb",
// 	}
// 	links := map[string]LinkDetail{"google": link0, "facebook": link1}
// 	status := Status{
// 		Code:    200,
// 		Message: "list all bookmarks.",
// 	}

// 	content, _ := BuildHATEOAS(links, status, nil, nil)
// 	fmt.Print(content)
// }
