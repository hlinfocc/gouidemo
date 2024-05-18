package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hlinfocc/gouidemo/assets"
	"github.com/hlinfocc/gouidemo/pkg/mytheme"
	"github.com/hlinfocc/gouidemo/pkg/web"
	"github.com/hlinfocc/gouidemo/pkg/websocket"
)

// func init() {
// 	//设置中文字体:解决中文乱码问题
// 	out, err := findfont.Find("./fonts/HarmonyOS_Sans_SC_Light.ttf")
// 	if err != nil {
// 		log.Fatal("无法加载字体文件:", err)
// 		fontPaths := findfont.List()
// 		for _, path := range fontPaths {
// 			log.Printf("字体文件路径: %s\n", path)
// 			if strings.Contains(path, "msyh.ttf") ||
// 				strings.Contains(path, "simhei.ttf") ||
// 				strings.Contains(path, "HarmonyOS_Sans_SC_Light.ttf") ||
// 				strings.Contains(path, "GB_ST_GB18030.ttf") ||
// 				strings.Contains(path, "DreamHanSerifCN-W10.ttf") ||
// 				strings.Contains(path, "simkai.ttf") {
// 				fontData, err := os.ReadFile(path)
// 				if err != nil {
// 					log.Fatal("无法加载字体文件:", err)
// 					panic(err)
// 				}
// 				_, err = opentype.Parse(fontData)
// 				if err != nil {
// 					log.Fatal("无法解析字体数据:", err)
// 					panic(err)
// 				}
// 				os.Setenv("FYNE_FONT", path)
// 				break
// 			}
// 		}
// 	} else {
// 		fontData, err := os.ReadFile(out)
// 		if err != nil {
// 			log.Fatal("无法加载字体文件:", err)
// 			panic(err)
// 		}
// 		_, err = opentype.Parse(fontData)
// 		if err != nil {
// 			log.Fatal("无法解析字体数据:", err)
// 			panic(err)
// 		}
// 		os.Setenv("FYNE_FONT", out)
// 	}
// }

func main() {
	myApp := app.NewWithID("net.hlinfo.gouidemo")

	myApp.SetIcon(assets.ResourceLogoPng)

	myWindow := myApp.NewWindow("golang UI之fyne示例")

	// 设置窗口的大小，例如宽度为600，高度为400
	myWindow.SetMaster()
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.CenterOnScreen()
	myWindow.FixedSize()

	myApp.Settings().SetTheme(mytheme.NewTheme())

	ok := "为天地立心，为生民立命，为往圣继绝学，为万事开太平。"

	hello := widget.NewLabel(ok)

	myWindow.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("道可道，非常道；名可名，非常名。无名天地之始；有名万物之母。故常无欲，以观其妙；常有欲，以观其徼。")
		}),
	))
	// myWindow.SetCloseIntercept(func() {
	// 	myWindow.Hide()
	// })
	// 开启web服务
	go web.StartWebServer()
	// 开启websocket服务
	go websocket.StartWebsocket()
	// 启动窗口
	myWindow.ShowAndRun()

}

// func myTheme() {
// 	// 加载自定义字体
// 	customFont, err := loadCustomFont("assets/your-font-file.ttf")
// 	if err != nil {
// 		fyne.LogError("Failed to load custom font", err)
// 		return

// 	}
// 	// 创建一个使用自定义字体的主题

// }

// // loadCustomFont 加载自定义字体文件

// func loadCustomFont(path string) (fyne.Resource, error) {

// 	return fyne.LoadResourceFromPath(path)

// }
