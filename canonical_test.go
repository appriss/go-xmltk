package xmltk

import "testing"
import "io/ioutil"
import "bytes"
import "fmt"


//Tests currently call "xmllint" in a dummy function. Replace function with real implementation.

func testCompareResult(file string, sample []byte ) (bool, error) {
    target, err := ioutil.ReadFile(file)
    if err != nil {return false, err}
    result := bytes.Equal(sample, target)
    return result, nil
}



//Test 1: http://www.w3.org/TR/xml-c14n#Example-OutsideDoc (Section 3.1)
//Function must demonstrate the following:
//Loss of XML declaration.
//Loss of DTD.
//Normalization of whitespace outside of document.
//Loss of whitespace between PI Target and its data.
//Retention of whitespace inside PI data.

func TestCanonical10Step31 (t *testing.T) {
   indoc , err := ioutil.ReadFile("testdata/test-canonical1.0-input-3.1.xml")
   if err != nil { t.Error(err.Error()); return }
   // Here's the function, we may refactor
   outdoc, err := Canonicalize(indoc, XML_C14N_10, true)
   if err != nil { t.Error(err.Error()); return }
   result, err := testCompareResult("testdata/test-canonical1.0-output-3.1.xml", outdoc)
   if err != nil { t.Error(err.Error()); return }
   if result != true { fmt.Printf("%s",outdoc)
      t.Error("Output does not match expected result, file test-canonical1.0-input-3.1.xml") }
}