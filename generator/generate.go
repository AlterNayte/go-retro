package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const baseTemplate = `// Code generated by GoRetro; DO NOT EDIT.
package goretro

type AuthType string

const (
    AuthNone AuthType = "none"
    AuthBasic AuthType = "basic"
    AuthAPIKey AuthType = "api_key"
    AuthBearer AuthType = "bearer"
)
`

const clientTemplate = `// Code generated by GoRetro; DO NOT EDIT.
package goretro

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"{{ if .HasQueryParams }}
	"net/url"{{ end }}
    "time"
	"{{ .ModulePath }}"
	{{ if .IncludeBytes }}"bytes"{{end}}
)

type {{ .ClientType }} struct {
    BaseURL       string
    HTTPClient    *http.Client
    AuthType      AuthType
    Username      string
    Password      string
    APIKey        string
    BearerToken   string
    CustomAuthFunc func() (string, error)
    CustomHeaders  map[string]string
    Timeout        time.Duration
    MaxRetries     int
}

func New{{ .ClientType }}(baseURL string) *{{ .ClientType }} {
    return &{{ .ClientType }}{
        BaseURL:      baseURL,
        HTTPClient:   &http.Client{},
        AuthType:     AuthNone,
        CustomHeaders: make(map[string]string),
        Timeout:      30 * time.Second,
        MaxRetries:   3,
    }
}

func (c *{{ .ClientType }}) SetBasicAuth(username, password string) {
    c.AuthType = AuthBasic
    c.Username = username
    c.Password = password
}

func (c *{{ .ClientType }}) SetAPIKeyAuth(apiKey string) {
    c.AuthType = AuthAPIKey
    c.APIKey = apiKey
}

func (c *{{ .ClientType }}) SetBearerAuth(token string) {
    c.AuthType = AuthBearer
    c.BearerToken = token
}

func (c *{{ .ClientType }}) SetCustomAuthFunc(authFunc func() (string, error)) {
    c.CustomAuthFunc = authFunc
}

func (c *{{ .ClientType }}) SetCustomHeader(key, value string) {
    c.CustomHeaders[key] = value
}

func (c *{{ .ClientType }}) doRequest(req *http.Request) (*http.Response, error) {
    for k, v := range c.CustomHeaders {
        req.Header.Set(k, v)
    }
    req.Header.Set("Content-Type", "application/json")

    for i := 0; i < c.MaxRetries; i++ {
        c.HTTPClient.Timeout = c.Timeout
        resp, err := c.HTTPClient.Do(req)
        if err != nil {
            if i == c.MaxRetries-1 {
                return nil, err
            }
            time.Sleep(2 * time.Second)
            continue
        }

        if resp.StatusCode >= 200 && resp.StatusCode < 300 {
            return resp, nil
        }

        if i == c.MaxRetries-1 {
            bodyBytes, _ := ioutil.ReadAll(resp.Body)
            return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
        }

        time.Sleep(2 * time.Second)
    }

    return nil, errors.New("max retries exceeded")
}

{{ range .Methods }}
func (c *{{ $.ClientType }}) {{ .Name }}({{ .Params }}) ({{ .Returns }}) {
    reqURL := c.BaseURL + "{{ .Path }}"
    {{ if .PathParams }}
    reqURL = fmt.Sprintf(reqURL, {{ .PathParams }})
    {{ end }}
    {{ if .QueryParams }}
    query := url.Values{}
    {{ range .QueryParams }}
    query.Add("{{ .Key }}", fmt.Sprintf("%v", {{ .Value }}))
    {{ end }}
    reqURL += "?" + query.Encode()
    {{ end }}

    req, err := http.NewRequest("{{ .Method }}", reqURL, nil)
    if err != nil {
        return nil, err
    }

    {{ if eq .Method "POST" }}
    bodyBytes, err := json.Marshal({{ .Body }})
    if err != nil {
        return nil, err
    }
    req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
    req.Header.Set("Content-Type", "application/json")
    {{ end }}

    {{ if eq .Auth "basic" }}
    req.SetBasicAuth(c.Username, c.Password)
    {{ else if eq .Auth "apiKey" }}
    req.Header.Set("Authorization", "Api-Key " + c.APIKey)
    {{ else if eq .Auth "bearer" }}
    if c.CustomAuthFunc != nil {
        token, err := c.CustomAuthFunc()
        if err != nil {
            return nil, err
        }
        req.Header.Set("Authorization", "Bearer " + token)
    } else {
        req.Header.Set("Authorization", "Bearer " + c.BearerToken)
    }
    {{ end }}

    log.Printf("Making request to %s", req.URL)

    resp, err := c.doRequest(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result {{ .ReturnType }}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return result, nil
}
{{ end }}
`

