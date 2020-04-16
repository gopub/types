package types

func init() {
	RegisterAnyType(&Image{})
	RegisterAnyType(&Video{})
	RegisterAnyType(&Audio{})
	RegisterAnyType(&WebPage{})
	RegisterAnyType(&File{})
}

type Image struct {
	URL       string `json:"url"`
	Width     int    `json:"w,omitempty"`
	Height    int    `json:"h,omitempty"`
	Format    string `json:"fmt,omitempty"`
	Size      int    `json:"size,omitempty"`
	Name      string `json:"name,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Data      []byte `json:"-"`
}

type Video struct {
	URL      string `json:"url"`
	Format   string `json:"fmt,omitempty"`
	Duration int    `json:"duration,omitempty"`
	Size     int    `json:"size,omitempty"`
	Image    *Image `json:"img,omitempty"`
	Name     string `json:"name,omitempty"`
	Data     []byte `json:"-"`
}

type Audio struct {
	URL      string `json:"url"`
	Format   string `json:"fmt,omitempty"`
	Duration int    `json:"duration,omitempty"`
	Size     int    `json:"size,omitempty"`
	Name     string `json:"name,omitempty"`
	Data     []byte `json:"-"`
}

type File struct {
	URL    string `json:"url"`
	Name   string `json:"name"`
	Size   int    `json:"size,omitempty"`
	Format string `json:"fmt,omitempty"`
	Data   []byte `json:"-"`
}

type WebPage struct {
	Title   string `json:"title,omitempty"`
	Summary string `json:"summary,omitempty"`
	Image   *Image `json:"image,omitempty"`
	URL     string `json:"url"`
}
