
package StringFormatter

func ShowTransactionId(trnxID string) string {
	if trnxID != "" {
		return trnxID[0:10] //RETURN 1ST 10 CHARACTER
	}

	return ""
}