type Method struct {
	Name        string
	Params      string
	Returns     string
	Method      string
	Path        string
	PathParams  string
	Body        string
	ReturnType  string
	Auth        string
	QueryParams []QueryParam
}

type QueryParam struct {
	Key   string
	Value string
}

type Interface struct {
	Name           string
	PackageName    string
	ModulePath     string
	ClientType     string
	IncludeBytes   bool
	HasQueryParams bool
	Methods        []Method
}

func Generate(outputDir string, inputDir string) error {
	var files []string
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	if len(files) > 0 {
		// Generate the base file
		packageName := "goretro" // Replace with actual package name
		err := generateBaseFile(outputDir, packageName)
		if err != nil {
			return err
		}
	}

	for _, inputFile := range files {
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, inputFile, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		for _, decl := range node.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				for _, spec := range genDecl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							ifaceName := typeSpec.Name.Name
							clientType := ifaceName + "Client"
							var methods []Method
							includeBytes := false
							hasQueryParams := false

							for _, field := range structType.Fields.List {
								if len(field.Names) == 0 {
									continue
								}
								methodName := field.Names[0].Name
								if field.Tag == nil || field.Tag.Value == "" {
									continue
								}
								methodTag := field.Tag.Value
								methodTag = strings.Trim(methodTag, "`")
								tags := strings.Split(methodTag, " ")
								if len(tags) < 2 {
									continue
								}
								method := strings.TrimPrefix(tags[0], `method:"`)
								method = strings.TrimSuffix(method, `"`)
								path := strings.TrimPrefix(tags[1], `path:"`)
								path = strings.TrimSuffix(path, `"`)
								if method == "" || path == "" {
									continue
								}
								auth := ""
								query := ""

								if len(tags) > 2 {
									if strings.HasPrefix(tags[2], `auth:"`) {
										auth = strings.TrimPrefix(tags[2], `auth:"`)
										auth = strings.TrimSuffix(auth, `"`)
									} else if strings.HasPrefix(tags[2], `query:"`) {
										query = strings.TrimPrefix(tags[2], `query:"`)
										query = strings.TrimSuffix(query, `"`)
									}
								}

								// Check the fourth tag in case the third was `query` and the fourth might be `auth`, or vice versa
								if len(tags) > 3 {
									if strings.HasPrefix(tags[3], `auth:"`) && auth == "" { // Only set auth if it hasn't been set yet
										auth = strings.TrimPrefix(tags[3], `auth:"`)
										auth = strings.TrimSuffix(auth, `"`)
									} else if strings.HasPrefix(tags[3], `query:"`) && query == "" { // Only set query if it hasn't been set yet
										query = strings.TrimPrefix(tags[3], `query:"`)
										query = strings.TrimSuffix(query, `"`)
									}
								}

								params := ""
								body := ""
								pathParams := ""

								queryParams := []QueryParam{}
								if funcType, ok := field.Type.(*ast.FuncType); ok {
									var pathParamReplacements []string
									var paramList []string

									for _, param := range funcType.Params.List {
										//paramName := param.Names[0].Name
										paramType := formatType(param.Type, node.Name.Name)

										// Handle multiple parameters sharing the same type
										for _, paramNameIdent := range param.Names {
											paramName := paramNameIdent.Name

											// Add each parameter to the list of function parameters
											paramList = append(paramList, fmt.Sprintf("%s %s", paramName, paramType))

											// Check if the path contains the parameter and replace it with the format placeholder
											if strings.Contains(path, "{"+paramName+"}") {
												pathParamReplacements = append(pathParamReplacements, paramName)
												path = strings.Replace(path, "{"+paramName+"}", "%s", -1)
											} else {
												// If the parameter is not in the path, it's considered a query parameter
												if strings.Contains(query, paramName) {
													queryParams = append(queryParams, QueryParam{
														Key:   paramName,
														Value: paramName,
													})
													hasQueryParams = true

												} else if method == "POST" && body == "" {
													includeBytes = true
													body = paramName
												}
											}
										}
									}

									params = strings.Join(paramList, ", ")

									// If pathParamReplacements are found, format the path string
									if len(pathParamReplacements) > 0 {
										pathParams = strings.Join(pathParamReplacements, ", ")
									}

									if method == "POST" {
										includeBytes = true
										body = funcType.Params.List[0].Names[0].Name
									}

									returnType := ""
									if funcType.Results.NumFields() > 0 {
										retType := funcType.Results.List[0].Type
										if _, isIdent := retType.(*ast.Ident); isIdent {
											returnType = formatType(retType, node.Name.Name)
										} else {
											returnType = fmt.Sprintf("%s", formatType(retType, node.Name.Name))
										}
									}

									methods = append(methods, Method{
										Name:        methodName,
										Params:      params,
										Returns:     fmt.Sprintf("%s, error", returnType),
										Method:      method,
										Path:        path,
										PathParams:  pathParams,
										Body:        body,
										ReturnType:  returnType,
										Auth:        auth,
										QueryParams: queryParams,
									})
								}
							}
							if len(methods) == 0 {
								continue // Skip if there are no valid methods
							}
							modulePath := deriveModulePath(inputFile)

							iface := Interface{
								Name:           ifaceName,
								PackageName:    node.Name.Name,
								ModulePath:     modulePath,
								ClientType:     clientType,
								IncludeBytes:   includeBytes,
								HasQueryParams: hasQueryParams,
								Methods:        methods,
							}

							tmpl, err := template.New("client").Parse(clientTemplate)
							if err != nil {
								return err
							}

							var buf bytes.Buffer
							err = tmpl.Execute(&buf, iface)
							if err != nil {
								return err
							}

							//finalOutputDir := fmt.Sprintf("%s/%s", outputDir, strings.ToLower(iface.PackageName))
							outputFile := fmt.Sprintf("%s/%s_client_gen.go", outputDir, strings.ToLower(ifaceName))
							os.MkdirAll(outputDir, os.ModePerm)
							err = os.WriteFile(outputFile, buf.Bytes(), 0644)
							if err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}

	return nil
}

