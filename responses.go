package yadisk

import (
	"fmt"
	"time"
)

type Disk struct {
	PaidMaxFileSize            int            `json:"paid_max_file_size,omitempty" yaml:"paid_max_file_size,omitempty"`
	MaxFileSize                int            `json:"max_file_size,omitempty" yaml:"max_file_size,omitempty"`
	TotalSpace                 int            `json:"total_space,omitempty" yaml:"total_space,omitempty"`
	TrashSize                  int            `json:"trash_size,omitempty" yaml:"trash_size,omitempty"`
	UsedSpace                  int            `json:"used_space,omitempty" yaml:"used_space,omitempty"`
	IsPaid                     bool           `json:"is_paid,omitempty" yaml:"is_paid,omitempty"`
	RegTime                    *time.Time     `json:"reg_time,omitempty" yaml:"reg_time,omitempty"`
	SystemFolders              *SystemFolders `json:"system_folders,omitempty" yaml:"system_folders,omitempty"`
	User                       *User          `json:"user,omitempty" yaml:"user,omitempty"`
	UnlimitedAutouploadEnabled bool           `json:"unlimited_autoupload_enabled,omitempty" yaml:"unlimited_autoupload_enabled,omitempty"`
	Revision                   int            `json:"revision,omitempty" yaml:"revision,omitempty"`
}

type SystemFolders struct {
	Odnoklassniki string `json:"odnoklassniki,omitempty" yaml:"odnoklassniki,omitempty"`
	Google        string `json:"google,omitempty" yaml:"google,omitempty"`
	Instagram     string `json:"instagram,omitempty" yaml:"instagram,omitempty"`
	Vkontakte     string `json:"vkontakte,omitempty" yaml:"vkontakte,omitempty"`
	Attach        string `json:"attach,omitempty" yaml:"attach,omitempty"`
	Mailru        string `json:"mailru,omitempty" yaml:"mailru,omitempty"`
	Downloads     string `json:"downloads,omitempty" yaml:"downloads,omitempty"`
	Applications  string `json:"applications,omitempty" yaml:"applications,omitempty"`
	Facebook      string `json:"facebook,omitempty" yaml:"facebook,omitempty"`
	Social        string `json:"social,omitempty" yaml:"social,omitempty"`
	Messenger     string `json:"messenger,omitempty" yaml:"messenger,omitempty"`
	Calendar      string `json:"calendar,omitempty" yaml:"calendar,omitempty"`
	Photostream   string `json:"photostream,omitempty" yaml:"photostream,omitempty"`
	Screenshots   string `json:"screenshots,omitempty" yaml:"screenshots,omitempty"`
	Scans         string `json:"scans,omitempty" yaml:"scans,omitempty"`
}

type User struct {
	RegTime     *time.Time `json:"reg_time,omitempty" yaml:"reg_time,omitempty"`
	DisplayName string     `json:"display_name,omitempty" yaml:"display_name,omitempty"`
	UID         string     `json:"uid,omitempty" yaml:"uid,omitempty"`
	Country     string     `json:"country,omitempty" yaml:"country,omitempty"`
	IsChild     bool       `json:"is_child,omitempty" yaml:"is_child,omitempty"`
	Login       string     `json:"login,omitempty" yaml:"login,omitempty"`
}

type FileResourceList struct {
	Items  []Resource `json:"items,omitempty" yaml:"items"`
	Limit  int        `json:"limit,omitempty" yaml:"limit"`
	Offset int        `json:"offset,omitempty" yaml:"offset"`
}

type LastUploadedResourceList struct {
	Items []Resource `json:"items,omitempty" yaml:"items,omitempty"`
	Limit int        `json:"limit,omitempty" yaml:"limit,omitempty"`
}

