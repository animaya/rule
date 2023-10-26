package rule

import (
	"fmt"
	"time"
	)

func WhatTimeIsItNow() {
	fmt.Println(time.Now().Format(time.RFC3339))
	
}
