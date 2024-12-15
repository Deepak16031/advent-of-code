package main

import (
	"slices"
	"strconv"
)

type Block struct {
	fileId   int
	position int
	size     int
}

func Solution2(input []string) int {
	var usedBlocks []*Block
	var freeBlocks []*Block
	fileId := 0
	pos := 0
	for i := 0; i < len(input); i++ {
		size, _ := strconv.Atoi(input[i])
		if i%2 == 0 {
			usedBlocks = append(usedBlocks, &Block{fileId, pos, size})
			fileId++
		} else {
			freeBlocks = append(freeBlocks, &Block{-1, pos, size})
		}
		pos += size
	}

	slices.Reverse(usedBlocks)

	for _, usedBlock := range usedBlocks {
		for _, freeBlock := range freeBlocks {
			if freeBlock.size >= usedBlock.size && freeBlock.position < usedBlock.position {
				usedBlock.position = freeBlock.position
				freeBlock.size -= usedBlock.size
				freeBlock.position += usedBlock.size
				break
			}
		}

	}

	sum := 0

	slices.SortFunc(usedBlocks, func(a, b *Block) int {
		return a.position - b.position
	})
	for _, usedBlock := range usedBlocks {
		start := usedBlock.position - 1
		end := usedBlock.position + usedBlock.size - 1
		fileSum := ((end * (end + 1) / 2) - (start * (start + 1) / 2)) * usedBlock.fileId
		//fmt.Printf("%+v, %v\n", usedBlock, fileSum)
		sum += fileSum
	}

	return sum
}
