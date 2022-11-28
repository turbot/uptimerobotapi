package uptimerobotapi

import "net/http"

type PSPService struct {
	apiClient *Client
}

type GetPSPsParams struct {
	// PSPs should be provided as IDs separated with hyphens '-' example: 123-4567-845
	PSPs   string `url:"psps,omitempty"`
	Offset int    `url:"offset,omitempty"`
	Limit  *int   `url:"limit,omitempty"`
}

type PSPsResp struct {
	Stat       string     `json:"stat"`
	Pagination Pagination `json:"pagination"`
	PSPs       []PSP      `json:"psps"`
}

type PSP struct {
	ID           int    `json:"id"`
	FriendlyName string `json:"friendly_name"`
	Monitors     int    `json:"monitors"`
	Sort         int    `json:"sort"`
	Status       int    `json:"status"`
	StandardURL  string `json:"standard_url"`
	CustomURL    string `json:"customURL"`
}

type NewPSPParams struct {
	Type         int    `url:"type"`
	FriendlyName string `url:"friendly_name"`
	// Monitors should be provided as hyphen separated IDs (or 0 for all)
	Monitors     string `url:"monitors"`
	CustomDomain string `url:"custom_domain,omitempty"`
	Password     string `url:"password,omitempty"`
	Sort         int    `url:"sort,omitempty"`
	HideURLLinks bool   `url:"hide_url_links,omitempty"`
	// Status should be 0 (paused) or 1 (active), other values may return an error
	Status int `url:"status,omitempty"`
}

type PSPResp struct {
	Stat string `json:"stat"`
	PSP  struct {
		ID int `json:"id"`
	} `json:"psp"`
}

type EditPSPParams struct {
	ID int `url:"id"`
	NewPSPParams
}

type DeletePSPParams struct {
	ID int `url:"id"`
}

// GetPSPs Get https://uptimerobot.com/#getPSPsWrap
func (ps *PSPService) GetPSPs(params GetPSPsParams) (*PSPsResp, error) {
	obj := &PSPsResp{}
	err := ps.apiClient.request(http.MethodPost, "getPSPs", params, obj)
	return obj, err
}

// NewPSP Get https://uptimerobot.com/#newPSPWrap
func (ps *PSPService) NewPSP(params NewPSPParams) (*PSPResp, error) {
	obj := &PSPResp{}
	err := ps.apiClient.request(http.MethodPost, "newPSP", params, &obj)
	return obj, err
}

// EditPSP Get https://uptimerobot.com/#editPSPWrap
func (ps *PSPService) EditPSP(params EditPSPParams) (*PSPResp, error) {
	obj := &PSPResp{}
	err := ps.apiClient.request(http.MethodPost, "editPSP", params, &obj)
	return obj, err
}

// DeletePSP Get https://uptimerobot.com/#deletePSPWrap
func (ps *PSPService) DeletePSP(id int) (*PSPResp, error) {
	obj := &PSPResp{}
	params := DeletePSPParams{ID: id}
	err := ps.apiClient.request(http.MethodPost, "deletePSP", params, obj)
	return obj, err
}
