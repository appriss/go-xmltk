package xmltk


import "os/exec"
import "io/ioutil"

const (
	XML_C14N_10 = iota
)

// The Canonicalize function takes an XML document and ensures that it is represented 
//in the Canonical XML form.
//Parameters: 
// doc: the document to be canonicalized.
// algo: the algorithm to be used for canonicalization. Valid values:
//   XML_C14N_10 - Canonical XML 1.0 (http://www.w3.org/TR/xml-c14n)
//   XML_C14N_11 - Canonical XML 1.1 (http://www.w3.org/TR/xml-c14n11/)
//   XML_C14N_20 - Canonical XML 2.0 (http://www.w3.org/TR/xml-c14n2/)
//  comments: Shall XML comments be excluded?
// Returns a new, Canonical XML document. If err if not nil, the document failed parsing and so
// is likely a malformed XML document.

func Canonicalize(doc []byte, algo int , comments bool) ([]byte, error) {
	// Stub to call xmllint to generate output
	tempfile, err  := ioutil.TempFile("","")
	if err != nil { return nil, err }
	_, err = tempfile.Write(doc)
	if err != nil { return nil, err }
	err = tempfile.Close()
    cmd := exec.Command("xmllint","-c14n", tempfile.Name())
    output, err := cmd.Output()
    return output, nil
}

