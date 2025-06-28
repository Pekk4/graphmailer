package graphmailer

type Attachment struct {
	DataType     string `json:"@odata.type"`
	Name         string `json:"name"`
	ContentBytes string `json:"contentBytes"`
	ContentType  string `json:"contentType"`
	Cid          string `json:"contentId"`
	Inline       string `json:"isInline"`
}

func NewAttachment(
	name string,
	content string,
	contentType string,
	cid string,
	isInline string,
) Attachment {
	attachment := Attachment{
		DataType:     "#microsoft.graph.fileAttachment",
		Name:         name,
		ContentBytes: content,
		ContentType:  contentType,
		Cid:          cid,
		Inline:       isInline,
	}

	return attachment
}
