//给你一个字符串 s，找到 s 中最长的回文子串。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
// 
//
// 示例 2： 
//
// 
//输入：s = "cbbd"
//输出："bb"
// 
//
// 示例 3： 
//
// 
//输入：s = "a"
//输出："a"
// 
//
// 示例 4： 
//
// 
//输入：s = "ac"
//输出："a"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 1000 
// s 仅由数字和英文字母（大写和/或小写）组成 
// 
// Related Topics 字符串 动态规划 👍 4083 👎 0

package GoStudy
//使用动态规划做 On2  ,所以使用 中心扩散法，找中心轴
//定义 dep[i][j]表示从第i个字符 到第j个字符这一段是不是一个会问穿
//回文串去掉头尾后还是一个回文串
//leetcode submit region begin(Prohibit modification and deletion)


func LongestPalindrome(s string) string {

	if s == "" {
		return ""
	}
	start ,end :=0,0
	for i,_:= range s{
		l1, r1 := ExpandAroundCenter(s, i, i)
		l2, r2 := ExpandAroundCenter(s, i, i+1)
		if r1-l1>end-start {
			start,end=l1,r1
		}
		if r2 -l2 >end-start{
			start,end=l2,r2
		}
	}
return s[start:end+1]

}
//中心扩散
func ExpandAroundCenter( s string, l,r int)  (int,int){

	for ; l>=0 && r< len(s) && s[l]== s[r]; l ,r=l-1,r+1 {}
 return l+1,r-1
}
//leetcode submit region end(Prohibit modification and deletion)
