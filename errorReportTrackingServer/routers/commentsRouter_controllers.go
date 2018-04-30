package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"] = append(beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"] = append(beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"] = append(beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"] = append(beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"] = append(beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"] = append(beego.GlobalControllerRouter["errorReportTrackingServer/controllers:TTrackingrecordController"],
		beego.ControllerComments{
			Method: "GetbyTime",
			Router: `/time`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})
}
