package filter

import "github.com/windrivder/crawlergo/pkg/model"

type FilterHandler interface {
	DoFilter(req *model.Request) bool
}
