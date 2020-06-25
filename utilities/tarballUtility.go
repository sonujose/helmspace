package utilities

import(
	"archive/tar"
    "compress/gzip"
    "net/http"
    "strings"
    "bytes"
    "log"
    "io"
)

// GetFileBlobFromTarBall - Get File Blob data from repository tarball gzip
func GetFileBlobFromTarBall(tarURL string, fileName string) (*string, error) {

    blobData := new(string)

    res, err := http.Get(tarURL)

	if err != nil {
        log.Fatalf("API call failed %s -: %s", tarURL, err.Error())
		return blobData, err
    }
    
	defer res.Body.Close()

    gzipStream, err := gzip.NewReader(res.Body)
    
    log.Printf("Response from server data %v", gzipStream)

	if err != nil {
		return blobData, err
    }
    
	defer gzipStream.Close()

    tarReader := tar.NewReader(gzipStream)

    for true {
        header, err := tarReader.Next()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
            return blobData, err
        }

        log.Printf("Validating file - %s", header.Name)

        if strings.EqualFold(header.Name, fileName) {
            log.Printf("Filename found - %s", fileName)
            var b bytes.Buffer
            io.Copy(&b, tarReader)
            *blobData = string(b.Bytes())
            break
        }

    }
    
    log.Printf("BlobData = %s", *blobData)
    
	return blobData, nil
}


