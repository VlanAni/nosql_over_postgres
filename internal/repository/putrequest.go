package repository

type PutRequest struct {
	collectionName string
	id             string
	payload        []byte
}

func NewPutRequest(collectionName, id string, payload []byte) *PutRequest {
	payloadCopy := make([]byte, len(payload))
	copy(payloadCopy, payload)

	return &PutRequest{
		collectionName: collectionName,
		id:             id,
		payload:        payloadCopy,
	}
}

func (pr *PutRequest) CollectionName() string {
	return pr.collectionName
}

func (pr *PutRequest) ID() string {
	return pr.id
}

func (pr *PutRequest) Payload() []byte {
	payloadCopy := make([]byte, len(pr.payload))
	copy(payloadCopy, pr.payload)
	return payloadCopy
}
