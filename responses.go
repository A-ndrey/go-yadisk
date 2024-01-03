package yadisk

import (
	"fmt"
	"time"
)

type Disk struct {
	PaidMaxFileSize            int            `json:"paid_max_file_size,omitempty"`
	MaxFileSize                int            `json:"max_file_size,omitempty"`
	TotalSpace                 int            `json:"total_space,omitempty"`
	TrashSize                  int            `json:"trash_size,omitempty"`
	UsedSpace                  int            `json:"used_space,omitempty"`
	IsPaid                     bool           `json:"is_paid,omitempty"`
	RegTime                    *time.Time     `json:"reg_time,omitempty"`
	SystemFolders              *SystemFolders `json:"system_folders,omitempty"`
	User                       *User          `json:"user,omitempty"`
	UnlimitedAutouploadEnabled bool           `json:"unlimited_autoupload_enabled,omitempty"`
	Revision                   int            `json:"revision,omitempty"`
}

type SystemFolders struct {
	Odnoklassniki string `json:"odnoklassniki,omitempty"`
	Google        string `json:"google,omitempty"`
	Instagram     string `json:"instagram,omitempty"`
	Vkontakte     string `json:"vkontakte,omitempty"`
	Attach        string `json:"attach,omitempty"`
	Mailru        string `json:"mailru,omitempty"`
	Downloads     string `json:"downloads,omitempty"`
	Applications  string `json:"applications,omitempty"`
	Facebook      string `json:"facebook,omitempty"`
	Social        string `json:"social,omitempty"`
	Messenger     string `json:"messenger,omitempty"`
	Calendar      string `json:"calendar,omitempty"`
	Photostream   string `json:"photostream,omitempty"`
	Screenshots   string `json:"screenshots,omitempty"`
	Scans         string `json:"scans,omitempty"`
}

type User struct {
	RegTime     *time.Time `json:"reg_time,omitempty"`
	DisplayName string     `json:"display_name,omitempty"`
	UID         string     `json:"uid,omitempty"`
	Country     string     `json:"country,omitempty"`
	IsChild     bool       `json:"is_child,omitempty"`
	Login       string     `json:"login,omitempty"`
}

type FileResourceList struct {
	Items  []Resource `json:"items,omitempty"`
	Limit  int        `json:"limit,omitempty"`
	Offset int        `json:"offset,omitempty"`
}

type LastUploadedResourceList struct {
	Items []Resource `json:"items,omitempty"`
	Limit int        `json:"limit,omitempty"`
}

type PublicResourceList struct {
	Items  []Resource `json:"items,omitempty"`
	Type   string     `json:"type,omitempty"`
	Limit  int        `json:"limit,omitempty"`
	Offset int        `json:"offset,omitempty"`
}

type Resource struct {
	AntivirusStatus  string         `json:"antivirus_status,omitempty"`
	ResourceID       string         `json:"resource_id,omitempty"`
	Share            *ShareInfo     `json:"share,omitempty"`
	File             string         `json:"file,omitempty"`
	Size             int            `json:"size,omitempty"`
	PhotosliceTime   *time.Time     `json:"photoslice_time,omitempty"`
	Embedded         *ResourceList  `json:"_embedded,omitempty"`
	Exif             *Exif          `json:"exif,omitempty"`
	CustomProperties map[string]any `json:"custom_properties,omitempty"`
	MediaType        string         `json:"media_type,omitempty"`
	Preview          string         `json:"preview,omitempty"`
	Type             string         `json:"type,omitempty"`
	MimeType         string         `json:"mime_type,omitempty"`
	Revision         int            `json:"revision,omitempty"`
	PublicURL        string         `json:"public_url,omitempty"`
	Path             string         `json:"path,omitempty"`
	MD5              string         `json:"md5,omitempty"`
	PublicKey        string         `json:"public_key,omitempty"`
	Sha256           string         `json:"sha256,omitempty"`
	Name             string         `json:"name,omitempty"`
	Created          *time.Time     `json:"created,omitempty"`
	Sizes            []Preview      `json:"sizes,omitempty"`
	Modified         *time.Time     `json:"modified,omitempty"`
	CommentIDs       *CommentIDs    `json:"comment_ids,omitempty"`
}

