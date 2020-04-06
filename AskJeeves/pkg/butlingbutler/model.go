package butlingbutler

type UserResponse struct {
	Timestamp     string `json:Timestamp`
	EventID       string `json:"EventID"`
	Username      string `json:Username`
	Location      string `json:Location`
	IPaddr        string `json:IPaddress`
	VPNhash       string `json:VPNhash`
	Device        string `json:Device`
	Hostname      string `json:Hostname`
	UserSelection string `json:"user_selection"`
}
