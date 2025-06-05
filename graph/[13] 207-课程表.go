package graph

// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表
// 示如果要学习课程 ai 则 必须 先学习课程 bi 。
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：numCourses = 2, prerequisites = [[1,0]]
// 输出：true
// 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//
// 示例 2：
//
// 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
// 输出：false
// 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
//
// 提示：
//
// 1 <= numCourses <= 2000
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// prerequisites[i] 中的所有课程对 互不相同
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildUpCourseGraph(numCourses, prerequisites)
	return !detectCircle(graph)
	// return !detectCircleBFS(graph)
}

func buildUpCourseGraph(numCourse int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourse)
	for _, prerequisite := range prerequisites {
		from, to := prerequisite[1], prerequisite[0]
		graph[from] = append(graph[from], to)
	}
	return graph
}

/**
1. 构建邻接表，graph[v] = w表示v -> w，即w依赖v；
2. 基于邻接表表示的图，进行环检测。
*/
