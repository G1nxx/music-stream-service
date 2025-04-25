package users

type Author struct {
	User
	ID            int64
	ReleaseIDs    []int64
	SubscriberIDs []int64
	Verified      bool
}
