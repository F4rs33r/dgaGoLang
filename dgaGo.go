package dga

import (
  "fmt"
  "strings"
  "crypto/md5"
  )
  
 type Gen struct {
      day, year, month, seed, i int
      dom string
      lock *sync.Mutex
      
     }
     
func New(day int,year, month, dom String) *Gen {
    return NewSeed(day,year,month, 0, dom)
}


func NewSeed(day, year, month, seed int, dom string ) *Gen {
  if !strings.HasPrefix(dom, "."){
      dom = "." + dom
    }
    
    return *Gen{
    day: day,
    year: year,
    month: month,
    dom: dom,
    seed: seed,
    lock: new(sync.Mutex),
    }
}


func (g *Gen) Next() string{
      g.lock.Lock()
      defer g.lock.Unlock()
      
      g.i++
      
      return fmt.Sprintf("%x%s", md5.Sum([]byte( fmt.Sprintf("%v%v%v%v%v", g.day, g.year,g.month, g.seed, g.i),)), g.dom)
}
}
