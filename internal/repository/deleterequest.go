package repository

type DeleteRequest struct {
	collectionName string
	id             string
}

func NewDeleteRequest(collectionName, id string) *DeleteRequest {
	return &DeleteRequest{
		collectionName: collectionName,
		id:             id,
	}
}

func (dr *DeleteRequest) CollectionName() string {
	return dr.collectionName
}

func (dr *DeleteRequest) ID() string {
	return dr.id
}
