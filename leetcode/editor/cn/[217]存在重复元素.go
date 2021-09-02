//给定一个整数数组，判断是否存在重复元素。 
//
// 如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。 
//
// 
//
// 示例 1: 
//
// 
//输入: [1,2,3,1]
//输出: true 
//
// 示例 2: 
//
// 
//输入: [1,2,3,4]
//输出: false 
//
// 示例 3: 
//
// 
//输入: [1,1,1,3,3,4,3,2,4,2]
//输出: true 
// Related Topics 数组 哈希表 排序 👍 461 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
package main

//排序，排好序再一一比较 在对数字从小到大排序之后，数组的重复元素一定出现在相邻位置中。
//因此，我们可以扫描已排序的数组，每次判断相邻的两个元素是否相等，如果相等则说明存在重复的元素。
/*func containsDuplicate(nums []int) bool {
sort.Ints(nums)  //排序
	for i := 1; i <len(nums) ; i++ {
		if nums[i]==nums[i-1]{
			return true
		}
	}
	return false
}*/
// 对于数组中每个元素，我们将它插入到哈希表中。如果插入一个元素时发现该元素已经存在于哈希表中，则说明存在重复的元素。
func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}
	for _,v :=range nums{
		if _, has:= m[v]; has {
			return true
		}
		m[v]= struct{}{}
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, has := set[v]; has {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)