package tools

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

/**
@description
@date: 05/21 23:59
@author Gk
**/

func RoundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*10000) / 10000
}

// getColumn 从二维切片（代表CSV数据）中提取指定列的数据
func GetColumn(data [][]string, columnIndex int) ([]string, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("输入数据为空")
	}
	if columnIndex < 0 {
		return nil, fmt.Errorf("列索引不能为负数")
	}

	var columnData []string
	for i, row := range data {
		if columnIndex >= len(row) {
			// 如果某行的列数少于请求的列索引，可以根据需求决定如何处理
			// 这里我们返回错误，也可以选择填充空字符串或跳过
			return nil, fmt.Errorf("第 %d 行的列数 (%d) 少于请求的列索引 (%d)", i+1, len(row), columnIndex)
		}
		columnData = append(columnData, row[columnIndex])
	}
	return columnData, nil
}

func ReadCsv(name string) ([][]string, error) {
	csvFile, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	csvReader.LazyQuotes = true
	csvReader.TrimLeadingSpace = true
	allRecords, err := csvReader.ReadAll()
	if err != nil {
		if parseErr, ok := err.(*csv.ParseError); ok {
			log.Fatalf("CSV Parse Error on line %d, column %d: %s",
				parseErr.Line, parseErr.Column, parseErr.Err)
		}
		log.Fatal("读取CSV错误: ", err)
		return nil, err
	}

	if len(allRecords) == 0 {
		fmt.Println("CSV 文件为空或无法读取任何记录。")
		return nil, err
	}

	return allRecords, nil
}

func WriteCsv(csvName string, data [][]string) error {
	if len(data) == 0 {
		return fmt.Errorf("CSV 文件为空或无法读取任何记录。")
	}

	path2 := path.Dir(csvName)

	err := CreateDirIfNotExist(path2)
	if err != nil {
		fmt.Printf("output: %v\n", err)
		return err
	}

	// 创建新的 CSV 文件来保存更新后的数据
	outputFile, err := os.Create(csvName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer outputFile.Close()

	// 创建 CSV 写入器
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	if err := writer.WriteAll(data); err != nil {
		fmt.Println("Error writing records:", err)
		return err
	}
	writer.Flush()
	return nil
}

func GetAllFiles(dirPath string) ([]string, error) {
	var files []string

	// 检查目录是否存在
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("目录不存在: %s", dirPath)
	}
	if err != nil {
		return nil, fmt.Errorf("获取目录信息失败: %s, 错误: %v", dirPath, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("提供的路径不是一个目录: %s", dirPath)
	}

	// filepath.WalkDir 会遍历指定目录下的所有文件和子目录
	// 第三个参数是一个回调函数，它会在访问每个文件或目录时被调用
	err = filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// 如果在遍历过程中遇到错误（例如权限问题），可以选择如何处理
			// 这里我们打印错误并继续遍历其他文件/目录
			fmt.Printf("访问路径 %q 时发生错误: %v\n", path, err)
			return err // 可以返回 filepath.SkipDir 来跳过整个目录
		}

		// 检查是否是普通文件 (而不是目录)
		if !d.IsDir() && strings.HasSuffix(path, "csv") {
			files = append(files, path)
		}
		return nil // 继续遍历
	})

	if err != nil {
		return nil, fmt.Errorf("遍历目录 %s 时出错: %v", dirPath, err)
	}

	return files, nil
}

func GetOneList(records [][]string) map[string]string {
	datamap := make(map[string]string)
	// 遍历现有记录并添加新列的数据
	for idx, record := range records[1:] {
		//res := f(record)
		_, ok := datamap[record[0]]
		if ok {
			fmt.Printf("output: 重复 %v,%v\n", idx, record[0])
		} else {
			datamap[record[0]] = record[2]
		}
		//newRow := append(record, res)
		//updatedRecords = append(updatedRecords, newRow)
	}
	return datamap
}

func MergeSimpleTypeCsv(records1, records2 [][]string) [][]string {
	var updatedRecords [][]string
	updatedRecords = append(updatedRecords, []string{"name", "avg", "max", "min", "sum", "up_max2", "max_sum"})

	addData := GetOneList(records2)

	// 遍历现有记录并添加新列的数据
	for _, record := range records1[1:] {
		//res := f(record)

		v, ok := addData[record[0]]
		if !ok {
			continue
		}

		vv, _ := strconv.ParseFloat(v, 10)
		r2, _ := strconv.ParseFloat(record[2], 10)

		record = append(record, v, fmt.Sprintf("%.2f", vv+r2))
		//newRow := append(record, res)
		updatedRecords = append(updatedRecords, record)
	}
	return updatedRecords
}

