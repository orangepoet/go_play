package play

type WxURLLinkRequest struct {
	Path           string `json:"path"`
	Query          string `json:"query"`
	IsExpire       bool   `json:"is_expire"`
	ExpireType     int    `json:"expire_type"`
	ExpireTime     int64  `json:"expire_time"`
	ExpireInterval int    `json:"expire_interval"`
}

type WxURLLinkResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	URLLink string `json:"url_link"`
}

type Sth struct {
	Name string
}

type Dao struct {
}

func NewDao() *Dao {
	return &Dao{}
}

func (d *Dao) GetData() (int, error) {
	return 1, nil
}

//go:noinline
func GetData1() (int, error) {
	return 1, nil
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetData() (int, error) {
	return NewDao().GetData()
}

type A struct {
	*B
	Name string `json:"name,omitempty"`
}

type B struct {
	P1 int    `json:"p1,omitempty"`
	P2 string `json:"p2,omitempty"`
}
