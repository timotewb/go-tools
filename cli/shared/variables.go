package shared

// List file names for the code to ignore
var IgnoreFiles = []string{".DS_Store", "._.DS_Store"}

// List of valid image extensions
var ImgExtComp = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", "heic"}
var ImgExtRaw = []string{".dng", ".cr2", ".cr3"}
var ImgExtAll = append(ImgExtComp, ImgExtRaw...)
