package main

// https://adventofcode.com/2024/day/9

import (
	"aoc/helper"
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := helper.ReadString("input.txt")
	disk := parseDisk(str)

	disk1 := disk.Clone()
	disk1.SplitBlocks()
	disk1.Compact()
	solution1 := disk1.Checksum()
	fmt.Println("-> part 1:", solution1)

	disk2 := disk.Clone()
	disk2.Compact()
	solution2 := disk2.Checksum()
	fmt.Println("-> part 2:", solution2)
}

const (
	FileIDNone FileID = -1
)

type FileID int

type Block struct {
	Pos  int
	Size int
	ID   FileID
}

type Disk struct {
	fileBlocks  []Block
	emptyBlocks []Block
	size        int
}

func (d *Disk) String() string {
	runes := []rune(strings.Repeat(" ", d.size))
	for _, b := range d.fileBlocks {
		for i := 0; i < b.Size; i++ {
			runes[b.Pos+i] = '0' + rune(b.ID%10)
		}
	}
	for _, b := range d.emptyBlocks {
		for i := 0; i < b.Size; i++ {
			runes[b.Pos+i] = '.'
		}
	}
	return string(runes)
}

func (d *Disk) Clone() *Disk {
	return &Disk{
		fileBlocks:  helper.CloneSlice(d.fileBlocks),
		emptyBlocks: helper.CloneSlice(d.emptyBlocks),
		size:        d.size,
	}
}

func parseDisk(str string) *Disk {
	fileBlocks := make([]Block, 0)
	emptyBlocks := make([]Block, 0)
	var pos int
	for i, r := range str {
		if r >= '0' && r <= '9' {
			blockLen := int(r - '0')
			if i%2 == 0 {
				fileBlocks = append(fileBlocks, Block{
					Pos:  pos,
					Size: blockLen,
					ID:   FileID(i / 2),
				})

			} else {
				if blockLen > 0 {
					emptyBlocks = append(emptyBlocks, Block{
						Pos:  pos,
						Size: blockLen,
						ID:   FileIDNone,
					})
				}
			}

			pos += blockLen
		}
	}
	return &Disk{
		fileBlocks:  fileBlocks,
		emptyBlocks: emptyBlocks,
		size:        pos,
	}
}

func (d *Disk) SplitBlocks() {
	splittedFiles := make([]Block, 0)
	for _, f := range d.fileBlocks {
		for i := 0; i < f.Size; i++ {
			splittedFiles = append(splittedFiles, Block{
				Pos:  f.Pos + i,
				Size: 1,
				ID:   f.ID,
			})
		}
	}
	d.fileBlocks = splittedFiles
}

func (d *Disk) Compact() {
	for i := len(d.fileBlocks) - 1; i >= 0; i-- {
		if d.fileBlocks[i].Size == 0 {
			// nothing to move
			continue
		}

		emptyIndex, ok := d.findEmptySpaceBlockIndex(d.fileBlocks[i].Size)
		if !ok {
			// cannot move
			continue
		}

		if d.emptyBlocks[emptyIndex].Pos >= d.fileBlocks[i].Pos {
			// already leftmost index
			continue
		}

		d.fileBlocks[i].Pos = d.emptyBlocks[emptyIndex].Pos
		if d.emptyBlocks[emptyIndex].Size > d.fileBlocks[i].Size {
			d.emptyBlocks[emptyIndex].Pos += d.fileBlocks[i].Size
			d.emptyBlocks[emptyIndex].Size -= d.fileBlocks[i].Size
		} else {
			d.emptyBlocks = helper.RemoveIndex(d.emptyBlocks, emptyIndex)
		}
	}

	sort.Slice(d.fileBlocks, func(i, j int) bool {
		return d.fileBlocks[i].Pos < d.fileBlocks[j].Pos
	})
}

func (d *Disk) findEmptySpaceBlockIndex(blockLen int) (int, bool) {
	for i := range d.emptyBlocks {
		if d.emptyBlocks[i].Size >= blockLen {
			return i, true
		}
	}
	return -1, false
}

func (d *Disk) Checksum() int64 {
	var checksum int64
	for _, f := range d.fileBlocks {
		for i := 0; i < f.Size; i++ {
			checksum += int64((f.Pos + i) * int(f.ID))
		}
	}
	return checksum
}
