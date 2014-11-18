package main

import (
        "fmt"
        "io/ioutil"
        "os"
        "syscall"
        "time"
)

func timespecToTime(ts syscall.Timespec) time.Time {
        return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func main() {
        fi, err := ioutil.ReadDir(os.Args[1])

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        for _, dir := range fi {
                fmt.Printf("Name: %v \n", dir.Name())
                fmt.Printf("Size: %v \n", dir.Size())
                fmt.Printf("Mode: %v \n", dir.Mode())
                fmt.Printf("ModTime: %v \n", dir.ModTime())
                fmt.Printf("IsDir: %v \n", dir.IsDir())
                fmt.Println()

                stat := dir.Sys().(*syscall.Stat_t)

                fmt.Printf("Sys: %T \n", stat)
                fmt.Printf("Stat_t.Dev: %v \n", stat.Dev)
                fmt.Printf("Stat_t.Ino: %v \n", stat.Ino)
                fmt.Printf("Stat_t.Mode: %v \n", stat.Mode)
                fmt.Printf("Stat_t.Nlink: %v \n", stat.Nlink)
                fmt.Printf("Stat_t.Uid: %v \n", stat.Uid)
                fmt.Printf("Stat_t.Gid: %v \n", stat.Gid)
                fmt.Printf("Stat_t.Rdev: %v \n", stat.Rdev)
                fmt.Printf("Stat_t.Size: %v \n", stat.Size)
                // Atim, Mtim, Ctim undefined for Darwin
                //fmt.Printf("Stat_t.Atim: %v \n", timespecToTime(stat.Atim))
                //fmt.Printf("Stat_t.Mtim: %v \n", timespecToTime(stat.Mtim))
                //fmt.Printf("Stat_t.Ctim: %v \n", timespecToTime(stat.Ctim))
                fmt.Printf("Stat_t.Blksize: %v \n", stat.Blksize)
                // Pad_cgo_0 undefined for Linux
                //fmt.Printf("Stat_t.Pad_cgo_0: %v \n", stat.Pad_cgo_0)
                fmt.Printf("Stat_t.Blocks: %v \n", stat.Blocks)
                // Fstype undefined for Darwin and Linux
                //fmt.Printf("Stat_t.Fstyep: %v \n", stat.Fstype)

                fmt.Println()
        }
}