// CreateDirIfNotExist 检查路径是否存在，如果不存在且是目录则创建它
func CreateDirIfNotExist(dirPath string) error {
	// 使用 os.Stat 获取文件/目录信息
	info, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		// 路径不存在，创建目录
		fmt.Printf("目录 '%s' 不存在，正在创建...\n", dirPath)
		// os.MkdirAll 会创建所有必要的父目录，类似 mkdir -p
		// 0755 是一种常见的目录权限 (rwxr-xr-x)
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("创建目录 '%s' 失败: %w", dirPath, err)
		}
		return nil
	}

	if err != nil {
		// 其他类型的错误 (例如权限问题)
		return fmt.Errorf("检查目录 '%s' 时发生错误: %w", dirPath, err)
	}

	if !info.IsDir() {
		// 路径存在，但不是一个目录
		return fmt.Errorf("路径 '%s' 已存在，但它不是一个目录", dirPath)
	}

	return nil
}

func ListFilesInCurrentDir(dirPath string) ([]string, error) {
	var files []string

	// 检查目录是否存在且确实是一个目录
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("目录不存在: %s", dirPath)
	}
	if err != nil {
		return nil, fmt.Errorf("获取目录 '%s' 信息失败: %w", dirPath, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("路径 '%s' 不是一个目录", dirPath)
	}

	// 读取目录内容
	// os.ReadDir 返回一个 []fs.DirEntry 切片
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("读取目录 '%s' 内容失败: %w", dirPath, err)
	}

	for _, entry := range entries {
		// 检查条目是否是文件 (而不是目录)
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), "csv") {
			// 构建文件的完整路径
			filePath := filepath.Join(dirPath, entry.Name())
			files = append(files, filePath)
		}
	}

	return files, nil
}

func ListFilesInCurrentDirAny(dirPath string) ([]string, error) {
	var files []string

	// 检查目录是否存在且确实是一个目录
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("目录不存在: %s", dirPath)
	}
	if err != nil {
		return nil, fmt.Errorf("获取目录 '%s' 信息失败: %w", dirPath, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("路径 '%s' 不是一个目录", dirPath)
	}

	// 读取目录内容
	// os.ReadDir 返回一个 []fs.DirEntry 切片
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("读取目录 '%s' 内容失败: %w", dirPath, err)
	}

	for _, entry := range entries {
		// 检查条目是否是文件 (而不是目录)
		if !strings.Contains(entry.Name(), "DS_Store") {
			// 构建文件的完整路径
			filePath := filepath.Join(dirPath, entry.Name())
			files = append(files, filePath)
		}
		//filePath := filepath.Join(dirPath, entry.Name())
		//files = append(files, filePath)
	}

	return files, nil
}

func ListFilesInCurrentDirFilter(dirPath, filter string) ([]string, error) {
	var files []string

	// 检查目录是否存在且确实是一个目录
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("目录不存在: %s", dirPath)
	}
	if err != nil {
		return nil, fmt.Errorf("获取目录 '%s' 信息失败: %w", dirPath, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("路径 '%s' 不是一个目录", dirPath)
	}

	// 读取目录内容
	// os.ReadDir 返回一个 []fs.DirEntry 切片
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("读取目录 '%s' 内容失败: %w", dirPath, err)
	}

	for _, entry := range entries {
		// 检查条目是否是文件 (而不是目录)
		if !strings.Contains(entry.Name(), "DS_Store") && strings.Contains(strings.ToLower(entry.Name()), filter) {
			// 构建文件的完整路径
			filePath := filepath.Join(dirPath, entry.Name())
			files = append(files, filePath)
		}
		//filePath := filepath.Join(dirPath, entry.Name())
		//files = append(files, filePath)
	}

	return files, nil
}

func ListFilesInCurrentDirFilter2(dirPath string, f func(entry os.DirEntry) bool) ([]string, error) {
	var files []string

	// 检查目录是否存在且确实是一个目录
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("目录不存在: %s", dirPath)
	}
	if err != nil {
		return nil, fmt.Errorf("获取目录 '%s' 信息失败: %w", dirPath, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("路径 '%s' 不是一个目录", dirPath)
	}

	// 读取目录内容
	// os.ReadDir 返回一个 []fs.DirEntry 切片
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("读取目录 '%s' 内容失败: %w", dirPath, err)
	}

	for _, entry := range entries {
		// 检查条目是否是文件 (而不是目录)
		if !strings.Contains(entry.Name(), "DS_Store") {

			if f(entry) {
				//if  strings.Contains(strings.ToLower(entry.Name()), filter) {
				// 构建文件的完整路径
				filePath := filepath.Join(dirPath, entry.Name())
				files = append(files, filePath)
			}

		}
		//filePath := filepath.Join(dirPath, entry.Name())
		//files = append(files, filePath)
	}

	return files, nil
}

