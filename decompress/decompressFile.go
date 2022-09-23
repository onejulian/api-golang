package decompress

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DecompressXML() (string, error) {

	resp, err := http.Get("https://github.com/Julian-sUsername/storage-fun/blob/main/elecciones.gz?raw=true")
	if err != nil {
		fmt.Println(err)
	}
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
	content, _ := io.ReadAll(data)

	result := string(content)
	
	defer file.Close()
	defer outFile.Close()

	return result, nil
}