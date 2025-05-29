package array

import (
	"math/rand"
	"time"
)

//实现RandomizedSet 类：
//
//
//
//
// RandomizedSet() 初始化 RandomizedSet 对象
// bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
// bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
// int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
//
//
//
//
// 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
//
//
//
// 示例：
//
//
//输入
//["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove",
//"insert", "getRandom"]
//[[], [1], [2], [2], [], [1], [2], []]
//输出
//[null, true, false, true, 2, true, false, 2]
//
//解释
//RandomizedSet randomizedSet = new RandomizedSet();
//randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
//randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
//randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
//randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
//randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
//randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
//randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
//
//
//
//
// 提示：
//
//
// -2³¹ <= val <= 2³¹ - 1
// 最多调用 insert、remove 和 getRandom 函数 2 * 10⁵ 次
// 在调用 getRandom 方法时，数据结构中 至少存在一个 元素。

type RandomizedSet struct {
	data     []int
	location map[int]int
}

func RandomizedSetConstructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		data:     make([]int, 0),
		location: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.location[val]; ok {
		return false
	}
	this.location[val] = len(this.data)
	this.data = append(this.data, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.location[val]
	if !ok {
		return false
	}
	lastIndex := len(this.data) - 1
	this.location[this.data[index]], this.location[this.data[lastIndex]] = lastIndex, index
	this.data[index], this.data[lastIndex] = this.data[lastIndex], this.data[index]
	delete(this.location, val)
	this.data = this.data[:lastIndex]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.data))
	return this.data[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

/**
数组：
- O(1)时间插入元素至末尾；
- O(1)时间交换两个元素；
- O(n)时间定位元素位置；

哈希表：
- O(1)时间定位元素位置。

结合两种数据结构的特性，
使用数组进行随机下标选取，且插入元素到数组末尾；
使用哈希表在删除特定元素时进行定位。
这样可满足O(1)时间复杂度插入、删除元素，且均等概率随机返回元素。
*/