func formatType(expr ast.Expr, packageName string) string {
	switch t := expr.(type) {
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", formatType(t.X, packageName))
	case *ast.Ident:
		if isPrimitive(t.Name) {
			return t.Name
		}
		return fmt.Sprintf("%s.%s", packageName, t.Name)
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", formatType(t.Elt, packageName))
	default:
		return ""
	}
}

// isPrimitive checks if a type is a Go primitive type.
func isPrimitive(typeName string) bool {
	primitives := map[string]bool{
		"bool":       true,
		"byte":       true,
		"complex64":  true,
		"complex128": true,
		"error":      true,
		"float32":    true,
		"float64":    true,
		"int":        true,
		"int8":       true,
		"int16":      true,
		"int32":      true,
		"int64":      true,
		"rune":       true,
		"string":     true,
		"uint":       true,
		"uint8":      true,
		"uint16":     true,
		"uint32":     true,
		"uint64":     true,
		"uintptr":    true,
	}
	return primitives[typeName]
}

// deriveModulePath derives the module path from the input file path.
func deriveModulePath(inputFile string) string {
	absPath, err := filepath.Abs(inputFile)
	if err != nil {
		return ""
	}

	// Find the directory containing the go.mod file.
	dir := absPath
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			break
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return ""
		}
		dir = parentDir
	}

	// Compute the module path.
	modFile, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		return ""
	}

	lines := strings.Split(string(modFile), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			modulePath := strings.TrimSpace(strings.TrimPrefix(line, "module "))
			relativePath := strings.TrimPrefix(absPath, dir)
			return filepath.ToSlash(filepath.Join(modulePath, filepath.Dir(relativePath)))
		}
	}

	return ""
}

func generateBaseFile(outputDir, packageName string) error {
	baseFilePath := filepath.Join(outputDir, "base_gen.go")
	os.MkdirAll(outputDir, os.ModePerm)

	baseContent := struct {
		PackageName string
	}{
		PackageName: packageName,
	}

	tmpl, err := template.New("base").Parse(baseTemplate)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, baseContent)
	if err != nil {
		return err
	}

	err = os.WriteFile(baseFilePath, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
