package main

import (
    "path/filepath"
    "fmt"
    "os"
    "flag"
    "reflect"
    "./pkg"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
    switch event.EventType {
    case oss.TransferStartedEvent:
        fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
            event.ConsumedBytes, event.TotalBytes)
    case oss.TransferDataEvent:
        fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
            event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
    case oss.TransferCompletedEvent:
        fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
            event.ConsumedBytes, event.TotalBytes)
    case oss.TransferFailedEvent:
        fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
            event.ConsumedBytes, event.TotalBytes)
    default:
    }
}

func main() {
    // 读取配置文件
    myConfig := new(conf.Config)
    myConfig.InitConfig("oss.config")

    readVariable := myConfig.Read("default", "readVariable")
    
    var Endpoint, AccessKeyId, AccessKeySecret, bucketName, localFile, remoteFolder string
    if readVariable == "1" {
        // 读取环境变量
        Endpoint = os.Getenv("Endpoint")
        AccessKeyId = os.Getenv("AccessKeyId")
        AccessKeySecret = os.Getenv("AccessKeySecret")

        bucketName = os.Getenv("bucketName")
        remoteFolder = os.Getenv("remoteFolder")
        localFile = os.Getenv("localFile")
    } else { // 读取配置文件
        Endpoint = myConfig.Read("oss", "Endpoint")
        AccessKeyId = myConfig.Read("oss", "AccessKeyId")
        AccessKeySecret = myConfig.Read("oss", "AccessKeySecret")

        bucketName = myConfig.Read("upload", "bucketName")
        remoteFolder = myConfig.Read("upload", "remoteFolder")
        localFile = myConfig.Read("upload", "localFile")
    }

    // 上传到根目录
    if remoteFolder == "/" {
        remoteFolder = ""
    }

    // 接受终端输入参数
    flag.Parse()
    para1 := flag.Arg(0)

    if !isEmpty(para1) {
        fmt.Printf("Input parameters: %s\n", para1);
        localFile = para1
    }

    if isEmpty(localFile) {
        fmt.Println("Please take parameters. Local file/folder path not define.")
        os.Exit(-1)
    }

    fmt.Println("OSS Go SDK Version: ", oss.Version)
    // 创建OSSClient
    client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    // 获取存储空间。
    bucket, err := client.Bucket(bucketName)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    // 获取要上传文件列表
    fmt.Println("Need to upload file lists：")
    fileList, err := getFilelist(localFile)
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
    for _, v := range fileList {
        fmt.Println(v)
    }

    fmt.Println("")

    // 上传
    fmt.Println("Uploading...")
    for _, v := range fileList {
        objectName := remoteFolder+v
        localFile = v
        err = bucket.PutObjectFromFile(objectName, localFile, oss.Progress(&OssProgressListener{}))
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(-1)
        }
    }
}

// 获取文件列表
func getFilelist(path string) ([]string, error) {
    var fileList []string
    err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
        if ( f == nil ) {
            return err
        }

        if f.IsDir() {
            return nil
        }
        fileList = append(fileList, path);
        return nil
    })
    return fileList, err
}

// 判断值是否为空
func isEmpty(a interface{}) bool {
    v := reflect.ValueOf(a)
    if v.Kind() == reflect.Ptr {
        v=v.Elem()
    }
    return v.Interface() == reflect.Zero(v.Type()).Interface()
}