package translator

import (
	"fmt"
	"strings"
)

type refCounterEntry struct {
	count int
	match bool
}

type refCounter struct {
	numRefs int
	m       map[string]*refCounterEntry
}

func NewRefCounter() *refCounter {
	return &refCounter{
		numRefs: 0,
		m:       make(map[string]*refCounterEntry),
	}
}

func (r *refCounter) addRef(id string) int {
	if r.m[id] == nil {
		r.numRefs++
		r.m[id] = &refCounterEntry{r.numRefs, false}
	}
	return r.m[id].count
}

func (r *refCounter) addDeref(id string) (int, error) {
	if r.m[id] == nil {
		return -1, fmt.Errorf("dereferencing unknown id %s", id)
	}
	r.m[id].match = true
	return r.m[id].count, nil
}

func (r *refCounter) allMatched() error {
	for id, e := range r.m {
		if !e.match {
			return fmt.Errorf("unmatched reference \"%s\" ", id)
		}
	}
	return nil
}

const MaxSection = 64 + 1

type sectionCounter struct {
	titles map[int]string
	//        a     => titles[a]
	subtitles map[int]map[int]string
	// 		  a.b   => subtitles[a][b]
	headings map[int]map[int]map[int]string
	// 	      a.b.c => headings[a][b][c]
	section [3]int
}

func NewSectionCounter() *sectionCounter {
	return &sectionCounter{
		titles:    make(map[int]string),
		subtitles: make(map[int]map[int]string),
		headings:  make(map[int]map[int]map[int]string),
	}
}

func (c *sectionCounter) GetSection() string {
	return fmt.Sprintf("%d.%d.%d", c.GetTitleNum(), c.GetSubtitleNum(), c.GetHeadingNum())
}

func (c *sectionCounter) GetSectionDisplayId() string {
	// removes the trailing dot and zeros
	// I feel so smart!
	return strings.TrimRight(c.GetSection(), "0.")
}

func (c *sectionCounter) GetSectionId() string {
	return "h" + c.GetSection()
}

func (c *sectionCounter) AddTitle(title string) {
	c.section[1] = 0
	c.section[2] = 0
	c.titles[c.section[0]+1] = title
	c.section[0]++
}

func (c *sectionCounter) GetTitleNum() int {
	return c.section[0]
}

func (c *sectionCounter) AddSubtitle(subtitle string) {
	c.section[2] = 0
	if c.subtitles[c.section[0]] == nil {
		c.subtitles[c.section[0]] = make(map[int]string)
	}
	c.subtitles[c.section[0]][c.section[1]+1] = subtitle
	c.section[1]++
}

func (c *sectionCounter) GetSubtitleNum() int {
	return c.section[1]
}

func (c *sectionCounter) AddHeading(heading string) {
	if c.headings[c.section[0]] == nil {
		c.headings[c.section[0]] = make(map[int]map[int]string)
	}
	if c.headings[c.section[0]][c.section[1]] == nil {
		c.headings[c.section[0]][c.section[1]] = make(map[int]string)
	}
	c.headings[c.section[0]][c.section[1]][c.section[2]+1] = heading
	c.section[2]++
}

func (c *sectionCounter) GetHeadingNum() int {
	return c.section[2]
}
