package rquid

import "github.com/google/uuid"

func CreateReqUid() (rqId string) {
	uid, err := uuid.NewRandom()
	if err != nil {
		rqId = uuid.Nil.String()
	} else {
		rqId = uid.String()
	}
	return rqId
}
