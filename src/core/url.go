package core

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/dkeng/pkg/convert"
)

// ShortURL 短链接
func ShortURL(url string) ([4]string, error) {
	//要使用生成URL的字符
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x",
		"y", "z", "0", "1", "2", "3", "4", "5",
		"6", "7", "8", "9", "A", "B", "C", "D",
		"E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z",
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(url))
	uhex := hex.EncodeToString(md5Ctx.Sum(nil))
	resultURL := [4]string{}
	for i := 0; i < 4; i++ {
		// 把加密字符按照8位一组16进制与0x3FFFFFFF进行位与运算
		num := convert.BytesToInt32([]byte(uhex[i*8 : i*8+8]))
		hexint := 0x3FFFFFFF & num
		outChars := ""
		for j := 0; j < 6; j++ {
			//把得到的值与0x0000003D进行位与运算，取得字符数组chars索引
			index := 0x0000003D & hexint
			//把取得的字符相加
			outChars += chars[index]
			//每次循环按位右移5位
			hexint = hexint >> 5
		}
		//把字符串存入对应索引的输出数组
		resultURL[i] = outChars
	}
	return resultURL, nil
}
