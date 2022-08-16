package decompress

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func DecompressXML() (string, error) {
	gzipFile, err := os.Open("./assets/elecciones.gz")

	if err != nil {
		fmt.Println(err)
	}

	file, err := gzip.NewReader(gzipFile)

	if err != nil {
		fmt.Println(err)
	}

	outFile, err := os.Create("decompresed")

	if err != nil {
		fmt.Println(err)
	}

	io.Copy(outFile, file)

	data, _ := os.Open("decompresed")
	content, _ := ioutil.ReadAll(data)

	result := string(content)
	
	defer file.Close()
	defer outFile.Close()
	defer os.Remove("decompresed")

	return result, nil
}