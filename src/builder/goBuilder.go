package builder

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

type Builder struct {
	moduleName     string
	ports          *bytes.Buffer
	inputs         []Port
	constructor    *bytes.Buffer
	subConstructor *bytes.Buffer
	observer       *bytes.Buffer
	assigns        *bytes.Buffer
	instances      *bytes.Buffer
	runMethod      *bytes.Buffer
	preAlways      *bytes.Buffer
	always         *bytes.Buffer
	function       *bytes.Buffer
	source         *bytes.Buffer
}

// GenerateSource brings together the sources in a module
func (b *Builder) generateSource() {
	b.source.WriteString("package generated\n")
	b.source.WriteString("import \"github.com/verilog2Go/src/variable\"\n")
	b.source.WriteString(b.ports.String())
	b.source.WriteString(b.constructor.String())
	b.source.WriteString(b.subConstructor.String())
	b.source.WriteString("func (" + moduleName + " *" + strings.Title(moduleName) + ") Exec() {\n" + b.assigns.String() + b.instances.String() + "}\n\n")
	b.source.WriteString(b.runMethod.String())
	b.source.WriteString(b.preAlways.String())
	b.source.WriteString(b.always.String())
	b.source.WriteString(b.function.String())
}

// WriteFile writes source to a file
func (builder *Builder) WriteFile(filename string) {
	os.Mkdir("tdd", 0777)
	file, err := os.Create("tdd/" + filename + ".go")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	builder.generateSource()
	file.Write(builder.source.Bytes())
}

var moduleName string

// StartModule initializes the module
func (b *Builder) StartModule(modName string) {
	moduleName = modName
}

//DeclarePorts declares ports
func (b *Builder) DeclarePorts(ports []Port, variables []Variable) {
	if len(ports) < 1 {
		return
	}
	b.ports.WriteString("type " + strings.Title(moduleName) + " struct{\n")
	dimensions := ""
	var tmp []string
	for i := 0; i < len(ports); i++ {
		if i < len(ports)-1 {
			if !ports[i].isDimension {
				// b.ports.WriteString(ports[i].id + ", ")
				tmp = append(tmp, ports[i].id)
			} else {
				dimensions += ports[i].id
			}
		} else {
			if !ports[i].isDimension {
				// b.ports.WriteString(ports[i].id)
				tmp = append(tmp, ports[i].id)
			} else {
				dimensions += ports[i].id
			}
		}
	}
	// b.ports = b.ports[:len(b.ports)-2]
	b.ports.WriteString(strings.Join(tmp, ","))
	b.ports.WriteString(" *variable.BitArray\n")
	if dimensions != "" {
		dimensions = dimensions[:len(dimensions)]
		b.ports.WriteString(dimensions + " []*variable.BitArray\n")
	}

	for i := 0; i < len(variables); i++ {
		if i < len(variables)-1 {
			b.ports.WriteString(variables[i].id + ", ")
		} else {
			b.ports.WriteString(variables[i].id)
			b.ports.WriteString(" int\n")
		}
	}

	b.ports.WriteString("}\n\n")
}

// DeclareInput creates an array of input signals
func (b *Builder) DeclareInput(input Port) {
	b.inputs = append(b.inputs, input)
}

