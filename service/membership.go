// 멤버십 구조체 선언

package service

// 멤버십이 가지는 정보
type Membership struct {
	Type int `json:"type"`
	Cost int `json:"cost"`
}

const (
	// 멤버십 종류
	MEMBERSHIP_NETFLIX_TYPE_NO = iota
	MEMBERSHIP_NETFLIX_TYPE_BASIC
	MEMBERSHIP_NETFLIX_TYPE_STANDARD
	MEMBERSHIP_NETFLIX_TYPE_PREMIUM

	MEMBERSHIP_WAVVE_TYPE_NO
	MEMBERSHIP_WAVVE_TYPE_BASIC
	MEMBERSHIP_WAVVE_TYPE_STANDARD
	MEMBERSHIP_WAVVE_TYPE_PREMIUM
	MEMBERSHIP_WAVVE_TYPE_FLO
	MEMBERSHIP_WAVVE_TYPE_BUGS
	MEMBERSHIP_WAVVE_TYPE_KB

	// 멤버십 종류별 가격
	NETFLIX_MEMBERSHIP_COST_NO       = 0
	NETFLIX_MEMBERSHIP_COST_BASIC    = 9_500
	NETFLIX_MEMBERSHIP_COST_STANDARD = 13_500
	NETFLIX_MEMBERSHIP_COST_PREMIUM  = 17_000

	WAVVE_MEMBERSHIP_COST_NO       = 0
	WAVVE_MEMBERSHIP_COST_BASIC    = 7_900
	WAVVE_MEMBERSHIP_COST_STANDARD = 10_900
	WAVVE_MEMBERSHIP_COST_PREMIUM  = 13_900
	WAVVE_MEMBERSHIP_COST_FLO      = 13_750
	WAVVE_MEMBERSHIP_COST_BUGS     = 13_750
	WAVVE_MEMBERSHIP_COST_KB       = 6_700
)
