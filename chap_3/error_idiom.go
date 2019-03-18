result, err := doSomething()
if err != nil {
	// エラー処理
}

if _, err := doSomething(); err != nil {
	// エラー時の処理
}

