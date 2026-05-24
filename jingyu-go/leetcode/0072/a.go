package leetcode
func minDistance(word1 string, word2 string) int {
    w1,w2 := len(word1),len(word2)
	s1 := make([][]int,w1+1)
	for i := range s1{
		s1[i] = make([]int,w2+1)
	}
	for i := range w2{
		s1[0][i+1] = i+1
	} 
	for i,x := range word1{
		s1[i+1][0] = i+1
		for j,y:= range word2{
			if x == y {
				s1[i+1][j+1] = s1[i][j]
			}else{
				s1[i+1][j+1] = min(s1[i][j+1],s1[i+1][j],s1[i][j])+1
			}
		}
	}
	return s1[w1][w2]
}