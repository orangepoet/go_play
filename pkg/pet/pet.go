package pet

import "time"

type Pet struct {
	Id int64
}

// step
// 1:[~], 2:[~], 3:[~],

type PetStep struct {
	Id    int64
	PetId int64
	Step  int32
}

type UserPet struct {
	UserId     int64
	CreateTime time.Time
	ModifyTime time.Time
	// IsDel int64
	Status    int32
	FeedCount int64
	//Step     int32
	Progress int64
}

// 先对鱼重构，再抽取
