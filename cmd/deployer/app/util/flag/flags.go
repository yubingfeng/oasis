/* 
Copyright 2018 hill
*/


package flag

import  (

        "strings"
	"github.com/golang/glog"
	"github.com/spf13/pflag"
)


func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {

        if strings.Contains(name, "_") {
                return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
        }
	return pflag.NormalizedName(name)
}

func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {

        if strings.Contains(name, "_") {
                nname := strings.Replace(name, "_", "-", -1)
		glog.Warning("%s is deprecated and will be removed in next version use %s instead.", name,  nname)
		return pflag.NormalizedName(name)
	}
	return pflag.NormalizedName(name)

}
