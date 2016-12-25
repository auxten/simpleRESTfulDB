package store

import (
    "github.com/bitly/go-simplejson"
    log "github.com/auxten/logrus"
    "io/ioutil"
    "encoding/json"
) 

func Dump(data map[string]interface{}) {
    j, err := json.Marshal(data)
    if err != nil {
        log.Errorf("Marshal json error: %s", err)
        return
    }
    ioutil.WriteFile("dump.db", j, 0644)
}

func Load() map[string]interface{}{
    j, err := ioutil.ReadFile("dump.db")
    if err != nil {
        log.Errorf("Load db error: %s", err)
        return nil
    }
    json_j, err := simplejson.NewJson(j)
    if err != nil {
        log.Errorf("Unmarshal error: %s", err)
        return nil
    }
    
    m, _ := json_j.Map()
    return m
}




