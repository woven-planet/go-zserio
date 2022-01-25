package generator

import (
	"strings"
)

// StripImports searches a *.go file for the import section, and checks if the
// imports are used.
func StripImports(code []byte) ([]byte, error) {
	content := string(code)

	// sarch the import section
	importKey := "import ("
	importIndex := strings.Index(content, importKey)
	if importIndex < 0 {
		// this file either has only one import, or no imports at all. The zserio
		// code generator always generated a full import block, so this behavior
		// can be accepted.
		return code, nil
	}
	// extract the import section
	prefix := content[0 : importIndex+len(importKey)]
	content = content[importIndex+len(importKey):]
	importEndIndex := strings.Index(content, ")")
	importSection := content[0:importEndIndex]
	content = content[importEndIndex:]
	imports := strings.Split(importSection, "\n")

	// Extract the last name of the import line, search in the code if it is used,
	// and delete the import if it cannot be found.
	usedImports := []string{}
	for _, currentImport := range imports {
		line := strings.TrimSpace(currentImport)
		if len(line) < 3 {
			continue
		}
		importName := ""
		if strings.Contains(line, " ") {
			// this string uses an alias
			importName = strings.Split(line, " ")[0]
		} else {
			lastSeparator := strings.LastIndex(line, "/")
			if lastSeparator < 0 {
				// The import name does not contain any slashes, it must be a
				// internal module, such as "errors"
				importName = line[1 : len(line)-1]
			} else {
				importName = line[lastSeparator+1 : len(line)-1]
			}
		}
		if strings.Contains(content, importName+".") {
			usedImports = append(usedImports, currentImport)
		}
	}
	importSection = strings.Join(usedImports, "\n")
	return []byte(prefix + importSection + content), nil

}
