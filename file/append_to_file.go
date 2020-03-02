/*
 * @Description: 追加内容到文件末尾
 * @Author: leo
 * @Date: 2020-03-01 20:09:41
 * @LastEditors: leo
 * @LastEditTime: 2020-03-01 20:11:19
 */
 package tools

import (
	"fmt"
	"io"
	"os"
)

 // AppendToFile 追加内容到文件末尾
func AppendToFile(fileName, content string) error {
	 // 以只写的模式，打开文件
	 file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	 if err != nil {
	 	fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	 } else {
	 	// 查找文件末尾的偏移量
	 	n, _ := file.Seek(0, io.SeekEnd)
	 	// 从末尾的偏移量开始写入内容
	 	_, err = file.WriteAt([]byte(content), n)
	 }
	 defer file.Close()
	return err
 }
