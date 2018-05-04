package drawing

import (
	"github.com/GUMI-golang/gumi/gcore"
)

var Drawing _Drawing

type _Drawing struct {
	Ruler
}

type Ruler struct {
	// Global Value
	RulerWidth float32
	//
	Graduation
	Grid
	Hint
}
type (
	Graduation struct{}
	Grid struct {}
	Hint struct {}
)

func init() {
	Drawing.Ruler.RulerWidth = 4
}

var interpreter = gcore.FunctionCaller{
	Fns:[]gcore.Function{
		{
			Name:"graduation.vertical",
			Args:[]gcore.FunctionArgument{gcore.FnArgFloat},
			Callback: func(args ...interface{}) []interface{} {
				return []interface{}{Drawing.Graduation.Vertical(float32(args[0].(float64)))}
			},
		},
		{
			Name:"graduation.horizontal",
			Args:[]gcore.FunctionArgument{gcore.FnArgFloat},
			Callback: func(args ...interface{}) []interface{} {
				return []interface{}{Drawing.Graduation.Horizontal(float32(args[0].(float64)))}
			},
		},
		{
			Name:"grid.vertical",
			Args:[]gcore.FunctionArgument{gcore.FnArgFloat},
			Callback: func(args ...interface{}) []interface{} {
				return []interface{}{Drawing.Grid.Vertical(float32(args[0].(float64)))}
			},
		},
		{
			Name:"grid.horizontal",
			Args:[]gcore.FunctionArgument{gcore.FnArgFloat},
			Callback: func(args ...interface{}) []interface{} {
				return []interface{}{Drawing.Grid.Horizontal(float32(args[0].(float64)))}
			},
		},
		{
			Name:"hint.vertical",
			Args:[]gcore.FunctionArgument{gcore.FnArgFloat},
			Callback: func(args ...interface{}) []interface{} {
				return []interface{}{Drawing.Hint.Vertical(float32(args[0].(float64)))}
			},
		},
		{
			Name:"hint.horizontal",
			Args:[]gcore.FunctionArgument{gcore.FnArgFloat},
			Callback: func(args ...interface{}) []interface{} {
				return []interface{}{Drawing.Hint.Horizontal(float32(args[0].(float64)))}
			},
		},
		{
			Name:"fps",
			Args:[]gcore.FunctionArgument{gcore.FnArgString},
			Callback: func(args ...interface{}) []interface{} {
				return []interface{}{Drawing.FPS(args[0].(string))}
			},
		},
	},
}