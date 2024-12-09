package main

// https://adventofcode.com/2024/day/9

import (
	"aoc/helper"
	"fmt"
	"sort"
)

func main() {
	str := helper.ReadString("example-1.txt")
	disk := parseDisk(str)

	disk1 := disk.Clone()
	disk1.Compact1()
	solution1 := disk1.Checksum()
	fmt.Println("-> part 1:", solution1)

	disk2 := disk.Clone()
	disk2.Compact2()
	solution2 := disk2.Checksum()
	fmt.Println("-> part 2:", solution2, "(MAYBE WRONG!)")

	// 8648109921625 is too high
	// 6458577412415 is too low
}

const (
	FileIDNone FileID = -1
)

type FileID int64

type File struct {
	ID   FileID
	Pos  int
	Size int
}

type Disk struct {
	blocks                  []FileID
	files                   []File
	leftmostEmptySpaceIndex int
	emptySpaces             map[int][]int
	maxEmptySpaceLen        int
}

func (d *Disk) Clone() *Disk {
	return &Disk{
		blocks:                  helper.CloneSlice(d.blocks),
		files:                   helper.CloneSlice(d.files),
		leftmostEmptySpaceIndex: d.leftmostEmptySpaceIndex,
		emptySpaces:             helper.CloneMapOfSlices(d.emptySpaces),
		maxEmptySpaceLen:        d.maxEmptySpaceLen,
	}
}

func parseDisk(str string) *Disk {
	blocks := make([]FileID, 0)
	files := make([]File, 0)
	emptySpaces := make(map[int][]int)
	var maxEmptySpaceLen int
	for i, r := range str {
		if r >= '0' && r <= '9' {
			blockLen := int(r - '0')
			fileID := FileIDNone
			if i%2 == 0 {
				fileID = FileID(i / 2)
				files = append(files, File{
					ID:   fileID,
					Pos:  len(blocks),
					Size: blockLen,
				})
			} else {
				if blockLen > 0 {
					if blockLen > maxEmptySpaceLen {
						maxEmptySpaceLen = blockLen
					}
					emptySpaces[blockLen] = append(emptySpaces[blockLen], len(blocks))
				}
			}
			for j := 0; j < blockLen; j++ {
				blocks = append(blocks, fileID)
			}
		}
	}
	return &Disk{
		blocks:                  blocks,
		files:                   files,
		leftmostEmptySpaceIndex: -1,
		emptySpaces:             emptySpaces,
		maxEmptySpaceLen:        maxEmptySpaceLen,
	}
}

func (d *Disk) Compact1() {
	for i := len(d.blocks) - 1; i >= 0; i-- {
		d.selectNextEmptySpaceIndex()

		if d.leftmostEmptySpaceIndex >= i {
			break
		}

		if d.blocks[i] != FileIDNone {
			d.blocks[d.leftmostEmptySpaceIndex] = d.blocks[i]
			d.blocks[i] = FileIDNone
		}
	}
}

func (d *Disk) selectNextEmptySpaceIndex() {
	for d.leftmostEmptySpaceIndex < len(d.blocks) {
		if d.leftmostEmptySpaceIndex >= 0 && d.blocks[d.leftmostEmptySpaceIndex] == FileIDNone {
			return
		}
		d.leftmostEmptySpaceIndex++
	}
	panic("no empty space left")
}

func (d *Disk) Checksum() int64 {
	var checksum int64
	for i, f := range d.blocks {
		if f != FileIDNone {
			checksum += int64(i) * int64(f)
		}
	}
	return checksum
}

func (d *Disk) Compact2() {
	for i := len(d.files) - 1; i >= 0; i-- {
		pos, emptyBlockLen, ok := d.findFittingEmptySpace(d.files[i].Size)
		if ok {
			if pos >= d.files[i].Pos {
				continue
			}

			for j := 0; j < d.files[i].Size; j++ {
				d.blocks[pos+j] = d.files[i].ID
				d.blocks[d.files[i].Pos+j] = FileIDNone
			}
			d.files[i].Pos = pos

			d.emptySpaces[emptyBlockLen] = d.emptySpaces[emptyBlockLen][1:]
			remainingBlockLen := emptyBlockLen - d.files[i].Size
			if remainingBlockLen > 0 {
				d.emptySpaces[remainingBlockLen] = append(d.emptySpaces[remainingBlockLen], pos+d.files[i].Size)
				sort.Ints(d.emptySpaces[remainingBlockLen])
				fmt.Println(remainingBlockLen, "->", d.emptySpaces[remainingBlockLen])
			}
		}
	}
}

func (d *Disk) findFittingEmptySpace(blockLen int) (int, int, bool) {
	for i := blockLen; i <= d.maxEmptySpaceLen; i++ {
		if arr, ok := d.emptySpaces[i]; ok && len(arr) > 0 {
			return arr[0], i, true
		}
	}
	return -1, -1, false
}
