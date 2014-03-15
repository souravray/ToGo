package app

import ( 
	"github.com/robfig/revel"
)

// var MyFilter = func(c *revel.Controller, fc []revel.Filter) {
// 	user:=models.User{}

// 	if username, ok := c.Session["user"]; ok {
// 		if err:=c.Orm.Where("name = ?", username).First(&user).Error;  err==nil {
// 			fmt.Println("==== Inside My filter ======", c.Session["user"])
// 		}
// 	}
// 	fc[0](c, fc[1:]) // Execute the next filter stage.
// }

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}
}
