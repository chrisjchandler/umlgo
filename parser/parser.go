package parser

import (
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "strings"
)

func ParseAndGenerateUML(code string) (string, error) {
    fset := token.NewFileSet()

    // Parse the Go code string
    node, err := parser.ParseFile(fset, "", code, parser.ParseComments)
    if err != nil {
        return "", fmt.Errorf("error parsing Go code: %v", err)
    }

    // Store function definitions
    funcDecls := make(map[string]*ast.FuncDecl)
    // Store function calls
    funcCalls := make(map[string][]string)

    // Traverse the AST to extract function definitions
    ast.Inspect(node, func(n ast.Node) bool {
        switch x := n.(type) {
        case *ast.FuncDecl:
            funcDecls[x.Name.Name] = x
        case *ast.CallExpr:
            if sel, ok := x.Fun.(*ast.SelectorExpr); ok {
                if ident, ok := sel.X.(*ast.Ident); ok {
                    if ident.Name == "http" { // Exclude standard library http calls
                        return true
                    }
                }
            }
            if ident, ok := x.Fun.(*ast.Ident); ok {
                funcName := ident.Name
                funcCalls[funcName] = append(funcCalls[funcName], funcName)
            }
        }
        return true
    })

    // Generate UML actions for function definitions
    var umlActions []string
    for fnName, fn := range funcDecls {
        var fnDescription string
        if fn.Doc != nil {
            fnDescription = fn.Doc.Text()
        }

        // Generate UML action based on function name and description
        umlAction := fmt.Sprintf("+ %s()", fnName)
        if fnDescription != "" {
            umlAction += " : " + strings.TrimSpace(fnDescription)
        }

        umlActions = append(umlActions, umlAction)
    }

    // Generate UML relations for function calls
    var umlRelations []string
    for caller, callees := range funcCalls {
        for _, callee := range callees {
            if _, exists := funcDecls[callee]; exists {
                umlRelation := fmt.Sprintf("%s --> %s", caller, callee)
                umlRelations = append(umlRelations, umlRelation)
            }
        }
    }

    // Generate final UML code with extracted actions and relations
    umlCode := fmt.Sprintf("@startuml\nclass DummyClass {\n    %s\n}\n\n%s\n@enduml",
        strings.Join(umlActions, "\n    "),
        strings.Join(umlRelations, "\n"))

    return umlCode, nil
}
