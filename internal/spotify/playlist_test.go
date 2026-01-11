package spotify

import (
	"context"
	"fmt"
	"testing"
)

func TestPlaylist(t *testing.T) {
	p := playlist{"https://api.spotify.com/v1", "BQBeIrGnpyUYHO0Gn1hNjmYkdN6y97ZvIQw582ImkyF9bNTlWCTzk6aiClJ4Y5yRaRYQMHlzdKHQKoOc0uz7gWWzeqQ01c3UT1uyaR9l9nOn6_BUoinDrt6ZH6jTNKHb6Qs4onAITbdFBPi8OKfEwELtNvrv59EhU-qc9Pl8A_qsnT1CQ8JFQdedcy0dmyzEIpuFufDY4SQNGrPOYrTBTe0NQca9SUJHX8he__TOAyrtLGv8kKlX5jio-C1I-fREdz23RQal33jZz541EhzENvXcTxTR-hp7GHT0aFT6brmSDEHXJv3E_JNRpyz4P4FPCqGvjZMFO18dnOA3orXb_aRAXwFuuqlbS0L1Syb3nsv9X90rZSOw18RlRY4-HZEjVuy1h3SxPDbcLzCMctqIkww"}

	ap, err := p.Get(context.Background(), "7egW6I9Z7iS2glkXu7QZ7t")
	fmt.Println(ap, err)
}
