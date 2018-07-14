

package controllers
import (
		"fmt"
		"log"
		"net/http"
		"os"
		"io"
		"strings"
		"bufio"
		"LogConverter/model"
		valid "github.com/asaskevich/govalidator"
		)

type FileLogController struct {
	BaseController
}

func (c *FileLogController)FileLog(){
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == http.MethodPost{
		var  dataSlice [][]string
		file,header,err :=r.FormFile("uploadfile")
			if err !=nil{
				log.Println("uploading error",err)
				http.Error(w,"error in uploading file",http.StatusInternalServerError)
				return
			}
			fmt.Println("inside json")
		if _, err := os.Stat("./testUploadJson/"); os.IsNotExist(err) {

			os.Mkdir("./testUploadJson/",os.ModePerm)
		}
		jsonFile, err := os.OpenFile("./testUploadJson/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		io.Copy(jsonFile, file)
		defer file.Close()

		fmt.Println("string contain")
		f, _ := os.Open("./testUploadJson/"+header.Filename)
		scanner := bufio.NewScanner(f)
		 var fileSlice []string


		for scanner.Scan() {
			line := scanner.Text()

			// Split the line on commas.
			parts := strings.Split(line, "\n")
			fileSlice=append(fileSlice,parts[0])

		}
		for i:=0;i<len(fileSlice);i++{
			if len(fileSlice[i])>0{
				var logDetails []string
				dataParts  :=strings.Split(fileSlice[i],"|")
				for j:=0;j<len(dataParts);j++{
					logDetails=append(logDetails,dataParts[j])
				}
				fmt.Println()
				firstDate :=logDetails[0]
				firstDataNumber :=firstDate[0:4]
				if valid.IsInt(firstDataNumber){
					logDetails=append(logDetails,header.Filename)
					logDetails=append(logDetails,"first_format")

				}else{
					logDetails=append(logDetails,header.Filename)
					logDetails=append(logDetails,"second_format")

				}
				dataSlice=append(dataSlice,logDetails)
			}
		}
		logData := model.LogDetails{}
       dbStatus :=logData.InsertIntoDb(dataSlice)
       switch dbStatus{
	   case true:
		   c.TplName ="templates/uploaded.html"
	   case false:
		   c.TplName ="templates/fileUpload.html"
	   }

		}else{
			c.TplName ="templates/fileUpload.html"
		}

	}



