package controllers

import (
	"github.com/simman/Image_Beego_Qiniu/Util"
	"github.com/simman/Image_Beego_Qiniu/models"
    "fmt"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.TplNames = "index.tpl"
}

func (this *HomeController) Hot() {

	pageSize := 20

	p := new(models.SmPhoto)

	_, pCount := p.GetPhotoCount()
	pager := this.SetPage(pageSize, pCount)

	err, s := p.GetAllPhotos(pager.Offset(), pageSize)

	if err == nil {
		this.Data["data"] = s
	} else {
		this.Ctx.WriteString("get photos data is error, error")
	}

	this.TplNames = "list.html"
}

func (this *HomeController) My() {
	this.TplNames = "404.tpl"
}

func (this *HomeController) Upload() {

	r, h, err := this.GetFile("files")

	defer func() {

		if rr := recover(); rr != nil {
			Logger.Warn("err = %v", err)

			errj := make(map[string]string)
			errj["error"] = err.Error()
			this.Data["json"] = errj
			this.ServeJson()
		}
	}()

	Logger.Debug("fileName : %s", h.Filename)

	qiniuManager := Util.GetQiniuManager(Configer.String("ACCESS_KEY"), Configer.String("SECRET_KEY"), Configer.String("BUCKET_NAME"))

	ret, err := qiniuManager.UploadIoFile("", r, h)

	if err != nil {
		Logger.Warning("Upload file is faild, err = %v", err)
	} else {
		Logger.Debug("ret = %#v", ret)

        sPhoto := new(models.SmPhoto)
        sPhoto.Name = ret.Fname
        sPhoto.Ext = ret.Ext
        sPhoto.Url = fmt.Sprintf("%s%s", Configer.String("BUCKET_URL"), ret.Fname)
        sPhoto.Type = ret.MimeType
        sPhoto.Size = int32(ret.Fsize)
        sPhoto.Sha1 = ret.Hash

        err := sPhoto.InsertOnPhoto()

        if err != nil {
            Logger.Warning("Insert Photo is Faild, err %v", err)
        }
	}

    this.Data["json"] = uploadSuccess(ret)
    this.ServeJson()
}

func uploadSuccess(ret Util.PutRets) map[string][]map[string]string {

	result := make(map[string][]map[string]string)
    result["files"] = make([]map[string]string, 10)
    result["files"][0] = make(map[string]string)

    result["files"][0]["deleteType"]    = "DELETE"
    result["files"][0]["deleteUrl"]     = ""
    result["files"][0]["name"]          = ret.Fname
    result["files"][0]["size"]          = fmt.Sprintf("%d", ret.Fsize)
    result["files"][0]["type"]          = ret.MimeType
    result["files"][0]["thumbnailUrl"]  = fmt.Sprintf("%s%s!thumb", Configer.String("BUCKET_URL"), ret.Fname)
    result["files"][0]["url"]           = fmt.Sprintf("%s%s", Configer.String("BUCKET_URL"), ret.Fname)

	return result
}
