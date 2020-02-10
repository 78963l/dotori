package main

import "strings"

// Tags2str 템플릿함수는 태그리스트를 공백으로 분리된 문자열로 만든다.
func Tags2str(tags []string) string {
	var newtags []string
	for _, tag := range tags {
		if tag != "" {
			newtags = append(newtags, tag)
		}
	}
	return strings.Join(newtags, " ")
}

// add함수는 입력받은 두 정수를 더한 값을 반환한다.
func add(a, b int) int {
	return (a + b)
}