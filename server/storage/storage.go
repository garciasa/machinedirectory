package storage

// Item Represent a row in database
type Item struct {
	ID         uint32 `json:"id"`
	IP         string `json:"ip"`
	DomainName string `json:"domainname"`
	Tags       string `json:"tags"`
	Deleted    bool   `json:"-"`
}
