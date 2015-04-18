package Util

import (
    . "github.com/qiniu/api/conf"
    "github.com/qiniu/api/rs"
    qio "github.com/qiniu/api/io"
    "io"
//    "github.com/qiniu/api/url"
    "fmt"
    "time"
    "mime/multipart"
    "path"
)

var qiniuManager *QiniuUploadManager
var Debug bool

type QiniuUploadManager struct {
    AccessKey string
    SecretKey string
    BucketName  string
}

// 定义返回值结构
type PutRets struct {
    qio.PutRet
    Bucket string `json:"bucket"`
    Fname string `json:"fname"`
    Fsize  float64 `json:"fsize"`
    MimeType string `json:"mimeType"`
    EndUser string `json:"endUser"`
    PersistentId string `json:"persistentId"`
//    Exif string `json:"exif"`
//    ImageInfo string `json:"imageInfo"`
    Ext string `json:"ext"`
    Uuid string `json:"uuid"`
}

func init() {
    Debug = false
}

// 获取七牛管理
func GetQiniuManager(accessKey string, secretKey string, bucketName string) *QiniuUploadManager {

    if accessKey == "" || secretKey == "" || bucketName == "" {
        return nil
    }

    if qiniuManager == nil {
        qiniuManager = new(QiniuUploadManager)
        qiniuManager.AccessKey = accessKey
        qiniuManager.SecretKey = secretKey
        qiniuManager.BucketName = bucketName

        ACCESS_KEY = accessKey
        SECRET_KEY = secretKey
    }
    return qiniuManager
}

// @Title 上传本地文件
// @Name UploadLocalFile
// @Param fileName string 文件名 如果为空,那么则使用默认的格式 2014-01-01_md5.ext
// @Param filePath string 文件路径
// @Return token string token
func (qn *QiniuUploadManager) UploadLocalFile(fileName string, filePath string) (PutRets ,error) {

    var err error
    var ret PutRets
    var extra = &qio.PutExtra{
        // Params:   params,
        // MimeType: mieType,
        // Crc32:    crc32,
        // CheckCrc: CheckCrc,
    }

    if fileName == "" {
        timeNow := time.Now().Format("2006-01-02")
        fileName = fmt.Sprintf("%s_%d", timeNow, time.Now().UnixNano())
    }

    err = qio.PutFile(nil, &ret, getUploadToken(qiniuManager.BucketName), fileName,filePath, extra)

    if err != nil {
        //上传产生错误
        fmt.Printf("io.PutFileWithoutKey is Failed: %v\n", err)
        return ret, err
    }

    fmt.Printf("ret result: %#v\n", ret)
    return ret, nil
}

func (qn *QiniuUploadManager) UploadIoFile(fileName string, r io.Reader, h *multipart.FileHeader) (PutRets ,error) {

    var err error
    var ret PutRets
    var extra = &qio.PutExtra{
        // Params:   params,
        // MimeType: mieType,
        // Crc32:    crc32,
        // CheckCrc: CheckCrc,
    }

    if fileName == "" {
        timeNow := time.Now().Format("2006-01-02")
        fileName = fmt.Sprintf("%s_%d%s", timeNow, time.Now().UnixNano(), path.Ext(h.Filename))
    }

    err = qio.Put(nil, &ret, getUploadToken(qiniuManager.BucketName), fileName, r, extra)

    if err != nil {
        //上传产生错误
        fmt.Printf("io.PutFileWithoutKey is Failed: %v\n", err)
        return ret, err
    }

    fmt.Printf("ret result: %#v\n", ret)
    return ret, nil
}

// @Title 获取UploadToken
// @Name getUploadToken
// @Param bucketName string 存储空间
// @Return token string token
func getUploadToken(bucketName string) string {
    putPolicy := rs.PutPolicy {
        Scope:         bucketName,
        //CallbackUrl: callbackUrl,
        //        CallbackBody: `{"name": $(fname)}`,
        //ReturnUrl:   returnUrl,
        ReturnBody:  `{"key": $(key),"bucket": $(bucket), "hash": $(etag), "fname": $(fname), "fsize": $(fsize),"mimeType": $(mimeType), "endUser": $(endUser), "persistentId": $(persistentId), "exif": $(exif),"imageInfo": $(imageInfo), "ext": $(ext), "uuid": $(uuid)}`,
        //AsyncOps:    asyncOps,
        //EndUser:     endUser,
        //Expires:     expires,
    }

    tokenStr := putPolicy.Token(nil)

    if  Debug {
        fmt.Printf("UploadToken : %s\n", tokenStr)
    }
    return tokenStr
}