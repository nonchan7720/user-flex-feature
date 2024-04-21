package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/i18n"
)

func main() {
	mode := os.Args[1]
	root := os.Args[2]
	exeDir := os.Args[3]
	switch strings.ToLower(mode) {
	case "original":
		originalMessage(root, exeDir)
	case "validation":
		validationMessage(strings.Split(root, ";"), exeDir)
	}
}

func originalMessage(root string, exeDir string) {
	allKeys := map[string]struct{}{}
	filenames := []string{"ja.yaml", "en.yaml"}
	for _, file := range filenames {
		result := i18n.TranslateLanguage{
			Keys: map[string]string{},
		}
		filename := filepath.Join(exeDir, file)
		if _, err := os.Stat(filename); err == nil {
			readFile(filename, &result)
		}
		fset := token.NewFileSet()

		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(info.Name(), "_test.go") {
				return nil
			}
			if strings.HasSuffix(info.Name(), ".go") {
				values := parseFile(fset, path)
				for _, value := range values {
					if _, ok := result.Keys[value]; !ok {
						result.Keys[value] = ""
					}
					allKeys[value] = struct{}{}
				}
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		keys := make([]string, 0, len(result.Keys))
		for key := range result.Keys {
			keys = append(keys, key)
		}
		for _, key := range keys {
			if _, ok := allKeys[key]; !ok {
				delete(result.Keys, key)
			}
		}
		writeFile(filename, result)
	}
}

func validationMessage(roots []string, exeDir string) {
	allKeys := map[string]struct{}{}
	filenames := []string{"validation_ja.yaml", "validation_en.yaml"}
	for _, file := range filenames {
		result := i18n.TranslateLanguage{
			Keys: map[string]string{},
		}
		filename := filepath.Join(exeDir, file)
		if _, err := os.Stat(filename); err == nil {
			readFile(filename, &result)
		}
		for _, root := range roots {
			fset := token.NewFileSet()
			err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if strings.Contains(info.Name(), "_test.go") {
					return nil
				}
				if strings.HasSuffix(info.Name(), ".go") {
					values := parseFileValidation(fset, path)
					for key, value := range values {
						if _, ok := result.Keys[key]; !ok {
							result.Keys[key] = value
						}
						allKeys[key] = struct{}{}
					}
				}
				return nil
			})
			if err != nil {
				panic(err)
			}
		}
		keys := make([]string, 0, len(result.Keys))
		for key := range result.Keys {
			keys = append(keys, key)
		}
		for _, key := range keys {
			if _, ok := allKeys[key]; !ok {
				delete(result.Keys, key)
			}
		}
		writeFile(filename, result)
	}
}

func readFile(filename string, result *i18n.TranslateLanguage) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(result); err != nil {
		panic(err)
	}
}

func writeFile(path string, result i18n.TranslateLanguage) {
	f, err := os.Create(filepath.Join(path, ""))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := yaml.NewEncoder(f).Encode(&result); err != nil {
		panic(err)
	}
}

func parseFile(fset *token.FileSet, path string) []string {
	results := []string{}
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		if fun, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
			ident, ok := fun.X.(*ast.Ident)
			if !ok {
				return true
			}
			if strings.Contains(ident.Name, "i18n") && fun.Sel.Name == "Translate" {
				arg := callExpr.Args[1]
				value := exprToString(arg)
				results = append(results, value)
			}
		}
		return true
	})
	return results
}

func parseFileValidation(fset *token.FileSet, path string) map[string]string {
	results := map[string]string{}
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(node, func(n ast.Node) bool {
		code, value := nodeIsCallExpr(n)
		if code != "" {
			results[code] = value
		}
		return true
	})
	return results
}

func nodeIsCallExpr(n ast.Node) (code string, value string) {
	setReturnValue := func(node *ast.CallExpr) {
		codeArgs := node.Args[0]
		code = exprToString(codeArgs)
		messageArgs := node.Args[1]
		value = exprToString(messageArgs)
	}
	switch node := n.(type) {
	case *ast.CallExpr:
		if ident, ok := node.Fun.(*ast.Ident); ok {
			if ident.Name == "NewError" {
				setReturnValue(node)
			}
		}
		if fun, ok := node.Fun.(*ast.SelectorExpr); ok {
			ident, ok := fun.X.(*ast.Ident)
			if !ok {
				return
			}
			if strings.Contains(ident.Name, "validation") && fun.Sel.Name == "NewError" {
				setReturnValue(node)
			}
		}
	case *ast.GenDecl:
		if node.Tok == token.VAR {
			for _, spec := range node.Specs {
				valSpec, ok := spec.(*ast.ValueSpec)
				if ok {
					for _, v := range valSpec.Values {
						return nodeIsCallExpr(v)
					}
				}
			}
		}
	}
	return
}

func exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			return strings.Trim(e.Value, "\"")
		}
		return e.Value
	case *ast.Ident:
		return e.Name
	default:
		return ""
	}
}
