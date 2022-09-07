package decompress

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func DecompressXML() (string, error) {

	resp, _ := http.Get("https://github.com/Julian-sUsername/storage-fun/blob/main/elecciones.gz?raw=true")
	contentFile, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	file, err := gzip.NewReader(bytes.NewReader(contentFile))

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

	return result, nil
}