package global

//文件信息格式
type FileInfo struct {
	// ID          int
	Name        string
	Size        int64
	CreateTime  string
	DownLandURL string
	Md5         string
	Branch      string
	Text        string
}

// func (f FileInfo) GetID() int {
// 	return f.ID
// }

func (f FileInfo) GetInfo() FileInfo {
	return f
}

func (f FileInfo) GetName() string {
	return f.Name
}

func (f FileInfo) GetSize() int64 {
	return f.Size
}

func (f FileInfo) GetCreateTime() string {
	return f.CreateTime
}

func (f FileInfo) GetDownLandURL() string {
	return f.DownLandURL
}

func (f FileInfo) GetBranch() string {
	return f.Branch
}

func (f FileInfo) GetText() string {
	return f.Text
}

// func (f FileInfo) GetDownLandNum() int {
// 	return f.DownLandNum
// }

func (f FileInfo) GetMD5() string {
	return f.Md5
}
