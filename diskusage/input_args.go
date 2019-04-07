package diskusage

import (
	"errors"
	"fmt"
	"time"
)

//LimitDefault - default Limit value
const LimitDefault = 10 //default value for a arglimit
//DepthDefault - default Depth value
const DepthDefault = 1 //default depth in results
//SortDefault - default Sort value
const SortDefault = "size_desc"

//CsvDefault - default value "nocsv" without generation csv file at end
const CsvDefault = "nocsv"

//InputArgs - input arguments
var InputArgs TInputArgs

//TInputArgs - the programm arguments
type TInputArgs struct {
	Path        string //analysed path
	Limit       int    //limit folders in results
	Depth       int    //depth of subfolders in results (-1 - all, 1 - only current, 2 and more - 2 and more)
	FixUnit     string //fixed size unit in a results presetation (b, Kb, Mb, ...). Has a upper priority than "argmaxunit". Must be in sizeUnits.
	Sort        string //results sorting by
	CsvFileName string //csv file name to export results
}

//SetPath - init Path field
func (t *TInputArgs) SetPath(inpath *string) error {
	newpath := CleanPath(inpath, true)
	if len(newpath) == 0 {
		return errors.New("Error! Argument 'path' could not be an empty string")
	}
	t.Path = newpath
	return nil
}

//SetLimit - init Limit field
func (t *TInputArgs) SetLimit(limit *int) error {
	if *limit < 0 {
		fmt.Printf("Argument 'limit' is negative (%d) and has been set to default value (%d)", *limit, LimitDefault)
		*limit = LimitDefault //set to default value
	}
	t.Limit = *limit
	return nil
}

//SetFixUnit - init FixUnit field
func (t *TInputArgs) SetFixUnit(fixunit *string) error {
	if _, ok := sizeUnits[*fixunit]; !ok && len(*fixunit) > 0 {
		return errors.New("Error! Argument 'fixunit' is not in allowable range {b, Kb, Mb, Gb, Tb, Pb}")
	}
	if len(*fixunit) > 0 {
		fmt.Printf("Results will be represented with fixed units style in '%s'\n", *fixunit)
	}
	t.FixUnit = *fixunit
	return nil
}

//SetDepth - init Depth field
func (t *TInputArgs) SetDepth(depth *int) error {
	if *depth < 0 {
		fmt.Printf("Argument 'depth' is negative (%d) and has been set to default value (%d)", *depth, DepthDefault)
	}
	t.Depth = *depth
	return nil
}

//SetSort - init Sort field
func (t *TInputArgs) SetSort(sort *string) error {
	if _, ok := sortValues[*sort]; !ok && len(*sort) > 0 {
		t.Sort = SortDefault
		return errors.New("Error! Argument 'sort' is not in allowable range {size_desc, name_asc} and replaced to :" + SortDefault)
	}

	t.Sort = *sort
	return nil
}

//SetCsvFileName - init SetCsvFileName field
func (t *TInputArgs) SetCsvFileName(csvfilename *string) error {
	if len(*csvfilename) == 0 {
		t.CsvFileName = getDefaultCsvFileName()
		fmt.Printf("Csv file for export to: '%s'\n", t.CsvFileName)
		return nil
	}
	t.CsvFileName = *csvfilename
	return nil
}

func getDefaultCsvFileName() string {
	tnow := time.Now()
	tnowstr := tnow.Format("20060102_150405")
	return fmt.Sprintf("./results/result_%s.csv", tnowstr)
}

//PrintArguments - print arguments
func (t TInputArgs) PrintArguments() {
	fmt.Println("\nArguments:")
	fmt.Printf("   path: %s\n", t.Path)
	fmt.Printf("   limit: %d\n", t.Limit)
	fmt.Printf("   fixunit: %s\n", t.FixUnit)
	fmt.Printf("   depth: %d\n", t.Depth)
	fmt.Printf("   sort: %s\n", t.Sort)
	fmt.Printf("   csv: %s\n", t.CsvFileName)
}
