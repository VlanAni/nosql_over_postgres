package repository

type GetRequest struct {
	collectionName string
	id             string
}

func NewGetRequest(collectionName, id string) *GetRequest {
	return &GetRequest{
		collectionName: collectionName,
		id:             id,
	}
}

func (gr *GetRequest) CollectionName() string {
	return gr.collectionName
}

func (gr *GetRequest) ID() string {
	return gr.id
}
