package xmltk

import "testing"
import "io/ioutil"
import "bytes"
import "fmt"


//Tests currently call "xmllint" in a dummy function. Replace function with real implementation.

func testXMLCompare(infile string, goldfile string) error {
   indoc , err := ioutil.ReadFile(infile)
   if err != nil { return err }
   golddoc, err := ioutil.ReadFile(goldfile)
   if err != nil { return err }
   // Here's the function, we may refactor
   outdoc, err := Canonicalize(indoc, XML_C14N_10, true)
   if err != nil { return err }
   exact := bytes.Equal(outdoc, golddoc)
   if exact != true {
   	err := fmt.Errorf("Output doc %s does not match success file %q", outdoc, goldfile)
   	return err
   }
   return nil
}



//Test 1: http://www.w3.org/TR/xml-c14n#Example-OutsideDoc (Section 3.1)
//Function must demonstrate the following:
//Loss of XML declaration.
//Loss of DTD.
//Normalization of whitespace outside of document.
//Loss of whitespace between PI Target and its data.
//Retention of whitespace inside PI data.
func TestCanonical10Step31 (t *testing.T) {
   err := testXMLCompare("testdata/test-canonical1.0-input-3.1.xml","testdata/test-canonical1.0-output-3.1.xml")
   if err != nil {
     t.Error(err.Error())
     return
   }
}

//Test 2: http://www.w3.org/TR/xml-c14n#Example-WhitespaceInContent
//Function must demonstrate the following:
//Retain all whitespace between consecutive start tags, clean or dirty
//Retain all whitespace between consecutive end tags, clean or dirty
//Retain all whitespace between end tag/start tag pair, clean or dirty
//Retain all whitespace in character content, clean or dirty
func TestCanonical10Step32 (t *testing.T) {

}