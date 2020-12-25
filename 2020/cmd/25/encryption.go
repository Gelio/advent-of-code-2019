package main

const subjectNum = 7
const moduloBy = 20201227

func getLoopSize(publicKey int) int {
	x := subjectNum

	loopSize := 1
	for x != publicKey {
		x = (x * subjectNum) % moduloBy
		loopSize++
	}

	return loopSize
}

func getEncryptionKey(subjecNum, loopSize int) int {
	x := subjecNum

	for i := 1; i < loopSize; i++ {
		x = (x * subjecNum) % moduloBy
	}

	return x
}
