package ir

import "github.com/IBAX-io/needle/grammar/ast"

type Function struct {
	ast *ast.FuncDef

	Id               int32
	TreeType         ast.TreeType
	Name             string
	Parameters       []*Parameter
	ReturnParameters []*ReturnParameter
	Tails            map[string]*FuncTail
	Variadic         bool
	Auto             []string
	Body             *Body
}

func (b *Builder) processFunction(ast *ast.FuncDef) *Function {
	f := &Function{
		ast:      ast,
		Id:       ast.Id,
		Name:     ast.Name,
		TreeType: ast.TreeType,
		Tails:    make(map[string]*FuncTail),
	}

	if ast.FuncSignature.ParameterList != nil {
		f.Variadic = ast.FuncSignature.ParameterList.Variadic
		for i, list := range ast.FuncSignature.ParameterList.Parameter.NameList {
			for _, s := range list.Name {
				f.Parameters = append(f.Parameters, &Parameter{
					Name:     s,
					TypeName: ast.FuncSignature.ParameterList.Parameter.TypeName[i],
					Index:    i,
				})
			}
		}
	}

	for _, t := range ast.FuncSignature.FuncTail {
		f.Tails[t.Name] = b.processFuncTail(f, t)
	}

	if ast.FuncSignature.ReturnParameters != nil {
		for _, list := range ast.FuncSignature.ReturnParameters.TypeName {
			f.ReturnParameters = append(f.ReturnParameters, &ReturnParameter{
				TypeName: list,
			})
		}
	}
	if ast.Body != nil {
		f.Body = b.processBody(ast.Body)
	}
	return f
}

type Parameter struct {
	Name     string
	TypeName string
	Index    int
}

type ReturnParameter struct {
	TypeName string
}

type FuncTail struct {
	ast *ast.FuncTail

	Id       int32
	Parent   Function
	Name     string
	Params   []*Parameter
	Offset   []int
	Variadic bool
}

func (b *Builder) processFuncTail(parent *Function, ast *ast.FuncTail) *FuncTail {
	f := &FuncTail{
		ast:  ast,
		Id:   ast.Id,
		Name: ast.Name,
	}

	if ast.ParameterList != nil {
		f.Variadic = ast.ParameterList.Variadic
		if ast.ParameterList.Parameter == nil {
			return f
		}
		for i, list := range ast.ParameterList.Parameter.NameList {
			for _, s := range list.Name {
				f.Params = append(f.Params, &Parameter{
					Name:     s,
					TypeName: ast.ParameterList.Parameter.TypeName[i],
					Index:    i,
				})
				f.Offset = append(f.Offset, i+parent.Parameters[len(parent.Parameters)-1].Index+1)
			}
		}
	}

	return f
}
