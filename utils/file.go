// utils/file.go
package utils

import (
	//"archive/tar"
	//"archive/zip"
	//"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// get filepath base name
func Basename(file string) string {
	return path.Base(file)
}

// get filepath dir name
func Dirname(file string) string {
	return path.Dir(file)
}

// get absolute filepath, based on built executable file
func RealPath(file string) (string, error) {
	if path.IsAbs(file) {
		return file, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, file), err
}

// get file modified time
func FileMTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// Copy copies file from source to target path.
// It returns false and error when error occurs in underlying functions.
func Copy(file string, to string) (bool, error) {
	sf, e := os.Open(file)
	if e != nil {
		return false, e
	}
	defer sf.Close()
	df, e2 := os.Create(to)
	if e2 != nil {
		return false, e2
	}
	defer df.Close()
	// buffer reader, do chunk transfer
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := sf.Read(buf)
		if err != nil && err != io.EOF {
			return false, err
		}
		if n == 0 {
			break
		}
		// write a chunk
		if _, err := df.Write(buf[:n]); err != nil {
			return false, err
		}
	}
	return true, nil
}

// move file to new path
func Move(file string, to string) (bool, error) {
	// copy
	r, e := Copy(file, to)
	if e != nil {
		return r, e
	}
	// then remove
	e = os.Remove(file)
	if e != nil {
		r = false
	}
	return r, e
}

// delete file
func Unlink(file string) error {
	return os.Remove(file)
}

// rename file name
func Rename(file string, to string) error {
	return os.Rename(file, to)
}

// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// SaveFile saves content type '[]byte' to file by given path.
// It returns error when fail to finish operation.
func SaveFile(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

// SaveFileS saves content type 'string' to file by given path.
// It returns error when fail to finish operation.
func SaveFileS(filePath string, s string) (int, error) {
	return SaveFile(filePath, []byte(s))
}

// ReadFile reads data type '[]byte' from file by given path.
// It returns error when fail to finish operation.
func ReadFile(filePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}

// ReadFileS reads data type 'string' from file by given path.
// It returns error when fail to finish operation.
func ReadFileS(filePath string) (string, error) {
	b, err := ReadFile(filePath)
	return string(b), err
}

//// Unzip unzips .zip file to 'destPath' and returns sub-directories.
//// It returns error when fail to finish operation.
//func Unzip(srcPath, destPath string) ([]string, error) {
//	// Open a zip archive for reading
//	r, err := zip.OpenReader(srcPath)
//	if err != nil {
//		return nil, err
//	}
//	defer r.Close()

//	dirs := make([]string, 0, 5)
//	// Iterate through the files in the archive
//	for _, f := range r.File {
//		// Get files from archive
//		rc, err := f.Open()
//		if err != nil {
//			return nil, err
//		}

//		dir := path.Dir(f.Name)
//		// Create directory before create file
//		os.MkdirAll(destPath+"/"+dir, os.ModePerm)
//		dirs = AppendStr(dirs, dir)

//		if f.FileInfo().IsDir() {
//			continue
//		}

//		// Write data to file
//		fw, _ := os.Create(path.Join(destPath, f.Name))
//		if err != nil {
//			return nil, err
//		}
//		_, err = io.Copy(fw, rc)

//		if fw != nil {
//			fw.Close()
//		}
//		if err != nil {
//			return nil, err
//		}
//	}
//	return dirs, nil
//}

//// UnTarGz ungzips and untars .tar.gz file to 'destPath' and returns sub-directories.
//// It returns error when fail to finish operation.
//func UnTarGz(srcFilePath string, destDirPath string) ([]string, error) {
//	// Create destination directory
//	os.Mkdir(destDirPath, os.ModePerm)

//	fr, err := os.Open(srcFilePath)
//	if err != nil {
//		return nil, err
//	}
//	defer fr.Close()

//	// Gzip reader
//	gr, err := gzip.NewReader(fr)
//	if err != nil {
//		return nil, err
//	}
//	defer gr.Close()

//	// Tar reader
//	tr := tar.NewReader(gr)

//	dirs := make([]string, 0, 5)
//	for {
//		hdr, err := tr.Next()
//		if err == io.EOF {
//			// End of tar archive
//			break
//		}

//		// Check if it is directory or file
//		if hdr.Typeflag != tar.TypeDir {
//			// Get files from archive
//			// Create directory before create file
//			dir := path.Dir(hdr.Name)
//			os.MkdirAll(destDirPath+"/"+dir, os.ModePerm)
//			dirs = AppendStr(dirs, dir)

//			// Write data to file
//			fw, _ := os.Create(destDirPath + "/" + hdr.Name)
//			if err != nil {
//				return nil, err
//			}
//			_, err = io.Copy(fw, tr)
//			if err != nil {
//				return nil, err
//			}
//		}
//	}
//	return dirs, nil
//}
