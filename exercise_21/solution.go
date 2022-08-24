package exercise_21

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var mergedNode *ListNode
	travM := mergedNode
	travL1 := list1
	travL2 := list2

	for travL1 != nil || travL2 != nil {
		if mergedNode == nil {
			if travL1 != nil && travL2 != nil {
				if travL1.Val > travL2.Val {
					mergedNode = &ListNode{
						Val:  travL2.Val,
						Next: nil,
					}
					travL2 = travL2.Next
				} else {
					mergedNode = &ListNode{
						Val:  travL1.Val,
						Next: nil,
					}
					travL1 = travL1.Next
				}
			} else {
				if travL1 == nil {
					mergedNode = &ListNode{
						Val:  travL2.Val,
						Next: nil,
					}
					travL2 = travL2.Next
				} else if travL2 == nil {
					mergedNode = &ListNode{
						Val:  travL1.Val,
						Next: nil,
					}
					travL1 = travL1.Next
				}
			}

			travM = mergedNode
		} else {
			if travL1 != nil && travL2 != nil {
				if travL1.Val > travL2.Val {
					travM.Next = &ListNode{
						Val:  travL2.Val,
						Next: nil,
					}

					travL2 = travL2.Next
				} else {
					travM.Next = &ListNode{
						Val:  travL1.Val,
						Next: nil,
					}

					travL1 = travL1.Next
				}
			} else {
				if travL1 == nil {
					travM.Next = &ListNode{
						Val:  travL2.Val,
						Next: nil,
					}
					travL2 = travL2.Next
				} else if travL2 == nil {
					travM.Next = &ListNode{
						Val:  travL1.Val,
						Next: nil,
					}
					travL1 = travL1.Next
				}
			}

			travM = travM.Next
		}
	}
	return mergedNode
}
