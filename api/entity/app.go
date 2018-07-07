package entity

type App struct {
	Name             string            `json:"name"`
	PackageName      string            `json:"package_name"`
	Labels           map[string]string `json:"labels"`
	OsName           string            `json:"os_name"`
	Path             string            `json:"path"`
	Revision         uint64            `json:"revision"`
	VersionCode      string            `json:"version_code"`
	VersionName      string            `json:"version_name"`
	SdkVersion       uint16            `json:"sdk_version"`
	RawSdkVersion    string            `json:"raw_sdk_version"`
	TargetSdkVersion uint16            `json:"target_sdk_version"`
	Signature        string            `json:"signature"`
	ShortMessage     string            `json:"message"`
	AppFileUrl       string            `json:"file"`
	MD5              string            `json:"md5"`
	AppIconUrl       string            `json:"icon"`
	AppOwner         LiteUser          `json:"user"`
}
