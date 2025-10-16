package array

import "strconv"

// 给你一个字符数组 chars ，请使用下述算法压缩：
//
// 从一个空字符串 s 开始。对于 chars 中的每组 连续重复字符 ：
//
// 如果这一组长度为 1 ，则将字符追加到 s 中。
// 否则，需要向 s 追加字符，后跟这一组的长度。
//
// 压缩后得到的字符串 s 不应该直接返回 ，需要转储到字符数组 chars 中。需要注意的是，如果组长度为 10 或 10 以上，则在 chars 数组中会
// 被拆分为多个字符。
//
// 请在 修改完输入数组后 ，返回该数组的新长度。
//
// 你必须设计并实现一个只使用常量额外空间的算法来解决此问题。
//
// 注意：数组中超出返回长度的字符无关紧要，应予忽略。
//
// 示例 1：
//
// 输入：chars = ["a","a","b","b","c","c","c"]
// 输出：返回 6 ，输入数组的前 6 个字符应该是：["a","2","b","2","c","3"]
// 解释："aa" 被 "a2" 替代。"bb" 被 "b2" 替代。"ccc" 被 "c3" 替代。
//
// 示例 2：
//
// 输入：chars = ["a"]
// 输出：返回 1 ，输入数组的前 1 个字符应该是：["a"]
// 解释：唯一的组是“a”，它保持未压缩，因为它是一个字符。
//
// 示例 3：
//
// 输入：chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
// 输出：返回 4 ，输入数组的前 4 个字符应该是：["a","b","1","2"]。
// 解释：由于字符 "a" 不重复，所以不会被压缩。"bbbbbbbbbbbb" 被 “b12” 替代。
//
// 提示：
//
// 1 <= chars.length <= 2000
// chars[i] 可以是小写英文字母、大写英文字母、数字或符号
func compress(chars []byte) int {
	slow, fast := 0, 0
	for fast < len(chars) {
		count := 0
		curChar := chars[fast]
		for fast < len(chars) && curChar == chars[fast] {
			count++
			fast++
		}
		chars[slow] = curChar
		slow++
		if count > 1 {
			digits := []byte(strconv.Itoa(count))
			for _, digit := range digits {
				chars[slow] = digit
				slow++
			}
		}
	}
	return slow
}

/**
思路：
双指针。
slow 指向当前写入位置，fast 指向当前统计位置。
每次固定一个字符curChar，对这个字符的重复出现次数进行统计，将统计结果拼接在curChar之后。
slow 在写入后自增，fast 在统计后自增。

最终 slow 会多自增一次，使得 slow == len(压缩字符串)，此时直接返回即可。

将数字转为字符串：
strconv.Itoa(count)
将数字转为字节数组，数组中的每个元素是字符串中的一个字符：
digits := []byte(strconv.Itoa(count))
*/
