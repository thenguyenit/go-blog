package router

import "errors"

type Route struct {
	key     string
	pattern string
	view    string
	handler string
}

type Router map[string]Route

func initRouter() Router {
	var router = make(Router)

	router["home"] = Route{
		pattern: "/",
		handler: "index",
	}

	router["about-me"] = Route{
		pattern: "/about-me/",
		handler: "ArticleDetail",
	}

	return router
}

func (router Router) search(pattern string) (Route, error) {
	for _, r := range router {
		if r.pattern == pattern {
			return r, nil
		}
	}

	var nilRoute Route
	err := errors.New("A route: " + pattern + " not found")
	return nilRoute, err
}
