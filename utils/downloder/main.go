package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/iikira/BaiduPCS-Go/pcsutil/converter"
	"github.com/iikira/BaiduPCS-Go/pcsverbose"
	"github.com/iikira/BaiduPCS-Go/requester"
	"github.com/iikira/BaiduPCS-Go/requester/downloader"
)

const (
	//StrDownloadInitError 初始化下载发生错误
	StrDownloadInitError = "初始化下载发生错误"
)

var (
	parallel       int
	cacheSize      int
	test           bool
	isPrintStatus  bool
	downloadSuffix = ".downloader_downloading"
)

func init() {
	flag.IntVar(&parallel, "p", 5, "download max parallel")
	flag.IntVar(&cacheSize, "c", 30000, "download cache size")
	flag.BoolVar(&pcsverbose.IsVerbose, "verbose", false, "verbose")
	flag.BoolVar(&test, "test", false, "test download")
	flag.BoolVar(&isPrintStatus, "status", false, "print status")
	flag.StringVar(&requester.UserAgent, "ua", "", "User-Agent")

	flag.Parse()
}

func download(id int, downloadURL, savePath string, client *requester.HTTPClient, newCfg downloader.Config) error {
	var (
		file     *os.File
		writerAt io.WriterAt
		err      error
		exitChan chan struct{}
	)

	if !newCfg.IsTest {
		newCfg.InstanceStatePath = savePath + downloadSuffix

		// 创建下载的目录
		dir := filepath.Dir(savePath)
		fileInfo, err := os.Stat(dir)
		if err != nil {
			err = os.MkdirAll(dir, 0777)
			if err != nil {
				return err
			}
		} else if !fileInfo.IsDir() {
			return fmt.Errorf("%s, path %s: not a directory", StrDownloadInitError, dir)
		}

		file, err = os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY, 0666)
		if file != nil {
			defer file.Close()
		}
		if err != nil {
			return fmt.Errorf("%s, %s", StrDownloadInitError, err)
		}

		// 空指针和空接口不等价
		if file != nil {
			writerAt = file
		}
	}

	download := downloader.NewDownloader(downloadURL, writerAt, &newCfg)
	download.SetClient(client)
	download.SetFirstCheckMethod("GET")

	exitChan = make(chan struct{})

	download.OnExecute(func() {
		if isPrintStatus {
			go func() {
				for {
					time.Sleep(1 * time.Second)
					select {
					case <-exitChan:
						return
					default:
						download.PrintAllWorkers()
					}
				}
			}()
		}

		if newCfg.IsTest {
			fmt.Printf("[%d] 测试下载开始\n\n", id)
		}

		var (
			ds                            = download.GetDownloadStatusChan()
			format                        = "\r[%d] ↓ %s/%s %s/s in %s, left %s ............"
			downloaded, totalSize, speeds int64
			leftStr                       string
		)
		for {
			select {
			case <-exitChan:
				return
			case v, ok := <-ds:
				if !ok { // channel 已经关闭
					return
				}

				downloaded, totalSize, speeds = v.Downloaded(), v.TotalSize(), v.SpeedsPerSecond()
				if speeds <= 0 {
					leftStr = "-"
				} else {
					leftStr = (time.Duration((totalSize-downloaded)/(speeds)) * time.Second).String()
				}

				fmt.Printf(format, id,
					converter.ConvertFileSize(v.Downloaded(), 2),
					converter.ConvertFileSize(v.TotalSize(), 2),
					converter.ConvertFileSize(v.SpeedsPerSecond(), 2),
					v.TimeElapsed()/1e7*1e7, leftStr,
				)
			}
		}
	})

	err = download.Execute()
	close(exitChan)
	fmt.Printf("\n")
	if err != nil {
		// 下载失败, 删去空文件
		if info, infoErr := file.Stat(); infoErr == nil {
			if info.Size() == 0 {
				pcsverbose.Verbosef("[%d] remove empty file: %s\n", id, savePath)
				os.Remove(savePath)
			}
		}
		return err
	}

	if !newCfg.IsTest {
		fmt.Printf("[%d] 下载完成, 保存位置: %s\n", id, savePath)
	} else {
		fmt.Printf("[%d] 测试下载结束\n", id)
	}

	return nil
}

func main() {
	if flag.NArg() == 0 {
		flag.Usage()
		if runtime.GOOS == "windows" {
			bufio.NewReader(os.Stdin).ReadByte()
		}

		return
	}

	for k := range flag.Args() {
		var (
			savePath string
			err      error
		)
		if !test {
			savePath, err = downloader.GetFileName(flag.Arg(k), nil)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		download(k, flag.Arg(k), savePath, nil, downloader.Config{
			MaxParallel:       parallel,
			CacheSize:         cacheSize,
			InstanceStatePath: savePath + downloadSuffix,
			IsTest:            test,
		})
	}
	fmt.Println()
}
