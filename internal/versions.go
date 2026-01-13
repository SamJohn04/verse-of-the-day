package internal

// the expansion of the most common versions of the Holy Bible
// returns the shortform if not found
func CommonBibleVersionNames(versionShortForm string) string {
	switch versionShortForm {
	case "NIV":
		return "New International Version"
	case "KJV":
		return "King James Version"
	case "NKJV":
		return "New King James Version"
	default:
		return versionShortForm
	}
}
