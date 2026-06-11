package solution

const MOD = 1e9 + 7

func assignEdgeWeights(edges [][]int) int {
	mp := map[int][]int{}
	for _, e := range edges {
		mp[e[0]] = append(mp[e[0]], e[1])
		mp[e[1]] = append(mp[e[1]], e[0]) // TODO: 添加反向边
	}

	k := 0 // TODO: 改为0
	var dfs func(int, int)
	dfs = func(i, dep int) {
		k = max(k, dep)
		for _, j := range mp[i] {
			dfs(j, dep + 1)
		}
	}
	dfs(1, 0) // TODO: 改为0

	return powMod(2, k-1, MOD) % MOD // TODO: 用快速幂替代math.Pow
}

// TODO: 添加快速幂函数
func powMod(a, n, mod int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = (res * a) % mod
		}
		a = (a * a) % mod
		n /= 2
	}
	return res
}
