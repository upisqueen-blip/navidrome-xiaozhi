package subsonic

import (
    "fmt"
    "net/http"
)

import "github.com/navidrome/navidrome/server/subsonic/responses"

func (api *Router) GetSongCountRaw(w http.ResponseWriter, r *http.Request) (*responses.Subsonic, error) {
    count, err := api.ds.MediaFile(r.Context()).CountAll()
    if err != nil {
        return nil, err
    }
    w.Header().Set("Content-Type", "application/json")
    _, err = w.Write([]byte(fmt.Sprintf("{\"count\": %d}", count)))
    if err != nil {
        return nil, err
    }
    return nil, nil
}

