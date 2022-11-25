package supercluster

type AllFiles map[string]file

type ListResponse struct {
	Files []file `json:"files"`
}

type CreatePayload struct {
	Name     string `json:"name"`
	Contents string `json:"contents"`
}

type CreateResponse struct {
	File file `json:"file"`
}

type ModifyPayload struct {
	Contents string `json:"contents"`
}

type ModifyResponse struct {
	File file `json:"file"`
}

type ResponseError struct {
	Error string `json:"error"`
}

type AddrsResponse struct {
	ID    string   `json:"id"`
	Addrs []string `json:"addrs"`
}
