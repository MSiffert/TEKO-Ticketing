package dao

type TicketPriority int

const (
	Low TicketPriority = iota
	Medium
	High
	Urgent
)

func (p TicketPriority) String() string {
	return [...]string{"Low", "Medium", "High", "Urgent"}[p]
}

func GetPriorityMap() map[int]string {
	priorityMap := make(map[int]string)
	for p := Low; p <= Urgent; p++ {
		priorityMap[int(p)] = p.String()
	}
	return priorityMap
}
