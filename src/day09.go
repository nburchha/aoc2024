package main

import (
	"os"
	"strings"
	"fmt"
	"strconv"
)

func parseDiskMap(diskMap string) []int {
	var disk []int
	fileID := 0
	isFile := true

	for i := 0; i < len(diskMap); i++ {
		blockSize, _ := strconv.Atoi(string(diskMap[i]))
		for j := 0; j < blockSize; j++ {
			if isFile {
				disk = append(disk, fileID)
			} else {
				disk = append(disk, -1)
			}
		}
		if isFile {
			fileID++
		}
		isFile = !isFile
	}
	return disk
}

func compactDisk1(disk []int) []int {
	writePos := 0
	for readPos := len(disk)-1; readPos >= 0; readPos-- {
		if disk[readPos] != -1 {
			for writePos < len(disk) -1 && disk[writePos] != -1 {
				writePos++
			}
			if readPos <= writePos { break }
			disk[writePos] = disk[readPos]
			if writePos != readPos {
				disk[readPos] = -1
			}
		}
	}
	return disk
}

func compactDisk2(disk []int) []int {
	fileRanges := make(map[int][2]int)

	currentFile := -1
	for i, block := range disk {
		if block != -1 {
			if block != currentFile {
				currentFile = block
				fileRanges[block] = [2]int{i, i}
			} else {
				rangeInfo := fileRanges[block]
				fileRanges[block] = [2]int{rangeInfo[0], i}
			}
		}
	}

	for fileID := len(fileRanges) - 1; fileID >= 0; fileID-- {
		if _, exists := fileRanges[fileID]; !exists {
			continue
		}
		start, end := fileRanges[fileID][0], fileRanges[fileID][1]
		fileLength := end - start + 1

		freeStart := -1
		freeLength := 0
		for i := 0; i < start; i++ {
			if disk[i] == -1 {
				if freeStart == -1 {
					freeStart = i
				}
				freeLength++
				if freeLength == fileLength {
					break
				}
			} else {
				freeStart = -1
				freeLength = 0
			}
		}

		if freeLength == fileLength {
			for i := 0; i < fileLength; i++ {
				disk[freeStart+i] = fileID
				disk[start+i] = -1
			}
		}
	}

	return disk
}

func calculateChecksum(disk []int) int {
	checksum := 0
	for pos, block := range disk {
		if block != -1 {
			checksum += pos * block
		}
	}
	return checksum
}

func main() {
	input, _ := os.ReadFile("input/day09.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	diskMap := split[0]
	intDiskMap:= parseDiskMap(diskMap)
	fmt.Println(intDiskMap)
	// intDiskMap = compactDisk1(intDiskMap)
	intDiskMap = compactDisk2(intDiskMap)
	fmt.Println(calculateChecksum(intDiskMap))
}