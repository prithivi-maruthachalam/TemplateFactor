package use

type UseTemplateConfig struct {
	TemplateName   string
	TargetDirPath  string
	NoFiles        bool
	NoFileContent  bool
	DirPermission  int
	FilePermission int
}
