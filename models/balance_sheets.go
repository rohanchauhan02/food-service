package models

type Lead struct {
	BussinessName string `json:"business_name"`
	BussinessGST  string `json:"bussiness_gst"`
	ContactNumber string `json:"contatc_number"`
	Provider      string `json:"provider"`
	LoanAmount    int    `json:"loan_amount"`
}

type LoanEnquiryRequest struct {
	BussinessName string `json:"business_name" required:"true"`
	BussinessGST  string `json:"bussiness_gst" required:"true"`
	ContactNumber int    `json:"contatc_number" required:"true"`
	Provider      string `json:"provider" required:"true"`
	LoanAmount    int    `json:"loan_amount" required:"true"`
}

type BussinessDetail struct {
	BussinessName string `json:"business_name"`
	BussinessGST  string `json:"bussiness_gst"`
}

type BalanceSheet struct {
	Year         int     `json:"year"`
	Month        int     `json:"month"`
	ProfitOrLoss float64 `json:"profit_or_loss"`
	AssetsValue  float64 `json:"assets_value"`
}

type LoanEnquiryResponse struct {
	Name          string  `json:"name"`
	PreAssessment float64 `json:"pre_assessment"`
}