// MoveFiles 批量移动匹配的文件到目标目录
func MoveFiles(pattern, destDir string) error {
	// 展开通配符，获取匹配的文件列表
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("解析通配符失败: %w", err)
	}

	if len(matches) == 0 {
		return fmt.Errorf("没有匹配的文件: %s", pattern)
	}

	// 确保目标目录存在
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return fmt.Errorf("创建目标目录失败: %w", err)
		}
	}

	for _, srcPath := range matches {
		if err := moveSingleFile(srcPath, destDir); err != nil {
			fmt.Printf("跳过文件 %s: %v\n", srcPath, err)
		} else {
			fmt.Printf("已移动: %s -> %s\n", srcPath, destDir)
		}
	}
	return nil

}

// moveSingleFile 移动单个文件，处理跨设备情况
func moveSingleFile(srcPath, destDir string) error {
	srcFileName := filepath.Base(srcPath)
	destPath := filepath.Join(destDir, srcFileName)

	// 检查源文件是否存在
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		return fmt.Errorf("源文件不存在: %w", err)
	}

	// 尝试直接重命名
	if err := os.Rename(srcPath, destPath); err == nil {
		return nil
	}

	// 重命名失败，尝试跨设备复制
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("打开源文件失败: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer destFile.Close()

	// 复制文件内容
	if _, err := io.Copy(destFile, srcFile); err != nil {
		return fmt.Errorf("复制文件内容失败: %w", err)
	}

	// 复制文件权限
	if err := os.Chmod(destPath, srcInfo.Mode()); err != nil {
		return fmt.Errorf("设置文件权限失败: %w", err)
	}

	// 删除源文件
	if err := os.Remove(srcPath); err != nil {
		return fmt.Errorf("删除源文件失败: %w", err)
	}

	return nil
}

func RenameDirectory(oldPath, newPath string) error {
	// 检查原目录是否存在
	if _, err := os.Stat(oldPath); os.IsNotExist(err) {
		return fmt.Errorf("原目录不存在: %s", oldPath)
	}

	// 执行重命名
	if err := os.Rename(oldPath, newPath); err != nil {
		return fmt.Errorf("重命名目录失败: %v", err)
	}
	if err := os.Remove(oldPath); err != nil {
		return fmt.Errorf("删除源文件失败: %w", err)
	}

	return nil
}

// 定义图片文件的扩展名
var imageExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}

// 检查文件是否为图片类型
func isImageFile(filename string) bool {
	ext := filepath.Ext(filename)
	for _, imageExt := range imageExtensions {
		if strings.ToLower(ext) == imageExt {
			return true
		}
	}
	return false
}

// 统计文件夹中的图片数量
func countImagesInDir(dir string) (int, error) {
	count := 0
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && isImageFile(path) {
			count++
		}
		return nil
	})
	return count, err
}

// 检查目录是否有子目录
func hasSubdirectories(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			return true, nil
		}
	}
	return false, nil
}

// 递归处理文件夹，只统计最深层级的文件夹
func ProcessDirs(root string) error {
	// 使用队列进行广度优先搜索，确保先处理父目录
	queue := []string{root}
	for len(queue) > 0 {
		dir := queue[0]
		queue = queue[1:]

		// 检查是否有子目录
		hasSubdirs, err := hasSubdirectories(dir)
		if err != nil {
			return err
		}

		if !hasSubdirs {
			// 如果没有子目录，统计图片数量
			count, err := countImagesInDir(dir)
			if err != nil {
				return err
			}
			if count > 10 {
				fmt.Printf("文件夹 %s 中的图片数量大于 10，数量为: %d\n", dir, count)
			}
		} else {
			// 如果有子目录，将子目录加入队列
			entries, err := os.ReadDir(dir)
			if err != nil {
				return err
			}
			for _, entry := range entries {
				if entry.IsDir() {
					queue = append(queue, filepath.Join(dir, entry.Name()))
				}
			}
		}
	}
	return nil
}
