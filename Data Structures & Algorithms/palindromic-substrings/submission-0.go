func countSubstrings(s string) int {
   amount := 0
//    var check func (i int) 
//    check = func (i int) {
// 			left := i-1
// 			right := i+1
// 		for {
// 			if left < 0 || right > len(s) - 1 {
// 				return
// 			}
// 			if s[left] == s[right] {
// 				amount++
// 				left--
// 				right++
// 				continue
// 			}
// 			break
// 		}
//    }


	var dfs func (l int, r int) 
	dfs = func (l int, r int){
		if l < 0 || r > len(s) - 1 {
			return
		}
		if s[l] != s[r]  {
			return
		}
		amount++
		dfs(l - 1, r + 1)
		// dfs(l, r + 1)
		// dfs(l-1, r)
	}


   for i := 0; i < len(s); i++ {
			dfs(i, i + 1)
			dfs(i, i)
   }
   return amount
}
