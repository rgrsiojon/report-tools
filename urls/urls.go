package urls

type (
	urls struct {
		CARD_REVIEW_PATH     string
		CARD_CHANGE_DUE_PATH string
	}
)

func ReturnURLS() urls {
	var url_patterns urls
	url_patterns.CARD_REVIEW_PATH = "/b/cards/review"
	url_patterns.CARD_CHANGE_DUE_PATH = "/b/cards/change-due"
	return url_patterns
}
