package set
// 实现一个自定义的集合 Set
type HashSet struct {
	m map[interface{}]bool
}

type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Elements() []interface{}
}

// NewHashSet 初始化一个set 类型
func NewHashSet() *HashSet {
	return &HashSet{make(map[interface{}]bool)}
}

// Add 添加元素
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

// Remove 删除元素
func (set *HashSet) Remove(e interface{})  {
	delete(set.m, e)
}

// Clear 清除所有元素
func (set *HashSet) Clear()  {
	set.m = make(map[interface{}]bool)
}

// Contains 判断是否包含某个元素值
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

// Len 获取元素的数量
func (set *HashSet) Len() int  {
	return len(set.m)
}

// Same 判断与其他 HashSet类型值是否相同
func (set *HashSet) Same(other *HashSet) bool {
	// 如果传入的是空的 Set false
	if other == nil {
		return false
	}
	// 如果两个 Set 长度不一样 false
	if set.Len() != other.Len() {
		return false
	}
	// 如果传入的Set 中有不包含原始set的 false
	for key := range set.m {
		if !other.m[key] {
			return false
		}
	}
	return true
}

// Elements 快照
func (set *HashSet) Elements() []interface{}  {
	// 获取HashSet中的 m的长度
	initLen := len(set.m)
	// 初始化一个切片 存储 m中的值
	snapshot := make([]interface{}, initLen)
	actualLen := 0
	// 循环将值存到 快照中
	for key := range set.m {
		if actualLen < initLen {
			snapshot[actualLen] = key
		} else {
			// m的值数量有所增加，实际迭代的次数大于 初始化的长度
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	// 如果实际迭代的次数 小于初始化时设置的 长度 就截取掉不需要的
	if actualLen < initLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}


// Union 并集运算
func (set *HashSet) Union(other *HashSet) *HashSet {
	// 如果set 和other 为nil 就为nil
	if set == nil && other == nil {
		return nil
	}
	unionedSet := NewHashSet()
	// 遍历set 中所有的值，然后将值赋给新的 set
	for _, v := range set.Elements() {
		unionedSet.Add(v)
	}
	if other.Len() == 0 {
		return unionedSet
	}
	// 再将 other中的元素的值也添加到 unionedSet
	for _, v := range other.Elements() {
		unionedSet.Add(v)
	}
	return unionedSet
}

// Intersect 交集运算
func (set *HashSet) Intersect(other *HashSet) *HashSet  {
	// 如果 set 和 other 有一个为nil 交集就为空
	if set == nil || other == nil {
		return nil
	}
	intersectSet := NewHashSet()
	if other.Len() == 0 {
		return intersectSet
	}
	if set.Len() < other.Len() {
		// 遍历数量少的
		for _, v := range set.Elements() {
			if other.Contains(v) {
				intersectSet.Add(v)
			}
		}
	} else {
		for _, v := range other.Elements() {
			if set.Contains(v) {
				intersectSet.Add(v)
			}

		}
	}
	return intersectSet
}

// Difference 差集
func (set *HashSet) Difference(other *HashSet) *HashSet {
	if set == nil || other == nil {
		return nil
	}
	differenceSet := NewHashSet()
	if other.Len() == 0 {
		for _, v := range set.Elements() {
			differenceSet.Add(v)
		}
		return differenceSet
	}
	for _, v := range set.Elements() {
		if !other.Contains(v) {
			differenceSet.Add(v)
		}
	}
	return differenceSet
}
