package XMLParser

func ParseWinnersXml() {
   // Open the xmlFile
   xmlFile, err := os.Open("/home/ivan/workspace-go/src/go-demo/data/winners.xml")
   // if os.Open returns an error then handle it
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println("\tSuccessfully opened winners.xml")
   // defer the closing of xmlFile so that we can parse it.
   defer xmlFile.Close()
   byteValue, _ := ioutil.ReadAll(xmlFile)
   // Unmarshal takes a []byte and fills the rss struct with the values found in the xmlFile
   rss := Rss{}
   xml.Unmarshal(byteValue, &rss)
   fmt.Println("Rss version: " + rss.Version)
   for i, item := range rss.Channel.Items {
      fmt.Println(item.Description)
   }
}
