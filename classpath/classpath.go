package classpath

import (
	"os"
	"path/filepath"
)

// Classpath include boot classpath, extend classpath, user classpath
type Classpath struct {
	bootClasspath Entry
	exClasspath   Entry
	userClasspath Entry
}

// ReadClass read the class file by class name
func (cp *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := cp.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := cp.exClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return cp.userClasspath.readClass(className)
}

func (cp *Classpath) String() string {
	return cp.userClasspath.String()

}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func (cp *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = newWildCardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.exClasspath = newWildCardEntry(jreExtPath)
}

func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClasspath = newEntry(cpOption)
}

// Parse parse the jre option and classath option
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
