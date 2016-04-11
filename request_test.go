package main

import "testing"

func TestQuery(t *testing.T) {
	query("I have a dog")
	query("我是一个苹果")
	query("word")
	query("天空")
}
