package types

func init() {
	RegisterAnyType(&Image{})
	RegisterAnyType(&Video{})
	RegisterAnyType(&Audio{})
	RegisterAnyType(&WebPage{})
	RegisterAnyType(&File{})
}

type Image struct {
	Link      string `json:"link"`
	Width     int    `json:"w,omitempty"`
	Height    int    `json:"h,omitempty"`
	Format    string `json:"fmt,omitempty"`
	Size      int    `json:"size,omitempty"`
	Name      string `json:"name,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Data      []byte `json:"-"`
}

type Video struct {
	Link     string `json:"link"`
	Format   string `json:"fmt,omitempty"`
	Duration int    `json:"duration,omitempty"`
	Size     int    `json:"size,omitempty"`
	Image    *Image `json:"img,omitempty"`
	Name     string `json:"name,omitempty"`
	Data     []byte `json:"-"`
}

type Audio struct {
	Link     string `json:"link"`
	Format   string `json:"fmt,omitempty"`
	Duration int    `json:"duration,omitempty"`
	Size     int    `json:"size,omitempty"`
	Name     string `json:"name,omitempty"`
	Data     []byte `json:"-"`
}

type File struct {
	Link   string `json:"link"`
	Name   string `json:"name"`
	Size   int    `json:"size,omitempty"`
	Format string `json:"fmt,omitempty"`
	Data   []byte `json:"-"`
}

type WebPage struct {
	Title   string `json:"title,omitempty"`
	Summary string `json:"summary,omitempty"`
	Image   *Image `json:"image,omitempty"`
	Link    string `json:"link"`
}