type PublicResourceList struct {
	Items  []Resource `json:"items"`
	Type   string     `json:"type"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
}

type Resource struct {
	//AntivirusStatus struct {} `json:"antivirus_status,omitempty" yaml:"antivirus_status,omitempty"`
	ResourceID     string        `json:"resource_id,omitempty" yaml:"resource_id,omitempty"`
	Share          *ShareInfo    `json:"share,omitempty" yaml:"share,omitempty"`
	File           string        `json:"file,omitempty" yaml:"file,omitempty"`
	Size           int           `json:"size,omitempty" yaml:"size,omitempty"`
	PhotosliceTime *time.Time    `json:"photoslice_time,omitempty" yaml:"photoslice_time,omitempty"`
	Embedded       *ResourceList `json:"_embedded,omitempty" yaml:"_embedded,omitempty"`
	Exif           *Exif         `json:"exif,omitempty" yaml:"exif,omitempty"`
	//CustomProperties struct {} `json:"custom_properties,omitempty" yaml:"custom_properties,omitempty"`
	MediaType  string      `json:"media_type,omitempty" yaml:"media_type,omitempty"`
	Preview    string      `json:"preview,omitempty" yaml:"preview,omitempty"`
	Type       string      `json:"type,omitempty" yaml:"type,omitempty"`
	MimeType   string      `json:"mime_type,omitempty" yaml:"mime_type,omitempty"`
	Revision   int         `json:"revision,omitempty" yaml:"revision,omitempty"`
	PublicURL  string      `json:"public_url,omitempty" yaml:"public_url,omitempty"`
	Path       string      `json:"path,omitempty" yaml:"path,omitempty"`
	MD5        string      `json:"md5,omitempty" yaml:"md5,omitempty"`
	PublicKey  string      `json:"public_key,omitempty" yaml:"public_key,omitempty"`
	Sha256     string      `json:"sha256,omitempty" yaml:"sha256,omitempty"`
	Name       string      `json:"name,omitempty" yaml:"name,omitempty"`
	Created    *time.Time  `json:"created,omitempty" yaml:"created,omitempty"`
	Sizes      []Preview   `json:"sizes,omitempty" yaml:"sizes,omitempty"`
	Modified   *time.Time  `json:"modified,omitempty" yaml:"modified,omitempty"`
	CommentIDs *CommentIDs `json:"comment_ids,omitempty" yaml:"comment_ids,omitempty"`
}

type ShareInfo struct {
	IsRoot  bool   `json:"is_root,omitempty" yaml:"is_root,omitempty"`
	IsOwned bool   `json:"is_owned,omitempty" yaml:"is_owned,omitempty"`
	Rights  string `json:"rights,omitempty" yaml:"rights,omitempty"`
}

type ResourceList struct {
	Sort   string     `json:"sort,omitempty" yaml:"sort,omitempty"`
	Items  []Resource `json:"items,omitempty" yaml:"items,omitempty"`
	Limit  int        `json:"limit,omitempty" yaml:"limit,omitempty"`
	Offset int        `json:"offset,omitempty" yaml:"offset,omitempty"`
	Path   string     `json:"path,omitempty" yaml:"path,omitempty"`
	Total  int        `json:"total,omitempty" yaml:"total,omitempty"`
}

type Preview struct {
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type Exif struct {
	DateTime *time.Time `json:"date_time,omitempty" yaml:"date_time,omitempty"`
	//GpsLongitude struct {} `json:"gps_longitude,omitempty" yaml:"gps_longitude,omitempty"`
	//GpsLatitude struct {} `json:"gps_latitude,omitempty" yaml:"gps_latitude,omitempty"`
}

type CommentIDs struct {
	PrivateResource string `json:"private_resource,omitempty" yaml:"private_resource,omitempty"`
	PublicResource  string `json:"public_resource,omitempty" yaml:"public_resource,omitempty"`
}

type ResourceUploadLink struct {
	OperationID string `json:"operation_id,omitempty" yaml:"operation_id,omitempty"`
	URL         string `json:"href,omitempty" yaml:"href,omitempty"`
	Method      string `json:"method,omitempty" yaml:"method,omitempty"`
	Templated   bool   `json:"templated,omitempty" yaml:"templated,omitempty"`
}

type Link struct {
	URL       string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

type Operation struct {
	Status string `json:"status"`
}

type Error struct {
	Code        int
	Message     string `json:"message"`
	Description string `json:"description"`
	Err         string `json:"error"`
}

func (e Error) Error() string {
	return fmt.Sprintf("(%d) %s", e.Code, e.Description)
}
