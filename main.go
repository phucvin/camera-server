8package main

import "math"
import "fmt"
import "net/http"
import "path/filepath"
import "os"
import "strings"
import "sort"
import "strconv"

type Pair struct {
        path string
        info os.FileInfo
}

const Prefix = "/home/tom/ftp/upload/"
const PageSize = 20

func min(a, b int) int {
        if a < b {
                return a
        }
        return b
}

func max(a, b int) int {
        if a < b {
                return b
        }
        return a
}

func indexh(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<html>")
        fmt.Fprintf(w, "<head>")
        fmt.Fprintf(w, "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">")
        fmt.Fprintf(w, "<link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css\" integrity=\"sha384-JcKb8q3iqJ61gNV9KGb8thSsNjpSL0n8PARn9HuZOnIxN0hoP+VmmDGMN5t9UJ0Z\" crossorigin=\"anonymous\">")
        fmt.Fprintf(w, "<style type=\"text/css\">a { color:#0caaed !important; } a:hover { color:#6603a8 !important; } a:visited { color:#6603a8 !important; }</style>")
        fmt.Fprintf(w, "<title>Vuon Cau Nguyen Camera History</title>")
        fmt.Fprintf(w, "</head>")
        fmt.Fprintf(w, "<h1>Vuon Cau Nguyen Camera History</h1>")

        var items []Pair
        filepath.Walk(Prefix,
                func(path string, info os.FileInfo, err error) error {
                        if err != nil {
                                return err
                        }
                        if !info.IsDir() && strings.HasSuffix(path, ".mp4") {
                                items = append(items, Pair{path, info})
                        }
                        return nil
                })
        sort.SliceStable(items, func(i, j int) bool {
                return items[i].info.ModTime().After(items[j].info.ModTime())
        })

        page, err := strconv.Atoi(r.URL.Query().Get("page"))
        if err != nil {
                page = 1
        }
        nPage := int(math.Ceil(float64(len(items)) / float64(PageSize)))
        page = min(max(page, 1), nPage)
        fmt.Fprintf(w, "<h3>Page %d / %d (newest first)</h3>", page, nPage)

        fmt.Fprintf(w, "<ol start=\"%d\">", 1+(page-1)*PageSize)
		if len(items) > 0 {
			for _, item := range items[(page-1)*PageSize : min(len(items), page*PageSize)] {
					fmt.Fprintf(w, "<li style=\"margin: 2em 0\"><a href=\"/view?f=%s\">%s</a></li>", item.path[len(Prefix):], item.path[len(Prefix):])
			}
		}
        fmt.Fprintf(w, "</ol>")

        fmt.Fprintf(w, "<div class=\"container\"><div class=\"row\">")
        if page > 1 {
                fmt.Fprintf(w, "<div class=\"col text-center\"><a href=\"/?page=1\">First</a></div>")
                fmt.Fprintf(w, "<div class=\"col text-center\"><a href=\"/?page=%d\">Prev</a></div>", page-1)
        } else {
                fmt.Fprintf(w, "<div class=\"col text-center\">First</div>")
                fmt.Fprintf(w, "<div class=\"col text-center\">Prev</div>")
        }
        if page < nPage {
                fmt.Fprintf(w, "<div class=\"col text-center\"><a href=\"/?page=%d\">Next</a></div>", page+1)
                fmt.Fprintf(w, "<div class=\"col text-center\"><a href=\"/?page=%d\">Last</a></div>", nPage)
        } else {
                fmt.Fprintf(w, "<div class=\"col text-center\">Next</div>")
                fmt.Fprintf(w, "<div class=\"col text-center\">Last</div>")
        }
        fmt.Fprintf(w, "</div></div>")
        fmt.Fprintf(w, "<br/>")

        fmt.Fprintf(w, "</html>")
}

func viewh(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, Prefix+r.URL.Query().Get("f"))
}

func main() {
        http.HandleFunc("/", indexh)
        http.HandleFunc("/view", viewh)
        panic(http.ListenAndServe("0.0.0.0:8000", nil))
}