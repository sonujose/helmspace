package utilities

import(
	"archive/tar"
    "compress/gzip"
    "strings"
    "bytes"
    "log"
    "io"
)


// GetFileBlobFromTarBall - Get File Blob data from repository tarball gzip
func GetFileBlobFromTarBall(gzipStream io.Reader, fileName string)(string, error) {
    var blobData string
    
    uncompressedStream, err := gzip.NewReader(gzipStream)
    
    if err != nil {
        log.Fatal("ExtractTarGz: NewReader failed")
        return blobData, nil
    }

    tarReader := tar.NewReader(uncompressedStream)

    for true {
        header, err := tarReader.Next()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
            return blobData, err
        }

        if strings.EqualFold(header.Name, fileName) {
            var b bytes.Buffer
            io.Copy(&b, tarReader)
            blobData = string(b.Bytes())
            break
        }

	}
	
	return blobData, nil
}
