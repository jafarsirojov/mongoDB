package status


type Status int

const (
	PENDING Status = iota
	APPROVED
	PROCESSING
	ACCEPTED
	FAILED
	CANCELED
)

func Validate(status string) Status {
	return statusMap[status]
}

func GetStr(status Status) string {
	return getMap[status]
}

var statusMap = map[string]Status{
	"pending":    PENDING,
	"approved":   APPROVED,
	"processing": PROCESSING,
	"accepted":   ACCEPTED,
	"failed":     FAILED,
	"canceled":   CANCELED,
}

var getMap = map[Status]string{
	PENDING:    "pending",
	APPROVED:   "approved",
	PROCESSING: "processing",
	ACCEPTED:   "accepted",
	FAILED:     "failed",
	CANCELED:   "canceled",
}

func ValidateStatus(value string) bool {
	_, ok := statusMap[value]
	return ok
}
