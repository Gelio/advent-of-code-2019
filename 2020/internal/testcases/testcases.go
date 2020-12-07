package testcases

// SplitTestCaseLines splits a single array of lines into lines for each test case
func SplitTestCaseLines(lines []string) (testCaseLines [][]string) {
	var testCase []string

	for _, line := range lines {
		if line == "" {
			testCaseLines = append(testCaseLines, testCase)
			testCase = make([]string, 0)
			continue
		}

		testCase = append(testCase, line)
	}

	if len(testCase) > 0 {
		testCaseLines = append(testCaseLines, testCase)
	}

	return
}
