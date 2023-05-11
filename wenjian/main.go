package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//打开文件
	file1, err := os.Open("src/practice/doc1.txt")
	if err != nil{
		fmt.Println("open file err=", err)
	}
	//输出文件，看看文件是什么，可以看出file就是一个指针*File 即file =&{0xc000108780}
	//fmt.Printf("file =%v", file)
	//关闭文件，函数退出时及时关闭文件
	defer file1.Close()

	//读文件方法一，显示在终端（带缓冲区方式）  只能读一行
	//创建一个*Reader， 带缓冲区  大小默认为4096
	reader := bufio.NewReader(file1)
	for{             //循环的读取文件的内容
		str, err := reader.ReadString('\n')//读到一个换行就结束  只能读一行
		if err == io.EOF{ //io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)//输出内容
	}
	fmt.Println("读文件结束")

	//读文件方法二，显示在终端，读入内存，适用于文件不大的情况
	filePath1 := "src/practice/doc1.txt"           //这种方式文件的Open和Close在ReadFile函数中自动进行，不需要自己写
	content, err := ioutil.ReadFile(filePath1)
	if err != nil{
		fmt.Println("read file err=", err)
	}
	//fmt.Printf("%v\n", content)//读取的内容显示到终端 [122 104... 51] [byte型]
	fmt.Printf("doc1文件的内容%v\n", string(content))

	//写文件方法一  bufio.NewWriter带缓存的
	filePath2 := "src/practice/doc1.txt"
	file2, err := os.OpenFile(filePath2, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("open file err=", err)
		return
	}
	defer file2.Close()
	str := "\nhello susu"
	writer := bufio.NewWriter(file2) //writer是带缓存的，因此
	writer.WriteString(str)  //WriteString方法是先写入到缓存中
	writer.Flush()           //Flush方法，将缓存中的数据写入到文件中

	//写文件方法二 ioutil.WriteFile     文件必须存在
	filePath3 := "src/practice/doc2.txt"
	err = ioutil.WriteFile(filePath3, content,0666)//将doc1中的内容写到doc2中去
	if err != nil{
		fmt.Println("write file err=", err)
	}
	fmt.Printf("doc2文件的内容%v\n", string(content))

	//判断文件是否存在
	PathExists(filePath1)

	//拷贝文件，将一张图片、电影、MP3拷贝到另一个文件
	srcFile := "src/1.jpg"
	dstFile := "src/2.jpg"
	_, err1 := CopyFile(dstFile, srcFile)
	if err1 != nil{
		fmt.Println("拷贝完成")
	}else {
		fmt.Println("拷贝错误")
	}

}

//判断文件是否存在
func PathExists(path string) (bool, error) {
	//判断文件是否存在
	_, err := os.Stat(path)
	if err == nil {  //文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err){//  如果os.IsNotExist()判断为true说明文件或文件夹不存在
		return false, nil
	}
	return false,err
}

//拷贝文件，将一张图片、电影、MP3拷贝到另一个文件
func CopyFile(dstFileName string, srcFileName string)(written int64, err error){
	srcFile, err := os.Open(srcFileName)
	if err != nil{
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcfile获取到Reader
	reader := bufio.NewReader(srcFile)
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil{
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//通过dstFile获取到Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}