type PublicResource struct {
	AntivirusStatus string                      `json:"antivirus_status,omitempty"`
	ViewsCount      int                         `json:"views_count,omitempty"`
	ResourceID      string                      `json:"resource_id,omitempty"`
	Share           *ShareInfo                  `json:"share,omitempty"`
	File            string                      `json:"file,omitempty"`
	Owner           *UserPublicInformation      `json:"owner,omitempty"`
	Size            int                         `json:"size,omitempty"`
	PhotosliceTime  *time.Time                  `json:"photoslice_time,omitempty"`
	Embedded        *EmbeddedPublicResourceList `json:"_embedded,omitempty"`
	Exif            *Exif                       `json:"exif,omitempty"`
	MediaType       string                      `json:"media_type,omitempty"`
	Preview         string                      `json:"preview,omitempty"`
	Type            string                      `json:"type,omitempty"`
	MimeType        string                      `json:"mime_type,omitempty"`
	Revision        int                         `json:"revision,omitempty"`
	PublicURL       string                      `json:"public_url,omitempty"`
	Path            string                      `json:"path,omitempty"`
	MD5             string                      `json:"md5,omitempty"`
	PublicKey       string                      `json:"public_key,omitempty"`
	SHA256          string                      `json:"sha256,omitempty"`
	Name            string                      `json:"name,omitempty"`
	Created         *time.Time                  `json:"created,omitempty"`
	Sizes           []Preview                   `json:"sizes,omitempty"`
	Modified        *time.Time                  `json:"modified,omitempty"`
	CommentIDs      *CommentIDs                 `json:"comment_ids,omitempty"`
}

type TrashResource struct {
	AntivirusStatus  string         `json:"antivirus_status,omitempty"`
	ResourceID       string         `json:"resource_id,omitempty"`
	Share            *ShareInfo     `json:"share,omitempty"`
	File             string         `json:"file,omitempty"`
	Size             int            `json:"size,omitempty"`
	PhotosliceTime   *time.Time     `json:"photoslice_time,omitempty"`
	Embedded         *ResourceList  `json:"_embedded,omitempty"`
	Exif             *Exif          `json:"exif,omitempty"`
	CustomProperties map[string]any `json:"custom_properties,omitempty"`
	OriginPath       string         `json:"origin_path,omitempty"`
	MediaType        string         `json:"media_type,omitempty"`
	Preview          string         `json:"preview,omitempty"`
	Type             string         `json:"type,omitempty"`
	MimeType         string         `json:"mime_type,omitempty"`
	Revision         int            `json:"revision,omitempty"`
	Deleted          *time.Time     `json:"deleted,omitempty"`
	PublicURL        string         `json:"public_url,omitempty"`
	Path             string         `json:"path,omitempty"`
	MD5              string         `json:"md5,omitempty"`
	PublicKey        string         `json:"public_key,omitempty"`
	Sha256           string         `json:"sha256,omitempty"`
	Name             string         `json:"name,omitempty"`
	Created          *time.Time     `json:"created,omitempty"`
	Sizes            []Preview      `json:"sizes,omitempty"`
	Modified         *time.Time     `json:"modified,omitempty"`
	CommentIDs       *CommentIDs    `json:"comment_ids,omitempty"`
}

type ShareInfo struct {
	IsRoot  bool   `json:"is_root,omitempty"`
	IsOwned bool   `json:"is_owned,omitempty"`
	Rights  string `json:"rights,omitempty"`
}

type ResourceList struct {
	Sort   string     `json:"sort,omitempty"`
	Items  []Resource `json:"items,omitempty"`
	Limit  int        `json:"limit,omitempty"`
	Offset int        `json:"offset,omitempty"`
	Path   string     `json:"path,omitempty"`
	Total  int        `json:"total,omitempty"`
}

type Preview struct {
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

type Exif struct {
	DateTime     *time.Time `json:"date_time,omitempty"`
	GpsLongitude float64    `json:"gps_longitude,omitempty"`
	GpsLatitude  float64    `json:"gps_latitude,omitempty"`
}

type CommentIDs struct {
	PrivateResource string `json:"private_resource,omitempty"`
	PublicResource  string `json:"public_resource,omitempty"`
}

type UserPublicInformation struct {
	Login       string `json:"login,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	UID         string `json:"uid,omitempty"`
}

type EmbeddedPublicResourceList struct {
	Sort      string           `json:"sort,omitempty"`
	PublicKey string           `json:"public_key,omitempty"`
	Items     []PublicResource `json:"items,omitempty"`
	Limit     int              `json:"limit,omitempty"`
	Offset    int              `json:"offset,omitempty"`
	Path      string           `json:"path,omitempty"`
	Total     int              `json:"total,omitempty"`
}

type TrashResourceList struct {
	Sort   string          `json:"sort,omitempty"`
	Items  []TrashResource `json:"items,omitempty"`
	Limit  int             `json:"limit,omitempty"`
	Offset int             `json:"offset,omitempty"`
	Path   string          `json:"path,omitempty"`
	Total  int             `json:"total,omitempty"`
}

type ResourceUploadLink struct {
	OperationID string `json:"operation_id,omitempty"`
	URL         string `json:"href,omitempty"`
	Method      string `json:"method,omitempty"`
	Templated   bool   `json:"templated,omitempty"`
}

type Link struct {
	URL       string `json:"href,omitempty"`
	Method    string `json:"method,omitempty"`
	Templated bool   `json:"templated,omitempty"`
}

type Operation struct {
	Status string `json:"status,omitempty"`
}

type Error struct {
	Code        int
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Err         string `json:"error,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("(%d) %s", e.Code, e.Description)
}
