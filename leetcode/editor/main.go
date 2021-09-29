package main



func main() {
	LongestPalindrome("bcbad")
}

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
