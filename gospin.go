package gospin

import (
	"fmt"
	"os"
	"time"
)

type Loder struct {
	loading bool
	delay   string
	pettern string
}

func CreateLoder(delay string, pettern string) Loder {
	loder := Loder{false, delay, pettern}
	return loder
}

func (l *Loder) StartLoading(str string) {
	l.loading = true
	fmt.Println(str)
	for l.loading {
		for _, s := range l.pettern {
			fmt.Printf("%c", s)
			complex, err := time.ParseDuration(l.delay)
			if err != nil {
				complex, _ = time.ParseDuration("500ms")
			}
			time.Sleep(complex)
			fmt.Fprint(os.Stdout, "\r")
		}
	}

}

func (l *Loder) StopLoading() {
	l.loading = false
	fmt.Println("")
}
