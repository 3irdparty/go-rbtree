package rbtree

import "testing"

func TestLeftRotate(t *testing.T) {
	var rbtree RBTree
	rbtree.Init()
	rbtree.Insert(10, 1)
	rbtree.Insert(20, 1)
	rbtree.Insert(30, 1)
	if rbtree.root.key != 20 || rbtree.root.left.key != 10 || rbtree.root.right.key != 30 {
		t.Errorf("Left Rotate Error")
	}
}

func TestRightRotate(t *testing.T) {
	var rbtree RBTree
	rbtree.Init()
	rbtree.Insert(30, 1)
	rbtree.Insert(20, 1)
	rbtree.Insert(10, 1)
	if rbtree.root.key != 20 || rbtree.root.left.key != 10 || rbtree.root.right.key != 30 {
		t.Errorf("Right Rotate Error")
	}
}

func TestMin(t *testing.T) {
	var rbtree RBTree
	rbtree.Init()
	rbtree.Insert(20, 1)
	rbtree.Insert(10, 1)
	rbtree.Insert(30, 1)
	if rbtree.Min().key != 10 {
		t.Errorf("Min Error")
	}
}

func TestMax(t *testing.T) {
	var rbtree RBTree
	rbtree.Init()
	rbtree.Insert(20, 1)
	rbtree.Insert(10, 1)
	rbtree.Insert(30, 1)
	if rbtree.Max().key != 30 {
		t.Errorf("Max Error")
	}
}

func TestPrevOf(t *testing.T) {
	var rbtree RBTree
	rbtree.Init()
	for i := 10; i <= 200; i = i + 10 {
		rbtree.Insert(int64(i), i)
	}

	for i := 20; i <= 200; i += 10 {
		if int64(i-10) != rbtree.PrevOf(rbtree.Search(int64(i))).Key() {
			t.Errorf("PrevOf Error")
		}
	}
	if rbtree.PrevOf(rbtree.Search(10)).Key() != -1 {
		t.Errorf("PrevOf Error")
	}
}

func TestNextOf(t *testing.T) {
	var rbtree RBTree
	rbtree.Init()
	for i := 10; i <= 200; i = i + 10 {
		rbtree.Insert(int64(i), i)
	}

	for i := 10; i < 200; i += 10 {
		if int64(i+10) != rbtree.NextOf(rbtree.Search(int64(i))).Key() {
			t.Errorf("NextOf Error", i+10)
		}
	}
	if rbtree.NextOf(rbtree.Search(200)).Key() != -1 {
		t.Errorf("NextOf Error")
	}
}

func TestDeep(t *testing.T) {
	var rbtree RBTree
	rbtree.Init()
	rbtree.Insert(20, 1)
	rbtree.Insert(10, 1)
	rbtree.Insert(30, 1)

	ldeep := rbtree.Deep(rbtree.root.left)
	rdeep := rbtree.Deep(rbtree.root.right)

	val := ldeep - rdeep
	if val < 0 {
		val *= -1
	}
	if val > 2 {
		t.Errorf("Deep error: leftdeep:", ldeep, ", rightdeep:", rdeep)
	}
}

func BenchmarkInsert(b *testing.B) {
	var rbtree RBTree
	rbtree.Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbtree.Insert(int64(i), i)
	}
}

func BenchmarkDelete(b *testing.B) {
	var rbtree RBTree
	rbtree.Init()
	for i := 0; i < b.N; i++ {
		rbtree.Insert(int64(i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbtree.Delete(int64(i))
	}
}
