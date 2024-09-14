package dao

type TicketStatus int

const (
	Open TicketStatus = iota
	InProgress
	Resolved
	Closed
)

func (s TicketStatus) String() string {
	return [...]string{"Open", "InProgress", "Resolved", "Closed"}[s]
}

func GetStatusMap() map[int]string {
	statusMap := make(map[int]string)
	for p := Open; p <= Closed; p++ {
		statusMap[int(p)] = p.String()
	}
	return statusMap
}