// CreateConstructor creates a constructor
func (b *Builder) CreateConstructor(funcName string, ports []Port, params []Param, variables []Variable) {
	if len(ports) < 1 {
		return
	}
	//First line of constructor
	b.constructor.WriteString("func New" + strings.Title(funcName) + "() " + strings.Title(moduleName) + "{\n")
	b.subConstructor.WriteString("func NewGoroutine" + strings.Title(funcName) + " (in []chan int, out []chan int) *" + strings.Title(moduleName) + "{\n")
	var tmp []string
	var contmp string
	var subcontmp string
	b.constructor.WriteString("args := &" + strings.Title(moduleName) + "{")
	b.subConstructor.WriteString(moduleName + " := &" + strings.Title(moduleName) + "{")
	for i := 0; i < len(ports); i++ {
		if !ports[i].isDimension {
			tmp = append(tmp, "variable.NewBitArray("+strconv.Itoa(ports[i].length)+")")
		}
	}
	for i := 0; i < len(ports); i++ {
		if ports[i].isDimension {
			tmp = append(tmp, "make([]*variable.BitArray, "+strconv.Itoa(ports[i].dimLength)+")")
			for j := 0; j < ports[i].dimLength; j++ {
				contmp += "args." + ports[i].id + "[" + strconv.Itoa(j) + "] = variable.NewBitArray(" + strconv.Itoa(ports[i].length) + ")\n"
				subcontmp += moduleName + "." + ports[i].id + "[" + strconv.Itoa(j) + "] = variable.NewBitArray(" + strconv.Itoa(ports[i].length) + ")\n"
			}
		}
	}
	for i := 0; i < len(variables); i++ {
		tmp = append(tmp, "0")
	}
	b.constructor.WriteString(strings.Join(tmp, ",") + "}\n")
	b.subConstructor.WriteString(strings.Join(tmp, ",") + "}\n")
	b.constructor.WriteString(contmp)
	b.subConstructor.WriteString(subcontmp)
	b.constructor.WriteString(b.observer.String())

	if len(params) > 0 {
		for _, v := range params {
			b.constructor.WriteString("args." + v.id + " = " + v.initiation + "\n")
		}
	}
	if hasInitial {
		b.constructor.WriteString("args.initial()\n")
		b.subConstructor.WriteString(moduleName + ".initial()\n")
	}
	b.subConstructor.WriteString("go " + moduleName + ".start(in, out)\n")
	b.constructor.WriteString("return *args\n}\n\n")
	b.subConstructor.WriteString("return " + moduleName + "\n}\n\n")
}

func (b *Builder) CreateRunMethod(ports []Port) {
	var inCounter, outCounter int
	b.runMethod.WriteString("func (" + moduleName + " *" + strings.Title(moduleName) + ") start(in []chan int, out []chan int) {\n")
	for _, v := range ports {
		if v.portType == "output" {
			b.runMethod.WriteString("defer close(out[" + strconv.Itoa(outCounter) + "])\n")
			outCounter++
		}
	}
	b.runMethod.WriteString("for {\n" + "select {\n")
	for _, v := range ports {
		if v.portType == "input" {
			b.runMethod.WriteString("case v, ok := <-in[" + strconv.Itoa(inCounter) + "]:\n")
			b.runMethod.WriteString("if ok {\n")
			b.runMethod.WriteString(moduleName + "." + v.id + ".Set(v)\n")
			// if AlwaysCounter != 0 {
			for i := 1; i <= AlwaysCounter; i++ {
				b.runMethod.WriteString("bitArrays" + strconv.Itoa(i) + " := " + moduleName + ".PreAlways" + strconv.Itoa(i) + "()\n")
			}
			for i := 1; i <= AlwaysCounter; i++ {
				b.runMethod.WriteString(moduleName + ".Always" + strconv.Itoa(i) + "(bitArrays" + strconv.Itoa(i) + ")\n")
			}
			// }
			b.runMethod.WriteString(moduleName + ".Exec()\n")
			var count int
			for _, v := range ports {
				if v.portType == "output" {
					b.runMethod.WriteString("out[" + strconv.Itoa(count) + "] <- " + moduleName + "." + v.id + ".ToInt()\n")
					count++
				}
			}
			b.runMethod.WriteString("} else {\n" + "return \n" + "}\n")
			inCounter++
		}
	}
	b.runMethod.WriteString("}\n}\n")
	b.runMethod.WriteString("}\n\n")
}

func (b *Builder) CreateAssign(id string, expression string) {
	b.assigns.WriteString(moduleName + "." + id + ".Assign(" + expression + ")\n")
}

// CreateInstance creates an instantiation
func (b *Builder) CreateInstance(instance Instance) {
	b.instances.WriteString(instance.instanceName + " := " + strings.Title(instance.moduleName) + "(&" + instance.moduleName + "{")
	var tmp string
	for _, exp := range instance.ports {
		tmp += exp + ", "
	}

	for key, exp := range instance.portMap {
		tmp += key + " : " + exp + ", "
	}
	tmp = tmp[:len(tmp)-2]
	b.instances.WriteString(tmp + "})\n" + instance.instanceName + ".Exec()\n")
}
