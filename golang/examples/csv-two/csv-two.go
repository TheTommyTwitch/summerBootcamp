package main

import (
  "encoding/csv"
  "fmt"
  "io"
  "os"
  "log"
  //"strconv"
)

type monumentInformation struct {
  columns map[string]int
}

type monument struct {
  id string
  link string
  elevation float64
}

func (info *monumentInformation) setColumns(record []string) {
  for idx, column := range record {
    info.columns[column] = idx
  }
}

func (info *monumentInformation) parseMonument(record []string) (*monument, error) {
  column := info.columns["id"]
  id := record[column]
  link := record[info.columns["link"]]
  elevation := record[info.columns["elevation"]]
  return &monument{
    id: id,
    link: link,
    elevation: elevation,
  }, nil
}

func main()  {
  f, err := os.Open("info.csv")
  if err != nil {
    log.Fatalln(err)
  }
  defer f.Close()

  monumentLookup := map[string]*monument{}

  info := &monumentInformation {
    columns: make(map[string]int),
  }

  csvReader := csv.NewReader(f)
  for rowCount := 0; ; rowCount++ {
    record, err := csvReader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      log.Fatalln(err)
    }

    if rowCount == 0 {
      info.setColumns(record)
    } else {
      monument, err := info.parseMonument(record)
      if err != nil {
        log.Fatalln(err)
      }
      monumentLookup[monument.id] = monument
    }
  }


  var monument monument

  fmt.Println(`
<html>
    <head></head>
    <body>
      <table>
        <tr>
          <th>Abbreviation</th>
          <th>Name</th>
        </tr>`)

	fmt.Println(`
        <tr>
          <td>` + monument.id + `</td>
          <td>` + monument.link + `</td>
        </tr>
    `)

	fmt.Println(`
      </table>
    </body>
</html>
    `)

}
