var hits = expvar.NewInt("hits")

func main() {
	http.HandleFunc("/fub", fub)

	log.Fatal(http.ListenAndServe(":9999", nil))
}

func fub(w http.ResponseWriter, r *http.Request) {
	hits.Add(1)

	for i := 0; i < 50000000; i++ {
		now := time.Now()
		seconds := now.Hour()*3600 + now.Minute()*60 + now.Second()

		var buf bytes.Buffer

		buf.WriteString("seconds so far: ")
		buf.WriteString(strconv.Itoa(seconds))

		buf.WriteTo(ioutil.Discard)
	}
	w.Write([]byte("fub done\n"))
}
