package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/vavilen84/gocommerce/auth"
	"github.com/vavilen84/gocommerce/models"

	_ "image/jpeg"
	_ "image/png"
	"os"
	"path"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) setAuthData() {
	token := auth.ParseToken(c.Ctx)
	c.Data["IsLoggedIn"] = token.IsLoggedIn
	if token.User != "" {
		var user models.User
		err := json.Unmarshal([]byte(token.User), &user)
		if err != nil {
			beego.Error(err)
		}

		c.Data["User"] = user

		o := orm.NewOrm()
		_, err = models.FindUserById(o, user.Id)
		if err == orm.ErrNoRows {
			auth.Logout(c.Ctx)
		}

		if user.Role == models.RoleAdmin {
			c.Data["IsAdmin"] = true

		} else {
			c.Data["IsAdmin"] = false
		}
	}
}

func (c *BaseController) setResponseData(title, templateName string) {
	c.setRenderData(title, templateName)
	c.setAuthData()
	c.Data["AWSS3Url"] = "https://fotokolo-local.s3.eu-central-1.amazonaws.com"
}

func (c *BaseController) setRenderData(title, templateName string) {
	c.Data["title"] = title
	c.Layout = "layout.html"
	c.TplName = templateName + ".html"
}

func (c *BaseController) saveFormFileImageToS3(imageFormName string) (filepath, originalFilename, uuid string) {
	file, header, err := c.GetFile(imageFormName)
	if err != nil {
		beego.Error(err)
		return "", "", ""
	}
	originalFilename = header.Filename
	ext, err := utils.GetImageExtension(file)
	if err != nil {
		beego.Error(err)
		return "", "", ""
	}
	uuid = utils.GenerateUUID()
	filename := uuid + "." + ext
	subDir := s3.GenerateSubfolderName(filename)
	tmpDir := utils.GetTmpDir(subDir)
	filepath = path.Join(tmpDir, filename)
	err = os.MkdirAll(tmpDir, 0775)
	if err != nil {
		beego.Error(err)
		return "", "", ""
	}
	err = c.SaveToFile(imageFormName, filepath)
	if err != nil {
		beego.Error(err)
		return "", "", ""
	}
	err = s3.SaveImageToS3(subDir, filename)
	if err != nil {
		beego.Error(err)
		return
	}
	err = os.Remove(filepath)
	if err != nil {
		beego.Error(err)
		return
	}
	filepath = path.Join(subDir, filename)
	return
}
