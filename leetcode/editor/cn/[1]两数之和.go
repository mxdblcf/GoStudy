//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。 
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。 
//
// 你可以按任意顺序返回答案。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
// 
//
// 示例 2： 
//
// 
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
// 
//
// 示例 3： 
//
// 
//输入：nums = [3,3], target = 6
//输出：[0,1]
// 
//
// 
//
// 提示： 
//
// 
// 2 <= nums.length <= 10⁴ 
// -10⁹ <= nums[i] <= 10⁹ 
// -10⁹ <= target <= 10⁹ 
// 只会存在一个有效答案 
// 
//
// 进阶：你可以想出一个时间复杂度小于 O(n²) 的算法吗？ 
// Related Topics 数组 哈希表 👍 12004 👎 0

//思路 On 扫描数组，对于每个元素，在map中能组合给定值的另一半数组，
package Gostudy
//leetcode submit region begin(Prohibit modification and deletion)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		//另一个值 = 12-1=11 ，我们要在map中找 key =11的东西，招到的话就返回 【】下标，
		another := target - nums[i]
		//判断这个key是否存在于map中，存在的话就执行下一步
		 if  _,ok:= m[another]; ok{
		 	return []int{m[another],i}
		 }
		 //找不到的话就把 数组元素 value 变成 key  , 下标变成value
		 m[nums[i]]=i
	}
	return nil
}
//leetcode submit region end(Prohibit modification and deletion)
