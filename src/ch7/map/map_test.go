package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2}
	t.Log(m1, len(m1))

	m2 := map[int]int{}
	t.Log(m2, len(m2))
	m2[4] = 16
	t.Log(m2, len(m2))

	m3 := make(map[int]int, 10)
	t.Log(m3, len(m3))
}

func TestMapAccess(t *testing.T) {
	m1 := map[int]int{}
	// 访问不存在的值
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if _, ok := m1[3]; ok {
		t.Log("Key 3 exist")
	} else {
		t.Log("key 3 is not exist")
	}
}

func TestMapTrave(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true

	// 判断元素是否存在
	if mySet[1] {
		t.Log("1 is exist")
	} else {
		t.Log("1 is not exist")
	}
	// 元素个数
	t.Log(len(mySet))
	// 删除元素
	delete(mySet, 1)

	t.Log(mySet)
}
