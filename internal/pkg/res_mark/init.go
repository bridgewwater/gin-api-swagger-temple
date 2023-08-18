package res_mark

import "fmt"

func init() {
	err := generateMarkGitHeadShort()
	if err != nil {
		fmt.Printf("generate MarkGitHeadShort, please error: %v\n", err)
	}
}
