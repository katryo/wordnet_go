package main

import (
	"fmt"
	"os"
)

func loadWordsWithLemmma(lemma string) ([]string) {

}

func main() {
    if len(os.Args) >= 3 {
        words = loadWordsWithLemma(sys.argv[1])
        if words:
            with SenseLoader() as sense_loader:
                senses = sense_loader.load_senses_with_synset(words[0])

            if len(sys.argv) >= 3:
                link = sys.argv[2]
            else:
                link = 'hypo'

            if len(sys.argv) == 4:
                lang = sys.argv[3]
            else:
                lang = 'jpn'

            print_synlinks_recursively(senses, link, lang)
            sys.exit()

        print("(nothing found)", file=sys.stderr)
        sys.exit()

    }
        // 例(example)
        //  python wn.py 夢
}

func add(x, y int) (sum int, later int) {
	sum = x + y
	later = x
	return
}
