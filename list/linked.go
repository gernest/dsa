package list

type LinkedList struct {
	Head  bool
	Next  *LinkedList
	Value interface{}
}

func Traverse(l *LinkedList, f func(*LinkedList) error) error {
	curr := l
	var err error
	for curr != nil {
		if err = f(curr); err != nil {
			return err
		}
		curr = curr.Next
	}
	return nil
}
