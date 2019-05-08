package addTwoNumbers

import (
	"reflect"
	"testing"
)

// ListNode Definition for singly-linked list.
// type ListNode struct {
// 	Val  uint8
// 	Next *ListNode
// }

func TestAddTwoNumbers1(t *testing.T) {
	listresult := addTwoNumbers(&ListNode{1, &ListNode{2, nil}}, &ListNode{2, &ListNode{3, nil}})
	if !reflect.DeepEqual(*listresult, ListNode{3, &ListNode{5, nil}}) {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}

}
